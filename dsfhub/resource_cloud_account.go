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

func resourceCloudAccount() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudAccountCreateContext,
		ReadContext:   resourceCloudAccountReadContext,
		UpdateContext: resourceCloudAccountUpdateContext,
		DeleteContext: resourceCloudAccountDeleteContext,
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
				Default:     nil,
			},
			"asset_id": {
				Type:        schema.TypeString,
				Description: "The following format is a suggestion that guarantees uniqueness and identifiability between AZURE assets using different applications/resource groups: 'directory_id/<directory_id>/subscription_id/<subscription_id>/<resource group name>/application_id/<application_id>'",
				Required:    true,
			},
			"asset_source": {
				Type:        schema.TypeString,
				Description: "The source platform/vendor/system of the asset data. Usually the service responsible for creating that asset document",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"asset_version": {
				Type:        schema.TypeFloat,
				Description: "Denotes the version of the asset",
				Required:    false,
				Optional:    true,
				Default:     nil,
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
				Required:    true,
				MinItems:    0,
				Set:         resourceCloudAccountConnectionHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_id": {
							Type:        schema.TypeString,
							Description: "The Access key ID of AWS secret access key used for authentication",
							Optional:    true,
							Default:     nil,
						},
						"access_key": {
							Type:        schema.TypeString,
							Description: "The Secret access key used for authentication",
							Optional:    true,
							Default:     nil,
							Required:    false,
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
						"application_id": {
							Type:        schema.TypeString,
							Description: "This is also referred to as the Client ID and it’s the unique identifier for the registered application being used to execute Python SDK commands against Azure’s API services. You can find this number under Azure Active Directory -> App Registrations -> Owned Applications",
							Optional:    true,
							Default:     nil,
						},
						"auth_mechanism": {
							Type:         schema.TypeString,
							Description:  "Specifies the auth mechanism used by the connection",
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"auth_file", "client_secret", "default", "iam_role", "key", "machine_role", "managed_identity", "profile", "service_account"}, false),
						},
						"ca_certs_path": {
							Type:        schema.TypeString,
							Description: "Certificate authority certificates path; what location should the sysetm look for certificate information from. Equivalent to --capath in a curl call",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"client_secret": {
							Type:        schema.TypeString,
							Description: "This a string containing a secret used by the application to prove its identity when requesting a token. You can get a secret by going to Azure Active Directory -> App Registrations -> Owned Applications, selecting the desired application and then going to Certificates & secrets -> Client secrets -> + New client secret",
							Required:    false,
							Optional:    true,
							Default:     nil,
							Sensitive:   true,
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
						"directory_id": {
							Type:        schema.TypeString,
							Description: "This is also referred to as the Tenant ID and is a GUID representing the Active Directory Tenant. It can be found in the Azure Active Directory page under the Azure portal",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"external_id": {
							Type:        schema.TypeString,
							Description: "External ID to use when assuming a role",
							Required:    false,
							Optional:    true,
							Default:     nil,
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
						"key_file": {
							Type:        schema.TypeString,
							Description: "Location on disk on the key to be used for authentication",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"project_id": {
							Type:        schema.TypeString,
							Description: "Used when running Sonar on a GCP hosted environment that doesn't have a service account linked to it.",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"reason": {
							Type:         schema.TypeString,
							Description:  "N/A",
							Required:     false,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"default"}, false),
						},
						"region": {
							Type:        schema.TypeString,
							Description: "Default region for this asset",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"role_name": {
							Type:        schema.TypeString,
							Description: "What role is used to get credentials from.",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"secret_key": {
							Type:        schema.TypeString,
							Description: "The Secret access key used for authentication",
							Required:    false,
							Optional:    true,
							Default:     nil,
							Sensitive:   true,
						},
						"session_token": {
							Type:        schema.TypeString,
							Description: "The session token used for authentication",
							Required:    false,
							Optional:    true,
							Default:     nil,
							Sensitive:   true,
						},
						"ssl": {
							Type:        schema.TypeBool,
							Description: "If true, use SSL when connecting",
							Required:    false,
							Optional:    true,
							Default:     false,
						},
						"subscription_id": {
							Type:        schema.TypeString,
							Description: "This is the Azure account subscription ID. You can find this number under the Subscriptions page on the Azure portal",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"username": {
							Type:        schema.TypeString,
							Description: "The name of a profile in /imperva/local/credentials/.aws/credentials to use for authenticating",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
					},
				},
			},
			"credentials_endpoint": {
				Type:        schema.TypeString,
				Description: "A specific sts endpoint to use",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"criticality": {
				Type:         schema.TypeInt,
				Description:  "The asset's importance to the business. These values are measured on a scale from \"Most critical\" (1) to \"Least critical\" (4). Allowed values: 1, 2, 3, 4",
				Required:     false,
				Optional:     true,
				ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4}),
				Default:      nil,
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
				Default:     nil,
				Computed:    true,
			},
			"location": {
				Type:        schema.TypeString,
				Description: "Current human-readable description of the physical location of the asset, or region.",
				Required:    false,
				Optional:    true,
				Default:     nil,
				Computed:    true,
			},
			"managed_by": {
				Type:        schema.TypeString,
				Description: "Email of the person who maintains the asset; can be different from the owner specified in the owned_by field. Defaults to admin_email.",
				Required:    false,
				Optional:    true,
				Default:     nil,
				Computed:    true,
			},
			"owned_by": {
				Type:        schema.TypeString,
				Description: "Email of Owner / person responsible for the asset; can be different from the person in the managed_by field. Defaults to admin_email.",
				Required:    false,
				Optional:    true,
				Default:     nil,
				Computed:    true,
			},
			"proxy": {
				Type:        schema.TypeString,
				Description: "Proxy to use for AWS calls if aws_proxy_config is populated the proxy field will get populated from the http value there",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"region": {
				Type:        schema.TypeString,
				Description: "For cloud systems with regions, the default region or region used with this asset",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"server_host_name": {
				Type:        schema.TypeString,
				Description: "Hostname (or IP if name is unknown)",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"server_ip": {
				Type:        schema.TypeString,
				Description: "IP address of the service where this asset is located. If no IP is available populate this field with other information that would identify the system e.g. hostname or AWS ARN, etc.",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"server_type": {
				Type:        schema.TypeString,
				Description: "The type of server or data service to be created as a data source. The list of available data sources is documented at: https://docs.imperva.com/bundle/v4.11-sonar-user-guide/page/84552.htm",
				Required:    true,
			},
			"server_port": {
				Type:        schema.TypeString,
				Description: "",
				Required:    false,
				Optional:    true,
				Default:     nil,
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
				Default:      nil,
			},
		},
	}
}

func resourceCloudAccountCreateContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Client)
	if isOk, err := checkResourceRequiredFields(requiredCloudAccountJson, ignoreCloudAccountParamsByServerType, d); !isOk {
		return diag.FromErr(err)
	}

	// check provided fields against schema
	cloudAccount := ResourceWrapper{}
	serverType := d.Get("server_type").(string)
	createResource(&cloudAccount, serverType, d)

	// create resource
	log.Printf("[INFO] Creating CloudAccount for serverType: %s and gatewayId: %s gatewayId: \n", serverType, cloudAccount.Data.GatewayID)
	createCloudAccountResponse, err := client.CreateCloudAccount(cloudAccount)
	if err != nil {
		log.Printf("[ERROR] adding CloudAccount for serverType: %s and gatewayId: %s | err: %s", serverType, cloudAccount.Data.GatewayID, err)
		return diag.FromErr(err)
	}

	// get asset_id
	assetId := d.Get("asset_id").(string)

	// wait for remoteSyncState
	err = waitForRemoteSyncState(ctx, dsfCloudAccountResourceType, assetId, m)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  fmt.Sprintf("Error while waiting for remoteSyncState = \"SYNCED\" for asset: %s", assetId),
			Detail:   fmt.Sprintf("Error: %s\n", err),
		})
	}

	// set ID
	cloudAccountId := createCloudAccountResponse.Data.AssetData.AssetID
	d.SetId(cloudAccountId)

	// Set the rest of the state from the resource read
	resourceCloudAccountReadContext(ctx, d, m)

	return diags
}

func resourceCloudAccountReadContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	cloudAccountId := d.Id()

	log.Printf("[INFO] Reading CloudAccount with cloudAccountId: %s\n", cloudAccountId)

	cloudAccountReadResponse, err := client.ReadCloudAccount(cloudAccountId)

	if err != nil {
		log.Printf("[ERROR] Reading cloudAccountReadResponse with cloudAccountId: %s | err: %s\n", cloudAccountId, err)
		return diag.FromErr(err)
	}

	if cloudAccountReadResponse != nil {
		log.Printf("[INFO] Reading CloudAcount with cloudAccountId: %s | err: %s\n", cloudAccountId, err)
	}

	log.Printf("[DEBUG] cloudAccountReadResponse: %s\n", cloudAccountReadResponse.Data.ID)
	// Set returned and computed values
	d.Set("admin_email", cloudAccountReadResponse.Data.AssetData.AdminEmail)
	d.Set("application", cloudAccountReadResponse.Data.AssetData.Application)
	d.Set("arn", cloudAccountReadResponse.Data.AssetData.Arn)
	d.Set("asset_display_name", cloudAccountReadResponse.Data.AssetData.AssetDisplayName)
	d.Set("asset_id", cloudAccountReadResponse.Data.AssetData.AssetID)
	d.Set("asset_source", cloudAccountReadResponse.Data.AssetData.AssetSource)
	d.Set("asset_version", cloudAccountReadResponse.Data.AssetData.Version)
	d.Set("available_regions", cloudAccountReadResponse.Data.AssetData.AvailableRegions)
	if cloudAccountReadResponse.Data.AssetData.CredentialsEndpoint != "" {
		d.Set("credentials_endpoint", cloudAccountReadResponse.Data.AssetData.CredentialsEndpoint)
	}
	d.Set("criticality", cloudAccountReadResponse.Data.AssetData.Criticality)
	d.Set("gateway_id", cloudAccountReadResponse.Data.GatewayID)
	d.Set("id", cloudAccountReadResponse.Data.ID)
	d.Set("jsonar_uid", cloudAccountReadResponse.Data.AssetData.JsonarUID)
	d.Set("location", cloudAccountReadResponse.Data.AssetData.Location)
	d.Set("managed_by", cloudAccountReadResponse.Data.AssetData.ManagedBy)
	d.Set("owned_by", cloudAccountReadResponse.Data.AssetData.OwnedBy)
	d.Set("proxy", cloudAccountReadResponse.Data.AssetData.Proxy)
	d.Set("region", cloudAccountReadResponse.Data.AssetData.Region)
	d.Set("server_host_name", cloudAccountReadResponse.Data.AssetData.ServerHostName) // TODO SR-4549
	d.Set("server_ip", cloudAccountReadResponse.Data.AssetData.ServerIP)              // TODO SR-4549
	d.Set("server_type", cloudAccountReadResponse.Data.ServerType)
	if cloudAccountReadResponse.Data.AssetData.ServerPort != nil {
		var serverPort string
		if serverPortNum, ok := cloudAccountReadResponse.Data.AssetData.ServerPort.(float64); ok {
			serverPort = fmt.Sprintf("%d", int(serverPortNum))
		} else {
			serverPort = cloudAccountReadResponse.Data.AssetData.ServerPort.(string)
		}
		d.Set("server_port", serverPort)
	}
	d.Set("used_for", cloudAccountReadResponse.Data.AssetData.UsedFor)

	// Handle Structs
	if cloudAccountReadResponse.Data.AssetData.AwsProxyConfig != nil {
		awsProxyConfig := &schema.Set{F: resourceAssetDataAWSProxyConfigHash}
		awsProxyConfigMap := map[string]interface{}{}
		awsProxyConfigMap["http"] = cloudAccountReadResponse.Data.AssetData.AwsProxyConfig.HTTP
		awsProxyConfigMap["https"] = cloudAccountReadResponse.Data.AssetData.AwsProxyConfig.HTTPS
		awsProxyConfig.Add(awsProxyConfigMap)
		d.Set("aws_proxy_config", awsProxyConfig)
	}

	if cloudAccountReadResponse.Data.AssetData.ServiceEndpoints != nil {
		serviceEndpoints := &schema.Set{F: resourceAssetDataServiceEndpointsHash}
		serviceEndpointsMap := map[string]interface{}{}
		serviceEndpointsMap["logs"] = cloudAccountReadResponse.Data.AssetData.ServiceEndpoints.Logs
		serviceEndpoints.Add(serviceEndpointsMap)
		d.Set("service_endpoints", serviceEndpoints)
	}

	connections := &schema.Set{F: resourceCloudAccountConnectionHash}
	for _, v := range cloudAccountReadResponse.Data.AssetData.Connections {
		connection := map[string]interface{}{}
		connection["access_id"] = v.ConnectionData.AccessID
		connection["access_key"] = v.ConnectionData.AccessKey
		connection["application_id"] = v.ConnectionData.ApplicationID
		connection["auth_mechanism"] = v.ConnectionData.AuthMechanism
		connection["ca_certs_path"] = v.ConnectionData.CaCertsPath
		connection["client_secret"] = v.ConnectionData.ClientSecret
		connection["directory_id"] = v.ConnectionData.DirectoryID
		connection["external_id"] = v.ConnectionData.ExternalID
		connection["key_file"] = v.ConnectionData.KeyFile
		connection["project_id"] = v.ConnectionData.ProjectID
		connection["reason"] = v.Reason
		connection["region"] = v.ConnectionData.Region
		connection["role_name"] = v.ConnectionData.RoleName
		connection["secret_key"] = v.ConnectionData.SecretKey
		connection["session_token"] = v.ConnectionData.SessionToken
		connection["ssl"] = v.ConnectionData.Ssl
		connection["subscription_id"] = v.ConnectionData.SubscriptionID
		connection["username"] = v.ConnectionData.Username

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

	log.Printf("[INFO] Finished reading CloudAccount with cloudAccountId: %s\n", cloudAccountId)

	return nil
}

func resourceCloudAccountUpdateContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Client)

	// check provided fields against schema
	cloudAccountId := d.Id()
	if isOk, err := checkResourceRequiredFields(requiredCloudAccountJson, ignoreCloudAccountParamsByServerType, d); !isOk {
		return diag.FromErr(err)
	}

	// convert provided fields into API payload
	cloudAccount := ResourceWrapper{}
	serverType := d.Get("server_type").(string)
	createResource(&cloudAccount, serverType, d)

	// update resource
	log.Printf("[INFO] Updating CloudAccount for serverType: %s and gatewayId: %s assetId: %s\n", cloudAccount.Data.ServerType, cloudAccount.Data.GatewayID, cloudAccount.Data.AssetData.AssetID)
	_, err := client.UpdateCloudAccount(cloudAccountId, cloudAccount)
	if err != nil {
		log.Printf("[ERROR] Updating CloudAccount for serverType: %s and gatewayId: %s assetId: %s | err:%s\n", cloudAccount.Data.ServerType, cloudAccount.Data.GatewayID, cloudAccount.Data.AssetData.AssetID, err)
		return diag.FromErr(err)
	}

	// get asset_id
	assetId := d.Get("asset_id").(string)

	// wait for remoteSyncState
	err = waitForRemoteSyncState(ctx, dsfCloudAccountResourceType, assetId, m)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  fmt.Sprintf("Error while waiting for remoteSyncState = \"SYNCED\" for asset: %s", assetId),
			Detail:   fmt.Sprintf("Error: %s\n", err),
		})
	}

	// set ID
	d.SetId(cloudAccountId)

	// Set the rest of the state from the resource read
	resourceCloudAccountReadContext(ctx, d, m)

	return diags
}

func resourceCloudAccountDeleteContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	cloudAccountId := d.Id()

	log.Printf("[INFO] Deleting secret manager with cloudAccountId: %s", cloudAccountId)

	cloudAccountDeleteResponse, err := client.DeleteCloudAccount(cloudAccountId)
	if cloudAccountDeleteResponse != nil {
		log.Printf("[INFO] CloudAccount has already been deleted with cloudAccountId: %s | err: %s\n", cloudAccountId, err)
	}

	return nil
}

func resourceCloudAccountConnectionHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})

	if v, ok := m["access_id"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["access_key"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["application_id"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["auth_mechanism"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["ca_certs_path"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["client_secret"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["directory_id"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["external_id"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["key_file"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["project_id"]; ok {
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

	if v, ok := m["session_token"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["ssl"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(bool)))
	}

	if v, ok := m["subscription_id"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["username"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	return PositiveHash(buf.String())
}
