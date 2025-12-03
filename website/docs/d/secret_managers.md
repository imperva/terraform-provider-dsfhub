---
subcategory: ""
layout: "dsfhub"
page_title: "DSFHUB Secret Managers - Terraform Data Source"
description: |-
  Provides dsfhub_secret_managers terraform data source.
---

# Data Source: dsfhub_secret_managers

Use this data source to get asset_ids of Secret Managers.

## Example Usage

___

All Secret Managers on a DSF HUB:

```hcl
data "dsfhub_secret_managers" "secret_managers" {}
```

Secret Managers filtered by asset_id regex:

```hcl
data "dsfhub_secret_managers" "secret_managers" {
  asset_id_regex = ".*us-east-2.*"  
}
```

Look up a Secret Manager by a specific asset_id to see if it exists:

```hcl
data "dsfhub_secret_managers" "secret_managers" {
  asset_id_regex = "arn:partition:service:region:account-id" 
}
```

## Argument Reference

This data source supports the following arguments:

- `asset_id_regex` (String) Optional - Regex string to apply to the Secret Managers list returned by DSFHUB. This allows for more advanced filtering not supported from the DSFHUB API. This filtering is done locally on what DSFHUB returns.

## Attribute Reference

The following attributes are exported:

- `asset_ids` - (String) Set of asset_ids of the matched Secret Managers.
