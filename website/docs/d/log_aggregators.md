---
subcategory: "Provider Reference"
layout: "dsfhub"
page_title: "DSFHUB Log Aggregators - Terraform Data Source"
description: |-
  Provides dsfhub_log_aggregators terraform data source.
---

# dsfhub_log_aggregators (Data Source)

Use this data source to get the asset_ids of LogAggregators.

### Example Usage

___

All LogAggregators on a DSF HUB

```hcl
data "dsfhub_log_aggregators" "log_aggregators" {}
```

LogAggregators filtered by asset_id regex

```hcl
data "dsfhub_log_aggregators" "log_aggregators" {
  asset_id_regex = ".*us-east-2.*"  
}
```

Look up a LogAggregator by specific asset_id to see if it exists

```hcl
data "dsfhub_log_aggregators" "log_aggregators" {
  asset_id_regex = "arn:partition:service:region:account-id" 
}
```

## Argument Reference

This data source supports the following arguments:

- `asset_id_regex` (String) Optional - Regex string to apply to the LogAggregators list returned by DSFHUB. This allows more advanced filtering not supported from the DSFHUB API. This filtering is done locally on what DSFHUB returns.

## Attribute Reference

The following attributes are exported:

- `asset_ids` (String) Set of asset_ids of the matched LogAggregators.
