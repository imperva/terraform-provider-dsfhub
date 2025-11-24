---
layout: "dsfhub"
page_title: "Provider: DSFHUB"
description: |-
  The DSFHUB provider is used to interact with Data Security Fabric Hub structured resources supported by Thales Group.
---

# DSFHUB Provider

The DSFHUB provider is used to interact with Data Security Fabric Hub structured resources supported by Thales Group. The provider needs to be configured with your DSF Hub endpoint and authentication token before it can be used.

Use the navigation to the left to read about the available resources and data sources. For integrated examples with cloud resource configuration and audit setup, please see the [DSFHub Agentless-Onboarding modules](https://github.com/imperva/terraform-dsfhub-agentless-onboarding).

## DSFHUB Provider Argument Reference

The following arguments are supported:

* `dsfhub_host` - (Required) The DSF Hub endpoint for [DSF HUB API](https://docs-cybersec.thalesgroup.com/bundle/v15.0-sonar-user-guide/page/84552.htm) operations. Example: 'https://yourDSFhostname:8443' or 'https://1.2.3.4:8443'.
* `dsfhub_token` - (Required) The [DSF API Token](https://docs-cybersec.thalesgroup.com/bundle/v15.0-sonar-user-guide/page/84555.htm) for API operations.
* `insecure_ssl` - (Optional) The boolean flag that instructs the provider to allow for insecure SSL API calls to a DSF Hub instance to support tests against instances with self-signed certificates.
* `sync_type` - (Optional) Determines whether to sync asset creation/update operations with the Agentless gateways. Defaults to SYNC_GW_BLOCKING. Available values: 
  - `SYNC_GW_BLOCKING`: The operation is synchronous and blocks until all gateways have been updated. This means that, if syncing the assets to Agentless Gateways fails, the provider will throw an error and not continue. This may result in a difference between the state of which Terraform is aware and the assets that were actually imported.
  - `SYNC_GW_NON_BLOCKING`: The operation is asynchronous and returns immediately.
  - `DO_NOT_SYNC_GW`: The operation is synchronous and does not update the gateways.

!> **Warning:** Hard-coded tokens and credentials are not recommended in any Terraform configuration and risks secret leakage should this file ever be committed to a public version control system.

Provider arguments can be set by adding an `dsfhub_host`, `dsfhub_token`, and optionally `insecure_ssl` and `sync_type`, to the `dsfhub` provider block.

Usage:
```hcl
terraform {
  required_providers {
    dsfhub = {
      source = "imperva/dsfhub"
    }
  }
}

provider "dsfhub" {
  dsfhub_host = "https://1.2.3.4:8443"
  dsfhub_token = "a1b2c3d4-e5f6-g8h9-wxyz-123456790"
  # insecure_ssl
  # sync_type
}
```

### Environment Variables
Provider arguments can be provided using the `DSFHUB_HOST`, `DSFHUB_TOKEN`, and optionally `INSECURE_SSL` or `SYNC_TYPE` environment variables.

For example:
```hcl
terraform {
  required_providers {
    dsfhub = {
      source = "imperva/dsfhub"
    }
  }
}

provider "dsfhub" {}
```

And on your terminal:
```bash
$ export DSFHUB_HOST="https://1.2.3.4:8443"
$ export DSFHUB_TOKEN="a1b2c3d4-e5f6-g8h9-wxyz-123456790"
$ export INSECURE_SSL=true
$ export SYNC_TYPE="SYNC_GW_NON_BLOCKING"
$ terraform plan
```
