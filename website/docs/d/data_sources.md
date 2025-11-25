---
subcategory: ""
layout: "dsfhub"
page_title: "DSFHUB Data Sources - Terraform Data Source"
description: |-
  Provides dsfhub_data_sources terraform data source.
---

# Data Source: dsfhub_data_sources

Use this data source to get asset_ids of DSF Data Sources.

## Example Usage

___

All DSF Data Sources on a DSF HUB:

```hcl
data "dsfhub_data_sources" "data_sources" {}
```

DSF Data Sources filtered by asset_id regex:

```hcl
data "dsfhub_data_sources" "data_sources" {
  asset_id_regex = ".*us-east-2.*"  
}
```

Look up a DSF Data Source by a specific asset_id to see if it exists:

```hcl
data "dsfhub_data_sources" "data_sources" {
  asset_id_regex = "arn:partition:service:region:account-id" 
}
```

## Argument Reference

This data source supports the following arguments:

- `asset_id_regex` (String) Optional - Regex string to apply to the Cloud Accounts list returned by DSFHUB. This allows for more advanced filtering not supported from the DSFHUB API. This filtering is done locally on what DSFHUB returns.

## Attribute Reference

The following attributes are exported:

- `asset_ids` - (String) Set of asset_ids of the matched DSF Data Sources.
