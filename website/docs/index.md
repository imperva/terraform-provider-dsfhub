---
subcategory: "IAM Access Analyzer"
layout: "dsfhub"
page_title: "Provider: DSFHUB"
sidebar_current: "docs-dsfhub-index"
description: |-
  The DSFHUB provider is used to interact with Data Security Fabric Hub resources supported by Imperva. The provider needs to be configured with the endpoint and authentication token before it can be used.
---

# DSFHUB Provider

The DSFHUB provider is used to interact with Data Security Fabric Hub resources supported by Imperva. The provider needs to be configured with the endpoint and authentication token before it can be used.

Use the navigation to the left to read about the available resources.

## DSFHUB Provider Argument Reference

The following arguments are supported:

* `dsfhub_host` - (Required) The DSF host endpoint for [DSF HUB API](https://docs.imperva.com/bundle/v4.13-sonar-user-guide/page/84552.htm) operations. Example: 'https://1.2.3.4:8443'. Can be set via `DSFHUB_HOST` shell [environment variable](https://en.wikipedia.org/wiki/Environment_variable).
* `dsfhub_token` - (Required) The [DSF API Token](https://docs.imperva.com/bundle/v4.13-sonar-user-guide/page/84555.htm) for API operations. You can retrieve this from the DSF management hub console. Can be set via `DSFHUB_TOKEN` shell [environment variable](https://en.wikipedia.org/wiki/Environment_variable).
* `insecure_ssl` - (Optional) The boolean flag that instructs the provider to allow for insecure SSL API calls to a DSF Hub instance to support tests against instances with self-signed certificates. Can be set via `INSECURE_SSL` shell [environment variable](https://en.wikipedia.org/wiki/Environment_variable).
* `sync_type` - (Optional) Determines whether to sync asset creation/update operations with the Agentless gateways. Defaults to SYNC_GW_BLOCKING. Can be set via `SYNC_TYPE` shell [environment variable](https://en.wikipedia.org/wiki/Environment_variable). Available values: 
  - SYNC_GW_BLOCKING: The operation is synchronous and blocks until all gateways have been updated. This means that, if syncing the assets to Agentless Gateways fails, the provider will throw an error and not continue. This may result in a difference between the state of which Terraform is aware and the assets that were actually imported.
  - SYNC_GW_NON_BLOCKING: The operation is asynchronous and returns immediately.
  - DO_NOT_SYNC_GW: The operation is synchronous and does not update the gateways.

For example,
```hcl
# Specify path for provider
terraform {
  required_providers {
    dsfhub = {
      source = "imperva/dsfhub"
    }
  }
}

provider "dsfhub" {
  dsfhub_host = "https://1.2.3.4:8443"
  dsfhub_token = "a1b2c3d4-e5f6-g8h9-wxyz-123456790"
}
```

### Environment Variables
Provider arguments can be set via environment variables as noted above. For example,
```hcl

# Specify path for provider
terraform {
  required_providers {
    dsfhub = {
      source = "imperva/dsfhub"
    }
  }
}

provider "dsfhub" {}
```
```bash
$ export DSFHUB_HOST="https://1.2.3.4:8443"
$ export DSFHUB_TOKEN="a1b2c3d4-e5f6-g8h9-wxyz-123456790"
$ export INSECURE_SSL=true
$ export SYNC_TYPE="SYNC_GW_NON_BLOCKING"
$ terraform plan
```

## Example Usage - dsfhub_cloud_account

The following is an example of creating a  [dsfhub_cloud_account](../r/cloud_account.md) resource used in this example to connect the DSFHUB to an AWS account. 

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

# Example dsfhub_cloud_account specific variables for AWS
variable "cloud_account_aws_asset_display_name" {
  default = "arn:partition:service:region:account-id"
}
variable "cloud_account_aws_asset_id" {
  default = "arn:partition:service:region:account-id"
}

# Example dsfhub_cloud_account usage for AWS
resource "dsfhub_cloud_account" "example_aws_cloud_account" {
  server_type = "AWS"
  admin_email = var.admin_email	# The email address to notify about this asset
  asset_display_name = var.cloud_account_aws_asset_display_name # User-friendly name of the asset, defined by user.
  asset_id = var.cloud_account_aws_asset_id # The unique identifier or resource name of the asset. For AWS, use arn, for Azure, use subscription ID, for GCP, use project ID
  gateway_id = var.gateway_id# The jsonarUid unique identifier of the agentless gateway. Example: '7a4af7cf-4292-89d9-46ec-183756ksdjd'
  asset_connection {
    auth_mechanism = "iam_role"
    reason = "default"
    region = var.region # For cloud systems with regions, the default region or region used with this asset
  }
}

```

## Example Usage - dsfhub_data_source

The following is an example of creating a  [dsfhub_data_source](../r/data_source.md) resource used to describe the database asset for the agentless gateway to consume audit from. The `dsfhub_cloud_account.asset_id` is referenced in the `dsfhub_data_source.parent_asset_id` param. 

```hcl
# Example dsfhub_data_source specific variables for AWS RDS MYSQL
variable "data_source_aws_rds_mysql_asset_display_name" {
  default = "arn:partition:service:region:account-id"
}
variable "data_source_aws_rds_mysql_asset_id" {
  default = "arn:partition:service:region:account-id"
}
variable "data_source_aws_rds_mysql_server_host_name" {
  default = "your-data-source-asset-id-here"
}
variable "data_source_aws_rds_mysql_username" {
  default = "your-db-username"
}
variable "data_source_aws_rds_mysql_password" {
  default = "your-db-password--here"
}

# Example dsfhub_data_source usage for AWS RDS MYSQL
resource "dsfhub_data_source" "aws_rds_mysql_password" {
  server_type = "AWS RDS MYSQL"
  admin_email = var.admin_email	
  asset_display_name = var.data_source_aws_rds_mysql_asset_display_name	
  asset_id = var.data_source_aws_rds_mysql_asset_id 
  gateway_id = var.gateway_id
  parent_asset_id = dsf_cloud_account.example_aws_cloud_account.asset_id
  server_host_name = var.data_source_aws_rds_mysql_server_host_name	
  asset_connection {
    auth_mechanism = "password"
    password = var.data_source_aws_rds_mysql_password 
    reason = "default" 
    username = var.data_source_aws_rds_mysql_username 
  }
}
```

## Example Usage - dsfhub_log_aggregator

The following is an example of creating a [dsfhub_log_aggregator](../r/log_aggregator.md) resource used as the mechanism for the agentless gateway to consume audit data in this case via `AWS LOG GROUP`. The `dsfhub_data_source.asset_id` is referenced in the `dsfhub_log_aggregator.parent_asset_id` param.

```hcl
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
  admin_email = var.admin_email
  asset_display_name = var.log_aggregator_aws_log_group_asset_display_name
  asset_id = var.log_aggregator_aws_log_group_asset_id
  gateway_id = var.gateway_id
  parent_asset_id = dsfhub_data_source.aws_rds_mysql_password.asset_id
  version = var.log_aggregator_aws_log_group_version
  asset_connection {
    auth_mechanism = "default"
    reason = "default"
    region = var.region
  }
}
```

