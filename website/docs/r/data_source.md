---
subcategory: ""
layout: "dsfhub"
page_title: "DSFHUB Data Source - Resource"
description: |-
  Provides a dsfhub_data_source terraform resource.
---

# Resource: dsfhub_data_source

Terraform resource for managing a DSFHub data source. 

The `dsf_data_source` resource supports the configuration parameters necessary to onboard a data source to DSF Hub, which allows for the management of its audit data via data analytics and management tools. Documentation for the underlying API used in this resource can be found at [Onboarding and Managing Data Sources](https://docs-cybersec.thalesgroup.com/bundle/v15.0-sonar-user-guide/page/84558.htm).

## Data Source Types

The following values for `server_type` are supported:

* AEROSPIKE
* ALIBABA APSARA MONGODB
* ALIBABA APSARA RDS MYSQL
* ALIBABA APSARA RDS POSTGRESQL
* ALIBABA MAX COMPUTE
* ALIBABA OSS
* AMBARI
* AWS ATHENA                     
* AWS DOCUMENTDB CLUSTER         
* AWS DOCUMENTDB                 
* AWS DYNAMODB                   
* AWS GLUE                       
* AWS LAKE FORMATION             
* AWS NEPTUNE CLUSTER            
* AWS NEPTUNE                    
* AWS OPENSEARCH                 
* AWS RDS AURORA MYSQL CLUSTER   
* AWS RDS AURORA MYSQL           
* AWS RDS AURORA POSTGRESQL CLUSTER
* AWS RDS AURORA POSTGRESQL  
* AWS RDS DB2   
* AWS RDS MARIADB                
* AWS RDS MS SQL SERVER          
* AWS RDS MYSQL                  
* AWS RDS ORACLE                 
* AWS RDS POSTGRESQL             
* AWS RDS POSTGRESQL CLUSTER     
* AWS REDSHIFT                   
* AWS REDSHIFT SERVERLESS        
* AWS S3
* AZURE COSMOSDB MONGO   
* AZURE COSMOSDB TABLE   
* AZURE COSMOSDB          
* AZURE DATABRICKS WORKSPACE
* AZURE DATA EXPLORER  
* AZURE MARIADB      
* AZURE MS SQL SERVER     
* AZURE MYSQL           
* AZURE MYSQL FLEXIBLE   
* AZURE POSTGRESQL 
* AZURE POSTGRESQL FLEXIBLE
* AZURE SQL MANAGED INSTANCE
* AZURE STORAGE ACCOUNT
* CASSANDRA
* CLICKHOUSE
* CLOUDANT
* CLOUDANT LOCAL
* CLOUDERA
* COCKROACHDB
* COUCHBASE
* DATASTAX
* DB2
* DRUID CLUSTER
* DRUID
* EDB POSTGRESQL
* ELASTICSEARCH
* ELOQUENCE
* EMR
* GAUSSDB
* GCP ALLOYDB POSTGRESQL
* GCP ALLOYDB POSTGRESQL CLUSTER
* GCP BIGQUERY
* GCP BIGTABLE
* GCP FIRESTORE
* GCP MS SQL SERVER
* GCP MYSQL
* GCP POSTGRESQL
* GCP SPANNER
* GEMFIRE
* GRAINITE
* GRIDGAIN IGNITE
* HBASE
* HDFS
* HIVE
* IMPALA
* INFORMIX
* IRIS
* MARIADB
* MAPR FS
* MAPR HBASE
* MARKLOGIC
* MONGODB
* MONGODB ATLAS
* MS SQL SERVER
* MYSQL
* NEO4J
* NETEZZA
* ORACLE
* PERCONA MONGODB
* PERCONA MYSQL
* POSTGRESQL
* PROGRESS OPENEDGE
* REDIS
* SAP HANA
* SAP IQ
* SCYLLADB
* SINGLESTORE
* SNOWFLAKE
* SPLUNK
* SYBASE
* TERADATA
* TIGERGRAPH
* VERTICA
* YUGABYTE CQL
* YUGABYTE SQL

## Example Usage

For integrated examples with cloud resource configuration and audit setup, please see the [DSFHub Agentless-Onboarding modules](https://github.com/imperva/terraform-dsfhub-agentless-onboarding).

## Argument Reference

The following arguments are required by all Data Source server types:

- `admin_email` - (String) The email address to notify about this asset
- `asset_display_name` - (String) User-friendly name of the asset, defined by user.
- `asset_id` - (String) The unique identifier or resource name of the asset. For most assets this should be a concatenation of Server Host Name + Server Type + Service Name + Server Port with “:” (colon) as separator, example: `mydbhost:MYSQL:my-db-service-name:3306`. For Cloud data sources, this value will be the resource name (e.g. AWS ARN) or resource ID.
- `gateway_id` - (String) The unique identifier of the Agentless Gateway that will own the asset. Example: "12345-abcde-12345-abcde-12345-abcde". You can find the value by connecting to SonarW and running 
```
db.getSiblingDB("lmrm__sonarg").asset.find(
  { "Server Type": "IMPERVA AGENTLESS GATEWAY", "Server Host Name": "your-hostname" },
  { jsonar_uid: 1, _id: 0 }
)
```
- `server_type` - (String) The type of server or data service to be created as a data source. See available list [above](#data-source-types) or at [Onboarding Databases to DSF Hub - Overview](https://docs-cybersec.thalesgroup.com/bundle/onboarding-databases-to-sonar-reference-guide/page/Onboarding-Databases-to-DSF-Hub---Overview_21077247.html).

The following arguments are optional, however some are only supported for certain server types. Please see the [asset specifications](https://docs-cybersec.thalesgroup.com/bundle/onboarding-databases-to-sonar-reference-guide/page/Asset-Specifications_35815461.html) for more details:

- `application` - (String) The Asset ID of the application asset that "owns" the asset.
- `asset_connection` - (Block) An `asset_connection` block as defined below.
- `asset_source` - (String) The source platform/vendor/system of the asset data. Usually the service responsible for creating that asset document
- `asset_version` - (Number) Denotes the database/service version of the asset
- `audit_info` - (Block) An `audit_info` block as defined below. Normally auto-populated when enabling the audit policy, it is a sub-document in JSON format containing configuration information for audit management. See documentation for values that can be added manually depending on asset type. Editing this value does NOT enable the audit policy.
- `audit_pull_enabled` - (Boolean) If true, sonargateway will collect the audit logs for the associated data source if it can, on the successful Connect Gateway playbook run. 
- `audit_type` - (String) Used to indicate what mechanism should be used to fetch logs on systems supporting multiple ways to get logs, see asset specific documentation for details
- `availability_zones` - (List of string) List of regions where the cluster is available from AWS data.
- `available_bucket_account_ids` - (List of string) A list of AWS Account IDs to use when pulling account specific audit logs from this bucket. eg: ['123456789012', ‘123456789013’]
- `available_regions` - (List of string) A list of regions to use in discovery actions that iterate through region
- `aws_proxy_config` - (Block) An `aws_proxy_config` block as defined below for an AWS proxy configuration.
- `bucket_account_id` - (String) AWS account number in the prefix of the files we are pulling. E.g. "123456789012" out of "AWSLogs/123456789012/service/us-east-1/2022/03/25/my_file.gz"
- `ca_certs_path` - (String) Certificate authority certificates path; what location should the sysetm look for certificate information from. Equivalent to --capath in a curl call
- `ca_file` - (String) Path to a certificate authority file to use with the call. Equivalent to --cacert in a curl call
- `cluster_engine` - (String) Cluster engine
- `cluster_id` - (String) Cluster identifier
- `cluster_member_id` - (String) The unique_id of the instance within the cluster
- `cluster_name` - (String) Cluster name
- `content_type` - (String) Content type should be set to the desired 'parent' asset server_type, which is the one that uses this asset as a destination for logs. NOTE: The content_type field will take precedence on the lookup for parent_asset_id field when checking which server is sending logs to this asset.
- `credentials_endpoint` - (String) A specific sts endpoint to use
- `criticality` - (Number) The asset's importance to the business. These values are measured on a scale from "Most critical" (1) to "Least critical" (4). Allowed values: 1, 2, 3, 4
- `database_name` - (String) Specifies the name of the database (or default DB) to connect to.
- `db_engine` - (String) Specifies the version of the engine being used by the database (e.g. oracle-ee, oracle-se, oracle-se1, oracle-se2)
- `db_instances_display_name` - (List of string) List of DB Cluster Members (instances)
- `duration_threshold` - (Number) How long (in milliseconds) a query's execution may take before it is flagged as slow, and output to the sonargd.slow_query collection.
- `enable_audit_management` - (Boolean) If true, Sonar is responsible for setting and updating the policies
- `enable_audit_monitoring` - (Boolean) If true, Sonar sends emails/alerts when the audit policies change.
- `enabled_logs_exports` - (String) List of Enabled Cloudwatch Logs Exports from AWS data
- `entitlement_enabled` - (Boolean) If true, Entitlement Management system is enabled.
- `gateway_service` - (String) The name of the gateway pull service (if any) used to retrieve logs for this source. Usually set by the connect gateway playbook.
- `host_timezone_offset` - (String) The offset value string is in the format "-/+hh:mm"
- `ignore_latest_of` - (String) A regex defining a group. From all the files with the same group, the latest one will be ignored, so that it isn't archived until server is done writing
- `is_cluster` - (Boolean) Indicates whether the asset is part of a cluster.
- `is_multi_zones` - (Boolean) True if the cluster is in multiple zones, False otherwise
- `jsonar_uid` - (String) Unique identifier (UID) attached to the Agentless Gateway controlling the asset
- `jsonar_uid_display_name` - (String) Unique identifier (UID) attached to the Agentless Gateway controlling the asset
- `location` - (String) Current human-readable description of the physical location of the asset, or region.
- `log_bucket_id` - (String) Asset ID of the S3 bucket which stores the logs for this server
- `logs_destination_asset_id` - (String) The asset name of the log aggregator that stores this asset's logs.
- `managed_by` - (String) Email of the person who maintains the asset; can be different from the owner specified in the owned_by field. Defaults to admin_email.
- `marker_alias` - (String) Cluster or System name for a DR pair or similar system where all nodes share a single log. All machines sharing a marker alias will use the same marker. This means that the log will be pulled once rather than once per machine.
- `max_concurrent_conn` - (String) Maximum number of concurrent connections that sensitive data management should use at once.
- `owned_by` - (String) Email of Owner / person responsible for the asset; can be different from the person in the managed_by field. Defaults to admin_email.
- `parent_asset_id` - (String) The name of an asset that this asset is part of (or related to). E.g. an AWS resource will generally have an AWS account asset as its parent. Also used to connect some log aggregating asset with the sources of their logs. E.g. An AWS LOG GROUP asset can have an AWS RDS data source as its parent, indicating that that is the log group for that RDS instance.
- `provider_url` - (String) URL for provider hosting the asset
- `proxy` - (String) Proxy to use for AWS calls. If aws_proxy_config is populated, the proxy field will get populated from the http value there.
- `pubsub_subscription` - (String) Pub/Sub subscription, e.g. "projects/my-project-name/subscriptions/my-subscription-name"
- `region` - (String) For cloud systems with regions, the default region or region used with this asset
- `resource_id` - (String) AWS Resource ID that the RDS Db2 audit logs will be stored under on S3. E.g. db-3TBJU4Y34IAVE2DQRQUWYOEX3I
- `sdm_enabled` - (Boolean) Sensitive data management (SDM) is enabled if this parameter is set to True.
- `searches` - (List of string) A list of searches
- `server_host_name` - (String) Hostname (or IP if name is unknown)
- `server_ip` - (String) IP address of the service where this asset is located. If no IP is available populate this field with other information that would identify the system e.g. hostname or AWS ARN, etc.
- `server_port` - (String) Port used by the source server
- `service_endpoint` - (String) Specify a particular endpoint for a given service
- `service_endpoints` - (Block) A `service_endpoints` block as defined below that specifies particular endpoints for a given service in the form of `<service name>: "endpoint"`.
- `service_name` - (String) Service name
- `subscription_id` - (String) This is the Azure account subscription ID. You can find this number under the Subscriptions page on the Azure portal
- `used_for` - (String) Designates how this asset is used / the environment that the asset is supporting.
- `virtual_hostname` - (String) Hostname of the endpoint of the cluster
- `virtual_ip` - (String) IP of the endpoint of the cluster
- `xel_directory` - (String) Absolute path of the XEL files location

### audit_info

The following argument is optional:

- `policy_template_name` - (String) Policy template name

### aws_proxy_config

The following arguments are optional:

- `http` - (String) HTTP endpoint for AWS proxy config
- `https` - (String) HTTPS endpoint for AWS proxy config

### service_endpoints

The following argument is optional:

- `logs` - (String) The log endpoint for a given service

### asset_connection

The following arguments are required:

- `auth_mechanism` - (String) Specifies the auth mechanism used by the connection
- `reason` - (String) Used to differentiate between connections belonging to the same asset. Use "default" or "sonargateway" for connections necessary for audit pull.

The following arguments are optional, however some are only supported for certain server types and authentication mechanism combinations. Please see the [asset specifications](https://docs-cybersec.thalesgroup.com/bundle/onboarding-databases-to-sonar-reference-guide/page/Asset-Specifications_35815461.html) for more details:

- `access_id` - (String) The Account Name/Access ID to use for authentication
- `access_key` - (String) The Secret access key used for authentication
- `account_name` - (String) The cloudant account name required when connecting a resource with IAM role
- `amazon_secret` - An `amazon_secret` block as defined below, to integrate the asset with AWS Secrets Manager.
- `api_key` - (String) IAM authentication API key
- `autocommit` - (Boolean) If true, Commit automatically don't wait for transaction to be explicitly committed
- `aws_connection_id` - (String) The parent AWS connection document_id
- `bucket` - (String)
- `ca_certs_path` - (String) Certificate authority certificates path; what location should the sysetm look for certificate information from. Equivalent to --capath in a curl call
- `ca_file` - (String) Use the specified certificate file to verify the peer. The file may contain multiple CA certificates.
- `cache_file` - (String) Holds Kerberos protocol credentials (for example, tickets, session keys and other identifying information).
- `cert_file` - (String) Use the specified client certificate file when getting a file with HTTPS, FTPS or another SSL-based protocol.
- `client_id` - (String) Azure client application ID
- `client_secret` - (String) Azure application client secret
- `cluster_id` - (String) Cluster identifier
- `cluster_member_id` - (String) The unique_id of the instance within the cluster
- `cluster_name` - (String) Cluster name
- `content_type` - (String) Content-Type to append to the HTTP headers in the curl call
- `crn` - (String) The CRN unique identifier of the resource
- `cyberark_secret` - A `cyberark_secret` block as defined below, to integrate the asset with CyberArk.
- `database_name` - (String) Specifies the name of the database (or default DB) to connect to.
- `db_role` - (String) The database role to use when connecting to this asset
- `dn` - (String) The distinguished name of a particular PKI certificate
- `dns_srv` - (Boolean)
- `driver` - (String) A path to a non-default driver location. If populated this driver will be used rather than the default
- `dsn` - (String) Data Source Name
- `external` - (Boolean)
- `external_id` - (String) External ID to use when assuming a role
- `extra_kinit_parameters` - (String)
- `hashicorp_secret` - A `hashicorp_secret` block as defined below, to integrate the asset with a HashiCorp Vault.
- `hive_server_type` - (String)
- `host_name_mismatch` - (Boolean)
- `hosts` - (String) Required for quering the logdna url. cloudantnosqldb in the case of a cloudant DB
- `httppath` - (String) Defaults to 'cliservice'
- `is_cluster` - (Boolean) Indicates whether the asset is part of a cluster.
- `jdbc_ssl_trust_server_certificate` - (Boolean) Whether to use a JDBC SSL Trust Server Certificate
- `jdbc_ssl_trust_store_location` - (String) JDBC SSL Trust Store Location
- `jdbc_ssl_trust_store_password` - (String) JDBC SSL Trust Store Password
- `kerberos_host_fqdn` - (String)
- `kerberos_kdc` - (String) The host name or IP Address of your Kerberos KDC machine
- `kerberos_retry_count` - (Number)
- `kerberos_service_kdc` - (String) Kerberos Service KDC
- `kerberos_service_realm` - (String) Kerberos Service Realm
- `kerberos_spn` - (String) The service and host of the Sybase Kerberos Principal. This will be the value prior to the '@' symbol of the principal value
- `key_file` - (String) Private key file name. Allows you to provide your private key in this separate file.
- `keytab_file` - (String) Specify a non-default keytab location
- `kinit_program_path` - (String)
- `net_service_name` - (String) Alias in tnsnames.ora replaces hostname, service name, and port in connection string
- `oauth_parameters` (Set of String) Additional parameters to pass when requesting a token
- `odbc_connection_string` - (String) Additional ODBC connection string parameters. This string will get added to the connection string
- `passphrase` - (String) Passphrase for the private key.
- `password` - (String) The password of the user being used for authentication
- `principal` - (String) The principal used for authentication
- `proxy_auto_detect` - (String)
- `proxy_password` - (String)
- `proxy_port` - (String)
- `proxy_server` - (String)
- `proxy_ssl_type` - (String)
- `redirect_uri` - (String)
- `region` - (String) The cloud geography/region/zone/data center that the resource resides
- `replica_set` - (String)
- `resource_id` - (String) Azure resource application ID
- `role_name` - (String) Role to use for authentication
- `schema` - (String) Schema name. A schema is a logical grouping of database objects
- `sec_before_operating_expired_token` (Number) How many more seconds should a token be valid for before the connections service will update it before returning a connection to a caller. Defaults to 300 seconds (5 minutes).
- `secret_key` - (String) The Secret access key used for authentication
- `self_signed` - (Boolean) Accept self-signed certificates
- `self_signed_cert` - (Boolean)
- `server_port` - (Number) Port used by the source server
- `service_key` - (String) The service key required in the logdna url query to connect to logdna and pull the logs
- `session_token` - (String) STS token used for session authentication
- `sid` - (String) SID used to connect, e.g. ORCL
- `snowflake_role` - (String) Role with which to log into Snowflake
- `ssl` - (Boolean) If true, use SSL when connecting
- `ssl_server_cert` - (String) Path to server certificate to use during authentication
- `tenant_id` - (String) Azure tenant ID
- `thrift_transport` - (Number) Defaults to 2.
- `tmp_user` - (Boolean) If true create a temporary user
- `token` - (String) Saved token to use for authentication
- `token_endpoint` - (String) URL of endpoint to query when requesting a token
- `transportmode` - (String)
- `use_keytab` - (Boolean) If true, authenticate using a key tab
- `username` - (String) The user to use when connecting
- `virtual_hostname` - (String) Hostname of the endpoint of the cluster
- `virtual_ip` - (String) IP of the endpoint of the cluster
- `wallet_dir` - (String) Path to the Oracle wallet directory
- `warehouse` - (String) The name of the warehouse to connect to

The following secret manager blocks are optional:

- `amazon_secret` - An `amazon_secret` block as defined below, to integrate the asset with AWS Secrets Manager.
- `cyberark_secret` - A `cyberark_secret` block as defined below, to integrate the asset with CyberArk.
- `hashicorp_secret` - A `hashicorp_secret` block as defined below, to integrate the asset with a HashiCorp Vault.

#### AWS Secret Manager: `asset_connection.amazon_secret`

A maximum of one block is supported.

The following arguments are optional:

- `field_mapping`- (Map of string) Field mapping for AWS secret
- `secret_asset_id` - (String) AWS secret manager asset_id
- `secret_name` - (String) AWS secret name

#### CyberArk Secret Manager: `asset_connection.cyberark_secret`

A maximum of one block is supported.

The following arguments are optional:

- `field_mapping` - (Map of string) Field mapping for CyberArk secret
- `secret_asset_id` - (String) CyberArk secret manager asset_id
- `secret_name` - (String) CyberArk secret name

#### HashiCorp Secret Manager: `asset_connection.hashicorp_secret`

The following arguments are optional:

- `field_mapping` (Map of string) Field mapping for HashiCorp secret
- `path` - (String) HashiCorp secret path
- `secret_asset_id` - (String) HashiCorp secret manager asset_id
- `secret_name` - (String) HashiCorp secret name

## Import

In Terraform v1.5.0 and later, use an import block to import DSF Data Sources using the `asset_id`. For example:

```
import {
  to = dsfhub_data_source.example_rds_oracle
  id = "arn:aws:rds:region:123456789012:db:my-rds-oracle"
}
```

Using terraform import, import DSF Data Sources using the `asset_id`. For example:

```
$ terraform import dsf_cloud_account.example "arn:aws:rds:region:123456789012:db:my-rds-oracle"
```

For detailed instructions on onboarding existing cloud resources to DSF using Terraform's import functionality, see [Importing and Onboarding Existing Data Sources with Terraform](https://docs-cybersec.thalesgroup.com/bundle/onboarding-databases-to-sonar-reference-guide/page/Importing-and-Onboarding-Existing-Data-Sources-with-Terraform_784990209.html).