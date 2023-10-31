---
layout: "dsfhub"
page_title: "dsfhub_secret_manager Data Source"
sidebar_current: "docs-dsfhub-data-source-dsfhub-secret-manager"
description: |-
Provides a dsfhub_secret_manager data source.  
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