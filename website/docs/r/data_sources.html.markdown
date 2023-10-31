---
layout: "dsfhub"
page_title: "dsfhub_data_source Resource"
sidebar_current: "docs-dsfhub-resource-dsfhub-data-source"
description: |-
Provides a dsfhub_data_source resource.  
---

# dsfhub_data_source (Resource)

Provides a data source resource.

The `dsf_data_source` resource contains the configuration parameters necessary for controlling and managing data via the data analytics and 
data source management tools, as well as the asset_connection information necessary for DSF Hub to connect to a particular asset.
Documentation for the underlying API used in this resource can be found at 
[Data Sources API Definition page](https://docs.imperva.com/bundle/v4.13-sonar-user-guide/page/84552.htm).

## Example Usage

```hcl
resource "dsfhub_data_source" "aws_rds_mysql_password" {
	server_type = "AWS RDS MYSQL"
	admin_email = var.admin_email	# The email address to notify about this asset
	arn = var.arn	# Amazon Resource Name - format is arn:partition:service:region:account-id:resource-type:resource-id and used as the asset_id
	asset_display_name = var.asset_display_name	# User-friendly name of the asset, defined by user.
	asset_id = var.asset_id	# Asset ID
	gateway_id = var.gateway_id	# Gateway ID
	server_host_name = var.server_host_name	# Hostname (or IP if name is unknown)
	asset_connection {
		auth_mechanism = "password"
		password = null # password description: "The password of the user being used to authenticate"
		reason = null # Example Values: "default", "sonargateway", "SDM", "audit management", "ad-hoc-query" # reason description: "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset"
		username = null # username description: ""
	}
}
```

## Data Source Types:
<ul>
    <li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/aerospike.tf">Aerospike</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/alibaba_apsara_mongodb.tf">Alibaba ApsaraDB for MongoDB</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/alibaba_apsara_rds_mysql.tf">Alibaba ApsaraDB RDS for MySQL</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/alibaba_apsara_rds_postgresql.tf">Alibaba ApsaraDB RDS for PostgreSQL</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/alibaba_oss.tf">Alibaba Object Storage Service</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/ambari.tf">Apache Ambari</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/aws_athena.tf">Amazon Athena</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/aws_documentdb_cluster.tf">Amazon DocumentDB Cluster</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/aws_documentdb.tf">Amazon DocumentDB</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/aws_dynamodb.tf">Amazon DynamoDB</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/aws_glue.tf">AWS Glue</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/aws_lake_formation.tf">AWS Lake Formation</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/aws_neptune_cluster.tf">AWS Neptune Cluster</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/aws_neptune.tf">AWS Neptune</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/aws_rds_aurora_mysql_cluster.tf">Amazon Aurora MySQL Cluster</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/aws_rds_aurora_mysql.tf">Amazon Aurora MySQL</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/aws_rds_aurora_postgresql_cluster.tf">Amazon Aurora PostgreSQL Cluster</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/aws_rds_aurora_postgresql.tf">Amazon Aurora PostgreSQL</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/aws_rds_mariadb.tf">Amazon RDS for MariaDB</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/aws_rds_ms_sql_server.tf">Amazon RDS for SQL Server</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/aws_rds_mysql.tf">Amazon RDS for MySQL</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/aws_rds_oracle.tf">Amazon RDS for Oracle</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/aws_rds_postgresql.tf">Amazon RDS for PostgreSQL</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/aws_redshift.tf">Amazon Redshift</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/aws_s3.tf">Amazon S3</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/azure_cosmosdb_mongo.tf">Azure Cosmos DB API for MongoDB</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/azure_cosmosdb_table.tf">Azure Cosmos DB Table API</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/azure_cosmosdb.tf">Azure Cosmos DB SQL API</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/azure_mariadb.tf">Azure Database for MariaDB</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/azure_ms_sql_server.tf">Azure SQL Server</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/azure_mysql.tf">Azure Database for MySQL</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/azure_postgresql.tf">Azure Database for PostgreSQL</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/azure_sql_managed_instance.tf">Azure SQL Managed Instance</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/azure_storage_account.tf">Azure Storage Account</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/cassandra.tf">Apache Cassandra</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/cloudant_local.tf">Cloudant Local</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/cloudant.tf">IBM Cloudant</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/cloudera.tf">Cloudera Data Platform</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/cockroachdb.tf">CockroachDB</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/couchbase.tf">Couchbase</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/datastax.tf">Datastax Enterprise </a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/db2.tf">IBM Db2</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/edb_postgresql.tf">EDB Postgres Advanced Server</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/elasticsearch.tf">Elasticsearch</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/eloquence.tf">Eloquence</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/emr.tf">EMR</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/gcp_bigquery.tf">Google Cloud BigQuery</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/gcp_bigtable.tf">Google Cloud Bigtable</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/gcp_ms_sql_server.tf">Google Cloud SQL for SQL Server</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/gcp_mysql.tf">Google Cloud SQL for MySQL</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/gcp_postgresql.tf">Google Cloud SQL for PostgreSQL</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/gcp_spanner.tf">Google Cloud Spanner</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/hbase.tf">Apache HBase</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/hdfs.tf">Hadoop Distributed File System</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/hive.tf">Apache Hive</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/hortonworks.tf">Hortonworks</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/impala.tf">Apache Impala</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/informix.tf">IBM Informix</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/iris.tf">InterSystems IRIS Data Platform</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/kinetica.tf">Kinetica</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/knox_gateway.tf">Apache Knox</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/mariadb.tf">MariaDB Server</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/marklogic.tf">MarkLogic Server</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/mongodb_atlas.tf">MongoDB Atlas</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/mongodb.tf">MongoDB Enterprise Server</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/ms_sql_server.tf">MsSQL Server</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/mysql.tf">MySQL</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/neo4j.tf">Neo4j</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/netezza.tf">Netezza Performance Server</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/oracle.tf">Oracle</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/percona_mongodb.tf">Percona Server for MongoDB</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/percona_mysql.tf">Percona Server for MySQL</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/postgresql.tf">PostgreSQL</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/progress_openedge.tf">Progress OpenEdge RDBMS</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/redis.tf">Redis</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/sap_hana.tf">SAP HANA</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/scylladb.tf">ScyllaDB</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/snowflake.tf">Snowflake for Data Warehouse</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/splunk.tf">Splunk</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/sybase.tf">Sybase</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/teradata.tf">Teradata Vantage</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/yarn.tf">Apache Hadoop YARN</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/yugabyte_cql.tf">Yugabyte CQL</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/yugabyte_sql.tf">Yugabyte SQL</a></li>
    <li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/data_sources/alibaba_max_compute.tf">Alibaba MaxCompute</a></li>
</ul>


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

## Import

DSF Data Source can be imported using the `asset_id`, e.g.:

```
$ terraform import dsf_data_source.example_aws_rds_mysql "arn:partition:service:region:account-id"
```
