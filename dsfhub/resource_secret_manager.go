package dsfhub

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSecretManager() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSecretManagerCreateContext,
		ReadContext:   resourceSecretManagerReadContext,
		UpdateContext: resourceSecretManagerUpdateContext,
		DeleteContext: resourceSecretManagerDeleteContext,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"admin_email": {
				Type:        schema.TypeString,
				Description: "The email address to notify about this asset",
				Required:    true,
			},
			"application": {
				Type:        schema.TypeString,
				Description: "The Asset ID of the application asset that \"owns\" the asset.",
				Required:    false,
				Optional:    true,
			},
			"arn": {
				Type:        schema.TypeString,
				Description: "Amazon Resource Name - format is arn:partition:service:region:account-id and used as the asset_id",
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     nil,
			},
			"asset_display_name": {
				Type:        schema.TypeString,
				Description: "User-friendly name of the asset, defined by user.",
				Required:    false,
				Optional:    true,
			},
			"asset_id": {
				Type:        schema.TypeString,
				Description: "The unique identifier or resource name of the asset. For most assets this should be a concatenation of Server Host Name + Server Type + Service Name + Server Port with “:” (colon) as separator",
				Required:    false,
				Optional:    true,
			},
			"asset_source": {
				Type:        schema.TypeString,
				Description: "The source platform/vendor/system of the asset data. Usually the service responsible for creating that asset document",
				Required:    false,
				Optional:    true,
			},
			"asset_version": {
				Type:        schema.TypeFloat,
				Description: "Denotes the version of the asset",
				Required:    false,
				Optional:    true,
			},
			"available_regions": {
				Type:        schema.TypeList,
				Description: "A list of regions to use in discovery actions that iterate through region",
				Required:    false,
				Optional:    true,
				Default:     nil,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"aws_proxy_config": {
				Type:        schema.TypeSet,
				Description: "AWS specific proxy configuration",
				Required:    false,
				Optional:    true,
				Default:     nil,
				MinItems:    0,
				Set:         resourceAssetDataAWSProxyConfigHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http": {
							Type:        schema.TypeString,
							Description: "HTTP endpoint for aws proxy config",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"https": {
							Type:        schema.TypeString,
							Description: "HTTPS endpoint for aws proxy config",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
					},
				},
			},
			"asset_connection": {
				Type:        schema.TypeSet,
				Description: "N/A",
				Required:    false,
				Optional:    true,
				MinItems:    0,
				Set:         resourceSecretManagerConnectionHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"reason": {
							Type:         schema.TypeString,
							Description:  "N/A",
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"default"}, false),
						},
						"access_id": {
							Type:        schema.TypeString,
							Description: "The Access key ID of AWS secret access key used for authentication",
							Required:    false,
							Optional:    true,
						},
						"amazon_secret": {
							Type:        schema.TypeSet,
							Description: "Configuration to integrate with AWS Secrets Manager",
							Required:    false,
							Optional:    true,
							Default:     nil,
							MinItems:    0,
							MaxItems:    1,
							Set:         resourceConnectionDataAmazonSecretHash,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"field_mapping": {
										Type:        schema.TypeMap,
										Description: "Field mapping for AWS secret",
										Required:    false,
										Optional:    true,
										Default:     nil,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"secret_asset_id": {
										Type:        schema.TypeString,
										Description: "AWS secret manager asset_id",
										Required:    false,
										Optional:    true,
										Default:     nil,
									},
									"secret_name": {
										Type:        schema.TypeString,
										Description: "AWS secret name",
										Required:    false,
										Optional:    true,
										Default:     nil,
									},
								},
							},
						},
						"auth_mechanism": {
							Type:         schema.TypeString,
							Description:  "Specifies the auth mechanism used by the connection",
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"default", "root_token", "iam_role", "app_role", "ec2", "profile", "key"}, false),
						},
						"aws_iam_server_id": {
							Type:        schema.TypeString,
							Description: "e.g. vault.example.com",
							Required:    false,
							Optional:    true,
						},
						"ca_certs_path": {
							Type:        schema.TypeString,
							Description: "Certificate authority certificates path; what location should the sysetm look for certificate information from. Equivalent to --capath in a curl call",
							Required:    false,
							Optional:    true,
						},
						"cert_file": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
						},
						"cyberark_secret": {
							Type:        schema.TypeSet,
							Description: "Configuration to integrate with CyberArk",
							Required:    false,
							Optional:    true,
							Default:     nil,
							MinItems:    0,
							MaxItems:    1,
							Set:         resourceConnectionDataCyberarkSecretHash,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"field_mapping": {
										Type:        schema.TypeMap,
										Description: "Field mapping for CyberArk secret",
										Required:    false,
										Optional:    true,
										Default:     nil,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"secret_asset_id": {
										Type:        schema.TypeString,
										Description: "CyberArk secret manager asset_id",
										Required:    false,
										Optional:    true,
										Default:     nil,
									},
									"secret_name": {
										Type:        schema.TypeString,
										Description: "CyberArk secret name",
										Required:    false,
										Optional:    true,
										Default:     nil,
									},
								},
							},
						},
						"external_id": {
							Type:        schema.TypeString,
							Description: "External ID to use when assuming a role",
							Required:    false,
							Optional:    true,
						},
						"hashicorp_secret": {
							Type:        schema.TypeSet,
							Description: "Configuration to integrate with HashiCorp Vault",
							Required:    false,
							Optional:    true,
							Default:     nil,
							MinItems:    0,
							Set:         resourceConnectionDataHashicorpSecretHash,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"field_mapping": {
										Type:        schema.TypeMap,
										Description: "Field mapping for HashiCorp secret",
										Required:    false,
										Optional:    true,
										Default:     nil,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"path": {
										Type:        schema.TypeString,
										Description: "HashiCorp secret path",
										Required:    false,
										Optional:    true,
										Default:     nil,
									},
									"secret_asset_id": {
										Type:        schema.TypeString,
										Description: "HashiCorp secret manager asset_id",
										Required:    false,
										Optional:    true,
										Default:     nil,
									},
									"secret_name": {
										Type:        schema.TypeString,
										Description: "HashiCorp secret name",
										Required:    false,
										Optional:    true,
										Default:     nil,
									},
								},
							},
						},
						"headers": {
							Type:        schema.TypeList,
							Description: "A list of headers to include in the request",
							Required:    false,
							Optional:    true,
							Default:     nil,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"key_file": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
						},
						"namespace": {
							Type:        schema.TypeString,
							Description: "Specifies which namespace to fetch credentials from if not root.",
							Required:    false,
							Optional:    true,
						},
						"nonce": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
						},
						"protocol": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
						},
						"query": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
						},
						"region": {
							Type:        schema.TypeString,
							Description: "Default AWS region for this asset",
							Required:    false,
							Optional:    true,
						},
						"role_name": {
							Type:        schema.TypeString,
							Description: "Role to use for authentication",
							Required:    false,
							Optional:    true,
						},
						"secret_key": {
							Type:        schema.TypeString,
							Description: "The Secret access key used for authentication",
							Required:    false,
							Optional:    true,
						},
						"self_signed": {
							Type:        schema.TypeBool,
							Description: "",
							Required:    false,
							Optional:    true,
						},
						"session_token": {
							Type:        schema.TypeString,
							Description: "STS token used for session authentication",
							Required:    false,
							Optional:    true,
							Sensitive:   true,
						},
						"ssl": {
							Type:        schema.TypeBool,
							Description: "If true, use SSL when connecting",
							Required:    false,
							Optional:    true,
						},
						"store_aws_credentials": {
							Type:        schema.TypeBool,
							Description: "",
							Required:    false,
							Optional:    true,
						},
						"username": {
							Type:        schema.TypeString,
							Description: "The name of a profile in /imperva/local/credentials/.aws/credentials to use for authenticating",
							Required:    false,
							Optional:    true,
						},
						"v2_key_engine": {
							Type:        schema.TypeBool,
							Description: "Use a KV2 secret engine",
							Required:    false,
							Optional:    true,
						},
					},
				},
			},
			"credentials_endpoint": {
				Type:        schema.TypeString,
				Description: "A specific sts endpoint to use",
				Required:    false,
				Optional:    true,
			},
			"criticality": {
				Type:         schema.TypeInt,
				Description:  "The asset's importance to the business. These values are measured on a scale from \"Most critical\" (1) to \"Least critical\" (4). Allowed values: 1, 2, 3, 4",
				Required:     false,
				Optional:     true,
				ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4}),
			},
			"gateway_id": {
				Type:        schema.TypeString,
				Description: "The jsonarUid unique identifier of the agentless gateway. Example: '7a4af7cf-4292-89d9-46ec-183756ksdjd'",
				Required:    true,
			},
			"id": {
				Type:        schema.TypeString,
				Description: "Unique identifier for the asset",
				Computed:    true,
			},
			"jsonar_uid": {
				Type:        schema.TypeString,
				Description: "Unique identifier (UID) attached to the Agentless Gateway controlling the asset",
				Required:    false,
				Optional:    true,
				Computed:    true,
			},
			"jsonar_uid_display_name": {
				Type:        schema.TypeString,
				Description: "Unique identifier (UID) attached to the Agentless Gateway controlling the asset",
				Required:    false,
				Optional:    true,
				Computed:    true,
			},
			"location": {
				Type:        schema.TypeString,
				Description: "Current human-readable description of the physical location of the asset, or region.",
				Required:    false,
				Optional:    true,
			},
			"managed_by": {
				Type:        schema.TypeString,
				Description: "Email of the person who maintains the asset; can be different from the owner specified in the owned_by field. Defaults to admin_email.",
				Required:    false,
				Optional:    true,
				Computed:    true,
			},
			"owned_by": {
				Type:        schema.TypeString,
				Description: "Email of Owner / person responsible for the asset; can be different from the person in the managed_by field. Defaults to admin_email.",
				Required:    false,
				Optional:    true,
				Computed:    true,
			},
			"proxy": {
				Type:        schema.TypeString,
				Description: "Proxy to use for AWS calls if aws_proxy_config is populated the proxy field will get populated from the http value there",
				Required:    false,
				Optional:    true,
			},
			"region": {
				Type:        schema.TypeString,
				Description: "For cloud systems with regions, the default region or region used with this asset",
				Required:    false,
				Optional:    true,
			},
			"server_host_name": {
				Type:        schema.TypeString,
				Description: "Hostname (or IP if name is unknown)",
				Required:    false,
				Optional:    true,
			},
			"server_ip": {
				Type:        schema.TypeString,
				Description: "IP address of the service where this asset is located. If no IP is available populate this field with other information that would identify the system e.g. hostname or AWS ARN, etc.",
				Required:    false,
				Optional:    true,
			},
			"server_port": {
				Type:        schema.TypeString,
				Description: "",
				Required:    false,
				Optional:    true,
			},
			"server_type": {
				Type:        schema.TypeString,
				Description: "The type of service to be created as a secret manager. Available values include AWS, CYBERARK, and HASHICORP.",
				Required:    true,
			},
			"service_endpoints": {
				Type:        schema.TypeSet,
				Description: "Specify particular endpoints for a given service in the form of <service name>: \"endpoint\"",
				Required:    false,
				Optional:    true,
				Default:     nil,
				MinItems:    0,
				Set:         resourceAssetDataServiceEndpointsHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logs": {
							Type:        schema.TypeString,
							Description: "The log endpoint for a given service",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
					},
				},
			},
			"used_for": {
				Type:         schema.TypeString,
				Description:  "Designates how this asset is used / the environment that the asset is supporting.",
				Required:     false,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Development", "Staging", "Test", "Disaster Recovery", "Demonstration", "Production", "QA", "Training"}, false),
			},
		},
	}
}

func resourceSecretManagerCreateContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Client)

	// check provided fields against schema
	if isOk, err := checkResourceRequiredFields(requiredSecretManagerFieldsJson, ignoreSecretManagerParamsByServerType, d); !isOk {
		return diag.FromErr(err)
	}

	// convert provided fields into API payload
	secretManager := ResourceWrapper{}
	serverType := d.Get("server_type").(string)
	createResource(&secretManager, serverType, d)

	// create resource
	log.Printf("[INFO] Creating SecretManager for serverType: %s and gatewayId: %s gatewayId: \n", serverType, secretManager.Data.GatewayID)
	createSecretManagerResponse, err := client.CreateSecretManager(secretManager)
	if err != nil {
		log.Printf("[ERROR] adding secret manager for serverType: %s and gatewayId: %s | err: %s\n", serverType, secretManager.Data.GatewayID, err)
		return diag.FromErr(err)
	}

	// get asset_id
	assetId := d.Get("asset_id").(string)

	// wait for remoteSyncState
	err = waitForRemoteSyncState(ctx, dsfSecretManagerResourceType, assetId, m)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  fmt.Sprintf("Error while waiting for remoteSyncState = \"SYNCED\" for asset: %s", assetId),
			Detail:   fmt.Sprintf("Error: %s\n", err),
		})
	}

	// set ID
	secretManagerId := createSecretManagerResponse.Data.ID
	d.SetId(secretManagerId)

	// Set the rest of the state from the resource read
	resourceSecretManagerReadContext(ctx, d, m)

	return diags
}

func resourceSecretManagerReadContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	secretManagerId := d.Id()

	log.Printf("[INFO] Reading secret manager with secretManagerId: %s\n", secretManagerId)

	secretManagerReadResponse, err := client.ReadSecretManager(secretManagerId)

	if err != nil {
		log.Printf("[ERROR] Reading secretManagerReadResponse with secretManagerId: %s | err: %s\n", secretManagerId, err)
		return diag.FromErr(err)
	}

	if secretManagerReadResponse != nil {
		log.Printf("[INFO] Reading SecretManager with secretManagerId: %s | err: %s\n", secretManagerId, err)
	}

	log.Printf("[DEBUG] secretManagerReadResponse: %s\n", secretManagerReadResponse.Data.AssetData.AssetID)
	// Set returned and computed values
	d.Set("admin_email", secretManagerReadResponse.Data.AssetData.AdminEmail)
	d.Set("application", secretManagerReadResponse.Data.AssetData.Application)
	d.Set("arn", secretManagerReadResponse.Data.AssetData.Arn)
	d.Set("asset_display_name", secretManagerReadResponse.Data.AssetData.AssetDisplayName)
	d.Set("asset_id", secretManagerReadResponse.Data.AssetData.AssetID)
	d.Set("asset_source", secretManagerReadResponse.Data.AssetData.AssetSource)
	d.Set("asset_version", secretManagerReadResponse.Data.AssetData.Version)
	d.Set("available_regions", secretManagerReadResponse.Data.AssetData.AvailableRegions)
	d.Set("credentials_endpoint", secretManagerReadResponse.Data.AssetData.CredentialsEndpoint)
	d.Set("criticality", secretManagerReadResponse.Data.AssetData.Criticality)
	d.Set("gateway_id", secretManagerReadResponse.Data.GatewayID)
	d.Set("id", secretManagerReadResponse.Data.ID)
	d.Set("jsonar_uid", secretManagerReadResponse.Data.AssetData.JsonarUID)
	d.Set("location", secretManagerReadResponse.Data.AssetData.Location)
	d.Set("managed_by", secretManagerReadResponse.Data.AssetData.ManagedBy)
	d.Set("owned_by", secretManagerReadResponse.Data.AssetData.OwnedBy)
	d.Set("proxy", secretManagerReadResponse.Data.AssetData.Proxy)
	d.Set("region", secretManagerReadResponse.Data.AssetData.Region)
	d.Set("server_host_name", secretManagerReadResponse.Data.AssetData.ServerHostName)
	d.Set("server_ip", secretManagerReadResponse.Data.AssetData.ServerIP)
	if secretManagerReadResponse.Data.AssetData.ServerPort != nil {
		var serverPort string
		if serverPortNum, ok := secretManagerReadResponse.Data.AssetData.ServerPort.(float64); ok {
			serverPort = fmt.Sprintf("%d", int(serverPortNum))
		} else {
			serverPort = secretManagerReadResponse.Data.AssetData.ServerPort.(string)
		}
		d.Set("server_port", serverPort)
	}
	d.Set("server_type", secretManagerReadResponse.Data.ServerType)
	d.Set("used_for", secretManagerReadResponse.Data.AssetData.UsedFor)

	if secretManagerReadResponse.Data.AssetData.AwsProxyConfig != nil {
		awsProxyConfig := &schema.Set{F: resourceAssetDataAWSProxyConfigHash}
		awsProxyConfigMap := map[string]interface{}{}
		awsProxyConfigMap["http"] = secretManagerReadResponse.Data.AssetData.AwsProxyConfig.HTTP
		awsProxyConfigMap["https"] = secretManagerReadResponse.Data.AssetData.AwsProxyConfig.HTTPS
		awsProxyConfig.Add(awsProxyConfigMap)
		d.Set("aws_proxy_config", awsProxyConfig)
	}

	if secretManagerReadResponse.Data.AssetData.ServiceEndpoints != nil {
		serviceEndpoints := &schema.Set{F: resourceAssetDataServiceEndpointsHash}
		serviceEndpointsMap := map[string]interface{}{}
		serviceEndpointsMap["logs"] = secretManagerReadResponse.Data.AssetData.ServiceEndpoints.Logs
		serviceEndpoints.Add(serviceEndpointsMap)
		d.Set("service_endpoints", serviceEndpoints)
	}

	connections := &schema.Set{F: resourceSecretManagerConnectionHash}
	for _, v := range secretManagerReadResponse.Data.AssetData.Connections {
		connection := map[string]interface{}{}
		connection["access_id"] = v.ConnectionData.AccessID
		connection["aws_iam_server_id"] = v.ConnectionData.AwsIamServerID
		connection["ca_certs_path"] = v.ConnectionData.CaCertsPath
		connection["cert_file"] = v.ConnectionData.CaFile
		connection["external_id"] = v.ConnectionData.ExternalID
		connection["headers"] = v.ConnectionData.Headers
		connection["key_file"] = v.ConnectionData.KeyFile
		connection["namespace"] = v.ConnectionData.Namespace
		connection["nonce"] = v.ConnectionData.Nonce
		connection["protocol"] = v.ConnectionData.Protocol
		connection["query"] = v.ConnectionData.Query
		connection["reason"] = v.Reason
		connection["region"] = v.ConnectionData.Region
		connection["role_name"] = v.ConnectionData.RoleName
		connection["secret_key"] = v.ConnectionData.SecretKey
		connection["self_signed"] = v.ConnectionData.SelfSigned
		connection["session_token"] = v.ConnectionData.SessionToken
		connection["ssl"] = v.ConnectionData.Ssl
		connection["store_aws_credentials"] = v.ConnectionData.StoreAwsCredentials
		connection["username"] = v.ConnectionData.Username
		connection["v2_key_engine"] = v.ConnectionData.V2KeyEngine

		// Handle structs
		if v.ConnectionData.AmazonSecret != nil {
			amazonSecret := &schema.Set{F: resourceConnectionDataAmazonSecretHash}
			amazonSecretMap := map[string]interface{}{}
			amazonSecretMap["secret_asset_id"] = v.ConnectionData.AmazonSecret.SecretAssetID
			amazonSecretMap["secret_name"] = v.ConnectionData.AmazonSecret.SecretName
			amazonSecret.Add(amazonSecretMap)
			connection["amazon_secret"] = amazonSecret
		}

		if v.ConnectionData.CyberarkSecret != nil {
			cyberarkSecret := &schema.Set{F: resourceConnectionDataCyberarkSecretHash}
			cyberarkSecretMap := map[string]interface{}{}
			cyberarkSecretMap["secret_asset_id"] = v.ConnectionData.CyberarkSecret.SecretAssetID
			cyberarkSecretMap["secret_name"] = v.ConnectionData.CyberarkSecret.SecretName
			cyberarkSecret.Add(cyberarkSecretMap)
			connection["cyberark_secret"] = cyberarkSecret
		}

		if v.ConnectionData.HashicorpSecret != nil {
			hashicorpSecret := &schema.Set{F: resourceConnectionDataHashicorpSecretHash}
			hashicorpSecretMap := map[string]interface{}{}
			hashicorpSecretMap["path"] = v.ConnectionData.HashicorpSecret.Path
			hashicorpSecretMap["secret_asset_id"] = v.ConnectionData.HashicorpSecret.SecretAssetID
			hashicorpSecretMap["secret_name"] = v.ConnectionData.HashicorpSecret.SecretName
			hashicorpSecret.Add(hashicorpSecretMap)
			connection["hashicorp_secret"] = hashicorpSecret
		}

		connections.Add(connection)
	}
	d.Set("asset_connection", connections)

	log.Printf("[INFO] Finished reading secret manager with secretManagerId: %s\n", secretManagerId)

	return nil
}

func resourceSecretManagerUpdateContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Client)

	// check provided fields against schema
	secretManagerId := d.Id()
	if isOk, err := checkResourceRequiredFields(requiredSecretManagerFieldsJson, ignoreSecretManagerParamsByServerType, d); !isOk {
		return diag.FromErr(err)
	}

	// convert provided fields into API payload
	secretManager := ResourceWrapper{}
	serverType := d.Get("server_type").(string)
	createResource(&secretManager, serverType, d)

	// update resource
	log.Printf("[INFO] Updating DSF data source for serverType: %s and gatewayId: %s assetId: %s\n", secretManager.Data.ServerType, secretManager.Data.GatewayID, secretManager.Data.AssetData.AssetID)
	_, err := client.UpdateSecretManager(secretManagerId, secretManager)
	if err != nil {
		log.Printf("[ERROR] Updating secret manager for serverType: %s and gatewayId: %s assetId: %s | err:%s\n", secretManager.Data.ServerType, secretManager.Data.GatewayID, secretManager.Data.AssetData.AssetID, err)
		return diag.FromErr(err)
	}

	// get asset_id
	assetId := d.Get("asset_id").(string)

	// wait for remoteSyncState
	err = waitForRemoteSyncState(ctx, dsfSecretManagerResourceType, assetId, m)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  fmt.Sprintf("Error while waiting for remoteSyncState = \"SYNCED\" for asset: %s", assetId),
			Detail:   fmt.Sprintf("Error: %s\n", err),
		})
	}

	// set ID
	d.SetId(secretManagerId)

	// Set the rest of the state from the resource read
	resourceSecretManagerReadContext(ctx, d, m)

	return diags
}

func resourceSecretManagerDeleteContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	secretManagerId := d.Id()

	log.Printf("[INFO] Deleting secret manager with secretManagerId: %s", secretManagerId)

	dsfDataSourceDeleteResponse, err := client.DeleteSecretManager(secretManagerId)
	if dsfDataSourceDeleteResponse != nil {
		log.Printf("[INFO] DSF secret manager has already been deleted with secretManagerId: %s | err: %s\n", secretManagerId, err)
	}

	return nil
}

func resourceSecretManagerConnectionHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})

	if v, ok := m["access_id"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["aws_iam_server_id"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["ca_certs_path"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["cert_file"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["external_id"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["headers"]; ok && v != nil {
		if headersSlice, ok := v.([]interface{}); ok {
			for _, header := range headersSlice {
				if headerStr, ok := header.(string); ok {
					buf.WriteString(fmt.Sprintf("%s-", headerStr))
				}
			}
		} else if headerStr, ok := v.(string); ok {
			// fallback if it's a single string
			buf.WriteString(fmt.Sprintf("%s-", headerStr))
		}
	}

	if v, ok := m["key_file"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["namespace"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["nonce"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["protocol"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["query"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["reason"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["region"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["role_name"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["secret_key"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["self_signed"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(bool)))
	}

	if v, ok := m["session_token"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["ssl"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(bool)))
	}

	if v, ok := m["store_aws_credentials"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(bool)))
	}

	if v, ok := m["username"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["v2_key_engine"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(bool)))
	}

	return PositiveHash(buf.String())
}
