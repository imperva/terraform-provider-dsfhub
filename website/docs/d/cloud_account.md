---
layout: "dsfhub"
page_title: "dsfhub_cloud_account Cloud Account"
sidebar_current: "docs-dsfhub-data-source-dsfhub-cloud-account"
description: |-
Provides a dsfhub_data_source data source.  
---

# dsfhub_cloud_account (Data Source)

Returns a CloudAccount resource configuration from a unique asset_id.

## Example Usage

```hcl
data "dsfhub_cloud_account" "example_aws_cloud_account" {
  asset_id = "arn:partition:service:region:account-id" # The value of the arn for aws resources
}
```

## Argument Reference

- `asset_id` (String) Current asset ID

## Attribute Reference

The following attributes are exported:

- `id` (String) The asset_id of this resource.
- `asset_id` (String) Current asset_id
