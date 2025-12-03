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

func resourceLogAggregator() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLogAggregatorCreateContext,
		ReadContext:   resourceLogAggregatorReadContext,
		UpdateContext: resourceLogAggregatorUpdateContext,
		DeleteContext: resourceLogAggregatorDeleteContext,
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
			"audit_data_type": {
				Type:        schema.TypeString,
				Description: "The type of audit data being collected",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"audit_info": {
				Type:        schema.TypeSet,
				Description: "Normally auto-populated when enabling the audit policy, it is a sub-document in JSON format containing configuration information for audit management. See documentation for values that can be added manually depending on asset type. Editing this value does NOT enable the audit policy.",
				Required:    false,
				Optional:    true,
				Default:     nil,
				MinItems:    0,
				Set:         resourceAssetDataAuditInfoHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_template_name": {
							Type:        schema.TypeString,
							Description: "Policy template name",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
					},
				},
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
			},
			"available_bucket_account_ids": {
				Type:        schema.TypeList,
				Description: "A list of S3 bucket Account IDs",
				Required:    false,
				Optional:    true,
				Default:     nil,
				Computed:    false,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
							Description: "The Access key ID of AWS secret access key used for authentication",
							Optional:    true,
							Default:     nil,
							Computed:    true,
						},
						"access_key": {
							Type:        schema.TypeString,
							Description: "The Secret access key used for authentication",
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
										Description: "AWS secret asset id",
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
							Computed:    true,
						},
						"auth_mechanism": {
							Type:         schema.TypeString,
							Description:  "Specifies the auth mechanism used by the connection",
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"default", "service_account", "azure_ad", "kerberos", "client_secret"}, false),
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
						"cache_file": {
							Type:        schema.TypeString,
							Description: "Holds Kerberos protocol credentials (for example, tickets, session keys and other identifying information).",
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
						"cyberark_secret": {
							Type:        schema.TypeSet,
							Description: "Configuration to integrate with CyberArk Secrets Manager",
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
						"db_role": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
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
							Description: "EventHub Name",
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
						"external": {
							Type:        schema.TypeBool,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
							Computed:    true,
						},
						"extra_kinit_parameters": {
							Type:        schema.TypeString,
							Description: "",
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
										Description: "HashiCorp secret asset_id",
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

						"kerberos_kdc": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},

						"kerberos_service_kdc": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"kerberos_service_realm": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"kerberos_spn": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"key_file": {
							Type:        schema.TypeString,
							Description: "Location on disk on the key to be used for authentication",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},

						"keytab_file": {
							Type:        schema.TypeString,
							Description: "Specify a non-default keytab location",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"kinit_program_path": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"passphrase": {
							Type:        schema.TypeString,
							Description: "Passphrase for the private key.",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"password": {
							Type:        schema.TypeString,
							Description: "The password of the user being used for authentication",
							Required:    false,
							Optional:    true,
							Default:     nil,
							Sensitive:   true,
						},
						"principal": {
							Type:        schema.TypeString,
							Description: "The principal used for authentication",
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
							Description: "The Secret access key used for authentication",
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
						"ssl_server_cert": {
							Type:        schema.TypeString,
							Description: "Path to server certificate to use during authentication",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"subscription_id": {
							Type:        schema.TypeString,
							Description: "This is the Azure account subscription ID. You can find this number under the Subscriptions page on the Azure portal",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"use_keytab": {
							Type:        schema.TypeBool,
							Description: "If true, authenticate using a key tab",
							Required:    false,
							Optional:    true,
							Default:     nil,
							Computed:    true,
						},
						"username": {
							Type:        schema.TypeString,
							Description: "The name of a profile in /imperva/local/credentials/.aws/credentials to use for authenticating",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"user_identity_client_id": {
							Type:        schema.TypeString,
							Description: "The client ID of a user-assigned managed identity.",
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
			"ca_certs_path": {
				Type:        schema.TypeString,
				Description: "Certificate authority certificates path; what location should the sysetm look for certificate information from. Equivalent to --capath in a curl call",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"ca_file": {
				Type:        schema.TypeString,
				Description: "Path to a certificate authority file to use with the call. Equivalent to --cacert in a curl call",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"consumer_group": {
				Type:        schema.TypeString,
				Description: "The name of the consumer group to use when pulling data from the logstore.",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"consumer_group_workers": {
				Type:        schema.TypeString,
				Description: "Only applies if Pull Type is consumer_group. The number of consumers that will be part of the consumer group. For best performance this should match the number of shards in your logstore.",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"consumer_worker_prefix": {
				Type:        schema.TypeString,
				Description: "The prefix to use for the consumer worker name. This can be useful if you are trying to pull from the same consumer group on different gateways",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"content_type": {
				Type:        schema.TypeString,
				Description: "content_type should be set to the desired <'parent' asset 'Server Type'>, which is the one that uses this asset as a destination for logs. NOTE: The content_type field will take precedence on the lookup for parent_asset_id field when checking which server is sending logs to this asset.",
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
			"database_name": {
				Type:        schema.TypeString,
				Description: "Specifies the name of the database (or default DB) to connect to.",
				Required:    false,
				Optional:    true,
				Default:     nil,
				Computed:    true,
			},
			"db_engine": {
				Type:        schema.TypeString,
				Description: "Specifies the version of the engine being used by the database (e.g. oracle-ee, oracle-se, oracle-se1, oracle-se2)",
				Required:    false,
				Optional:    true,
				Default:     nil,
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
			"logstore": {
				Type:        schema.TypeString,
				Description: "Unit that is used to collect, store and query logs",
				Required:    false,
				Optional:    true,
				Default:     nil,
				Computed:    true,
			},
			"logs_destination_asset_id": {
				Type:        schema.TypeString,
				Description: "The asset name of the log aggregator that stores this asset's logs.",
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
			"max_concurrent_conn": {
				Type:        schema.TypeString,
				Description: "Maximum number of concurrent connections that sensitive data management should use at once.",
				Required:    false,
				Optional:    true,
				Default:     nil,
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
			"pull_type": {
				Type:        schema.TypeString,
				Description: "The method used to pull data from the logstore.",
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
			"sdm_enabled": {
				Type:        schema.TypeBool,
				Description: "Sensitive data management (SDM) is enabled if this parameter is set to True.",
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
		},
	}
}

func resourceLogAggregatorCreateContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Client)

	// check provided fields against schema
	if isOk, err := checkResourceRequiredFields(requiredLogAggregatorJson, ignoreLogAggregatorParamsByServerType, d); !isOk {
		return diag.FromErr(err)
	}
	// convert provided fields into API payload
	logAggregator := ResourceWrapper{}
	serverType := d.Get("server_type").(string)
	createResource(&logAggregator, serverType, d)

	// auditPullEnabled set to false as connect/disconnect logic handled below
	logAggregator.Data.AssetData.AuditPullEnabled = false

	// create resource
	log.Printf("[INFO] Creating LogAggregator for serverType: %s and gatewayId: %s\n", logAggregator.Data.ServerType, logAggregator.Data.GatewayID)
	createLogAggregatorResponse, err := client.CreateLogAggregator(logAggregator)
	if err != nil {
		log.Printf("[ERROR] adding LogAggregator for serverType: %s and gatewayId: %s | err: %s", serverType, logAggregator.Data.GatewayID, err)
		return diag.FromErr(err)
	}

	// get asset_id
	assetId := d.Get("asset_id").(string)

	// wait for remoteSyncState
	err = waitForRemoteSyncState(ctx, dsfLogAggregatorResourceType, assetId, m)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  fmt.Sprintf("Error while waiting for remoteSyncState = \"SYNCED\" for asset: %s", assetId),
			Detail:   fmt.Sprintf("Error: %s\n", err),
		})
	}

	// Connect/disconnect asset to gateway
	err = connectDisconnectGateway(ctx, d, dsfLogAggregatorResourceType, m)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  fmt.Sprintf("Error while updating audit state for asset: %s", assetId),
			Detail:   fmt.Sprintf("Error: %s\n", err),
		})
	}

	// Set ID
	logAggregatorId := createLogAggregatorResponse.Data.ID
	d.SetId(logAggregatorId)

	// Set the rest of the state from the resource read
	log.Printf("[DEBUG] Writing log aggregator asset details to state")
	resourceLogAggregatorReadContext(ctx, d, m)

	return diags
}

func resourceLogAggregatorReadContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	logAggregatorId := d.Id()

	log.Printf("[INFO] Reading LogAggregator with logAggregatorId: %s\n", logAggregatorId)

	logAggregatorReadResponse, err := client.ReadLogAggregator(logAggregatorId)

	if err != nil {
		log.Printf("[ERROR] Reading logAggregatorReadResponse with logAggregatorId: %s | err: %s\n", logAggregatorId, err)
		return diag.FromErr(err)
	}

	if logAggregatorReadResponse != nil {
		log.Printf("[INFO] Reading LogAggregator with logAggregatorId: %s | err: %s\n", logAggregatorId, err)
	}

	log.Printf("[DEBUG] logAggregatorReadResponse: %s\n", logAggregatorReadResponse.Data.ID)
	// Set returned and computed values
	d.Set("admin_email", logAggregatorReadResponse.Data.AssetData.AdminEmail)
	d.Set("application", logAggregatorReadResponse.Data.AssetData.Application)
	d.Set("arn", logAggregatorReadResponse.Data.AssetData.Arn)
	d.Set("asset_display_name", logAggregatorReadResponse.Data.AssetData.AssetDisplayName)
	d.Set("asset_id", logAggregatorReadResponse.Data.AssetData.AssetID)
	d.Set("asset_source", logAggregatorReadResponse.Data.AssetData.AssetSource)
	d.Set("asset_version", logAggregatorReadResponse.Data.AssetData.Version)
	d.Set("audit_data_type", logAggregatorReadResponse.Data.AssetData.AuditDataType)
	d.Set("audit_pull_enabled", logAggregatorReadResponse.Data.AssetData.AuditPullEnabled)
	d.Set("audit_type", logAggregatorReadResponse.Data.AssetData.AuditType)
	d.Set("available_bucket_account_ids", logAggregatorReadResponse.Data.AssetData.AvailableBucketAccountIds)
	d.Set("available_regions", logAggregatorReadResponse.Data.AssetData.AvailableRegions)
	d.Set("bucket_account_id", logAggregatorReadResponse.Data.AssetData.BucketAccountId)
	d.Set("ca_certs_path", logAggregatorReadResponse.Data.AssetData.CaCertsPath)
	d.Set("ca_file", logAggregatorReadResponse.Data.AssetData.CaFile)
	d.Set("consumer_group", logAggregatorReadResponse.Data.AssetData.ConsumerGroup)
	d.Set("consumer_group_workers", logAggregatorReadResponse.Data.AssetData.ConsumerGroupWorkers)
	d.Set("consumer_worker_prefix", logAggregatorReadResponse.Data.AssetData.ConsumerWorkerPrefix)
	d.Set("content_type", logAggregatorReadResponse.Data.AssetData.ContentType)
	if logAggregatorReadResponse.Data.AssetData.CredentialsEndpoint != "" {
		d.Set("credentials_endpoint", logAggregatorReadResponse.Data.AssetData.CredentialsEndpoint)
	}
	d.Set("criticality", logAggregatorReadResponse.Data.AssetData.Criticality)
	d.Set("database_name", logAggregatorReadResponse.Data.AssetData.DatabaseName)
	d.Set("db_engine", logAggregatorReadResponse.Data.AssetData.DbEngine)
	// enable_audit_management / enable_audit_monitoring
	d.Set("endpoint", logAggregatorReadResponse.Data.AssetData.Endpoint)
	// entitlement_enabled
	d.Set("gateway_id", logAggregatorReadResponse.Data.GatewayID)
	d.Set("gateway_service", logAggregatorReadResponse.Data.AssetData.GatewayService)
	d.Set("id", logAggregatorReadResponse.Data.ID)
	d.Set("jsonar_uid", logAggregatorReadResponse.Data.AssetData.JsonarUID)
	d.Set("location", logAggregatorReadResponse.Data.AssetData.Location)
	d.Set("logs_destination_asset_id", logAggregatorReadResponse.Data.AssetData.LogsDestinationAssetID)
	d.Set("logstore", logAggregatorReadResponse.Data.AssetData.Logstore)
	d.Set("managed_by", logAggregatorReadResponse.Data.AssetData.ManagedBy)
	d.Set("max_concurrent_conn", logAggregatorReadResponse.Data.AssetData.MaxConcurrentConn)
	d.Set("owned_by", logAggregatorReadResponse.Data.AssetData.OwnedBy)
	d.Set("parent_asset_id", logAggregatorReadResponse.Data.AssetData.ParentAssetID)
	// prefix
	d.Set("project", logAggregatorReadResponse.Data.AssetData.Project)
	d.Set("proxy", logAggregatorReadResponse.Data.AssetData.Proxy)
	d.Set("pubsub_subscription", logAggregatorReadResponse.Data.AssetData.PubsubSubscription)
	d.Set("pull_type", logAggregatorReadResponse.Data.AssetData.PullType)
	d.Set("region", logAggregatorReadResponse.Data.AssetData.Region)
	d.Set("s3_provider", logAggregatorReadResponse.Data.AssetData.S3Provider) // TODO: may not be supported in all DSF versions
	d.Set("server_host_name", logAggregatorReadResponse.Data.AssetData.ServerHostName)
	d.Set("server_ip", logAggregatorReadResponse.Data.AssetData.ServerIP)
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
	d.Set("sdm_enabled", logAggregatorReadResponse.Data.AssetData.SdmEnabled)
	// subscribers
	// unmask
	d.Set("used_for", logAggregatorReadResponse.Data.AssetData.UsedFor)

	if logAggregatorReadResponse.Data.AssetData.AuditInfo != nil {
		auditInfo := &schema.Set{F: resourceAssetDataAuditInfoHash}
		auditInfoMap := map[string]interface{}{}
		auditInfoMap["policy_template_name"] = logAggregatorReadResponse.Data.AssetData.AuditInfo.PolicyTemplateName
		auditInfo.Add(auditInfoMap)
		d.Set("audit_info", auditInfo)
	}

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
		connection["access_id"] = v.ConnectionData.AccessID   // TODO SR-4549
		connection["access_key"] = v.ConnectionData.AccessKey // TODO SR-4549
		connection["application_id"] = v.ConnectionData.ApplicationID
		connection["auth_mechanism"] = v.ConnectionData.AuthMechanism
		connection["azure_storage_account"] = v.ConnectionData.AzureStorageAccount
		connection["azure_storage_container"] = v.ConnectionData.AzureStorageContainer
		connection["azure_storage_secret_key"] = v.ConnectionData.AzureStorageSecretKey
		connection["cache_file"] = v.ConnectionData.CacheFile
		connection["ca_certs_path"] = v.ConnectionData.CaCertsPath // TODO SR-4549
		connection["client_secret"] = v.ConnectionData.ClientSecret
		connection["db_role"] = v.ConnectionData.DbRole
		connection["directory_id"] = v.ConnectionData.DirectoryID
		connection["eventhub_access_key"] = v.ConnectionData.EventhubAccessKey
		connection["eventhub_access_policy"] = v.ConnectionData.EventhubAccessPolicy
		connection["eventhub_name"] = v.ConnectionData.EventhubName
		connection["eventhub_namespace"] = v.ConnectionData.EventhubNamespace
		connection["external"] = v.ConnectionData.External
		connection["external_id"] = v.ConnectionData.ExternalID // TODO SR-4549
		connection["extra_kinit_parameters"] = v.ConnectionData.ExtraKinitParameters
		connection["format"] = v.ConnectionData.Format
		connection["kerberos_kdc"] = v.ConnectionData.KerberosKdc
		connection["kerberos_service_realm"] = v.ConnectionData.KerberosServiceRealm
		connection["kerberos_service_kdc"] = v.ConnectionData.KerberosServiceKdc
		connection["kerberos_spn"] = v.ConnectionData.KerberosSpn
		connection["key_file"] = v.ConnectionData.KeyFile
		connection["keytab_file"] = v.ConnectionData.KeytabFile
		connection["kinit_program_path"] = v.ConnectionData.KinitProgramPath
		connection["passphrase"] = v.ConnectionData.Passphrase
		connection["password"] = v.ConnectionData.Password
		connection["principal"] = v.ConnectionData.Principal
		connection["reason"] = v.Reason
		connection["region"] = v.ConnectionData.Region
		connection["role_name"] = v.ConnectionData.RoleName // TODO SR-4549
		connection["secret_key"] = v.ConnectionData.SecretKey
		// sftp_test TODO SR-4549
		connection["ssl"] = v.ConnectionData.Ssl
		connection["ssl_server_cert"] = v.ConnectionData.SslServerCert
		connection["subscription_id"] = v.ConnectionData.SubscriptionID
		connection["use_keytab"] = v.ConnectionData.UseKeytab
		connection["user_identity_client_id"] = v.ConnectionData.UserIdentityClientID
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

	log.Printf("[INFO] Finished reading logAggregator with logAggregatorId: %s\n", logAggregatorId)

	return nil
}

func resourceLogAggregatorUpdateContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Client)

	// check provided fields against schema
	logAggregatorId := d.Id()
	if isOk, err := checkResourceRequiredFields(requiredLogAggregatorJson, ignoreLogAggregatorParamsByServerType, d); !isOk {
		return diag.FromErr(err)
	}

	// convert provided fields into API payload
	logAggregator := ResourceWrapper{}
	serverType := d.Get("server_type").(string)
	createResource(&logAggregator, serverType, d)

	// auditPullEnabled set to current value from state
	auditPullEnabled, _ := d.GetChange("audit_pull_enabled")
	logAggregator.Data.AssetData.AuditPullEnabled = auditPullEnabled.(bool)

	// update resource
	log.Printf("[INFO] Updating LogAggregator for serverType: %s and gatewayId: %s assetId: %s\n", logAggregator.Data.ServerType, logAggregator.Data.GatewayID, logAggregator.Data.AssetData.AssetID)
	_, err := client.UpdateLogAggregator(logAggregatorId, logAggregator)
	if err != nil {
		log.Printf("[ERROR] Updating LogAggregator for serverType: %s and gatewayId: %s assetId: %s | err:%s\n", logAggregator.Data.ServerType, logAggregator.Data.GatewayID, logAggregator.Data.AssetData.AssetID, err)
		return diag.FromErr(err)
	}

	// get asset_id
	assetId := d.Get("asset_id").(string)

	// wait for remoteSyncState
	err = waitForRemoteSyncState(ctx, dsfLogAggregatorResourceType, assetId, m)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  fmt.Sprintf("Error while waiting for remoteSyncState = \"SYNCED\" for asset: %s", assetId),
			Detail:   fmt.Sprintf("Error: %s\n", err),
		})
	}

	// Connect/disconnect asset to gateway
	err = connectDisconnectGateway(ctx, d, dsfLogAggregatorResourceType, m)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  fmt.Sprintf("Error while updating audit state for asset: %s", assetId),
			Detail:   fmt.Sprintf("Error: %s\n", err),
		})
	}

	// Set ID
	d.SetId(logAggregatorId)

	// Set the rest of the state from the resource read
	log.Printf("[DEBUG] Writing log aggregator asset details to state")
	resourceLogAggregatorReadContext(ctx, d, m)

	return diags
}

func resourceLogAggregatorDeleteContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	logAggregatorId := d.Id()

	log.Printf("[INFO] Deleting log aggregator with logAggregatorId: %s", logAggregatorId)

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

	if v, ok := m["cache_file"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["ca_certs_path"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["client_secret"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["db_role"]; ok {
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

	if v, ok := m["external"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(bool)))
	}

	if v, ok := m["extra_kinit_parameters"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["format"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["key_file"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["kerberos_kdc"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["kerberos_service_kdc"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["kerberos_service_realm"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["kerberos_spn"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["keytab_file"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["kinit_program_path"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["passphrase"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["password"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["principal"]; ok {
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

	if v, ok := m["ssl_server_cert"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["subscription_id"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["use_keytab"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(bool)))
	}

	if v, ok := m["username"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["user_identity_client_id"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	return PositiveHash(buf.String())
}
