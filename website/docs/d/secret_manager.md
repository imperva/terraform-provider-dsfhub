---
subcategory: ""
layout: "dsfhub"
page_title: "DSFHUB Secret Manager - Terraform Data Source"
description: |-
  Provides a dsfhub_secret_manager terraform data source.
---

# Data Source: dsfhub_secret_manager

Use this data source to get a Secret Manager resource given a unique asset_id.

## Example Usage

```hcl
data "dsfhub_secret_manager" "example_hashicorp" {
  asset_id = "my.hashicorp.vault.server.com" 
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
  
  asset_connection {
    auth_mechanism  = "password"
    password        = var.admin_password
    reason          = "default"
    username        = var.admin_username

    hashicorp_secret {
      secret_asset_id = data.dsfhub_secret_manager.example_hashicorp.asset_id
      secret_name     = "my-secret"
      path            = "secret/data/myapp/dbcreds"
      field_mapping = {
        username = "my-remote-user"
        password = "my-remote-password"
      }
    }
  }
}
```

## Argument Reference

- `asset_id` - (String) The asset_id of the asset.

## Attribute Reference

The following attributes are exported:

- `id` - (String) The asset_id of the asset.
- `asset_id` - (String) The asset_id of the asset.