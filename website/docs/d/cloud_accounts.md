---
subcategory: "Provider Reference"
layout: "dsfhub"
page_title: "DSFHUB Cloud Accounts - Terraform Data Source"
description: |-
  Provides dsfhub_cloud_accounts terraform data source.
---

# dsfhub_cloud_accounts (Data Source)

Use this data source to get the asset_ids of CloudAccounts.

### Example Usage

___

All CloudAccounts on a DSF HUB

```hcl
data "dsfhub_cloud_accounts" "cloud_accounts" {}
```

CloudAccounts filtered by asset_id regex

```hcl
data "dsfhub_cloud_accounts" "cloud_accounts" {
  asset_id_regex = ".*12345.*/your-user-name"  
}
```

Look up a cloudAccount by specific asset_id to see if it exists

```hcl
data "dsfhub_cloud_accounts" "cloud_accounts" {
  asset_id_regex = "arn:aws:iam::1234567890:user/your-user-name"  
}
```

## Argument Reference

This data source supports the following arguments:

- `asset_id_regex` (String) Optional - Regex string to apply to the CloudAccount list returned by DSFHUB. This allows more advanced filtering not supported from the DSFHUB API. This filtering is done locally on what DSFHUB returns.

## Attribute Reference

The following attributes are exported:

- `asset_ids` (String) Set of asset_ids of the matched CloudAccounts.
