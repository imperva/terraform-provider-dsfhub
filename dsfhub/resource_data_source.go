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

func resourceDSFDataSource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDSFDataSourceCreateContext,
		ReadContext:   resourceDSFDataSourceReadContext,
		UpdateContext: resourceDSFDataSourceUpdateContext,
		DeleteContext: resourceDSFDataSourceDeleteContext,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"admin_email": {
				Type:        schema.TypeString,
				Description: "The email address to notify about this asset",
				Required:    true,
			},
			// "appliance_type": {
			// 	Type:        schema.TypeString,
			// 	Description: "Appliance type",
			// 	Required:    false,
			// 	Optional:    true,
			// 	Computed:    true,
			// },
			"application": {
				Type:        schema.TypeString,
				Description: "The Asset ID of the application asset that \"owns\" the asset.",
				Required:    false,
				Optional:    true,
			},
			// "archive": {
			// 	Type:        schema.TypeBool,
			// 	Description: "If True archive files in the asset after being processed by sonargd. Defaults to True if field isn't present",
			// 	Required:    false,
			// 	Optional:    true,
			// },
			"arn": {
				Type:        schema.TypeString,
				Description: "Amazon Resource Name - format is arn:partition:service:region:account-id:resource-type:resource-id and used as the asset_id",
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     nil,
			},
			"asset_connection": {
				Type:        schema.TypeSet,
				Description: "N/A",
				Required:    false,
				Optional:    true,
				MinItems:    0,
				Set:         resourceDataSourceConnectionHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_id": {
							Type:        schema.TypeString,
							Description: "The Account Name/Access ID to use when authenticating to Snowflake",
							Required:    false,
							Optional:    true,
							Default:     nil,
							Computed:    true,
						},
						"account_name": {
							Type:        schema.TypeString,
							Description: "The cloudant account name required when connecting a resource with IAM role",
							Required:    false,
							Optional:    true,
							Default:     nil,
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
						"api_key": {
							Type:        schema.TypeString,
							Description: "IAM authentication API key",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"auth_mechanism": {
							Type:         schema.TypeString,
							Description:  "Specifies the auth mechanism used by the connection",
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"aws_credentials", "default", "ec2", "iam_role", "kerberos", "key_file", "key", "oauth-azure-ad", "oauth", "oauth2", "oracle_wallet", "password", "profile", "service_account", "ssl"}, false),
						},
						"autocommit": {
							Type:        schema.TypeBool,
							Description: "If true, Commit automatically don't wait for transaction to be explicitly committed",
							Required:    false,
							Optional:    true,
							Default:     nil,
							Computed:    true,
						},
						"aws_connection_id": {
							Type:        schema.TypeString,
							Description: "The parent AWS connection document_id",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"bucket": {
							Type:        schema.TypeString,
							Description: "",
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
							Description: "Use the specified certificate file to verify the peer. The file may contain multiple CA certificates.",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"cache_file": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"cert_file": {
							Type:        schema.TypeString,
							Description: "Use the specified client certificate file when getting a file with HTTPS, FTPS or another SSL-based protocol.",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"client_id": {
							Type:        schema.TypeString,
							Description: "Azure client application ID",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"client_secret": {
							Type:        schema.TypeString,
							Description: "Azure application client secret",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"cluster_id": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"cluster_member_id": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"cluster_name": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"content_type": {
							Type:        schema.TypeString,
							Description: "Content-Type to append to the HTTP headers in the curl call",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"crn": {
							Type:        schema.TypeString,
							Description: "The CRN unique identifier of the resource",
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
						"database_name": {
							Type:        schema.TypeString,
							Description: "Specifies the name of the database (or default DB) to connect to.",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"db_role": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"dn": {
							Type:        schema.TypeString,
							Description: "The distinguished name of a particular PKI certificate",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"dns_srv": {
							Type:        schema.TypeBool,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
							Computed:    true,
						},
						"driver": {
							Type:        schema.TypeString,
							Description: "A path to a non-default driver location. If populated this driver will be used rather than the default",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"dsn": {
							Type:        schema.TypeString,
							Description: "Data Source Name",
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
						"external_id": {
							Type:        schema.TypeString,
							Description: "External ID to use when assuming a role",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"extra_kinit_parameters": {
							Type:        schema.TypeString,
							Description: "",
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
						"hive_server_type": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"host_name_mismatch": {
							Type:        schema.TypeBool,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
							Computed:    true,
						},
						"hosts": {
							Type:        schema.TypeString,
							Description: "Required for quering the logdna url. cloudantnosqldb in the case of a cloudant DB",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"httppath": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"is_cluster": {
							Type:        schema.TypeBool,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
							Computed:    true,
						},
						"jdbc_ssl_trust_server_certificate": {
							Type:        schema.TypeBool,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
							Computed:    true,
						},
						"jdbc_ssl_trust_store_location": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"jdbc_ssl_trust_store_password": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"kerberos_host_fqdn": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"kerberos_kdc": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"kerberos_retry_count": {
							Type:        schema.TypeInt,
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
							Description: "Private key file name. Allows you to provide your private key in this separate file.",
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
						"net_service_name": {
							Type:        schema.TypeString,
							Description: "Alias in tnsnames.ora replaces hostname, service name, and port in connection string",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"oauth_parameters": {
							Type:        schema.TypeSet,
							Description: "Additional parameters to pass when requesting a token",
							Required:    false,
							Optional:    true,
							Default:     nil,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"odbc_connection_string": {
							Type:        schema.TypeString,
							Description: "Additional ODBC connection string parameters. This string will get added to the connection string",
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
						"proxy_auto_detect": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"proxy_password": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"proxy_port": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"proxy_server": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"proxy_ssl_type": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"reason": {
							Type:        schema.TypeString,
							Description: "N/A",
							Required:    true,
							//ValidateFunc: validation.StringInSlice([]string{"default", "sonargateway", "ad-hoc-query", "audit management", "SDM"}, false),
						},
						"redirect_uri": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"region": {
							Type:        schema.TypeString,
							Description: "The cloud geography/region/zone/data center that the resource resides",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"replica_set": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"resource_id": {
							Type:        schema.TypeString,
							Description: "Azure resource application ID",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"role_name": {
							Type:        schema.TypeString,
							Description: "Role to use for authentication",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"schema": {
							Type:        schema.TypeString,
							Description: "Schema name. A schema is a logical grouping of database objects",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"secret_key": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"self_signed": {
							Type:        schema.TypeBool,
							Description: "Accept self-signed certificates",
							Required:    false,
							Optional:    true,
							Default:     nil,
							Computed:    true,
						},
						"self_signed_cert": {
							Type:        schema.TypeBool,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
							Computed:    true,
						},
						"server_port": {
							Type:        schema.TypeInt,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"sec_before_operating_expired_token": {
							Type:        schema.TypeInt,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"service_key": {
							Type:        schema.TypeString,
							Description: "The service key required in the logdna url query to connect to logdna and pull the logs",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"sid": {
							Type:        schema.TypeString,
							Description: "SID used to connect, e.g. ORCL",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"snowflake_role": {
							Type:        schema.TypeString,
							Description: "Role with which to log into Snowflake",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"ssl": {
							Type:        schema.TypeBool,
							Description: "If true, use SSL when connecting",
							Required:    false,
							Optional:    true,
							Default:     nil,
							// Computed:    true,
						},
						"ssl_server_cert": {
							Type:        schema.TypeString,
							Description: "Path to server certificate to use during authentication",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"tenant_id": {
							Type:        schema.TypeString,
							Description: "Azure tenant ID",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"thrift_transport": {
							Type:        schema.TypeInt,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"tmp_user": {
							Type:        schema.TypeBool,
							Description: "If true create a temporary user",
							Required:    false,
							Optional:    true,
							Default:     nil,
							Computed:    true,
						},
						"token": {
							Type:        schema.TypeString,
							Description: "Saved token to use for authentication",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"token_endpoint": {
							Type:        schema.TypeString,
							Description: "URL of endpoint to query when requesting a token",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"transportmode": {
							Type:        schema.TypeString,
							Description: "",
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
							Description: "The user to use when connecting",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"virtual_hostname": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"virtual_ip": {
							Type:        schema.TypeString,
							Description: "",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"wallet_dir": {
							Type:        schema.TypeString,
							Description: "Path to the Oracle wallet directory",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
						"warehouse": {
							Type:        schema.TypeString,
							Description: "The name of the warehouse to connect to",
							Required:    false,
							Optional:    true,
							Default:     nil,
						},
					},
				},
			},
			"asset_display_name": {
				Type:        schema.TypeString,
				Description: "User-friendly name of the asset, defined by user.",
				Required:    true,
			},
			"asset_id": {
				Type:        schema.TypeString,
				Description: "The unique identifier or resource name of the asset. The recommended format is /subscriptions/my-subscription-id/resourceGroups/my-resource-group/providers/Microsoft.DocumentDb/databaseAccounts/my-cosmos-table",
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
				Computed:    true,
				Default:     nil,
			},
			// "audit_state": {
			// 	Type:         schema.TypeString,
			// 	Description:  "Audit state",
			// 	Required:     false,
			// 	Optional:     true,
			// 	ValidateFunc: validation.StringInSlice([]string{"COSMOS_TABLE"}, false),
			// 	Computed:     true,
			// },
			"audit_type": {
				Type:        schema.TypeString,
				Description: "Used to indicate what mechanism should be used to fetch logs on systems supporting multiple ways to get logs, see asset specific documentation for details",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"availability_zones": {
				Type:        schema.TypeList,
				Description: "List of regions where the cluster is available",
				Required:    false,
				Optional:    true,
				Default:     nil,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"cluster_engine": {
				Type:        schema.TypeString,
				Description: "",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"cluster_id": {
				Type:        schema.TypeString,
				Description: "",
				Required:    false,
				Optional:    true,
				Default:     nil,
				Computed:    true,
			},
			"cluster_member_id": {
				Type:        schema.TypeString,
				Description: "",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"cluster_name": {
				Type:        schema.TypeString,
				Description: "",
				Required:    false,
				Optional:    true,
				Default:     nil,
				Computed:    true,
			},
			"content_type": {
				Type:        schema.TypeString,
				Description: "Content type should be set to the desired <'parent' asset 'Server Type'>, which is the one that uses this asset as a destination for logs. NOTE: The content_type field will take precedence on the lookup for parent_asset_id field when checking which server is sending logs to this asset.",
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
			"db_instances_display_name": {
				Type:        schema.TypeList,
				Description: "",
				Required:    false,
				Optional:    true,
				Default:     nil,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"duration_threshold": {
				Type:        schema.TypeInt,
				Description: "",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"enable_audit_management": {
				Type:        schema.TypeBool,
				Description: "If true, Sonar is responsible for setting and updating the policies",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"enable_audit_monitoring": {
				Type:        schema.TypeBool,
				Description: "If true, Sonar sends emails/alerts when the audit policies change.",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"enabled_logs_exports": {
				Type:        schema.TypeList,
				Description: "",
				Required:    false,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"entitlement_enabled": {
				Type:        schema.TypeBool,
				Description: "If true, Entitlement Management system is enabled.",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"gateway_id": {
				Type:        schema.TypeString,
				Description: "The jsonarUid unique identifier of the agentless gateway. Example: '7a4af7cf-4292-89d9-46ec-183756ksdjd'",
				Required:    true,
			},
			// "gateway_name": {
			// 	Type:        schema.TypeString,
			// 	Description: "The jsonarUid unique identifier of the agentless gateway. Example: '7a4af7cf-4292-89d9-46ec-183756ksdjd'",
			// 	Required:    false,
			// 	Optional:    true,
			// 	Computed:    true,
			// },
			"gateway_service": {
				Type:        schema.TypeString,
				Description: "The name of the gateway pull service (if any) used to retrieve logs for this source. Usually set by the connect gateway playbook.",
				Required:    false,
				Optional:    true,
				Default:     nil,
				Computed:    true,
			},
			"host_timezone_offset": {
				Type:        schema.TypeString,
				Description: "The offset value string is in the format \"-/+hh:mm\"",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"id": {
				Type:        schema.TypeString,
				Description: "Unique identifier for the asset",
				Computed:    true,
			},
			"ignore_latest_of": {
				Type:        schema.TypeString,
				Description: "A regex defining a group. From all the files with the same group, the latest one will be ignored, so that it isn't archived until server is done writing",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"is_cluster": {
				Type:        schema.TypeBool,
				Description: "",
				Required:    false,
				Optional:    true,
				Default:     nil,
				Computed:    true,
			},
			"is_multi_zones": {
				Type:        schema.TypeBool,
				Description: "",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"jsonar_uid": {
				Type:        schema.TypeString,
				Description: "Unique identifier (UID) attached to the Agentless Gateway controlling the asset",
				Required:    false,
				Optional:    true,
				Default:     nil,
				Computed:    true,
			},
			"jsonar_uid_display_name": {
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
			"log_bucket_id": {
				Type:        schema.TypeString,
				Description: "Asset ID of the S3 bucket which stores the logs for this server",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"logs_destination_asset_id": {
				Type:        schema.TypeString,
				Description: "The asset name of the log aggregator that stores this asset's logs.",
				Required:    false,
				Optional:    true,
				Default:     nil,
				Computed:    true,
			},
			"marker_alias": {
				Type:        schema.TypeString,
				Description: "Cluster or System name for a DR pair or similar system where all nodes share a single log. All machines sharing a marker alias will use the same marker. This means that the log will be pulled once rather than once per machine.",
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
			"provider_url": {
				Type:        schema.TypeString,
				Description: "",
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
				Description: "",
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
			"resource_id": {
				Type:        schema.TypeString,
				Description: "AWS Resource ID that the audit logs will be stored under on S3. E.g. db-3TBJU4Y34IAVE2DQRQUWYOEX3I",
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
			"searches": {
				Type:        schema.TypeList,
				Description: "",
				Required:    false,
				Optional:    true,
				Default:     nil,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"server_port": {
				Type:        schema.TypeString,
				Description: "Port used by the source server",
				Required:    false,
				Optional:    true,
				Default:     nil,
				Computed:    true,
			},
			"server_type": {
				Type:        schema.TypeString,
				Description: "The type of server or data service to be created as a data source. The list of available data sources is documented at: https://docs.imperva.com/bundle/v4.11-sonar-user-guide/page/84552.htm",
				Required:    true,
			},
			"service_endpoint": {
				Type:        schema.TypeString,
				Description: "Specify a particular endpoint for a given service",
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
			"service_name": {
				Type:        schema.TypeString,
				Description: "",
				Required:    false,
				Optional:    true,
				Default:     nil,
				Computed:    true,
			},
			"ssl": {
				Type:        schema.TypeBool,
				Description: "",
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
			"used_for": {
				Type:         schema.TypeString,
				Description:  "Designates how this asset is used / the environment that the asset is supporting.",
				Required:     false,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Development", "Staging", "Test", "Disaster Recovery", "Demonstration", "Production", "QA", "Training"}, false),
				Default:      nil,
			},
			"virtual_hostname": {
				Type:        schema.TypeString,
				Description: "",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"virtual_ip": {
				Type:        schema.TypeString,
				Description: "",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
			"xel_directory": {
				Type:        schema.TypeString,
				Description: "Absolute path of the XEL files location",
				Required:    false,
				Optional:    true,
				Default:     nil,
			},
		},
	}
}

func resourceDSFDataSourceCreateContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Client)

	// check provided fields against schema
	if isOk, err := checkResourceRequiredFields(requiredDataSourceFieldsJson, ignoreDataSourceParamsByServerType, d); !isOk {
		return diag.FromErr(err)
	}

	// convert provided fields into API payload
	dsfDataSource := ResourceWrapper{}
	serverType := d.Get("server_type").(string)
	createResource(&dsfDataSource, serverType, d)

	// auditPullEnabled set to false as connect/disconnect logic handled below
	dsfDataSource.Data.AssetData.AuditPullEnabled = false

	// create resource
	log.Printf("[INFO] Creating DSF data source for serverType: %s and gatewayId: %s \n", dsfDataSource.Data.ServerType, dsfDataSource.Data.GatewayID)
	dsfDataSourceResponse, err := client.CreateDSFDataSource(dsfDataSource)
	if err != nil {
		log.Printf("[INFO] Creating DSF data source for serverType: %s and gatewayId: %s assetId: %s\n", dsfDataSource.Data.ServerType, dsfDataSource.Data.GatewayID, dsfDataSource.Data.AssetData.AssetID)
		return diag.FromErr(err)
	}

	// get asset_id
	assetId := d.Get("asset_id").(string)

	// wait for remoteSyncState
	err = waitForRemoteSyncState(ctx, dsfDataSourceResourceType, assetId, m)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  fmt.Sprintf("Error while waiting for remoteSyncState = \"SYNCED\" for asset: %s", assetId),
			Detail:   fmt.Sprintf("Error: %s\n", err),
		})
	}

	// Connect/disconnect asset to gateway
	err = connectDisconnectGateway(ctx, d, dsfDataSourceResourceType, m)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  fmt.Sprintf("Error while updating audit state for asset: %s", assetId),
			Detail:   fmt.Sprintf("Error: %s\n", err),
		})
	}

	// Set ID
	dsfDataSourceId := dsfDataSourceResponse.Data.AssetData.AssetID
	d.SetId(dsfDataSourceId)

	// Set the rest of the state from the resource read
	log.Printf("[DEBUG] Writing data source asset details to state")
	resourceDSFDataSourceReadContext(ctx, d, m)

	return diags
}

func resourceDSFDataSourceReadContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	dsfDataSourceId := d.Id()

	log.Printf("[INFO] Reading DSF data source with dsfDataSourceId: %s\n", dsfDataSourceId)
	dsfDataSourceReadResponse, err := client.ReadDSFDataSource(dsfDataSourceId)
	if err != nil {
		log.Printf("[ERROR] Reading dsfDataSourceReadResponse | err: %s\n", err)
		return diag.FromErr(err)
	}

	if dsfDataSourceReadResponse != nil {
		log.Printf("[INFO] Reading DSF data source with dsfDataSourceId: %s | err: %s\n", dsfDataSourceId, err)
	}

	log.Printf("[DEBUG] dsfDataSourceReadResponse: %s\n", dsfDataSourceReadResponse.Data.ID)

	// Set returned and computed values
	d.Set("admin_email", dsfDataSourceReadResponse.Data.AssetData.AdminEmail)
	d.Set("application", dsfDataSourceReadResponse.Data.AssetData.Application)
	//d.Set("archive", dsfDataSourceReadResponse.Data.AssetData.Archive)

	if dsfDataSourceReadResponse.Data.AssetData.Arn != "" {
		d.Set("arn", dsfDataSourceReadResponse.Data.AssetData.Arn)
	}
	d.Set("asset_display_name", dsfDataSourceReadResponse.Data.AssetData.AssetDisplayName)
	d.Set("asset_id", dsfDataSourceReadResponse.Data.AssetData.AssetID)
	d.Set("asset_source", dsfDataSourceReadResponse.Data.AssetData.AssetSource)
	d.Set("asset_version", dsfDataSourceReadResponse.Data.AssetData.Version)
	d.Set("audit_pull_enabled", dsfDataSourceReadResponse.Data.AssetData.AuditPullEnabled)
	d.Set("audit_type", dsfDataSourceReadResponse.Data.AssetData.AuditType)
	d.Set("availability_zones", dsfDataSourceReadResponse.Data.AssetData.AvailabilityZones)
	d.Set("available_bucket_account_ids", dsfDataSourceReadResponse.Data.AssetData.AvailableBucketAccountIds)
	d.Set("available_regions", dsfDataSourceReadResponse.Data.AssetData.AvailableRegions)
	d.Set("bucket_account_id", dsfDataSourceReadResponse.Data.AssetData.BucketAccountId)
	d.Set("ca_certs_path", dsfDataSourceReadResponse.Data.AssetData.CaCertsPath)
	d.Set("ca_file", dsfDataSourceReadResponse.Data.AssetData.CaFile)
	d.Set("cluster_engine", dsfDataSourceReadResponse.Data.AssetData.ClusterEngine)
	d.Set("cluster_id", dsfDataSourceReadResponse.Data.AssetData.ClusterID)
	d.Set("cluster_member_id", dsfDataSourceReadResponse.Data.AssetData.ClusterMemberID)
	d.Set("cluster_name", dsfDataSourceReadResponse.Data.AssetData.ClusterName)
	d.Set("content_type", dsfDataSourceReadResponse.Data.AssetData.ContentType)
	d.Set("credentials_endpoint", dsfDataSourceReadResponse.Data.AssetData.CredentialsEndpoint)
	d.Set("criticality", dsfDataSourceReadResponse.Data.AssetData.Criticality)
	d.Set("database_name", dsfDataSourceReadResponse.Data.AssetData.DatabaseName)
	d.Set("db_engine", dsfDataSourceReadResponse.Data.AssetData.DbEngine)
	d.Set("db_instances_display_name", dsfDataSourceReadResponse.Data.AssetData.DbInstancesDisplayName)
	// deduplication_filter
	// driver
	d.Set("duration_threshold", dsfDataSourceReadResponse.Data.AssetData.DurationThreshold)
	d.Set("enable_audit_management", dsfDataSourceReadResponse.Data.AssetData.EnableAuditManagement)
	d.Set("enable_audit_monitoring", dsfDataSourceReadResponse.Data.AssetData.EnableAuditMonitoring)
	d.Set("enabled_logs_exports", dsfDataSourceReadResponse.Data.AssetData.EnabledLogsExports)
	d.Set("entitlement_enabled", dsfDataSourceReadResponse.Data.AssetData.EntitlementEnabled)
	d.Set("gateway_id", dsfDataSourceReadResponse.Data.GatewayID)
	d.Set("gateway_service", dsfDataSourceReadResponse.Data.AssetData.GatewayService)
	d.Set("host_timezone_offset", dsfDataSourceReadResponse.Data.AssetData.HostTimezoneOffset)
	d.Set("id", dsfDataSourceReadResponse.Data.ID)
	d.Set("ignore_latest_of", dsfDataSourceReadResponse.Data.AssetData.IgnoreLatestOf)
	d.Set("is_cluster", dsfDataSourceReadResponse.Data.AssetData.IsCluster)
	d.Set("is_multi_zones", dsfDataSourceReadResponse.Data.AssetData.IsMultiZones)
	d.Set("jsonar_uid", dsfDataSourceReadResponse.Data.AssetData.JsonarUID)
	d.Set("location", dsfDataSourceReadResponse.Data.AssetData.Location)
	d.Set("log_bucket_id", dsfDataSourceReadResponse.Data.AssetData.LogBucketID)
	d.Set("logs_destination_asset_id", dsfDataSourceReadResponse.Data.AssetData.LogsDestinationAssetID)
	d.Set("managed_by", dsfDataSourceReadResponse.Data.AssetData.ManagedBy)
	d.Set("marker_alias", dsfDataSourceReadResponse.Data.AssetData.MarkerAlias)
	d.Set("max_concurrent_conn", dsfDataSourceReadResponse.Data.AssetData.MaxConcurrentConn)
	// namespace_name
	d.Set("owned_by", dsfDataSourceReadResponse.Data.AssetData.OwnedBy)
	d.Set("parent_asset_id", dsfDataSourceReadResponse.Data.AssetData.ParentAssetID)
	// prefix
	d.Set("provider_url", dsfDataSourceReadResponse.Data.AssetData.ProviderUrl)
	d.Set("proxy", dsfDataSourceReadResponse.Data.AssetData.Proxy)
	d.Set("pubsub_subscription", dsfDataSourceReadResponse.Data.AssetData.PubsubSubscription)
	d.Set("region", dsfDataSourceReadResponse.Data.AssetData.Region)
	d.Set("resource_id", dsfDataSourceReadResponse.Data.AssetData.ResourceID)
	d.Set("sdm_enabled", dsfDataSourceReadResponse.Data.AssetData.SdmEnabled)
	d.Set("searches", dsfDataSourceReadResponse.Data.AssetData.Searches)
	d.Set("server_host_name", dsfDataSourceReadResponse.Data.AssetData.ServerHostName)
	if dsfDataSourceReadResponse.Data.AssetData.ServerIP != "" {
		d.Set("server_ip", dsfDataSourceReadResponse.Data.AssetData.ServerIP)
	}
	if dsfDataSourceReadResponse.Data.AssetData.ServerPort != nil {
		var serverPort string
		if serverPortNum, ok := dsfDataSourceReadResponse.Data.AssetData.ServerPort.(float64); ok {
			serverPort = fmt.Sprintf("%d", int(serverPortNum))
		} else {
			serverPort = dsfDataSourceReadResponse.Data.AssetData.ServerPort.(string)
		}
		d.Set("server_port", serverPort)
	}
	d.Set("server_type", dsfDataSourceReadResponse.Data.ServerType)
	d.Set("service_endpoint", dsfDataSourceReadResponse.Data.AssetData.ServiceEndpoint)
	d.Set("service_name", dsfDataSourceReadResponse.Data.AssetData.ServiceName)
	d.Set("ssl", dsfDataSourceReadResponse.Data.AssetData.Ssl)
	d.Set("subscription_id", dsfDataSourceReadResponse.Data.AssetData.SubscriptionID)
	// unmask
	d.Set("used_for", dsfDataSourceReadResponse.Data.AssetData.UsedFor)
	d.Set("virtual_hostname", dsfDataSourceReadResponse.Data.AssetData.VirtualHostname)
	d.Set("virtual_ip", dsfDataSourceReadResponse.Data.AssetData.VirtualIp)
	d.Set("xel_directory", dsfDataSourceReadResponse.Data.AssetData.XelDirectory)

	if dsfDataSourceReadResponse.Data.AssetData.AuditInfo != nil {
		auditInfo := &schema.Set{F: resourceAssetDataAuditInfoHash}
		auditInfoMap := map[string]interface{}{}
		auditInfoMap["policy_template_name"] = dsfDataSourceReadResponse.Data.AssetData.AuditInfo.PolicyTemplateName
		auditInfo.Add(auditInfoMap)
		d.Set("audit_info", auditInfo)
	}

	if dsfDataSourceReadResponse.Data.AssetData.AwsProxyConfig != nil {
		awsProxyConfig := &schema.Set{F: resourceAssetDataAWSProxyConfigHash}
		awsProxyConfigMap := map[string]interface{}{}
		awsProxyConfigMap["http"] = dsfDataSourceReadResponse.Data.AssetData.AwsProxyConfig.HTTP
		awsProxyConfigMap["https"] = dsfDataSourceReadResponse.Data.AssetData.AwsProxyConfig.HTTPS
		awsProxyConfig.Add(awsProxyConfigMap)
		d.Set("aws_proxy_config", awsProxyConfig)
	}

	if dsfDataSourceReadResponse.Data.AssetData.ServiceEndpoints != nil {
		serviceEndpoints := &schema.Set{F: resourceAssetDataServiceEndpointsHash}
		serviceEndpointsMap := map[string]interface{}{}
		serviceEndpointsMap["logs"] = dsfDataSourceReadResponse.Data.AssetData.ServiceEndpoints.Logs
		serviceEndpoints.Add(serviceEndpointsMap)
		d.Set("service_endpoints", serviceEndpoints)
	}

	connections := &schema.Set{F: resourceDataSourceConnectionHash}
	for _, v := range dsfDataSourceReadResponse.Data.AssetData.Connections {
		connection := map[string]interface{}{}
		connection["reason"] = v.Reason

		connection["access_id"] = v.ConnectionData.AccessID // TODO SR-4549 - add access_ID
		connection["access_key"] = v.ConnectionData.AccessKey
		connection["account_name"] = v.ConnectionData.AccountName
		connection["api_key"] = v.ConnectionData.ApiKey
		connection["auth_mechanism"] = v.ConnectionData.AuthMechanism
		connection["autocommit"] = v.ConnectionData.Autocommit
		connection["aws_connection_id"] = v.ConnectionData.AwsConnectionID
		connection["bucket"] = v.ConnectionData.Bucket
		connection["ca_certs_path"] = v.ConnectionData.CaCertsPath
		connection["ca_file"] = v.ConnectionData.CaFile
		connection["cache_file"] = v.ConnectionData.CacheFile
		connection["cert_file"] = v.ConnectionData.CaFile
		connection["client_id"] = v.ConnectionData.ClientID
		connection["client_secret"] = v.ConnectionData.ClientSecret
		connection["cluster_id"] = v.ConnectionData.ClusterID              // TODO SR-4549
		connection["cluster_member_id"] = v.ConnectionData.ClusterMemberID // TODO SR-4549
		connection["cluster_name"] = v.ConnectionData.ClusterName          // TODO SR-4549
		connection["content_type"] = v.ConnectionData.ContentType
		connection["crn"] = v.ConnectionData.Crn
		connection["database_name"] = v.ConnectionData.DatabaseName
		connection["db_role"] = v.ConnectionData.DbRole
		connection["dn"] = v.ConnectionData.Dn
		connection["dns_srv"] = v.ConnectionData.DnsSrv
		connection["driver"] = v.ConnectionData.Driver
		connection["dsn"] = v.ConnectionData.Dsn
		connection["external"] = v.ConnectionData.External
		connection["external_id"] = v.ConnectionData.ExternalID
		connection["extra_kinit_parameters"] = v.ConnectionData.ExtraKinitParameters
		connection["hive_server_type"] = v.ConnectionData.HiveServerType
		connection["host_name_mismatch"] = v.ConnectionData.HostNameMismatch
		connection["hosts"] = v.ConnectionData.Hosts
		connection["httppath"] = v.ConnectionData.Httppath
		connection["is_cluster"] = v.ConnectionData.IsCluster
		connection["jdbc_ssl_trust_server_certificate"] = v.ConnectionData.JdbcSslTrustServerCertificate
		connection["jdbc_ssl_trust_store_location"] = v.ConnectionData.JdbcSslTrustStoreLocation
		connection["jdbc_ssl_trust_store_password"] = v.ConnectionData.JdbcSslTrustStorePassword
		connection["kerberos_host_fqdn"] = v.ConnectionData.KerberosHostFqdn
		connection["kerberos_kdc"] = v.ConnectionData.KerberosServiceKdc
		connection["kerberos_retry_count"] = v.ConnectionData.KerberosRetryCount
		connection["kerberos_service_kdc"] = v.ConnectionData.KerberosServiceKdc
		connection["kerberos_service_realm"] = v.ConnectionData.KerberosServiceRealm
		connection["kerberos_spn"] = v.ConnectionData.KerberosSpn
		connection["key_file"] = v.ConnectionData.KeyFile
		connection["keytab_file"] = v.ConnectionData.KeytabFile
		connection["kinit_program_path"] = v.ConnectionData.KinitProgramPath
		connection["net_service_name"] = v.ConnectionData.NetServiceName
		connection["odbc_connection_string"] = v.ConnectionData.OdbcConnectionString
		connection["passphrase"] = v.ConnectionData.Passphrase
		connection["password"] = v.ConnectionData.Password
		connection["principal"] = v.ConnectionData.Principal
		connection["proxy_auto_detect"] = v.ConnectionData.ProxyAutoDetect
		connection["proxy_password"] = v.ConnectionData.ProxyPassword
		connection["proxy_port"] = v.ConnectionData.ProxyPort
		connection["proxy_server"] = v.ConnectionData.ProxyServer
		connection["proxy_ssl_type"] = v.ConnectionData.ProxySslType
		connection["redirect_uri"] = v.ConnectionData.RedirectUri
		connection["region"] = v.ConnectionData.Region
		connection["replica_set"] = v.ConnectionData.ReplicaSet
		connection["resource_id"] = v.ConnectionData.ResourceID
		connection["role_name"] = v.ConnectionData.RoleName // TODO SR-4549
		connection["schema"] = v.ConnectionData.Schema
		connection["sec_before_operating_expired_token"] = v.ConnectionData.SecBeforeOperatingExpiredToken
		connection["secret_key"] = v.ConnectionData.SecretKey
		connection["self_signed_cert"] = v.ConnectionData.SelfSignedCert
		connection["self_signed"] = v.ConnectionData.SelfSigned
		if v.ConnectionData.ServerIp != "" {
			connection["server_ip"] = v.ConnectionData.ServerIp
		}
		connection["server_port"] = v.ConnectionData.ServerPort
		connection["service_key"] = v.ConnectionData.ServiceKey
		connection["session_token"] = v.ConnectionData.SessionToken
		connection["sid"] = v.ConnectionData.Sid
		connection["snowflake_role"] = v.ConnectionData.SnowflakeRole
		connection["ssl_server_cert"] = v.ConnectionData.SslServerCert
		connection["ssl"] = v.ConnectionData.Ssl
		if v.ConnectionData.SubscriptionID != "" {
			connection["subscription_id"] = v.ConnectionData.SubscriptionID
		}
		connection["tenant_id"] = v.ConnectionData.TenantID
		connection["thrift_transport"] = v.ConnectionData.ThriftTransport
		connection["tmp_user"] = v.ConnectionData.TmpUser
		connection["token_endpoint"] = v.ConnectionData.TokenEndpoint
		connection["token"] = v.ConnectionData.Token
		connection["transportmode"] = v.ConnectionData.Transportmode
		connection["use_keytab"] = v.ConnectionData.UseKeytab
		connection["username"] = v.ConnectionData.Username
		connection["virtual_hostname"] = v.ConnectionData.VirtualHostname // TODO SR-4549
		connection["virtual_ip"] = v.ConnectionData.VirtualIp             // TODO SR-4549
		connection["wallet_dir"] = v.ConnectionData.WalletDir
		connection["warehouse"] = v.ConnectionData.Warehouse

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

		if v.ConnectionData.OauthParameters != nil {
			oauthParameters := &schema.Set{F: resourceConnectionDataOauthParametersHash}
			oauthParametersMap := map[string]interface{}{}
			oauthParametersMap["parameter"] = v.ConnectionData.OauthParameters.Parameter
			oauthParameters.Add(oauthParametersMap)
			connection["oauth_parameters"] = oauthParameters
		}
		connections.Add(connection)
	}
	d.Set("asset_connection", connections)

	log.Printf("[INFO] Finished reading DSF data source with dsfDataSourceId: %s\n", dsfDataSourceId)

	return nil
}

func resourceDSFDataSourceUpdateContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Client)

	// check provided fields against schema
	dsfDataSourceId := d.Id()
	if isOk, err := checkResourceRequiredFields(requiredDataSourceFieldsJson, ignoreDataSourceParamsByServerType, d); !isOk {
		return diag.FromErr(err)
	}

	// convert provided fields into API payload
	dsfDataSource := ResourceWrapper{}
	serverType := d.Get("server_type").(string)
	createResource(&dsfDataSource, serverType, d)

	// auditPullEnabled set to current value from state
	auditPullEnabled, _ := d.GetChange("audit_pull_enabled")
	dsfDataSource.Data.AssetData.AuditPullEnabled = auditPullEnabled.(bool)

	// get asset_id
	assetId := d.Get("asset_id").(string)

	// wait for remoteSyncState
	err := waitForRemoteSyncState(ctx, dsfDataSourceResourceType, assetId, m)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  fmt.Sprintf("Error while waiting for remoteSyncState = \"SYNCED\" for asset: %s", assetId),
			Detail:   fmt.Sprintf("Error: %s\n", err),
		})
	}

	// update resource
	log.Printf("[INFO] Updating DSF data source for serverType: %s and gatewayId: %s assetId: %s\n", dsfDataSource.Data.ServerType, dsfDataSource.Data.GatewayID, dsfDataSource.Data.AssetData.AssetID)
	_, err = client.UpdateDSFDataSource(dsfDataSourceId, dsfDataSource)
	if err != nil {
		log.Printf("[ERROR] Updating data source for serverType: %s and gatewayId: %s assetId: %s | err:%s\n", dsfDataSource.Data.ServerType, dsfDataSource.Data.GatewayID, dsfDataSource.Data.AssetData.AssetID, err)
		return diag.FromErr(err)
	}

	// Connect/disconnect asset to gateway
	err = connectDisconnectGateway(ctx, d, dsfDataSourceResourceType, m)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  fmt.Sprintf("Error while updating audit state for asset: %s", d.Get("asset_id")),
			Detail:   fmt.Sprintf("Error: %s\n", err),
		})
	}

	// Set ID
	d.SetId(dsfDataSourceId)

	// Set the rest of the state from the resource read
	log.Printf("[DEBUG] Writing data source asset details to state")
	resourceDSFDataSourceReadContext(ctx, d, m)

	return diags
}

func resourceDSFDataSourceDeleteContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	dsfDataSourceId := d.Id()

	_, err := client.DeleteDSFDataSource(dsfDataSourceId)
	// if an error is returned, assume it has already been deleted
	if err != nil {
		log.Printf("[INFO] DSF data source has already been deleted with dsfDataSourceId: %s | err: %s\n", dsfDataSourceId, err)
	}
	return nil
}

func resourceDataSourceConnectionHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})

	//assetSchema := getSchema()
	//for _, field := range assetSchema.Connections {
	//	log.Printf("[DEBUG] Checking for field to ignore: '%v'\n", field.ID)
	//	if _, found := ignoreConnectionFields[field.ID]; !found {
	//		if v, ok := m[field.ID]; ok {
	//			switch value := reflect.TypeOf(v); value.Kind() {
	//			case reflect.Int:
	//				log.Printf("[DEBUG] resourceDataSourceConnectionHash Field.ID %v reflect.Int", field.ID)
	//				buf.WriteString(fmt.Sprintf("%v-", v.(int)))
	//			case reflect.Float64:
	//				log.Printf("[DEBUG] resourceDataSourceConnectionHash Field.ID %v reflect.Float64", field.ID)
	//				buf.WriteString(fmt.Sprintf("%v-", v.(float64)))
	//			case reflect.String:
	//				log.Printf("[DEBUG] resourceDataSourceConnectionHash Field.ID %v reflect.String", field.ID)
	//				buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	//			case reflect.Bool:
	//				log.Printf("[DEBUG] resourceDataSourceConnectionHash Field.ID %v reflect.Bool", field.ID)
	//				buf.WriteString(fmt.Sprintf("%v-", v.(bool)))
	//			//case reflect.Slice:
	//			//	// Handle slices or arrays here
	//			//case reflect.Map:
	//			//	// Handle maps here
	//			default:
	//				log.Printf("[DEBUG] resourceDataSourceConnectionHash Unknown type for field.ID:%v", field.ID)
	//			}
	//		}
	//	} else {
	//		log.Printf("[DEBUG] resourceDataSourceConnectionHash Ignoring field: '%s'\n", field.ID)
	//	}
	//}

	if v, ok := m["reason"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["access_id"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["access_key"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["account_name"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["api_key"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	//if v, ok := m["application_id"]; ok {
	//	buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	//}

	if v, ok := m["asset_source"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["auth_mechanism"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["autocommit"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(bool)))
	}

	if v, ok := m["aws__id"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["bucket"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["bucket_account_id"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["ca_certs_path"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["ca_file"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["cache_file"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["cert_file"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["client_id"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["client_secret"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["cluster_id"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["cluster_member_id"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["cluster_name"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["content_type"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["crn"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["database_name"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["db_role"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["dn"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["dns_srv"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(bool)))
	}

	if v, ok := m["driver"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["dsn"]; ok {
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

	if v, ok := m["hive_server_type"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["host_name_mismatch"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(bool)))
	}

	if v, ok := m["hosts"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["httppath"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["is_cluster"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(bool)))
	}

	if v, ok := m["jdbc_ssl_trust_server_certificate"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(bool)))
	}

	if v, ok := m["jdbc_ssl_trust_store_location"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["jdbc_ssl_trust_store_password"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["kerberos_host_fqdn"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["kerberos_kdc"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["kerberos_retry_count"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(int)))
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

	if v, ok := m["key_file"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["keytab_file"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["kinit_program_path"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["net_service_name"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if _, ok := m["oauth_parameters"]; ok {
		log.Printf("[DEBUG] m[\"oauth_parameters\"] %v", m["oauth_parameters"])
		//buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["odbc__string"]; ok {
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

	if v, ok := m["proxy_auto_detect"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["proxy_password"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["proxy_port"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["proxy_server"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["proxy_ssl_type"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["redirect_uri"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["region"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["replica_set"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["resource_id"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["role_name"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["schema"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["sec_before_operating_expired_token"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(int)))
	}

	if v, ok := m["secret_key"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["self_signed_cert"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(bool)))
	}

	if v, ok := m["self_signed"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(bool)))
	}

	if v, ok := m["server_ip"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["server_port"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(int)))
	}

	if v, ok := m["service_key"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["session_token"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["sid"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["snowflake_role"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["ssl_server_cert"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["ssl"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(bool)))
	}

	if v, ok := m["subscription_id"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["tenant_id"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["thrift_transport"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(int)))
	}

	if v, ok := m["tmp_user"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(bool)))
	}

	if v, ok := m["token_endpoint"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["token"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["transportmode"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["use_keytab"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(bool)))
	}

	if v, ok := m["username"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["virtual_hostname"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["virtual_ip"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["wallet_dir"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}

	if v, ok := m["warehouse"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(string)))
	}
	return PositiveHash(buf.String())
}
