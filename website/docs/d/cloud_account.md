---
layout: "dsfhub"
page_title: "dsfhub_cloud_account Cloud Account"
sidebar_current: "docs-dsfhub-data-source-dsfhub-cloud-account"
description: |-
Provides a dsfhub_data_source data source.  
---

# dsfhub_cloud_account (Data Source)

Returns a CloudAccount resource configuration from a unique asset_id.

## Example Usage

```hcl
data "dsfhub_cloud_account" "example_aws_cloud_account" {
  asset_id = "arn:partition:service:region:account-id" # The value of the arn for aws resources
}

resource "dsfhub_data_source" "rds_mysql_db" {
  server_type = "AWS RDS MYSQL"

  admin_email = var.admin_email	# The email address to notify about this asset
  arn = var.arn	# Amazon Resource Name - format is arn:partition:service:region:account-id:resource-type:resource-id and used as the asset_id
  asset_display_name = var.asset_display_name	# User-friendly name of the asset, defined by user.
  asset_id = var.asset_id	# Asset ID
  gateway_id = var.gateway_id	# Gateway ID
  server_host_name = var.server_host_name	# Hostname (or IP if name is unknown)
  region              = "us-east-2"
  server_port         = 3306
  version             = 8.0
  
  parent_asset_id     = dsfhub_cloud_account.example_aws_cloud_account.asset_id

  asset_connection {
    auth_mechanism  = "password"
    password        = var.admin_password
    reason          = "default"
    username        = var.admin_username
  }
}

```

## Argument Reference

- `asset_id` (String) Current asset ID

## Attribute Reference

The following attributes are exported:

- `id` (String) The asset_id of this resource.
- `asset_id` (String) Current asset_id
