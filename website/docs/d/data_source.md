---
subcategory: ""
layout: "dsfhub"
page_title: "DSFHUB Data Source - Terraform Data Source"
description: |-
  Provides a dsfhub_data_source terraform data source.
---

# Data Source: dsfhub_data_source

Use this data source to get a Data Source resource given a unique asset_id.

## Example Usage

```hcl
data "dsfhub_data_source" "example_aws_rds_mysql" {
  asset_id = "arn:rds:mysql:db:region:account-id"
}

resource "dsfhub_log_aggregator" "example_aws_log_group" {
  server_type        = "AWS LOG GROUP"
  admin_email        = "your@email.com"
  asset_display_name = "arn:partition:service:region:account-id"
  asset_id           = "arn:partition:service:region:account-id"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  parent_asset_id    = data.dsf_data_source.example_aws_rds_mysql.asset_id
  asset_connection {
    auth_mechanism = "default"
    reason         = "default"
    region         = "us-east-2"
  }
}
```

## Argument Reference

- `asset_id` - (String) The asset_id of the asset.

## Attribute Reference

The following attributes are exported:

- `id` - (String) The asset_id of the asset.
- `asset_id` - (String) The asset_id of the asset.
