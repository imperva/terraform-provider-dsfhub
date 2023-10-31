# dsfhub_data_source (Resource)

Provides a dsfhub_data_source resource for DATASTAX.

## Example usage:

```hcl
# ### DSF Provider ###
provider "dsfhub" {
	dsfhub_token = var.dsfhub_token # TF_VAR_dsfhub_token env variable
	dsfhub_host = var.dsfhub_host # TF_VAR_dsfhub_host env variable
	#insecure_ssl = false
}

# ### Resource example for DATASTAX ###
resource "dsfhub_data_source" "example_datastax" {
	server_type = "DATASTAX"
	# ### required ### 
	admin_email = var.admin_email	# The email address to notify about this asset
	asset_display_name = var.asset_display_name	# User-friendly name of the asset, defined by user.
	asset_id = var.asset_id	# Asset ID
	gateway_id = var.gateway_id	# Gateway ID
	server_host_name = var.server_host_name	# Hostname (or IP if name is unknown)
	server_ip = var.server_ip	# IP address of the service where this asset is located. If no IP is available populate this field with other information that would identify the system e.g. hostname or AWS ARN, etc.
	server_port = var.server_port	# Port used by the source server

	# ### optional ### 
	# application = var.application	# The Asset ID of the application asset that \"owns\" the asset.
	# asset_source = var.asset_source	# The source platform/vendor/system of the asset data. Usually the service responsible for creating that asset document
	# audit_info = var.audit_info	# Normally auto-populated when enabling the audit policy, it is a sub-document in JSON format containing configuration information for audit management. See documentation for values that can be added manually depending on asset type. Editing this value does NOT enable the audit policy.
	# audit_pull_enabled = var.audit_pull_enabled	# If true, sonargateway will collect the audit logs for this system if it can.
	# audit_type = var.audit_type # Example Values: "RSYSLOG", "TABLE"	# Used to indicate what mechanism should be used to fetch logs on systems supporting multiple ways to get logs, see asset specific documentation for details
	# cluster_id = var.cluster_id	# 
	# cluster_member_id = var.cluster_member_id	# 
	# cluster_name = var.cluster_name	# 
	# criticality = var.criticality # Example Values: "1", "2", "3", "4"	# The asset's importance to the business. These values are measured on a scale from \"Most critical\" (1) to \"Least critical\" (4). Allowed values: 1, 2, 3, 4
	# database_name = var.database_name	# Specifies the name of the database (or default DB) to connect to.
	# db_engine = var.db_engine	# Specifies the version of the engine being used by the database (e.g. oracle-ee, oracle-se, oracle-se1, oracle-se2)
	# enable_audit_management = var.enable_audit_management	# If true, Sonar is responsible for setting and updating the policies
	# enable_audit_monitoring = var.enable_audit_monitoring	# If true, Sonar sends emails/alerts when the audit policies change.
	# entitlement_enabled = var.entitlement_enabled	# If true, Entitlement Management system is enabled.
	# is_cluster = var.is_cluster	# 
	# jsonar_uid = var.jsonar_uid	# Unique identifier (UID) attached to the Sonar machine controlling the asset
	# location = var.location	# Current human-readable description of the physical location of the asset, or region.
	# managed_by = var.managed_by	# Email of the person who maintains the asset; can be different from the owner specified in the owned_by field. Defaults to admin_email.
	# owned_by = var.owned_by	# Email of Owner / person responsible for the asset; can be different from the person in the managed_by field. Defaults to admin_email.
	# region = var.region	# For cloud systems with regions, the default region or region used with this asset
	# used_for = var.used_for # Example Values: "Production", "Test", "Development", "Demonstration", "QA", "Staging", "Training", "Disaster Recovery"	# Designates how this asset is used / the environment that the asset is supporting.
	# version = var.version	# Denotes the version of the asset
	# virtual_hostname = var.virtual_hostname	# 
	# virtual_ip = var.virtual_ip	# 
	asset_connection {
		auth_mechanism = "password"
		# ### required ### 
		password = null # password description: "The password of the user being used to authenticate"
		reason = null # Example Values: "default", "sonargateway", "ad-hoc-query" # reason description: "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset"
		username = null # username description: "The user to use when connecting"
		# ### optional ### 
		# amazon_secret = null # amazon_secret description: "Configuration to integrate with AWS Secrets Manager"
		# cyberark_secret = null # cyberark_secret description: "Configuration to integrate with CyberArk Vault"
		# hashicorp_secret = null # hashicorp_secret description: "Configuration to integrate with HashiCorp Vault"
		# ssl = null # ssl description: "If true, use SSL when connecting"
	}
}
```


## Argument Reference

### Required

- `admin_email` (String) The email address to notify about this asset
- `asset_display_name` (String) User-friendly name of the asset, defined by user.
- `asset_id` (String) The unique identifier or resource name of the asset. For most assets this should be a concatenation of Server Host Name + Server Type + Service Name + Server Port with “:” (colon) as separator, example: `mydbhost:mysql:my-db-service-name:3306`.  For AWS data stores, this value will be the arn. For Azure data stores, the recommended format is `/subscriptions/my-subscription-id/resourceGroups/my-resource-group/`.
- `gateway_id` (String) The jsonarUid unique identifier of the agentless gateway. Example: '7a4af7cf-4292-89d9-46ec-183756ksdjd'
- `server_type` (String) The type of server or data service to be created as a data source. The list of available data sources is documented at: https://docs.imperva.com/bundle/v4.11-sonar-user-guide/page/84552.htm

### Optional

- `ad_provider` (String) The type of AWS RDS instance that the S3 asset is receiving audit logs from
- `asset_connection` (Block Set) N/A (see [below for nested schema](#nestedblock--asset_connection))
- `asset_source` (String) The source platform/vendor/system of the asset data. Usually the service responsible for creating that asset document
- `audit_info` (Block Set) Normally auto-populated when enabling the audit policy, it is a sub-document in JSON format containing configuration information for audit management. See documentation for values that can be added manually depending on asset type. Editing this value does NOT enable the audit policy. (see [below for nested schema](#nestedblock--audit_info))
- `audit_pull_enabled` (Boolean) If true, sonargateway will collect the audit logs for this system if it can.
- `audit_type` (String) Used to indicate what mechanism should be used to fetch logs on systems supporting multiple ways to get logs, see asset specific documentation for details
- `availability_zones` (String)
- `available_regions` (String) A list of regions to use in discovery actions that iterate through region
- `aws_proxy_config` (Block Set) AWS specific proxy configuration (see [below for nested schema](#nestedblock--aws_proxy_config))
- `ca_certs_path` (String) Certificate authority certificates path; what location should the sysetm look for certificate information from. Equivalent to --capath in a curl call
- `ca_file` (String) Path to a certificate authority file to use with the call. Equivalent to --cacert in a curl call
- `cluster_engine` (String)
- `cluster_id` (String)
- `cluster_member_id` (String)
- `cluster_name` (String)
- `content_type` (String) Content type should be set to the desired <'parent' asset 'Server Type'>, which is the one that uses this asset as a destination for logs. NOTE: The content_type field will take precedence on the lookup for parent_asset_id field when checking which server is sending logs to this asset.
- `credentials_endpoint` (String) A specific sts endpoint to use
- `criticality` (Number) The asset's importance to the business. These values are measured on a scale from "Most critical" (1) to "Least critical" (4). Allowed values: 1, 2, 3, 4
- `database_name` (String) Specifies the name of the database (or default DB) to connect to.
- `db_engine` (String) Specifies the version of the engine being used by the database (e.g. oracle-ee, oracle-se, oracle-se1, oracle-se2)
- `db_instances_display_name` (String)
- `duration_threshold` (Number)
- `enable_audit_management` (Boolean) If true, Sonar is responsible for setting and updating the policies
- `enable_audit_monitoring` (Boolean) If true, Sonar sends emails/alerts when the audit policies change.
- `enabled_logs_exports` (String)
- `entitlement_enabled` (Boolean) If true, Entitlement Management system is enabled.
- `gateway_service` (String) The name of the gateway pull service (if any) used to retrieve logs for this source. Usually set by the connect gateway playbook.
- `host_timezone_offset` (String) The offset value string is in the format "-/+hh:mm"
- `ignore_latest_of` (String) A regex defining a group. From all the files with the same group, the latest one will be ignored, so that it isn't archived until server is done writing
- `is_cluster` (Boolean)
- `is_multi_zones` (Boolean)
- `jsonar_uid` (String) Unique identifier (UID) attached to the Sonar machine controlling the asset
- `jsonar_uid_display_name` (String) Unique identifier (UID) attached to the Sonar machine controlling the asset
- `location` (String) Current human-readable description of the physical location of the asset, or region.
- `log_bucket_id` (String) Asset ID of the S3 bucket which stores the logs for this server
- `logs_destination_asset_id` (String) The asset name of the log aggregator that stores this asset's logs.
- `managed_by` (String) Email of the person who maintains the asset; can be different from the owner specified in the owned_by field. Defaults to admin_email.
- `max_concurrent_conn` (String) Maximum number of concurrent connections that sensitive data management should use at once.
- `owned_by` (String) Email of Owner / person responsible for the asset; can be different from the person in the managed_by field. Defaults to admin_email.
- `parent_asset_id` (String) The name of an asset that this asset is part of (/related to). E.g. an AWS resource will generally have an AWS account asset as its parent. Also used to connect some log aggregating asset with the sources of their logs. E.g. An AWS LOG GROUP asset can have an AWS RDS as its parent, indicating that that is the log group for that RDS.
- `provider_url` (String)
- `proxy` (String) Proxy to use for AWS calls if aws_proxy_config is populated the proxy field will get populated from the http value there
- `pubsub_subscription` (String)
- `region` (String) For cloud systems with regions, the default region or region used with this asset
- `sdm_enabled` (Boolean) Sensitive data management (SDM) is enabled if this parameter is set to True.
- `searches` (String)
- `server_host_name` (String) Hostname (or IP if name is unknown)
- `server_ip` (String) IP address of the service where this asset is located. If no IP is available populate this field with other information that would identify the system e.g. hostname or AWS ARN, etc.
- `server_port` (String) Port used by the source server
- `service_endpoint` (String) Specify a particular endpoint for a given service
- `service_endpoints` (Block Set) Specify particular endpoints for a given service in the form of <service name>: "endpoint" (see [below for nested schema](#nestedblock--service_endpoints))
- `service_name` (String)
- `ssl` (Boolean)
- `subscription_id` (String) This is the Azure account subscription ID. You can find this number under the Subscriptions page on the Azure portal
- `used_for` (String) Designates how this asset is used / the environment that the asset is supporting.
- `version` (Number) Denotes the version of the asset
- `virtual_hostname` (String)
- `virtual_ip` (String)
- `xel_directory` (String) Absolute path of the XEL files location

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--asset_connection"></a>
### Nested Schema for `asset_connection`

Required:

- `auth_mechanism` (String) Specifies the auth mechanism used by the connection
- `reason` (String) N/A

Optional:

- `access_id` (String) The Account Name/Access ID to use when authenticating to Snowflake
- `account_name` (String) The cloudant account name required when connecting a resource with IAM role
- `amazon_secret` (Block Set, Max: 1) Configuration to integrate with AWS Secrets Manager (see [below for nested schema](#nestedblock--asset_connection--amazon_secret))
- `api_key` (String) IAM authentication API key
- `autocommit` (Boolean) If true, Commit automatically don't wait for transaction to be explicitly committed
- `aws_connection_id` (String) The parent AWS connection document_id
- `bucket` (String)
- `ca_certs_path` (String) Certificate authority certificates path; what location should the sysetm look for certificate information from. Equivalent to --capath in a curl call
- `ca_file` (String) Use the specified certificate file to verify the peer. The file may contain multiple CA certificates.
- `cache_file` (String)
- `cert_file` (String) Use the specified client certificate file when getting a file with HTTPS, FTPS or another SSL-based protocol.
- `client_id` (String) Azure client application ID
- `client_secret` (String) Azure application client secret
- `cluster_id` (String)
- `cluster_member_id` (String)
- `cluster_name` (String)
- `content_type` (String) Content-Type to append to the HTTP headers in the curl call
- `credential_fields` (Block Set) Document containing values to build a profile from. Filling this will create a profile using the given profile name (see [below for nested schema](#nestedblock--asset_connection--credential_fields))
- `crn` (String) The CRN unique identifier of the resource
- `cyberark_secret` (Block Set, Max: 1) Configuration to integrate with AWS Secrets Manager (see [below for nested schema](#nestedblock--asset_connection--cyberark_secret))
- `database_name` (String) Specifies the name of the database (or default DB) to connect to.
- `db_role` (String)
- `dn` (String) The distinguished name of a particular PKI certificate
- `dns_srv` (Boolean)
- `driver` (String) A path to a non-default driver location. If populated this driver will be used rather than the default
- `dsn` (String) Data Source Name
- `external` (Boolean)
- `external_id` (String) External ID to use when assuming a role
- `extra_kinit_parameters` (String)
- `hashicorp_secret` (Block Set) Configuration to integrate with HashiCorp Vault (see [below for nested schema](#nestedblock--asset_connection--hashicorp_secret))
- `hive_server_type` (String)
- `host_name_mismatch` (Boolean)
- `hosts` (String) Required for quering the logdna url. cloudantnosqldb in the case of a cloudant DB
- `httppath` (String)
- `is_cluster` (Boolean)
- `jdbc_ssl_trust_server_certificate` (Boolean)
- `jdbc_ssl_trust_store_location` (String)
- `jdbc_ssl_trust_store_password` (String)
- `kerberos_host_fqdn` (String)
- `kerberos_kdc` (String)
- `kerberos_retry_count` (Number)
- `kerberos_service_kdc` (String)
- `kerberos_service_realm` (String)
- `kerberos_spn` (String)
- `key_file` (String) Private key file name. Allows you to provide your private key in this separate file.
- `keytab_file` (String) Specify a non-default keytab location
- `kinit_program_path` (String)
- `net_service_name` (String) Alias in tnsnames.ora replaces hostname, service name, and port in connection string
- `oauth_parameters` (Set of String) Additional parameters to pass when requesting a token
- `odbc_connection_string` (String) Additional ODBC connection string parameters. This string will get added to the connection string
- `passphrase` (String) Passphrase for the private key.
- `password` (String) The password of the user being used to authenticate
- `principal` (String) The principal used to authenticate
- `proxy_auto_detect` (String)
- `proxy_password` (String)
- `proxy_port` (String)
- `proxy_server` (String)
- `proxy_ssl_type` (String)
- `redirect_uri` (String)
- `region` (String) The cloud geography/region/zone/data center that the resource resides
- `replica_set` (String)
- `resource_id` (String) Azure resource application ID
- `role_name` (String) Role to use for authentication
- `schema` (String) Schema name. A schema is a logical grouping of database objects
- `secret_key` (String)
- `self_signed` (Boolean) Accept self-signed certificates
- `self_signed_cert` (Boolean)
- `server_port` (Number)
- `service_key` (String) The service key required in the logdna url query to connect to logdna and pull the logs
- `snowflake_role` (String) Role with which to log into Snowflake
- `ssl` (Boolean) If true, use SSL when connecting
- `ssl_server_cert` (String) Path to server certificate to use during authentication
- `tenant_id` (String) Azure tenant ID
- `thrift_transport` (Number)
- `tmp_user` (Boolean) If true create a temporary user
- `token` (String) Saved token to use to authenticate
- `token_endpoint` (String) URL of endpoint to query when requesting a token
- `transportmode` (String)
- `use_keytab` (Boolean) If true, authenticate using a key tab
- `username` (String) The user to use when connecting
- `virtual_hostname` (String)
- `virtual_ip` (String)
- `wallet_dir` (String) Path to the Oracle wallet directory
- `warehouse` (String) The name of the warehouse to connect to

<a id="nestedblock--asset_connection--amazon_secret"></a>
### Nested Schema for `asset_connection.amazon_secret`

Optional:

- `field_mapping` (Map of String) Field mapping for amazon secret
- `secret_asset_id` (String) Amazon secret asset id
- `secret_name` (String) Amazon secret mane


<a id="nestedblock--asset_connection--credential_fields"></a>
### Nested Schema for `asset_connection.credential_fields`

Optional:

- `credential_source` (String) HashiCorp secret asset id
- `role_arn` (String) HashiCorp secret mane


<a id="nestedblock--asset_connection--cyberark_secret"></a>
### Nested Schema for `asset_connection.cyberark_secret`

Optional:

- `field_mapping` (Map of String) Field mapping for amazon secret
- `secret_asset_id` (String) Amazon secret asset id
- `secret_name` (String) Amazon secret mane

<a id="nestedblock--asset_connection--hashicorp_secret"></a>
### Nested Schema for `asset_connection.hashicorp_secret`

Optional:

- `field_mapping` (Map of String) Field mapping for HashiCorp secret
- `path` (String) HashiCorp secret path
- `secret_asset_id` (String) HashiCorp secret asset id
- `secret_name` (String) HashiCorp secret mane

<a id="nestedblock--audit_info"></a>
### Nested Schema for `audit_info`

Optional:

- `policy_template_name` (String) Policy template name

<a id="nestedblock--aws_proxy_config"></a>
### Nested Schema for `aws_proxy_config`

Optional:

- `http` (String) HTTP endpoint for aws proxy config
- `https` (String) HTTPS endpoint for aws proxy config

<a id="nestedblock--service_endpoints"></a>
### Nested Schema for `service_endpoints`

Optional:

- `logs` (String) The log endpoint for a given service