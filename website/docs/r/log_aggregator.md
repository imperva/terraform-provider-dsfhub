---
subcategory: ""
layout: "dsfhub"
page_title: "DSFHUB Log Aggregator - Resource"
description: |-
  Provides a dsfhub_log_aggregator terraform resource.
---

# Resource: dsfhub_log_aggregator

Terraform resource for managing a DSFHub log aggregator. 

A log aggregator is a resource that serves as a log destination for one or more data sources in the DSF Hub.

The `dsfhub_log_aggregator` resource supports the configuration parameters necessary for onboarding log aggregators that store logs for one or more data sources in the DSF Hub. Documentation for the underlying API used in this resource can be found at [Onboarding and Managing Log Aggregators](https://docs-cybersec.thalesgroup.com/bundle/v15.0-sonar-user-guide/page/84559.htm).

## Log Aggregator Types

<ul>
	<li><a href="https://github.com/imperva/terraform-dsfhub-agentless-onboarding/blob/main/modules/dsfhub-alibaba-logstore/README.md">Alibaba Logstore</li>
	<li><a href="https://github.com/imperva/terraform-dsfhub-agentless-onboarding/blob/main/modules/dsfhub-aws-kinesis/README.md">AWS Kinesis</a></li>
	<li><a href="https://github.com/imperva/terraform-dsfhub-agentless-onboarding/blob/main/modules/dsfhub-aws-log-group/README.md">AWS Log Group</a></li>
	<li><a href="https://github.com/imperva/terraform-dsfhub-agentless-onboarding/blob/main/modules/dsfhub-aws-s3-bucket-la/README.md">Amazon S3</a></li>
	<li><a href="https://github.com/imperva/terraform-dsfhub-agentless-onboarding/blob/main/modules/dsfhub-azure-eventhub/README.md">Azure EventHub</a></li>
	<li><a href="https://github.com/imperva/terraform-dsfhub-agentless-onboarding/blob/main/modules/dsfhub-gcp-cloud-storage-bucket/README.md">GCP Cloud Storage Bucket</a></li>
	<li><a href="https://github.com/imperva/terraform-dsfhub-agentless-onboarding/blob/main/modules/dsfhub-gcp-pubsub/README.md">Google Cloud Pub/Sub</a></li>
	<li>SSH</li>
</ul>

## Example Usage

For integrated examples with cloud resource configuration and audit setup, please see the [DSFHub Agentless-Onboarding modules](https://github.com/imperva/terraform-dsfhub-agentless-onboarding).

### Basic Alibaba Log Aggregator

Example Alibaba LogStore Log Aggregator with the `default` authentication mechanism:

```hcl
resource "dsfhub_log_aggregator" "example_alibaba_log_aggregator" {
  server_type        = "ALIBABA LOGSTORE"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name" 
  asset_id           = "arn:acs:logstore:region:123456789012:project/slsaudit-center-123456789012-region:logstore/rds_log"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  parent_asset_id    = "arn:acs:123456789012"
  asset_connection {
    auth_mechanism = "default"
    reason         = "default"
  }
}
```

### Basic AWS Log Aggregator

Example AWS Kinesis Data Stream Log Aggregator with the `default` authentication mechanism. It relies on its parent AWS cloud account asset's authentication mechanism.

```hcl
resource "dsfhub_log_aggregator" "example_aws_log_aggregator" {
  server_type        = "AWS KINESIS"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name"
  asset_id           = "arn:aws:kinesis:region:123456789012:stream/aws-rds-das-cluster-XXXXXXX"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  parent_asset_id    = "arn:partition:service:region:account-id" 
  asset_connection {
    auth_mechanism = "default"
    reason         = "default"
    region         = var.region
  }
}
```

Example AWS Log Group Log Aggregator with the `default` authentication mechanism. It relies on its parent AWS cloud account asset's authentication mechanism.

```hcl
resource "dsfhub_log_aggregator" "example_aws_log_aggregator" {
  server_type        = "AWS LOG GROUP"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name"
  asset_id           = "arn:aws:logs:us-east-2:123456789012:log-group:/aws/rds/instance/my-database/audit:*"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  parent_asset_id    = "arn:partition:service:region:account-id" 
  asset_connection {
    auth_mechanism = "default"
    reason         = "default"
    region         = var.region
  }
}
```

Example AWS S3 Log Aggregator with the `default` authentication mechanism. It relies on its parent AWS cloud account asset's authentication mechanism.

```hcl
resource "dsfhub_log_aggregator" "example_aws_log_aggregator" {
  server_type        = "AWS S3"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name"
  asset_id           = "aarn:aws:s3:::your-bucket-name"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  parent_asset_id    = "arn:partition:service:region:account-id"
  server_host_name   = "your-bucket-name"
  asset_connection {
    auth_mechanism = "default"
    reason         = "default"
  }
}
```

### Basic Azure Log Aggregator

Example Azure Event Hub Log Aggregator with the `azure_ad` authentication mechanism:

```hcl
resource "dsfhub_log_aggregator" "example_azure_log_aggregator" {
  server_type        = "AZURE EVENTHUB"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name"
  asset_id           = "/subscriptions/11111111-2222-3333-4444-123456789012/resourceGroups/your-resource-group/providers/Microsoft.EventHub/namespaces/your-namespace/eventhubs/your-eventhub"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  server_host_name   = "your-namespace.servicebus.windows.net"
  asset_connection {
    auth_mechanism          = "azure_ad"
    reason                  = "default"
    azure_storage_account   = "yourazstorageaccount"
    azure_storage_container = "your-az-storage-container"
    eventhub_name           = "your-eventhub-name"
    eventhub_namespace      = "your-namespace"
    format                  = "Sql"
    # user_identity_client_id = "your-user-assigned-managed-identity"
  }
}
```

Example Azure Event Hub Log Aggregator with the `client_secret` authentication mechanism:

```hcl
resource "dsfhub_log_aggregator" "example_azure_log_aggregator" {
  server_type        = "AZURE EVENTHUB"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name"
  asset_id           = "/subscriptions/11111111-2222-3333-4444-123456789012/resourceGroups/your-resource-group/providers/Microsoft.EventHub/namespaces/your-namespace/eventhubs/your-eventhub"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  server_host_name   = "your-namespace.servicebus.windows.net"
  asset_connection {
    auth_mechanism          = "client_secret"
    reason                  = "default"
    application_id          = "11111111-2222-3333-4444-123456789012"
    azure_storage_account   = "yourazstorageaccount"
    azure_storage_container = "your-az-storage-container"
    client_secret           = "your-secret"
    directory_id            = "11111111-2222-3333-4444-123456789012"
    eventhub_name           = "your-eventhub-name"
    eventhub_namespace      = "your-namespace"
    format                  = "Sql"
    subscription_id         = "11111111-2222-3333-4444-123456789012"
  }
}
```

Example Azure Event Hub Log Aggregator with the `default` authentication mechanism:

```hcl
resource "dsfhub_log_aggregator" "example_azure_log_aggregator" {
  server_type        = "AZURE EVENTHUB"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name"
  asset_id           = "/subscriptions/11111111-2222-3333-4444-123456789012/resourceGroups/your-resource-group/providers/Microsoft.EventHub/namespaces/your-namespace/eventhubs/your-eventhub"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  server_host_name   = "your-namespace.servicebus.windows.net"
  asset_connection {
    auth_mechanism           = "default"
    reason                   = "default"
    azure_storage_account    = "yourazstorageaccount"
    azure_storage_container  = "your-az-storage-container"
    azure_storage_secret_key = "your-az-storage-secret-key"
    eventhub_access_key      = "your-eh-access-key"
    eventhub_access_policy   = "RootManageSharedAccessKey"
    eventhub_name            = "your-eventhub-name"
    eventhub_namespace       = "your-namespace"
    format                   = "Sql"
  }
}
```

### Basic GCP Log Aggregator

Example GCP Cloud Storage Bucket Log Aggregator with the `default` authentication mechanism:

```hcl
resource "dsfhub_log_aggregator" "example_gcp_log_aggregator" {
  server_type        = "GCP CLOUD STORAGE BUCKET"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name"
  asset_id           = "projects:my_bucket"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  parent_asset_id    = "my_service_account@project-name.iam.gserviceaccount.com:project-name"
  server_host_name   = "storage.googleapis.com"
  server_ip          = "storage.googleapis.com"
  server_port        = "443"
  asset_connection {
    auth_mechanism = "default"
    reason         = "default"
  }
}
```

Example GCP Cloud Storage Bucket Log Aggregator with the `service_account` authentication mechanism:

```hcl
resource "dsfhub_log_aggregator" "example_gcp_log_aggregator" {
  server_type        = "GCP CLOUD STORAGE BUCKET"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name"
  asset_id           = "projects:my_bucket"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  parent_asset_id    = "my_service_account@project-name.iam.gserviceaccount.com:project-name"
  server_host_name   = "storage.googleapis.com"
  server_ip          = "storage.googleapis.com"
  server_port        = "443"
  asset_connection {
    auth_mechanism = "service_account"
    reason         = "default"
  }
}
```

Example GCP Pub/Sub Log Aggregator with the `default` authentication mechanism:

```hcl
resource "dsfhub_log_aggregator" "example_gcp_log_aggregator" {
  server_type        = "GCP PUBSUB"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name"
  asset_id           = "projects/my-project/subscriptions/my-subscription"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  parent_asset_id    = "my_service_account@project-name.iam.gserviceaccount.com:project-name"
  pubsub_subscription = "projects/my-project/subscriptions/my-subscription"
  server_host_name   = "pubsub.googleapis.com
  server_ip          = "pubsub.googleapis.com"
  server_port        = "443"
  asset_connection {
    auth_mechanism = "default"
    reason         = "default"
  }
}
```

Example GCP Pub/Sub Log Aggregator with the `service_account` authentication mechanism:

```hcl
resource "dsfhub_log_aggregator" "example_gcp_log_aggregator" {
  server_type        = "GCP PUBSUB"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name"
  asset_id           = "projects/my-project/subscriptions/my-subscription"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  parent_asset_id    = "my_service_account@project-name.iam.gserviceaccount.com:project-name"
  pubsub_subscription = "projects/my-project/subscriptions/my-subscription"
  server_host_name   = "pubsub.googleapis.com"
  server_ip          = "pubsub.googleapis.com"
  server_port        = "443"
  asset_connection {
    auth_mechanism  = "service_account"
    reason          = "default"
    key_file        = "/path/to/gcp_service_file.json"
  }
}
```

### Basic SSH Log Aggregator

Example SSH Log Aggregator with the `default` authentication mechanism:

```hcl
resource "dsfhub_log_aggregator" "example_ssh_log_aggregator" {
  server_type        = "SSH"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name"
  asset_id           = "my.hostname:SSH::22"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  server_host_name   = "my.hostname"
  server_ip          = "my.hostname"
  asset_connection {
    auth_mechanism  = "default"
    reason          = "default"
  }
}
```

Example SSH Log Aggregator with the `kerberos` authentication mechanism:

```hcl
resource "dsfhub_log_aggregator" "example_ssh_log_aggregator" {
  server_type        = "SSH"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name"
  asset_id           = "my.hostname:SSH::22"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  server_host_name   = "my.hostname"
  server_ip          = "my.hostname"
  asset_connection {
    auth_mechanism  = "kerberos"
    reason          = "default"
  }
}
```

## Argument Reference

The following arguments are required by all Log Aggregator server types:

- `admin_email` - (String) The email address to notify about this asset
- `asset_connection` - (Block) An `asset_connection` block as defined below.
- `asset_id` - (String) The unique identifier of the asset.
- `gateway_id` - (String) The unique identifier of the Agentless Gateway that will own the asset. Example: "12345-abcde-12345-abcde-12345-abcde". You can find the value by connecting to SonarW and running 
```
db.getSiblingDB("lmrm__sonarg").asset.find(
  { "Server Type": "IMPERVA AGENTLESS GATEWAY", "Server Host Name": "your-hostname" },
  { jsonar_uid: 1, _id: 0 }
)
```
- `server_type` - (String) The type of cloud platform or service to be created as a log aggregator. The available values are `ALIBABA LOGSTORE`, `AWS KINESIS`, `AWS LOG GROUP`, `AWS S3`, `AZURE EVENTHUB`, `GCP CLOUD STORAGE BUCKET`, `GCP PUBSUB` and `SSH`.

The following arguments are optional, however some are only supported for certain server types. Please see the [asset specifications](https://docs-cybersec.thalesgroup.com/bundle/onboarding-databases-to-sonar-reference-guide/page/Asset-Specifications_35815461.html) for more details:

- `application` - (String) The Asset ID of the application asset that "owns" the asset.
- `asset_display_name` - (String) User-friendly name of the asset, defined by user.
- `asset_source` - (String) The source platform/vendor/system of the asset data. Usually the service responsible for creating that asset document
- `asset_version` - (Number) Denotes the database/service version of the asset
- `audit_data_type` - (String) The type of audit data being collected
- `audit_info` - (Block) An `audit_info` block as defined below. Normally auto-populated when enabling the audit policy, it is a sub-document in JSON format containing configuration information for audit management. See documentation for values that can be added manually depending on asset type. Editing this value does NOT enable the audit policy.
- `audit_pull_enabled` (Boolean) If true, sonargateway will collect the audit logs for the associated data source if it can, on the successful Connect Gateway playbook run. 
- `audit_type` - (String) Used to indicate what mechanism should be used to fetch logs on systems supporting multiple ways to get logs, see asset specific documentation for details. Example: "BIGQUERY","BIGTABLE","BUCKET","MSSQL","MYSQL","POSTGRESQL","SPANNER".
- `available_bucket_account_ids` - (List of string) A list of AWS Account IDs to use when pulling account specific audit logs from this bucket. eg: ['123456789012', ‘123456789013’]
- `available_regions` - (List of string) A list of regions to use in discovery actions that iterate through region
- `aws_proxy_config` - (Block) An `aws_proxy_config` block as defined below for an AWS proxy configuration.
- `bucket_account_id` - (String) AWS Account ID where the bucket resides. e.g: "123456789012" our of "redshift/AWSLogs/123456789012/redshift/us-east-1/2022/03/25/my_file.gz". Mandatory for audit_type `DYNAMODB`, `ORACLE` and `REDSHIFT`. 
- `ca_certs_path` - (String) Certificate authority certificates path; what location should the sysetm look for certificate information from. Equivalent to --capath in a curl call
- `ca_file` - (String) Path to a certificate authority file to use with the call. Equivalent to --cacert in a curl call
- `consumer_group` - (String) The name of the consumer group to use for the pull. Only applies to pull_type: consumer_group. Supported in DSF version 4.19+.
- `consumer_group_workers` - (String) The number of workers. An integer between 1 and 64 (inclusive) or the string "dynamic" which will automatically retrieve the number of shards and set the workers to that number. Defaults to 2. Only applies when `pull_type` is `consumer_group`.
- `consumer_worker_prefix` - (String) The prefix to use for the consumer group workers. By default the name of the consumer group is used. Supported in DSF version 4.19+.
- `content_type` - (String) Content type should be set to the desired <'parent' asset "Server Type">, which is the one that uses this asset as a destination for logs. Note: The content_type field will take precedence on the lookup for parent_asset_id field when checking which server is sending logs to this asset.
- `credentials_endpoint` - (String) A specific sts endpoint to use
- `criticality` - (Number) The asset's importance to the business. These values are measured on a scale from "Most critical" (1) to "Least critical" (4). Allowed values: 1, 2, 3, 4
- `database_name` - (String) Specifies the name of the database (or default DB) to connect to.
- `db_engine` - (String) Specifies the version of the engine being used by the database (e.g. oracle-ee, oracle-se, oracle-se1, oracle-se2)
- `endpoint` - (String) Logstore's endpoint
- `gateway_service` - (String) `gateway-aws@<DB type>.service` Not necessary to be set manually on the asset. Will be set by the Connect Gateway playbook.
- `jsonar_uid` - (String) Unique identifier (UID) attached to the Agentless Gateway controlling the asset
- `location` - (String) Current human-readable description of the physical location of the asset, or region.
- `logstore` - (String) Unit that is used to collect, store and query logs
- `logs_destination_asset_id` - (String) The asset name of the log aggregator that stores this asset's logs.
- `managed_by` - (String) Email of the person who maintains the asset; can be different from the owner specified in the owned_by field. Defaults to admin_email.
- `max_concurrent_conn` - (String) Maximum number of concurrent connections that sensitive data management should use at once.
- `owned_by` - (String) Email of Owner / person responsible for the asset; can be different from the person in the managed_by field. Defaults to admin_email.
- `parent_asset_id` - (String) The name of an asset that this asset is part of (or related to). E.g. an AWS resource will generally have an AWS account asset as its parent. Also used to connect some log aggregating asset with the sources of their logs. E.g. An AWS LOG GROUP asset can have an AWS RDS data source as its parent, indicating that that is the log group for that RDS instance.
- `project` - (String) Project separates different resources of multiple users and control access to specific resources
- `proxy` - (String) Proxy to use for AWS calls if aws_proxy_config is populated the proxy field will get populated from the http value there
- `pubsub_subscription` - (String) Pub/Sub subscription, e.g. "projects/my-project-name/subscriptions/my-subscription-name"
- `pull_type` - (String) The method used to pull data from an Alibaba logstore. Possible values: "log_client", "consumer_group". Defaults to "log_client".
- `region` - (String) For cloud systems with regions, the default region or region used with this asset
- `s3_provider` - (String) The type of AWS RDS instance that the S3 asset is receiving audit logs from. Accepted value: \"aws-rds-mssql\", required only for AWS RDS MS SQL SERVER auditing workflow up to DSF version 4.19.
- `sdm_enabled` - (Boolean) Sensitive data management (SDM) is enabled if this parameter is set to True.
- `server_host_name` - (String) Hostname (or IP if name is unknown)
- `server_ip` - (String) IP address of the service where this asset is located. If no IP is available populate this field with other information that would identify the system e.g. hostname or AWS ARN, etc.
- `server_port` - (String) Port used by the source server, or "443" for services reached over HTTPS.
- `service_endpoints` - (Block) A `service_endpoints` block as defined below that specifies particular endpoints for a given service in the form of `<service name>: "endpoint"`.
- `used_for` - (String) Designates how this asset is used / the environment that the asset is supporting.

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

- `access_id` - (String) The Access key ID of AWS secret access key used for authentication
- `access_key` - (String) The Secret access key used for authentication
- `amazon_secret` - An `amazon_secret` block as defined below, to integrate the asset with AWS Secrets Manager.
- `application_id` - (String) This is also referred to as the Client ID and it’s the unique identifier for the registered application being used to execute Python SDK commands against Azure’s API services. You can find this number under Azure Active Directory -> App Registrations -> Owned Applications
- `azure_storage_account` - (String) The name of the azure storage account. The field can contain only lowercase letters and numbers. Name must be between 3 and 24 characters.
- `azure_storage_container` - (String) Location where a given EventHub’s processing is stored (One storage container per EventHub). This name may only contain lowercase letters, numbers, and hyphens, and must begin with a letter or a number. Each hyphen must be preceded and followed by a non-hyphen character. The name must also be between 3 and 63 characters long.
- `azure_storage_secret_key` - (String) The secret key for the storage account associated with this audit setup.
- `cache_file` - (String) Holds Kerberos protocol credentials (for example, tickets, session keys and other identifying information).
- `ca_certs_path` - (String) Certificate authority certificates path; what location should the sysetm look for certificate information from. Equivalent to --capath in a curl call
- `client_secret` - (String) This a string containing a secret used by the application to prove its identity when requesting a token. You can get a secret by going to Azure Active Directory -> App Registrations -> Owned Applications, selecting the desired application and then going to Certificates & secrets -> Client secrets -> + New client secret
- `cyberark_secret` - A `cyberark_secret` block as defined below, to integrate the asset with CyberArk.
- `db_role` - (String) The database role to use when connecting to this asset
- `directory_id` - (String) This is also referred to as the Tenant ID and is a GUID representing the Active Directory Tenant. It can be found in the Azure Active Directory page under the Azure portal
- `eventhub_access_key` - (String) EventHub access key for this eventhub.
- `eventhub_access_policy` - (String) Authorization policy that will allow Sonar to access this specific EventHub. Example: 'RootManageSharedAccessKey'.
- `eventhub_name` - (String) EventHub name without additional resource ID information.
- `eventhub_namespace` - (String) The name for the management container that the EventHub belongs to, one namespace can contain multiple EventHubs. The namespace can contain only letters, numbers, and hyphens. The namespace must start with a letter, and it must end with a letter or number. The value must be between 6 and 50 characters long.
- `external` - (Boolean)
- `external_id` - (String) External ID to use when assuming a role
- `extra_kinit_parameters` - (String)
- `format` - (String) The type of audit data being sent to EventHub. Please see the asset specifications of Azure Event Hubs for an up-to-date list of accepted values.
- `hashicorp_secret` - A `hashicorp_secret` block as defined below, to integrate the asset with a HashiCorp Vault.
- `key_file` - (String) Location on disk on the key to be used for authentication
- `kerberos_kdc` - (String) The host name or IP Address of your Kerberos KDC machine
- `kerberos_service_kdc` - (String) Kerberos Service KDC
- `kerberos_service_realm` - (String) Kerberos Service Realm
- `kerberos_spn` - (String) The service and host of the Sybase Kerberos Principal. This will be the value prior to the '@' symbol of the principal value
- `keytab_file` - (String) Specify a non-default keytab location
- `kinit_program_path` - (String)
- `passphrase` - (String) Passphrase for the private key.
- `password` - (String) The password of the user being used for authentication
- `principal` - (String) The principal used for authentication
- `region` - (String) Default AWS region for this asset
- `role_name` - (String) What role is used to get credentials from.
- `secret_key` - (String) The Secret access key used for authentication
- `ssl` (Boolean) If true, use SSL when connecting
- `ssl_server_cert` - (String) Path to server certificate to use during authentication
- `subscription_id` - (String) This is the Azure account subscription ID. You can find this number under the Subscriptions page on the Azure portal
- `use_keytab` - (Boolean) If true, authenticate using a key tab
- `username` - (String) The name of a profile in /imperva/local/credentials/.aws/credentials to use for authenticating
- `user_identity_client_id` - (String) The client ID of a user-assigned managed identity. Defaults to the value of the environment variable AZURE_CLIENT_ID, if any. If not specified, a system-assigned identity will be used.

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

In Terraform v1.5.0 and later, use an import block to import Log Aggregators using the `asset_id`. For example:

```
import {
  to = dsf_log_aggregator.example_aws_log_group
  id = "arn:aws:logs:us-east-2:123456789012:log-group:/aws/rds/instance/my-database/audit:*"
}
```

Using terraform import, import Log Aggregators using the `asset_id`. For example:

```
$ terraform import dsf_log_aggregator.example_aws_log_group "arn:aws:logs:us-east-2:123456789012:log-group:/aws/rds/instance/my-database/audit:*"
```

For detailed instructions on onboarding existing cloud resources to DSF using Terraform's import functionality, see [Importing and Onboarding Existing Data Sources with Terraform](https://docs-cybersec.thalesgroup.com/bundle/onboarding-databases-to-sonar-reference-guide/page/Importing-and-Onboarding-Existing-Data-Sources-with-Terraform_784990209.html).
