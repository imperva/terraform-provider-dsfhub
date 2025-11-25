---
subcategory: ""
layout: "dsfhub"
page_title: "DSFHUB Cloud Account - Terraform Data Source"
description: |-
  Provides a dsfhub_cloud_account terraform data source.
---

# Data Source: dsfhub_cloud_account

Use this data source to get a Cloud Account resource given a unique asset_id.

## Example Usage

```hcl
data "dsfhub_cloud_account" "example_aws_cloud_account" {
  asset_id = "arn:partition:service:region:account-id"
}

resource "dsfhub_data_source" "aws_rds_mysql_asset" {
  server_type = "AWS RDS MYSQL"

  admin_email        = var.admin_email
  asset_display_name = var.asset_display_name
  asset_id           = var.asset_id	
  gateway_id         = var.gateway_id
  region             = "us-east-2"
  server_host_name   = var.server_host_name	
  server_port        = "3306"
  
  parent_asset_id = data.dsfhub_cloud_account.example_aws_cloud_account.asset_id

  asset_connection {
    auth_mechanism  = "password"
    password        = var.admin_password
    reason          = "default"
    username        = var.admin_username
  }
}

```

## Argument Reference

- `asset_id` - (String) The asset_id of the asset.

## Attribute Reference

The following attributes are exported:

- `id` - (String) The asset_id of the asset.
- `asset_id` - (String) The asset_id of the asset.
