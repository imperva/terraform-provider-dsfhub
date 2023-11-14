---
subcategory: "DSFHUB Provider Reference"
layout: "dsfhub"
page_title: "DSFHUB Log Aggregator - Terraform Data Source"
description: |-
  Provides a dsfhub_log_aggregator terraform data source.
---

# dsfhub_log_aggregator (Data Source)

Provides LogAggregator resource configuration from a unique asset_id.

## Example Usage

```hcl
data "dsfhub_log_aggregator" "example_aws_log_group_default" {
  asset_id = "arn:partition:service:region:account-id" # Use arn for aws resources
}
```

## Argument Reference

- `asset_id` (String) Current asset ID

## Attribute Reference

The following attributes are exported:

- `id` (String) The asset_id of this resource.
- `asset_id` (String) Current asset_id