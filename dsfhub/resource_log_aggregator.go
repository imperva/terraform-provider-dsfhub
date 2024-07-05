package dsfhub

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"log"
)

func resourceLogAggregator() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogAggregatorCreate,
		Read:   resourceLogAggregatorRead,
		Update: resourceLogAggregatorUpdate,
		Delete: resourceLogAggregatorDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"admin_email": {
				Type:        schema.TypeString,
				Description: "The email address to notify about this asset",
				Required:    true,
			},
			//"application": {
			//	Type:        schema.TypeString,
			//	Description: "The Asset ID of the application asset that \"owns\" the asset.",
			//	Required:    false,
			//	Optional:    true,
			//},
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
			"audit_pull_enabled": {
				Type:        schema.TypeBool,
				Description: "If true, sonargateway will collect the audit logs for this system if it can.",
				Required:    false,
				Optional:    true,
				Default:     nil,
				Computed:    true,
			},
			"audit_type": {
				Type:        schema.TypeString,
				Description: "Used to indicate what mechanism should be used to fetch logs on systems supporting multiple ways to get logs, see asset specific documentation for details.  Example: \"BIGQUERY\",\"BIGTABLE\",\"BUCKET\",\"MSSQL\",\"MYSQL\",\"POSTGRESQL\",\"SPANNER\"",
				Required:    false,
				Optional:    true,
				Default:     nil,
				// ValidateFunc: validation.StringInSlice([]string{"BIGQUERY", "BIGTABLE", "BUCKET", "MSSQL", "MYSQL", "POSTGRESQL", "SPANNER"}, false),
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
				Set:         resourceLogAggregatorConnectionHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_id": {
							Type:        schema.TypeString,
							Description: "The Access key ID of AWS secret access key used to authenticate",
							Optional:    true,
							Default:     nil,
							Computed:    true,
						},
						"access_key": {
							Type:        schema.TypeString,
							Description: "The Secret access key used to authenticate",
							Optional:    true,
							Default:     nil,
							Computed:    true,
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
										Description: "Field mapping for amazon secret",
										Required:    false,
										Optional:    true,
										Default:     nil,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"secret_asset_id": {
										Type:        schema.TypeString,
										Description: "Amazon secret asset id",
										Required:    false,
										Optional:    true,
										Default:     nil,
									},
									"secret_name": {
										Type:        schema.TypeString,
										Description: "Amazon secret mane",
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
							Computed:    true,
						},
						"auth_mechanism": {
							Type:         schema.TypeString,
							Description:  "Specifies the auth mechanism used by the connection",
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"default", "service_account", "azure_ad", "kerberos"}, false),
						},
						"azure_storage_account": {
							Type:        schema.TypeString,
							Description: "The name of the unique namespace where the EventHub is located. The field can contain only lowercase letters and numbers. Name must be between 3 and 24 characters.",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"azure_storage_container": {
							Type:        schema.TypeString,
							Description: "Location where a given EventHub\u2019s processing is stored (One storage container per EventHub). This name may only contain lowercase letters, numbers, and hyphens, and must begin with a letter or a number. Each hyphen must be preceded and followed by a non-hyphen character. The name must also be between 3 and 63 characters long.",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"azure_storage_secret_key": {
							Type:        schema.TypeString,
							Description: "Azure Storage Secret Key",
							Required:    false,
							Optional:    true,
							Default:     nil,
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
						},
						"credential_fields": {
							Type:        schema.TypeSet,
							Description: "Document containing values to build a profile from. Filling this will create a profile using the given profile name",
							Required:    false,
							Optional:    true,
							Default:     nil,
							MinItems:    0,
							Set:         resourceConnectionDataCredentialFieldsHash,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"credential_source": {
										Type:        schema.TypeString,
										Description: "HashiCorp secret asset id",
										Required:    false,
										Optional:    true,
										Default:     nil,
									},
									"role_arn": {
										Type:        schema.TypeString,
										Description: "HashiCorp secret mane",
										Required:    false,
										Optional:    true,
										Default:     nil,
									},
								},
							},
						},
						"cyberark_secret": {
							Type:        schema.TypeSet,
							Description: "Configuration to integrate with AWS Secrets Manager",
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
										Description: "Field mapping for amazon secret",
										Required:    false,
										Optional:    true,
										Default:     nil,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"secret_asset_id": {
										Type:        schema.TypeString,
										Description: "Amazon secret asset id",
										Required:    false,
										Optional:    true,
										Default:     nil,
									},
									"secret_name": {
										Type:        schema.TypeString,
										Description: "Amazon secret mane",
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
						"eventhub_access_key": {
							Type:        schema.TypeString,
							Description: "Eventhub Access Key",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"eventhub_access_policy": {
							Type:        schema.TypeString,
							Description: "EventHub Access Policy",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"eventhub_name": {
							Type:        schema.TypeString,
							Description: "EventHub Namespace",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"eventhub_namespace": {
							Type:        schema.TypeString,
							Description: "The name for the management container that the EventHub belongs to, one namespace can contain multiple EventHubs. The namespace can contain only letters, numbers, and hyphens. The namespace must start with a letter, and it must end with a letter or number. The value must be between 6 and 50 characters long.",
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
						"format": {
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
										Description: "HashiCorp secret asset id",
										Required:    false,
										Optional:    true,
										Default:     nil,
									},
									"secret_name": {
										Type:        schema.TypeString,
										Description: "HashiCorp secret mane",
										Required:    false,
										Optional:    true,
										Default:     nil,
									},
								},
							},
						},
						"key_file": {
							Type:        schema.TypeString,
							Description: "Location on disk on the key to be used to authenticate",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"reason": {
							Type:         schema.TypeString,
							Description:  "N/A",
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"default"}, false),
						},
						"region": {
							Type:        schema.TypeString,
							Description: "Default AWS region for this asset",
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
							Description: "The Secret access key used to authenticate",
							Required:    false,
							Optional:    true,
							Default:     nil,
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
			"bucket_account_id": {
				Type:        schema.TypeString,
				Description: "S3 bucket Account ID",
				Required:    false,
				Optional:    true,
				Default:     nil,
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
			"endpoint": {
				Type:        schema.TypeString,
				Description: "Logstore's endpoint",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"gateway_id": {
				Type:        schema.TypeString,
				Description: "The jsonarUid unique identifier of the agentless gateway. Example: '7a4af7cf-4292-89d9-46ec-183756ksdjd'",
				Required:    true,
			},
			"gateway_service": {
				Type:        schema.TypeString,
				Description: "The name of the gateway pull service (if any) used to retrieve logs for this source. Usually set by the connect gateway playbook.",
				Required:    false,
				Optional:    true,
				Default:     nil,
				Computed:    true,
			},
			"jsonar_uid": {
				Type:        schema.TypeString,
				Description: "Unique identifier (UID) attached to the Sonar machine controlling the asset",
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
			"logstore": {
				Type:        schema.TypeString,
				Description: "Unit that is used to collect, store and query logs",
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
			"parent_asset_id": {
				Type:        schema.TypeString,
				Description: "The name of an asset that this asset is part of (/related to). E.g. an AWS resource will generally have an AWS account asset as its parent. Also used to connect some log aggregating asset with the sources of their logs. E.g. An AWS LOG GROUP asset can have an AWS RDS as its parent, indicating that that is the log group for that RDS.",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"project": {
				Type:        schema.TypeString,
				Description: "Project separates different resources of multiple users and control access to specific resources",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"proxy": {
				Type:        schema.TypeString,
				Description: "Proxy to use for AWS calls if aws_proxy_config is populated the proxy field will get populated from the http value there",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"pubsub_subscription": {
				Type:        schema.TypeString,
				Description: "Pub/Sub Subscription. Example: projects/my-project-name/subscriptions/my-subscription-name",
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
			"s3_provider": {
				Type:        schema.TypeString,
				Description: "Accepted value: \"aws-rds-mssql\", required only for AWS RDS MS SQL SERVER auditing workflow.",
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
			"server_type": {
				Type:        schema.TypeString,
				Description: "The type of server or data service to be created as a data source. The list of available data sources is documented at: https://docs.imperva.com/bundle/v4.11-sonar-user-guide/page/84552.htm",
				Required:    true,
			},
			"server_ip": {
				Type:        schema.TypeString,
				Description: "IP address of the service where this asset is located. If no IP is available populate this field with other information that would identify the system e.g. hostname or AWS ARN, etc.",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"server_port": {
				Type:        schema.TypeString,
				Description: "",
				Required:    false,
				Optional:    true,
				Default:     nil,
				Computed:    true,
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
			"version": {
				Type:        schema.TypeFloat,
				Description: "Denotes the version of the asset",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
		},
	}
}

func resourceLogAggregatorCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	if isOk, err := checkResourceRequiredFields(requiredLogAggregatorJson, ignoreLogAggregatorParamsByServerType, d); !isOk {
		return err
	}
	logAggregator := ResourceWrapper{}
	serverType := d.Get("server_type").(string)
	createResource(&logAggregator, serverType, d)
	// auditPullEnabled set to false as connect/disconnect logic handled below
	logAggregator.Data.AssetData.AuditPullEnabled = false
	log.Printf("[INFO] Creating LogAggregator for serverType: %s and gatewayId: %s\n", logAggregator.Data.ServerType, logAggregator.Data.GatewayID)
	createLogAggregatorResponse, err := client.CreateLogAggregator(logAggregator)

	if err != nil {
		log.Printf("[ERROR] adding LogAggregator for serverType: %s and gatewayId: %s | err: %s", serverType, logAggregator.Data.GatewayID, err)
		return err
	}

	// Connect/disconnect asset to gateway
	connectDisconnectGateway(d, logAggregator, m)

	// Set ID
	logAggregatorId := createLogAggregatorResponse.Data.ID
	d.SetId(logAggregatorId)

	// Set the rest of the state from the resource read
	return resourceLogAggregatorRead(d, m)
}

func resourceLogAggregatorRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	logAggregatorId := d.Id()

	log.Printf("[INFO] Reading LogAggregator with logAggregatorId: %s\n", logAggregatorId)

	logAggregatorReadResponse, err := client.ReadLogAggregator(logAggregatorId)

	if err != nil {
		log.Printf("[ERROR] Reading logAggregatorReadResponse with logAggregatorId: %s | err: %s\n", logAggregatorId, err)
		return err
	}

	if logAggregatorReadResponse != nil {
		log.Printf("[INFO] Reading LogAggregator with logAggregatorId: %s | err: %s\n", logAggregatorId, err)
	}

	log.Printf("[DEBUG] logAggregatorReadResponse: %s\n", logAggregatorReadResponse.Data.ID)
	// Set returned and computed values
	d.Set("admin_email", logAggregatorReadResponse.Data.AssetData.AdminEmail)
	//d.Set("application", logAggregatorReadResponse.Data.AssetData.Application)
	d.Set("arn", logAggregatorReadResponse.Data.AssetData.Arn)
	d.Set("asset_display_name", logAggregatorReadResponse.Data.AssetData.AssetDisplayName)
	d.Set("asset_id", logAggregatorReadResponse.Data.AssetData.AssetID)
	d.Set("asset_source", logAggregatorReadResponse.Data.AssetData.AssetSource)
	d.Set("audit_pull_enabled", logAggregatorReadResponse.Data.AssetData.AuditPullEnabled)
	d.Set("audit_type", logAggregatorReadResponse.Data.AssetData.AuditType)
	d.Set("available_regions", logAggregatorReadResponse.Data.AssetData.AvailableRegions)
	d.Set("bucket_account_id", logAggregatorReadResponse.Data.AssetData.BucketAccountId)
	d.Set("credential_endpoint", logAggregatorReadResponse.Data.AssetData.CredentialsEndpoint)
	d.Set("criticality", logAggregatorReadResponse.Data.AssetData.Criticality)
	d.Set("gateway_id", logAggregatorReadResponse.Data.GatewayID)
	d.Set("gateway_service", logAggregatorReadResponse.Data.AssetData.GatewayService)
	d.Set("jsonar_uid", logAggregatorReadResponse.Data.AssetData.JsonarUID)
	d.Set("location", logAggregatorReadResponse.Data.AssetData.Location)
	d.Set("managed_by", logAggregatorReadResponse.Data.AssetData.ManagedBy)
	d.Set("owned_by", logAggregatorReadResponse.Data.AssetData.OwnedBy)
	d.Set("proxy", logAggregatorReadResponse.Data.AssetData.Proxy)
	d.Set("pubsub_subscription", logAggregatorReadResponse.Data.AssetData.PubsubSubscription)
	d.Set("region", logAggregatorReadResponse.Data.AssetData.Region)
	d.Set("s3_provider", logAggregatorReadResponse.Data.AssetData.S3Provider)
	d.Set("server_host_name", logAggregatorReadResponse.Data.AssetData.ServerHostName)
	d.Set("server_type", logAggregatorReadResponse.Data.ServerType)
	if logAggregatorReadResponse.Data.AssetData.ServerPort != nil {
		var serverPort string
		if serverPortNum, ok := logAggregatorReadResponse.Data.AssetData.ServerPort.(float64); ok {
			serverPort = fmt.Sprintf("%d", int(serverPortNum))
		} else {
			serverPort = logAggregatorReadResponse.Data.AssetData.ServerPort.(string)
		}
		d.Set("server_port", serverPort)
	}
	d.Set("used_for", logAggregatorReadResponse.Data.AssetData.UsedFor)
	d.Set("version", logAggregatorReadResponse.Data.AssetData.Version)

	if logAggregatorReadResponse.Data.AssetData.AwsProxyConfig != nil {
		awsProxyConfig := &schema.Set{F: resourceAssetDataAWSProxyConfigHash}
		awsProxyConfigMap := map[string]interface{}{}
		awsProxyConfigMap["http"] = logAggregatorReadResponse.Data.AssetData.AwsProxyConfig.HTTP
		awsProxyConfigMap["https"] = logAggregatorReadResponse.Data.AssetData.AwsProxyConfig.HTTPS
		awsProxyConfig.Add(awsProxyConfigMap)
		d.Set("aws_proxy_config", awsProxyConfig)
	}

	if logAggregatorReadResponse.Data.AssetData.ServiceEndpoints != nil {
		serviceEndpoints := &schema.Set{F: resourceAssetDataServiceEndpointsHash}
		serviceEndpointsMap := map[string]interface{}{}
		serviceEndpointsMap["logs"] = logAggregatorReadResponse.Data.AssetData.ServiceEndpoints.Logs
		serviceEndpoints.Add(serviceEndpointsMap)
		d.Set("service_endpoints", serviceEndpoints)
	}

	connections := &schema.Set{F: resourceLogAggregatorConnectionHash}
	for _, v := range logAggregatorReadResponse.Data.AssetData.Connections {
		connection := map[string]interface{}{}
		connection["access_id"] = v.ConnectionData.AccessID
		connection["access_key"] = v.ConnectionData.AccessKey
		connection["application_id"] = v.ConnectionData.ApplicationID
		connection["auth_mechanism"] = v.ConnectionData.AuthMechanism
		connection["azure_storage_account"] = v.ConnectionData.AzureStorageAccount
		connection["azure_storage_container"] = v.ConnectionData.AzureStorageContainer
		connection["azure_storage_secret_key"] = v.ConnectionData.AzureStorageSecretKey
		connection["ca_certs_path"] = v.ConnectionData.CaCertsPath
		connection["client_secret"] = v.ConnectionData.ClientSecret
		connection["cyberark_secret"] = v.ConnectionData.CyberarkSecret
		connection["directory_id"] = v.ConnectionData.DirectoryID
		connection["eventhub_access_key"] = v.ConnectionData.EventhubAccessKey
		connection["eventhub_access_policy"] = v.ConnectionData.EventhubAccessPolicy
		connection["eventhub_name"] = v.ConnectionData.EventhubName
		connection["eventhub_namespace"] = v.ConnectionData.EventhubNamespace
		connection["external_id"] = v.ConnectionData.ExternalID
		connection["format"] = v.ConnectionData.Format
		connection["hashicorp_secret"] = v.ConnectionData.HashicorpSecret
		connection["key_file"] = v.ConnectionData.KeyFile
		connection["reason"] = v.Reason
		connection["region"] = v.ConnectionData.Region
		connection["role_name"] = v.ConnectionData.RoleName
		connection["secret_key"] = v.ConnectionData.SecretKey
		connection["ssl"] = v.ConnectionData.Ssl
		connection["subscription_id"] = v.ConnectionData.SubscriptionID
		connection["username"] = v.ConnectionData.Username

		// Handle structs
		if v.ConnectionData.AmazonSecret != nil {
			amazonSecret := &schema.Set{F: resourceConnectionDataAmazonSecretHash}
			amazonSecretMap := map[string]interface{}{}
			//amazonSecretMap["field_mapping"] = v.ConnectionData.AmazonSecret.FieldMapping
			amazonSecretMap["secret_asset_id"] = v.ConnectionData.AmazonSecret.SecretAssetID
			amazonSecretMap["secret_name"] = v.ConnectionData.AmazonSecret.SecretName
			amazonSecret.Add(amazonSecretMap)
			connection["amazon_secret"] = amazonSecret
		}

		if v.ConnectionData.CredentialFields != nil {
			credentialFields := &schema.Set{F: resourceConnectionDataCredentialFieldsHash}
			credentialFieldsMap := map[string]interface{}{}
			credentialFieldsMap["credential_source"] = v.ConnectionData.CredentialFields.CredentialSource
			credentialFieldsMap["role_arn"] = v.ConnectionData.CredentialFields.RoleArn
			credentialFields.Add(credentialFieldsMap)
			connection["credential_fields"] = credentialFields
		}

		if v.ConnectionData.CyberarkSecret != nil {
			amazonSecret := &schema.Set{F: resourceConnectionDataCyberarkSecretHash}
			amazonSecretMap := map[string]interface{}{}
			//amazonSecretMap["field_mapping"] = v.ConnectionData.AmazonSecret.FieldMapping
			amazonSecretMap["secret_asset_id"] = v.ConnectionData.CyberarkSecret.SecretAssetID
			amazonSecretMap["secret_name"] = v.ConnectionData.CyberarkSecret.SecretName
			amazonSecret.Add(amazonSecretMap)
			connection["cyberark_secret"] = amazonSecret
		}

		if v.ConnectionData.HashicorpSecret != nil {
			hashicorpSecret := &schema.Set{F: resourceConnectionDataHashicorpSecretHash}
			hashicorpSecretMap := map[string]interface{}{}
			//hashicorpSecretMap["field_mapping"] = v.ConnectionData.HashicorpSecret.Path
			hashicorpSecretMap["path"] = v.ConnectionData.HashicorpSecret.Path
			hashicorpSecretMap["secret_asset_id"] = v.ConnectionData.HashicorpSecret.SecretAssetID
			hashicorpSecretMap["secret_name"] = v.ConnectionData.HashicorpSecret.SecretName
			hashicorpSecret.Add(hashicorpSecretMap)
			connection["hashicorp_secret"] = hashicorpSecret
		}

		connections.Add(connection)
	}
	d.Set("ca_connection", connections)

	log.Printf("[INFO] Finished reading logAggregator with logAggregatorId: %s\n", logAggregatorId)

	return nil
}

func resourceLogAggregatorUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	logAggregatorId := d.Id()
	if isOk, err := checkResourceRequiredFields(requiredLogAggregatorJson, ignoreLogAggregatorParamsByServerType, d); !isOk {
		return err
	}
	logAggregator := ResourceWrapper{}
	serverType := d.Get("server_type").(string)
	createResource(&logAggregator, serverType, d)

	// Do not change audit_pull_enabled in initial update payload as it is handled below
	auditPullEnabledChanged := d.HasChange("audit_pull_enabled")
	if auditPullEnabledChanged {
		log.Printf("[DEBUG] audit_pull_enabled value changed, setting to original value in initial update payload")
		origAuditPullEnabled, _ := d.GetChange("audit_pull_enabled")
		logAggregator.Data.AssetData.AuditPullEnabled = origAuditPullEnabled.(bool)
		log.Printf("[DEBUG] AuditPullEnabled value in logAggregator update payload is: '%v'", logAggregator.Data.AssetData.AuditPullEnabled)
	}

	log.Printf("[INFO] Updating LogAggregator for serverType: %s and gatewayId: %s assetId: %s\n", logAggregator.Data.ServerType, logAggregator.Data.GatewayID, logAggregator.Data.AssetData.AssetID)
	_, err := client.UpdateLogAggregator(logAggregatorId, logAggregator)

	if err != nil {
		log.Printf("[ERROR] Updating LogAggregator for serverType: %s and gatewayId: %s assetId: %s | err:%s\n", logAggregator.Data.ServerType, logAggregator.Data.GatewayID, logAggregator.Data.AssetData.AssetID, err)
		return err
	}

	// Connect/disconnect asset to gateway
	connectDisconnectGateway(d, logAggregator, m)

	// Set ID
	d.SetId(logAggregatorId)

	// Set the rest of the state from the resource read
	return resourceLogAggregatorRead(d, m)
}

func resourceLogAggregatorDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	logAggregatorId := d.Id()

	log.Printf("[INFO] Deleting secret manager with logAggregatorId: %s", logAggregatorId)

	logAggregatorDeleteResponse, err := client.DeleteLogAggregator(logAggregatorId)
	if logAggregatorDeleteResponse != nil {
		log.Printf("[INFO] LogAggregator has already been deleted with logAggregatorId: %s | err: %s\n", logAggregatorId, err)
	}

	return nil
}

func resourceLogAggregatorConnectionHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})

	if v, ok := m["access_id"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["access_key"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["application_id"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["auth_mechanism"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["azure_storage_account"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["azure_storage_container"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["azure_storage_secret_key"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["ca_certs_path"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["client_secret"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["directory_id"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["eventhub_access_key"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["eventhub_access_policy"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["eventhub_name"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["eventhub_namespace"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["external_id"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["format"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["key_file"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["reason"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["region"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["role_name"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["secret_key"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["ssl"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(bool)))
	}

	if v, ok := m["subscription_id"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["username"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	return PositiveHash(buf.String())
}
