---
subcategory: "Provider Reference"
layout: "dsfhub"
page_title: "DSFHUB Log Aggregator Resource"
description: |-
  Provides a dsfhub_log_aggregator resource.
---

# dsfhub_log_aggregator (Resource)

Provides a log aggregator resource.  A log aggregator is a resource that serves as a log destination for one or multiple data sources in the DSF HUB.

The `dsfhub_log_aggregator` resource contains the configuration parameters necessary to aggregate logs for one or multiple data sources in the DSF HUB platform.
Documentation for the underlying API used in this resource can be found at
[Log Aggregators API Definition page](https://docs.imperva.com/bundle/v4.13-sonar-user-guide/page/84552.htm).

## Example Usage

```hcl
# Example generic variable reference:
variable "admin_email" {
  default = "your@email.com"
}
variable "gateway_id" {
  default = "7a4af7cf-4292-89d9-46ec-183756ksdjd"
}
variable "region" {
  default = "us-east-1"
}

# Example dsfhub_log_aggregator specific variables for AWS LOG GROUP
variable "log_aggregator_aws_log_group_asset_display_name" {
  default = "arn:partition:service:region:account-id"
}
variable "log_aggregator_aws_log_group_asset_id" {
  default = "arn:partition:service:region:account-id"
}
variable "log_aggregator_parent_data_source_asset_id" {
  default = "your-data-source-asset-id-here"
}
variable "log_aggregator_aws_log_group_version" {
  default = 1.0
}

# Example dsfhub_log_aggregator usage for AWS LOG GROUP
resource "dsfhub_log_aggregator" "example_aws_log_group_default" {
  server_type = "AWS LOG GROUP"
  admin_email = var.admin_email	# The email address to notify about this asset
  asset_display_name = var.log_aggregator_aws_log_group_asset_display_name # User-friendly name of the asset, defined by user.
  asset_id = var.log_aggregator_aws_log_group_asset_id # The unique identifier or resource name of the asset. For AWS, use arn, for Azure, use subscription ID, for GCP, use project ID
  gateway_id = var.gateway_id # The jsonarUid unique identifier of the agentless gateway. Example: '7a4af7cf-4292-89d9-46ec-183756ksdjd'
  parent_asset_id = var.log_aggregator_parent_data_source_asset_id # The name of an asset that this asset is part of (/related to). E.g. an AWS resource will generally have an AWS account asset as its parent. Also used to connect some log aggregating asset with the sources of their logs. E.g. An AWS LOG GROUP asset can have an AWS RDS as its parent, indicating that that is the log group for that RDS.
  version = var.log_aggregator_aws_log_group_version # Denotes the version of the asset
  asset_connection {
    auth_mechanism = "default"
    reason = "default"
    region = var.region # For cloud systems with regions, the default region or region used with this asset
  }
}

```

## Log Aggregator Types
<ul>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/log_aggregators/alibaba_logstore.md">Alibaba Logstore</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/log_aggregators/aws_kinesis.md">AWS Kinesis</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/log_aggregators/aws_log_group.md">AWS Log Group</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/log_aggregators/aws_s3.md">Amazon S3</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/log_aggregators/azure_eventhub.md">Azure EventHub</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/log_aggregators/gcp_cloud_storage_bucket.md">GCP Cloud Storage Bucket</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/log_aggregators/gcp_pubsub.md">Google Cloud Pub/Sub</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/log_aggregators/ssh.md">SSH</a></li>
</ul>


## Argument Reference

### Required

- `admin_email` (String) The email address to notify about this asset
- `asset_connection` (Block Set, Min: 1) N/A (see [below for nested schema](#nestedblock--asset_connection))
- `asset_id` (String) The unique identifier or resource name of the asset.
- `gateway_id` (String) The jsonarUid unique identifier of the agentless gateway. Example: '7a4af7cf-4292-89d9-46ec-183756ksdjd'
- `server_type` (String) The type of server or data service to be created as a data source. The list of available data sources is documented at: https://docs.imperva.com/bundle/v4.11-sonar-user-guide/page/84552.htm

### Optional

- `asset_display_name` (String) User-friendly name of the asset, defined by user.
- `asset_source` (String) The source platform/vendor/system of the asset data. Usually the service responsible for creating that asset document
- `audit_type` (String) Used to indicate what mechanism should be used to fetch logs on systems supporting multiple ways to get logs, see asset specific documentation for details.  Example: "BIGQUERY","BIGTABLE","BUCKET","MSSQL","MYSQL","POSTGRESQL","SPANNER".
- `available_regions` (String) A list of regions to use in discovery actions that iterate through region
- `aws_proxy_config` (Block Set) AWS specific proxy configuration (see [below for nested schema](#nestedblock--aws_proxy_config))
- `credentials_endpoint` (String) A specific sts endpoint to use
- `criticality` (Number) The asset's importance to the business. These values are measured on a scale from "Most critical" (1) to "Least critical" (4). Allowed values: 1, 2, 3, 4
- `endpoint` (String) Logstore's endpoint
- `jsonar_uid` (String) Unique identifier (UID) attached to the Sonar machine controlling the asset
- `location` (String) Current human-readable description of the physical location of the asset, or region.
- `logstore` (String) Unit that is used to collect, store and query logs
- `managed_by` (String) Email of the person who maintains the asset; can be different from the owner specified in the owned_by field. Defaults to admin_email.
- `owned_by` (String) Email of Owner / person responsible for the asset; can be different from the person in the managed_by field. Defaults to admin_email.
- `parent_asset_id` (String) The name of an asset that this asset is part of (/related to). E.g. an AWS resource will generally have an AWS account asset as its parent. Also used to connect some log aggregating asset with the sources of their logs. E.g. An AWS LOG GROUP asset can have an AWS RDS as its parent, indicating that that is the log group for that RDS.
- `project` (String) Project separates different resources of multiple users and control access to specific resources
- `proxy` (String) Proxy to use for AWS calls if aws_proxy_config is populated the proxy field will get populated from the http value there
- `region` (String) For cloud systems with regions, the default region or region used with this asset
- `server_host_name` (String) Hostname (or IP if name is unknown)
- `server_ip` (String) IP address of the service where this asset is located. If no IP is available populate this field with other information that would identify the system e.g. hostname or AWS ARN, etc.
- `server_port` (String) Port used by the source server
- `service_endpoints` (Block Set) Specify particular endpoints for a given service in the form of <service name>: "endpoint" (see [below for nested schema](#nestedblock--service_endpoints))
- `used_for` (String) Designates how this asset is used / the environment that the asset is supporting.
- `version` (Number) Denotes the version of the asset

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--asset_connection"></a>
### Nested Schema for `asset_connection`

Required:

- `auth_mechanism` (String) Specifies the auth mechanism used by the connection
- `reason` (String) N/A

Optional:

- `access_id` (String) The Access key ID of AWS secret access key used to authenticate
- `access_key` (String) The Secret access key used to authenticate
- `amazon_secret` (Block Set, Max: 1) Configuration to integrate with AWS Secrets Manager (see [below for nested schema](#nestedblock--asset_connection--amazon_secret))
- `application_id` (String) This is also referred to as the Client ID and it’s the unique identifier for the registered application being used to execute Python SDK commands against Azure’s API services. You can find this number under Azure Active Directory -> App Registrations -> Owned Applications
- `ca_certs_path` (String) Certificate authority certificates path; what location should the sysetm look for certificate information from. Equivalent to --capath in a curl call
- `client_secret` (String) This a string containing a secret used by the application to prove its identity when requesting a token. You can get a secret by going to Azure Active Directory -> App Registrations -> Owned Applications, selecting the desired application and then going to Certificates & secrets -> Client secrets -> + New client secret
- `credential_fields` (Block Set) Document containing values to build a profile from. Filling this will create a profile using the given profile name (see [below for nested schema](#nestedblock--asset_connection--credential_fields))
- `cyberark_secret` (Block Set, Max: 1) Configuration to integrate with AWS Secrets Manager (see [below for nested schema](#nestedblock--asset_connection--cyberark_secret))
- `directory_id` (String) This is also referred to as the Tenant ID and is a GUID representing the Active Directory Tenant. It can be found in the Azure Active Directory page under the Azure portal
- `external_id` (String) External ID to use when assuming a role
- `hashicorp_secret` (Block Set) Configuration to integrate with HashiCorp Vault (see [below for nested schema](#nestedblock--asset_connection--hashicorp_secret))
- `key_file` (String) Location on disk on the key to be used to authenticate
- `region` (String) Default AWS region for this asset
- `role_name` (String) What role is used to get credentials from.
- `secret_key` (String) The Secret access key used to authenticate
- `ssl` (Boolean) If true, use SSL when connecting
- `subscription_id` (String) This is the Azure account subscription ID. You can find this number under the Subscriptions page on the Azure portal
- `username` (String) The name of a profile in /imperva/local/credentials/.aws/credentials to use for authenticating

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

Log Aggregators can be imported using the `asset_id`, e.g.:

```
$ terraform import dsf_log_aggregator.example_aws_log_group_default "arn:partition:service:region:account-id"
```
