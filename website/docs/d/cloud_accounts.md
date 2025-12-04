---
subcategory: ""
layout: "dsfhub"
page_title: "DSFHUB Cloud Accounts - Terraform Data Source"
description: |-
  Provides dsfhub_cloud_accounts terraform data source.
---

# Data Source: dsfhub_cloud_accounts

Use this data source to get asset_ids of Cloud Accounts.

## Example Usage

All Cloud Accounts on a DSF HUB:

```hcl
data "dsfhub_cloud_accounts" "cloud_accounts" {}
```

Cloud Accounts filtered by asset_id regex:

```hcl
data "dsfhub_cloud_accounts" "cloud_accounts" {
  asset_id_regex = ".*12345.*/your-user-name"  
}
```

Look up a Cloud Account by a specific asset_id to see if it exists:

```hcl
data "dsfhub_cloud_accounts" "cloud_accounts" {
  asset_id_regex = "arn:aws:iam::1234567890:user/your-user-name"  
}
```

## Argument Reference

This data source supports the following arguments:

- `asset_id_regex` (String) Optional - Regex string to apply to the Cloud Accounts list returned by DSFHUB. This allows for more advanced filtering not supported from the DSFHUB API. This filtering is done locally on what DSFHUB returns.

## Attribute Reference

The following attributes are exported:

- `asset_ids` - (String) Set of asset_ids of the matched Cloud Accounts.
