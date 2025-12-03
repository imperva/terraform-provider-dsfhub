---
subcategory: ""
layout: "dsfhub"
page_title: "DSFHUB Log Aggregator - Terraform Data Source"
description: |-
  Provides a dsfhub_log_aggregator terraform data source.
---

# Data Source: dsfhub_log_aggregator

Use this data source to get a Log Aggregator resource given a unique asset_id.

## Example Usage

```hcl
data "dsfhub_log_aggregator" "example_aws_log_group" {
  asset_id = "arn:aws:logs:us-east-2:123456789012:log-group:/aws/rds/instance/my-database/audit:*"
}

resource "dsfhub_data_source" "aws_rds_mysql_asset" {
  server_type = "AWS RDS MYSQL"

  admin_email               = var.admin_email
  asset_display_name        = var.asset_display_name
  asset_id                  = var.asset_id	
  gateway_id                = var.gateway_id
  logs_destination_asset_id = data.dsfhub_log_aggregator.example_aws_log_group.asset_id
  region                    = "us-east-2"
  server_host_name          = var.server_host_name	
  server_port               = "3306"
}
```

## Argument Reference

- `asset_id` - (String) The asset_id of the asset.

## Attribute Reference

The following attributes are exported:

- `id` - (String) The asset_id of the asset.
- `asset_id` - (String) The asset_id of the asset.