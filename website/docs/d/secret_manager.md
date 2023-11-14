---
subcategory: "Provider Reference"
layout: "dsfhub"
page_title: "DSFHUB Secret Manager - Terraform Data Source"
description: |-
  Provides a dsfhub_secret_manager terraform data source.
---

# dsfhub_secret_manager (Data Source)

Provides SecretManager resource configuration from a unique asset_id.

## Example Usage

```hcl
data "dsfhub_secret_manager" "example_hashicorp" {
  asset_id = "my.hashicorp.vault.server.com" 
}
```

## Argument Reference

- `asset_id` (String) Current asset ID

## Attribute Reference

The following attributes are exported:

- `id` (String) The asset_id of this resource.
- `asset_id` (String) Current asset_id