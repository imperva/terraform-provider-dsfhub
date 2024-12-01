package dsfhub

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func createResource(dsfDataSource *ResourceWrapper, serverType string, d *schema.ResourceData) {
	assetSchema := getSchema()
	//  Iterate through dsfDataSourceData.Data struct fields, retrieve value from d.get() using schema field.id
	structDataFieldsAry := reflect.Indirect(reflect.ValueOf(&dsfDataSource.Data))
	structDataFieldKeys := reflect.ValueOf(&dsfDataSource.Data).Elem()
	for i := 0; i < structDataFieldsAry.NumField(); i++ {
		curStructField := structDataFieldsAry.Type().Field(i)
		log.Printf("[DEBUG] checking for Data field in assetSchema: %v\n", curStructField.Name)
		if schemaField, found := assetSchema.Details[curStructField.Name]; found {
			log.Printf("[DEBUG] Data field curStructField.Name '%v' present in assetSchema\n", curStructField.Name)
			if curStructField.Name != "AssetData" {
				//Check to see if field value is set in tf input
				if _, found := d.GetOk(schemaField.ID); found {
					log.Printf("[DEBUG] Get schema Data field by schemaField.ID (%v)\n", schemaField.ID)
					structField := structDataFieldKeys.FieldByName(curStructField.Name)
					populateStructField(&structField, schemaField, d)
				} else {
					log.Printf("[DEBUG] Data field %v not provided in terraform config - not found by d.GetOk(%v)\n", schemaField.ID, schemaField.ID)
				}
			}
		} else {
			log.Printf("[DEBUG] Field not found in Data assetSchema, assetSchema.Details[%v]: %v", curStructField.Name, assetSchema.Details[curStructField.Name])
		}
	}

	//  Iterate through dsfDataSourceData.Data.AssetData struct fields, retrieve value from d.get() using schema field.id
	structAssetDataFieldsAry := reflect.Indirect(reflect.ValueOf(&dsfDataSource.Data.AssetData))
	structAssetDataFieldKeys := reflect.ValueOf(&dsfDataSource.Data.AssetData).Elem()
	for i := 0; i < structAssetDataFieldsAry.NumField(); i++ {
		curStructField := structAssetDataFieldsAry.Type().Field(i)
		log.Printf("[DEBUG] checking for field in assetSchema: %v\n", curStructField.Name)
		if schemaField, found := assetSchema.Details[curStructField.Name]; found {
			log.Printf("[DEBUG] field curStructField.Name '%v' present in assetSchema\n", curStructField.Name)
			if curStructField.Name != "Connections" {
				//Check to see if field value is set in tf input
				if _, found := d.GetOk(schemaField.ID); found {
					structField := structAssetDataFieldKeys.FieldByName(curStructField.Name)
					log.Printf("[DEBUG] Get schema field by schemaField.ID (%v)\n", schemaField.ID)
					if structField.Kind() == reflect.Ptr {
						switch schemaField.ID {
						case "audit_info":
							inputVal := d.Get(schemaField.ID).(*schema.Set)
							log.Printf("[DEBUG] audit_info: %v, %v", schemaField, inputVal)
							for _, schemaFieldInt := range inputVal.List() {
								ai := AuditInfo{}
								schemaField := schemaFieldInt.(map[string]interface{})
								for fieldName, fieldObjInt := range schemaField {
									fieldObj := fieldObjInt.(interface{})
									switch fieldName {
									case "policy_template_name":
										ai.PolicyTemplateName = fieldObj.(string)
									}
								}
								dsfDataSource.Data.AssetData.AuditInfo = &ai
							}
						case "aws_proxy_config":
							inputVal := d.Get(schemaField.ID).(*schema.Set)
							log.Printf("[DEBUG] aws_proxy_config: %v, %v", schemaField, inputVal)
							for _, schemaFieldInt := range inputVal.List() {
								apc := AwsProxyConfig{}
								schemaField := schemaFieldInt.(map[string]interface{})
								for fieldName, fieldObjInt := range schemaField {
									fieldObj := fieldObjInt.(interface{})
									switch fieldName {
									case "http":
										apc.HTTP = fieldObj.(string)
									case "https":
										apc.HTTPS = fieldObj.(string)
									}
								}
								dsfDataSource.Data.AssetData.AwsProxyConfig = &apc
							}
						}
					} else {
						if schemaField.ID == "server_port" {
							log.Printf("[DEBUG] Setting AssetData server_port interface as string (%v)\n", schemaField.ID)
							dsfDataSource.Data.AssetData.ServerPort = d.Get("server_port").(string)
						} else {
							populateStructField(&structField, schemaField, d)
						}
					}
				} else {
					log.Printf("[DEBUG] AssetData field %v not provided in terraform config - not found by d.GetOk(%v)\n", schemaField.ID, schemaField.ID)
				}
			}
		} else {
			log.Printf("[DEBUG] Field not found in AssetData assetSchema, assetSchema.Details[%v]: %v", curStructField.Name, assetSchema.Details[curStructField.Name])
		}
	}

	// Check to see if AWS serverType, populate arn from asset_id
	arn_state := d.Get("arn").(string)
	isAwsServerType := strings.HasPrefix(dsfDataSource.Data.ServerType, "AWS")
	// fmt.Println(arn_state)
	if isAwsServerType {
		if arn_state == "" {
			dsfDataSource.Data.AssetData.Arn = dsfDataSource.Data.AssetData.AssetID
		} else {
			dsfDataSource.Data.AssetData.Arn = arn_state
		}
	}

	//  Iterate through asset_connection blocks in resource input
	var connectionsAry = make([]AssetConnection, 0)
	connections := d.Get("asset_connection").(*schema.Set)
	for _, conn := range connections.List() {
		connection := conn.(map[string]interface{})
		curConnection := AssetConnection{}
		curConnection.Reason = connection["reason"].(string)
		//  Iterate through dsfDataSourceData.Data.AssetData.Connections struct fields, retrieve value from d.get() using schema field.id
		structConnDataFieldsAry := reflect.Indirect(reflect.ValueOf(&curConnection.ConnectionData))
		structConnDataFieldKeys := reflect.ValueOf(&curConnection.ConnectionData).Elem()
		for i := 0; i < structConnDataFieldsAry.NumField(); i++ {
			curStructField := structConnDataFieldsAry.Type().Field(i)
			if schemaField, found := assetSchema.Connections[curStructField.Name]; found {
				// // Check to see if field value is set in tf input
				if _, found := connection[schemaField.ID]; found {
					log.Printf("Check field type and assign to connection, connection[%v]: %v", schemaField.ID, connection[schemaField.ID])
					structField := structConnDataFieldKeys.FieldByName(curStructField.Name)
					paramVal := connection[schemaField.ID]
					if reflect.TypeOf(paramVal) == reflect.TypeOf(&schema.Set{}) {
						switch schemaField.ID {
						case "amazon_secret":
							inputVal := connection[schemaField.ID].(*schema.Set)
							if inputVal.Len() > 0 {
								for _, schemaFieldInt := range inputVal.List() {
									as := Secret{}
									schemaField := schemaFieldInt.(map[string]interface{})
									for fieldName, fieldObjInt := range schemaField {
										fieldObj := fieldObjInt.(interface{})
										switch fieldName {
										case "field_mapping":
											as.FieldMapping = make(map[string]string)
											fieldObj := fieldObj.(map[string]interface{})
											for fmFieldName, fmFieldObj := range fieldObj {
												as.FieldMapping[fmFieldName] = string(fmFieldObj.(string))
											}
										case "secret_asset_id":
											as.SecretAssetID = fieldObj.(string)
										case "secret_name":
											as.SecretName = fieldObj.(string)
										}
									}
									curConnection.ConnectionData.AmazonSecret = &as
								}
							}
						case "credential_fields":
							inputVal := connection[schemaField.ID].(*schema.Set)
							if inputVal.Len() > 0 {
								for _, schemaFieldInt := range inputVal.List() {
									cf := CredentialFields{}
									schemaField := schemaFieldInt.(map[string]interface{})
									for fieldName, fieldObjInt := range schemaField {
										fieldObj := fieldObjInt.(interface{})
										switch fieldName {
										case "secret_asset_id":
											cf.CredentialSource = fieldObj.(string)
										case "secret_name":
											cf.RoleArn = fieldObj.(string)
										}
									}
									curConnection.ConnectionData.CredentialFields = &cf
								}
							}
						case "cyberark_secret":
							inputVal := connection[schemaField.ID].(*schema.Set)
							if inputVal.Len() > 0 {
								for _, schemaFieldInt := range inputVal.List() {
									hs := Secret{}
									schemaField := schemaFieldInt.(map[string]interface{})
									for fieldName, fieldObjInt := range schemaField {
										fieldObj := fieldObjInt.(interface{})
										switch fieldName {
										case "field_mapping":
											hs.FieldMapping = make(map[string]string)
											fieldObj := fieldObj.(map[string]interface{})
											for fmFieldName, fmFieldObj := range fieldObj {
												hs.FieldMapping[fmFieldName] = string(fmFieldObj.(string))
											}
										case "path":
											hs.Path = fieldObj.(string)
										case "secret_asset_id":
											hs.SecretAssetID = fieldObj.(string)
										case "secret_name":
											hs.SecretName = fieldObj.(string)
										}
									}
									curConnection.ConnectionData.CyberarkSecret = &hs
								}
							}
						case "hashicorp_secret":
							inputVal := connection[schemaField.ID].(*schema.Set)
							if inputVal.Len() > 0 {
								for _, schemaFieldInt := range inputVal.List() {
									hs := Secret{}
									schemaField := schemaFieldInt.(map[string]interface{})
									for fieldName, fieldObjInt := range schemaField {
										fieldObj := fieldObjInt.(interface{})
										switch fieldName {
										case "field_mapping":
											hs.FieldMapping = make(map[string]string)
											fieldObj := fieldObj.(map[string]interface{})
											for fmFieldName, fmFieldObj := range fieldObj {
												hs.FieldMapping[fmFieldName] = string(fmFieldObj.(string))
											}
										case "path":
											hs.Path = fieldObj.(string)
										case "secret_asset_id":
											hs.SecretAssetID = fieldObj.(string)
										case "secret_name":
											hs.SecretName = fieldObj.(string)
										}
									}
									curConnection.ConnectionData.HashicorpSecret = &hs
								}
							}
						case "oauth_parameters":
							inputVal := connection[schemaField.ID].(*schema.Set)
							if inputVal.Len() > 0 {
								for _, schemaFieldInt := range inputVal.List() {
									op := OauthParameters{}
									schemaField := schemaFieldInt.(map[string]interface{})
									for fieldName, fieldObjInt := range schemaField {
										fieldObj := fieldObjInt.(interface{})
										switch fieldName {
										case "parameter":
											op.Parameter = fieldObj.(string)
										}
									}
									curConnection.ConnectionData.OauthParameters = &op
								}
							}
						}
					} else {
						if reflect.TypeOf(paramVal) != nil {
							switch value := reflect.TypeOf(paramVal); value.Kind() {
							case reflect.Int:
								log.Printf("[DEBUG] schemaField.ID %v, Type=Int: %v\n", schemaField.ID, value)
								value := connection[schemaField.ID].(int)
								structField.SetInt(int64(value))
							case reflect.Float64:
								log.Printf("[DEBUG] schemaField.ID %v, Type=Float: %v\n", schemaField.ID, value)
								value := connection[schemaField.ID].(float64)
								structField.SetFloat(value)
							case reflect.String:
								log.Printf("[DEBUG] schemaField.ID %v, Type=String: %v\n", schemaField.ID, value)
								value := connection[schemaField.ID].(string)
								structField.SetString(value)
							case reflect.Bool:
								log.Printf("[DEBUG] schemaField.ID %v, Type=Bool: %v\n", schemaField.ID, value)
								value := connection[schemaField.ID].(bool)
								structField.SetBool(value)
							//case reflect.Slice:
							//	value := d.Get(schemaField.ID)
							//	structField.SetBool(value)
							//	log.Printf("Slice: %v\n", value)
							//	// Handle slices or arrays here
							//case reflect.Map:
							//	log.Printf("Map: %v\n", value)
							//	// Handle maps here
							default:
								log.Printf("[DEBUG] Unknown type for field %v connection[schemaField.ID]: Type:%v\n", schemaField.ID, reflect.TypeOf(paramVal))
							}
						}
					}
				} else {
					log.Printf("[DEBUG] Connection field not found, connection[%v]: %v ", schemaField.ID, connection[schemaField.ID])
				}
			} else {
				log.Printf("[DEBUG] Parsing connection fields, assetSchema.Connections[%v] not found: %v", curStructField.Name, assetSchema.Connections[curStructField.Name])
			}
		}
		connectionsAry = append(connectionsAry, curConnection)
	}
	dsfDataSource.Data.AssetData.Connections = connectionsAry
}

func checkResourceRequiredFields(requiredFieldsJson string, ignoreParamsByServerType map[string]map[string]bool, d *schema.ResourceData) (bool, error) {
	missingParams := []string{}
	var requiredFields RequiredFieldsMap
	err := json.Unmarshal([]byte(requiredFieldsJson), &requiredFields)
	if err != nil {
		log.Printf("[DEBUG] json.Unmarshal([]byte(requiredFieldsJson), &requiredFields) %s:\n", err)
		panic(err)
	}

	serverType := d.Get("server_type").(string)
	serverTypeObj, found := requiredFields.ServerType[serverType]
	if !found {
		return false, fmt.Errorf("unsupported serverType: %s\n", serverType)
	}
	for _, field := range serverTypeObj.Required {
		curField := d.Get(field)
		log.Printf("[DEBUG] Checking for field: '%v', curField: %v, reflect.TypeOf() '%v'\n", field, curField, reflect.ValueOf(d.Get(field)))
		if _, ok := d.GetOk(field); !ok {
			if _, found := ignoreParamsByServerType[serverType][field]; !found {
				missingParams = append(missingParams, field)
				log.Printf("[DEBUG] ERROR: Missing required field '%s' for serverType '%s'\n", field, serverType)
			} else {
				log.Printf("[INFO] Ignoring missing required field '%s' for serverType '%s'\n", field, serverType)
			}
		}
	}

	connections := d.Get("asset_connection").(*schema.Set)
	for _, conn := range connections.List() {
		connection := conn.(map[string]interface{})
		authMechanism := connection["auth_mechanism"].(string)
		log.Printf("[DEBUG] Checking for authMechanism: %s\n", authMechanism)
		authMechanismFields, found := serverTypeObj.AuthMechanisms[authMechanism]
		if !found {
			return false, fmt.Errorf("unsupported authMechanism '%v' for serverType '%v'\n", authMechanism, serverType)
		}
		for _, field := range authMechanismFields {
			log.Printf("[DEBUG] Checking for field: '%s', value: '%s'\n", field, connection[field])
			val := fmt.Sprintf("%v", connection[field])
			if _, found := connection[field]; !found || strings.Trim(val, " ") == "" {
				if _, found := ignoreParamsByServerType[serverType][field]; !found {
					missingParams = append(missingParams, field)
					log.Printf("[DEBUG] Missing required connection field '%s' for serverType '%s' with auth_mechanism '%s'\n", field, serverType, authMechanism)
				} else {
					log.Printf("[INFO] Ignoring missing required connection field '%s' for serverType '%s' with auth_mechanism '%s'\n", field, serverType, authMechanism)
				}
			}
		}
	}
	if len(missingParams) > 0 {
		return false, fmt.Errorf("missing required fields for dsfhub_data_source with serverType '%s', missing fields: %s\n", serverType, "\""+strings.Join(missingParams, ", ")+"\"")
	} else {
		return true, nil
	}
}

func populateStructField(structField *reflect.Value, schemaField SchemaField, d *schema.ResourceData) {
	//log.Printf("structField: %v, d.get: %v", schemaField.ID, d.Get(schemaField.ID))
	if structField.IsValid() {
		if structField.CanSet() {
			switch structField.Kind() {
			case reflect.Int:
				value := d.Get(schemaField.ID).(int)
				log.Printf("[DEBUG] populateStructField value for schemaField.ID: %v int '%v' Type:%v\n", schemaField.ID, value, reflect.TypeOf(value))
				structField.SetInt(int64(value))
			case reflect.Float64:
				value := d.Get(schemaField.ID).(float64)
				log.Printf("[DEBUG] populateStructField value for schemaField.ID: %v float64 '%v' Type:%v\n", schemaField.ID, value, reflect.TypeOf(value))
				structField.SetFloat(value)
			case reflect.String:
				value := d.Get(schemaField.ID).(string)
				log.Printf("[DEBUG] populateStructField value for schemaField.ID: %v string '%v' Type:%v\n", schemaField.ID, value, reflect.TypeOf(value))
				structField.SetString(value)
			case reflect.Bool:
				value := d.Get(schemaField.ID).(bool)
				log.Printf("[DEBUG] populateStructField value for schemaField.ID: %v bool '%v' Type:%v\n", schemaField.ID, value, reflect.TypeOf(value))
				structField.SetBool(value)
			//case reflect.Interface:
			//	value := d.Get(schemaField.ID).(string)
			//	log.Printf("[DEBUG] reflect.Interface %v Type:%v\n", schemaField.ID, reflect.TypeOf(value))
			//	structField.SetString(string(value))
			case reflect.Slice:
				value := d.Get(schemaField.ID).([]interface{})
				for _, v := range value {
					log.Printf("[DEBUG] slice value v: %v\n", v)
				}
				structField.Set(reflect.ValueOf(value))
			//case reflect.Map:
			//	// Handle maps here
			default:
				log.Printf("[DEBUG] populateStructField unknown type for field schemaField.ID:%v Type:%v\n", schemaField.ID, structField.Kind())
			}
		} else {
			log.Printf("[DEBUG] Schema field can not be set, !structField.CanSet(): %v ", schemaField.ID)
		}
	} else {
		log.Printf("[DEBUG] Schema field invalid, structField.IsValid(): %v", schemaField.ID)
	}
}

func getSchema() AssetSchema {
	var assetSchema AssetSchema
	err := json.Unmarshal([]byte(assetSchemaJson), &assetSchema)
	if err != nil {
		log.Printf("[DEBUG] json.Unmarshal([]byte(assetSchemaJson), &assetSchema) %s:\n", err)
		panic(err)
	}
	return assetSchema
}

func contains(l []string, x string) bool {
	for _, a := range l {
		if a == x {
			return true
		}
	}
	return false
}

// readAsset reads an asset of any resource type
func readAsset(client Client, resourceType string, assetId string) (*ResourceWrapper, error) {
	var result *ResourceWrapper
	var err error

	switch resourceType {
	case dsfDataSourceResourceType:
		{
			log.Printf("[INFO] reading data_source asset %v", assetId)
			result, err = client.ReadDSFDataSource(assetId)
		}
	case dsfLogAggregatorResourceType:
		{
			log.Printf("[INFO] reading log_aggregator asset %v", assetId)
			result, err = client.ReadLogAggregator(assetId)
		}
	case dsfCloudAccountResourceType:
		{
			log.Printf("[INFO] reading cloud_account asset %v", assetId)
			result, err = client.ReadSecretManager(assetId)
		}
	case dsfSecretManagerResourceType:
		{
			log.Printf("[INFO] reading secret_manager asset %v", assetId)
			result, err = client.ReadLogAggregator(assetId)
		}
	default:
		{
			return nil, fmt.Errorf("invalid resourceType: %v", resourceType)
		}
	}

	if err != nil {
		return result, err
	}

	return result, nil
}

// waitUntilAuditState reads an asset periodically to check the status of audit_pull_enabled
func waitUntilAuditState(ctx context.Context, desiredState bool, resourceType string, assetId string, m interface{}) error {
	client := m.(*Client)

	pendingState := strconv.FormatBool(!desiredState)
	targetState := strconv.FormatBool(desiredState)

	stateChangeConf := &retry.StateChangeConf{
		Pending: []string{
			pendingState,
		},
		Target: []string{
			targetState,
		},
		Refresh:    auditStateRefreshFunc(*client, resourceType, assetId),
		Timeout:    5 * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 5 * time.Second,
	}

	_, err := stateChangeConf.WaitForStateContext(ctx)
	if err != nil {
		log.Printf("[ERROR] error waiting for audit collection state to update to %v for asset %v", desiredState, assetId)
		return err
	}

	return nil
}

// auditStateRefreshFunc reads an asset to check the status of audit_pull_enabled
func auditStateRefreshFunc(client Client, resourceType string, assetId string) retry.StateRefreshFunc {
	return func() (any, string, error) {
		var result *ResourceWrapper
		var err error

		result, err = readAsset(client, resourceType, assetId)
		if err != nil {
			return 0, "", err
		}

		return result, strconv.FormatBool(result.Data.AssetData.AuditPullEnabled), nil
	}
}

// waitForRemoteSyncState reads an asset periodically to check the status of remoteSyncState becomes "SYNCED"
// posible values for remoteSyncState = ["SYNCED", "NOT_SYNCED", "UNKNOWN"]
func waitForRemoteSyncState(ctx context.Context, resourceType string, assetId string, m interface{}) error {
	client := m.(*Client)

	stateChangeConf := &retry.StateChangeConf{
		Pending: []string{
			"NOT_SYNCED",
			"UNKNOWN",
		},
		Target: []string{
			"SYNCED",
		},
		Refresh:    remoteSyncStateRefreshFunc(*client, resourceType, assetId),
		Timeout:    5 * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 5 * time.Second,
	}

	_, err := stateChangeConf.WaitForStateContext(ctx)
	if err != nil {
		log.Printf("[ERROR] error while waiting for remoteSyncState = \"SYNCED\" for asset %v", assetId)
		return err
	}

	return nil
}

// remoteSyncStateRefreshFunc reads an asset to check the status of remoteSyncState
func remoteSyncStateRefreshFunc(client Client, resourceType string, assetId string) retry.StateRefreshFunc {
	return func() (any, string, error) {
		var result *ResourceWrapper
		var err error

		result, err = readAsset(client, resourceType, assetId)
		if err != nil {
			return 0, "", err
		}

		return result, result.Data.RemoteSyncState, nil
	}
}

// checkAuditState reads an asset to check the status of audit_pull_enabled
func checkAuditState(ctx context.Context, m interface{}, assetId string, resourceType string) (bool, error) {
	client := m.(*Client)
	var auditPullState bool
	var result *ResourceWrapper
	var err error

	result, err = readAsset(*client, resourceType, assetId)
	if err != nil {
		return false, err
	}

	auditPullState = result.Data.AssetData.AuditPullEnabled

	return auditPullState, nil
}

// connectDisconnectGateway determines whether an asset should be connected to or disconnected from gateway
func connectDisconnectGateway(ctx context.Context, d *schema.ResourceData, resourceType string, m interface{}) error {
	assetId := d.Get("asset_id").(string)
	auditPullEnabled := d.Get("audit_pull_enabled").(bool)
	auditType := d.Get("audit_type").(string)
	auditPullEnabledChanged := d.HasChange("audit_pull_enabled")
	auditTypeChanged := d.HasChange("audit_type")

	log.Printf("[DEBUG] connectDisconnectGateway - assetId: %v", assetId)
	log.Printf("[DEBUG] connectDisconnectGateway - auditPullEnabled: %v", auditPullEnabled)
	log.Printf("[DEBUG] connectDisconnectGateway - auditType: %v", auditType)
	log.Printf("[DEBUG] connectDisconnectGateway - auditPullEnabledChanged: %v", auditPullEnabledChanged)
	log.Printf("[DEBUG] connectDisconnectGateway - auditTypeChanged: %v", auditTypeChanged)

	// if audit_pull_enabled has been changed, connect/disconnect from gateway as needed
	if auditPullEnabledChanged {
		if auditPullEnabled {
			err := connectGateway(ctx, m, assetId, resourceType)
			if err != nil {
				return err
			}
		} else {
			err := disconnectGateway(ctx, m, assetId, resourceType)
			if err != nil {
				return err
			}
		}
		// if asset is already connected, check whether relevant fields have been updated and reconnect to gateway
	} else if auditPullEnabled {
		if auditTypeChanged {
			origAuditType, newAuditType := d.GetChange("audit_type")
			log.Printf("[INFO] auditType value has changed from %s to %s, reconnecting asset to gateway\n", origAuditType, newAuditType)
			err := reconnectGateway(ctx, m, assetId, resourceType)
			if err != nil {
				return err
			}
		}
	} else {
		log.Printf("[INFO] Asset %s does not need to be connected to or disconnected from gateway", assetId)
	}
	return nil
}

// connectGateway connects an asset to gateway
func connectGateway(ctx context.Context, m interface{}, assetId string, resourceType string) error {
	client := m.(*Client)
	_, err := client.EnableAuditDSFDataSource(assetId)
	if err != nil {
		log.Printf("[INFO] Error enabling audit for assetId: %s\n", assetId)
		return err
	}

	// ensure asset is synced to gateway
	err = waitForRemoteSyncState(ctx, resourceType, assetId, m)
	if err != nil {
		return err
	}

	// confirm asset is connected to gateway
	isAuditPullEnabled, err := checkAuditState(ctx, m, assetId, resourceType)
	if err != nil || !isAuditPullEnabled {
		return err
	}

	return nil
}

// disconnectGateway disconnects an asset from gateway
func disconnectGateway(ctx context.Context, m interface{}, assetId string, resourceType string) error {
	client := m.(*Client)
	_, err := client.DisableAuditDSFDataSource(assetId)
	if err != nil {
		log.Printf("[INFO] Error disabling audit for assetId: %s\n", assetId)
		return err
	}

	// ensure asset is synced to gateway
	err = waitForRemoteSyncState(ctx, resourceType, assetId, m)
	if err != nil {
		log.Printf("[INFO] Error while waiting for audit state to update for assetId: %s\n", assetId)
		return err
	}

	// confirm asset is disconnected from gateway
	isAuditPullEnabled, err := checkAuditState(ctx, m, assetId, resourceType)
	if err != nil || isAuditPullEnabled {
		return err
	}

	return nil
}

// reconnectGateway first disconnects and then reconnects an asset to gateway
func reconnectGateway(ctx context.Context, m interface{}, assetId string, resourceType string) error {
	log.Printf("[INFO] Re-enabling audit for assetId: %s\n", assetId)

	err := disconnectGateway(ctx, m, assetId, resourceType)
	if err != nil {
		return err
	}

	err = connectGateway(ctx, m, assetId, resourceType)
	if err != nil {
		return err
	}

	return nil
}

// ConnectionData resource hash functions
func resourceConnectionDataAmazonSecretHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	//if v, ok := m["field_mapping"]; ok {
	//	buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	//}
	if v, ok := m["path"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}
	if v, ok := m["secret_asset_id"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}
	if v, ok := m["secret_name"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}
	return PositiveHash(buf.String())
}

func resourceConnectionDataCredentialFieldsHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if v, ok := m["credential_source"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}
	if v, ok := m["role_arn"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}
	return PositiveHash(buf.String())
}

func resourceConnectionDataCyberarkSecretHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	//if v, ok := m["field_mapping"]; ok {
	//	buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	//}
	if v, ok := m["path"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}
	if v, ok := m["secret_asset_id"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}
	if v, ok := m["secret_name"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}
	return PositiveHash(buf.String())
}

func resourceConnectionDataHashicorpSecretHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	//if v, ok := m["field_mapping"]; ok {
	//	buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	//}
	if v, ok := m["path"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}
	if v, ok := m["secret_asset_id"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}
	if v, ok := m["secret_name"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}
	return PositiveHash(buf.String())
}

func resourceConnectionDataOauthParametersHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if v, ok := m["parameter"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}
	return PositiveHash(buf.String())
}

// AssetData resource hash functions
func resourceAssetDataAuditInfoHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if v, ok := m["policy_template_name"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}
	return PositiveHash(buf.String())
}

func resourceAssetDataAWSProxyConfigHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if v, ok := m["http"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}
	if v, ok := m["https"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}
	return PositiveHash(buf.String())
}

func resourceAssetDataServiceEndpointsHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if v, ok := m["logs"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}
	return PositiveHash(buf.String())
}

// testAccParseResourceAttributeReference parses a terraform field and
// determines whether it is a reference to another resource. If the field is
// a reference, return the input string and if not, return it wrapped in
// double-quotes.
func testAccParseResourceAttributeReference(field string) string {
	var regExpr string = `dsfhub_[A-Za-z0-9_-].+\.[A-Za-z0-9_-].+` //e.g. dsfhub_cloud_account.my-cloud-account, dsfhub_cloud_account.my-cloud-account.asset_id
	var parsedField string

	isReference, _ := regexp.Match(regExpr, []byte(field))
	if isReference {
		parsedField = field
	} else {
		parsedField = fmt.Sprintf("\"%s\"", field)
	}
	return parsedField
}

// ConfigCompose can be called to concatenate multiple strings to build test
// configurations
func ConfigCompose(config ...string) string {
	var str strings.Builder

	for _, conf := range config {
		str.WriteString(conf)
	}

	return str.String()
}

// ignoreChangesBlock creates a lifecycle block to be added to a resource config
// containing the ignore_changes feature - an array specifying a list of attribute
// names that may change in the future, but should not affect said resource after
// its creation.
//
// See the following for more details
// https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle#ignore_changes
func createIgnoreChangesBlock(ignoredAttributes []string) string {
	var ignoreChangesBlock, ignoredFields string

	// build list like
	// ignore_changes = [ attribute1, attribute2 ]
	ignoredFields = `[ ` + strings.Join(ignoredAttributes, `, `) + ` ]`
	log.Printf("[INFO] creating ignore_changes list with ignored_fields: %s", ignoredFields)
	ignoreChangesBlock = fmt.Sprintf(`
  lifecycle {
    ignore_changes = %[1]s
  }`, ignoredFields)

	return ignoreChangesBlock
}

// ignoreAssetConnectionBlock builds the lifecycle meta argument block to ignore
// fields masked in the connection.
// e.g. client_secret   = "*****", password = "*****"
func ignoreAssetConnectionChangesBlock() string {
	var lifecycleBlock string
	ignoredAttributes := []string{`asset_connection`}
	lifecycleBlock = createIgnoreChangesBlock(ignoredAttributes)
	return lifecycleBlock
}

// checkGatewayId checks that the GATEWAY_ID environment variable is set correctly
// for acceptance tests
func checkGatewayId(t *testing.T) string {
	gatewayId := os.Getenv("GATEWAY_ID")
	if gatewayId == "" {
		t.Fatal("GATEWAY_ID environment variable must be set")
	}
	return gatewayId
}

// flipAuditPullEnabledConfig takes a config string and flips any instances of
// the "audit_pull_enabled" field
func flipAuditPullEnabledConfig(s string) string {
	regtrue := regexp.MustCompile(`(audit_pull_enabled\s*=\s*)(true)`)
	regfalse := regexp.MustCompile(`(audit_pull_enabled\s*=\s*)(false)`)
	regTRUE := regexp.MustCompile(`(audit_pull_enabled\s*=\s*)(TRUE)`)

	s = regtrue.ReplaceAllString(s, "${1}TRUE") // get em outta the way
	s = regfalse.ReplaceAllString(s, "${1}true")
	s = regTRUE.ReplaceAllString(s, "${1}false")

	return s
}

// validateImportStep returns a test step for validating import on a resource
func validateImportStep(resourceTypeAndName string) resource.TestStep {
	step := resource.TestStep{
		ResourceName:      resourceTypeAndName,
		ImportState:       true,
		ImportStateVerify: true,
	}
	return step
}

// createResourceAttrChecks creates a TestCheckFunc for a map of resources and
// fields to check for each resource
// accepts values like:
// [
//
//	 "resource1": [
//	 	 "field1": "value1",
//		 "field2": "value2",
//	 ],
//	 ...,
//
// ]
func createResourceAttrChecks(attrChecks map[string]map[string]string) resource.TestCheckFunc {
	var checks []resource.TestCheckFunc
	for resourceName, resourceChecks := range attrChecks {
		for field, value := range resourceChecks {
			check := resource.TestCheckResourceAttr(resourceName, field, value)
			checks = append(checks, check)
		}
	}
	return resource.ComposeTestCheckFunc(checks...)
}

// disconnectedResourceAttrChecks updates a list of resource attribute checks
// to account for the resource being disconnected from gateway.
//
// For example, updating "audit_pull_enabled" to "false", or "gateway_service"
// to an empty string.
func disconnectedResourceAttrChecks(attrChecks map[string]map[string]string) resource.TestCheckFunc {
	var checks []resource.TestCheckFunc
	for resourceName, resourceChecks := range attrChecks {
		for field, value := range resourceChecks {
			switch field {
			case "audit_pull_enabled":
				value = "false"
			case "gateway_service":
				value = ""
			}
			check := resource.TestCheckResourceAttr(resourceName, field, value)
			checks = append(checks, check)
		}
	}
	return resource.ComposeTestCheckFunc(checks...)
}

// connectDisconnectGatewaySteps returns test steps to connect a resource to
// gateway and then disconnect it.
func connectDisconnectGatewaySteps(config string, initialAttrChecks map[string]map[string]string, refreshAttrChecks map[string]map[string]string, validate bool, resourceTypeAndName string) []resource.TestStep {
	var steps []resource.TestStep

	// connect to gateway and check attributes
	step := resource.TestStep{
		Config: config,
		Check:  createResourceAttrChecks(initialAttrChecks),
	}
	steps = append(steps, step)

	// refresh the state and check attributes if specified
	if refreshAttrChecks != nil {
		step = resource.TestStep{
			RefreshState: true,
			Check:        createResourceAttrChecks(refreshAttrChecks),
		}
		steps = append(steps, step)
	}

	// disconnect gateway and check atttributes
	var flippedConfig = flipAuditPullEnabledConfig(config)
	var flippedChecks = disconnectedResourceAttrChecks(initialAttrChecks)
	step = resource.TestStep{
		Config: flippedConfig,
		Check:  flippedChecks,
	}
	steps = append(steps, step)

	// refresh the state and check attributes if specified
	if refreshAttrChecks != nil {
		var flippedChecks = disconnectedResourceAttrChecks(initialAttrChecks)
		step = resource.TestStep{
			RefreshState: true,
			Check:        flippedChecks,
		}
		steps = append(steps, step)
	}

	// if specified, validate import
	if validate {
		steps = append(steps, validateImportStep(resourceTypeAndName))
	}
	return steps
}
