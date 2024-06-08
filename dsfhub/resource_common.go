package dsfhub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"reflect"
	"strings"
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
					log.Printf("[DEBUG] Unknown type for Data field d.GetOk(%v)\n", schemaField.ID)
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
					log.Printf("[DEBUG] Unknown type for AssetData field d.GetOk(%v)\n", schemaField.ID)
				}
			}
		} else {
			log.Printf("[DEBUG] Field not found in AssetData assetSchema, assetSchema.Details[%v]: %v", curStructField.Name, assetSchema.Details[curStructField.Name])
		}
	}

	// Check to see if AWS serverType, populate arn from asset_id
	if ok := strings.HasPrefix(dsfDataSource.Data.ServerType, "AWS"); ok {
		dsfDataSource.Data.AssetData.Arn = dsfDataSource.Data.AssetData.AssetID
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
		return false, fmt.Errorf("[DEBUG] Unsupported serverType: %s\n", serverType)
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
			return false, fmt.Errorf("[DEBUG] Unsupported authMechanism '%v' for serverType '%v'\n", authMechanism, serverType)
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
		return false, fmt.Errorf("[DEBUG] Missing required fields for dsf_data_source with serverType '%s', missing fields: %s\n", serverType, "\""+strings.Join(missingParams, ", ")+"\"")
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

// AssetData  resource hash functions
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
