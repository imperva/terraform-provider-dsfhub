---
layout: "dsfhub"
page_title: "dsfhub_cloud_account Resource"
sidebar_current: "docs-dsfhub-resource-dsfhub-cloud-account"
description: |-
Provides a dsfhub_cloud_account resource.
---

# dsfhub_cloud_account (Resource)

Provides a cloud account resource.

The `dsfhub_cloud_account` resource contains the configuration parameters necessary to access [Cloud Accounts](https://docs.imperva.com/bundle/v4.13-sonar-user-guide/page/80357.htm) 
(AWS, GCP, Azure, etc), from the Unified Settings Console in the DSF HUB platform. 
Documentation for the underlying API used in this resource can be found at
[Cloud Account API Definition page](https://docs.imperva.com/bundle/v4.13-sonar-user-guide/page/84552.htm).

## Example Usage

```hcl
resource "dsfhub_cloud_account" "example_aws_cloud_account" {
  server_type = "AWS"
  admin_email = "your@email.com"
  asset_display_name = "arn:partition:service:region:account-id" # User-friendly name of the asset
  asset_id = "arn:partition:service:region:account-id" # Also populates arn field for aws
  gateway_id = "12345-abcde-12345-abcde-12345-abcde"
  asset_connection {
    auth_mechanism = "iam_role"
    reason = "default"
    region = "us-east-1"
  }
}
```

## Cloud Account Types:
<ul>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/cloud_accounts/aws.md>Amazon Account</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/cloud_accounts/alibaba.md>Alibaba Cloud Account</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/cloud_accounts/azure.md>Azure</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/cloud_accounts/gcp.md>Google Cloud</a></li>
</ul>

## Argument Reference

### Required

- `admin_email` (String) The email address to notify about this asset
- `asset_connection` (Block Set, Min: 1) N/A (see [below for nested schema](#nestedblock--asset_connection))
- `asset_id` (String) (String) The unique identifier or resource name of the asset.
- `gateway_id` (String) The jsonarUid unique identifier of the agentless gateway. Example: '7a4af7cf-4292-89d9-46ec-183756ksdjd'
- `server_type` (String) The type of server or data service to be created as a data source. The list of available data sources is documented at: https://docs.imperva.com/bundle/v4.11-sonar-user-guide/page/84552.htm

### Optional

- `asset_display_name` (String) User-friendly name of the asset, defined by user.
- `asset_source` (String) The source platform/vendor/system of the asset data. Usually the service responsible for creating that asset document
- `available_regions` (String) A list of regions to use in discovery actions that iterate through region
- `aws_proxy_config` (Block Set) AWS specific proxy configuration (see [below for nested schema](#nestedblock--aws_proxy_config))
- `credentials_endpoint` (String) A specific sts endpoint to use
- `criticality` (Number) The asset's importance to the business. These values are measured on a scale from "Most critical" (1) to "Least critical" (4). Allowed values: 1, 2, 3, 4
- `jsonar_uid` (String) Unique identifier (UID) attached to the Sonar machine controlling the asset
- `location` (String) Current human-readable description of the physical location of the asset, or region.
- `managed_by` (String) Email of the person who maintains the asset; can be different from the owner specified in the owned_by field. Defaults to admin_email.
- `owned_by` (String) Email of Owner / person responsible for the asset; can be different from the person in the managed_by field. Defaults to admin_email.
- `proxy` (String) Proxy to use for AWS calls if aws_proxy_config is populated the proxy field will get populated from the http value there
- `region` (String) For cloud systems with regions, the default region or region used with this asset
- `server_host_name` (String) Hostname (or IP if name is unknown)
- `server_port` (String)
- `service_endpoints` (Block Set) Specify particular endpoints for a given service in the form of <service name>: "endpoint" (see [below for nested schema](#nestedblock--service_endpoints))
- `used_for` (String) Designates how this asset is used / the environment that the asset is supporting.
- `version` (Number) Denotes the version of the asset

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--asset_connection"></a>
### Nested Schema for `asset_connection`

Required:

- `auth_mechanism` (String) Specifies the auth mechanism used by the connection
- `reason` (String) N/A
- `region` (String) Default AWS region for this asset

Optional:

- `access_id` (String) The Access key ID of AWS secret access key used to authenticate
- `access_key` (String) The Secret access key used to authenticate
- `amazon_secret` (Block Set, Max: 1) Configuration to integrate with AWS Secrets Manager (see [below for nested schema](#nestedblock--asset_connection--amazon_secret))
- `application_id` (String) This is also referred to as the Client ID and it’s the unique identifier for the registered application being used to execute Python SDK commands against Azure’s API services. You can find this number under Azure Active Directory -> App Registrations -> Owned Applications
- `ca_certs_path` (String) Certificate authority certificates path; what location should the sysetm look for certificate information from. Equivalent to --capath in a curl call
- `client_secret` (String) This a string containing a secret used by the application to prove its identity when requesting a token. You can get a secret by going to Azure Active Directory -> App Registrations -> Owned Applications, selecting the desired application and then going to Certificates & secrets -> Client secrets -> + New client secret
- `credential_fields` (Block Set) Document containing values to build a profile from. Filling this will create a profile using the given profile name (see [below for nested schema](#nestedblock--asset_connection--credential_fields))
- `cyberark_secret` (Block Set, Max: 1) Configuration to integrate with AWS Secrets Manager (see [below for nested schema](#nestedblock--asset_connection--cyberark_secret))
- `directory_id` (String) This is also referred to as the Tenant ID and is a GUID representing the Active Directory Tenant. It can be found in the Azure Active Directory page under the Azure portal
- `external_id` (String) External ID to use when assuming a role
- `hashicorp_secret` (Block Set) Configuration to integrate with HashiCorp Vault (see [below for nested schema](#nestedblock--asset_connection--hashicorp_secret))
- `key_file` (String) Location on disk on the key to be used to authenticate
- `role_name` (String) What role is used to get credentials from.
- `secret_key` (String) The Secret access key used to authenticate
- `ssl` (Boolean) If true, use SSL when connecting
- `subscription_id` (String) This is the Azure account subscription ID. You can find this number under the Subscriptions page on the Azure portal
- `username` (String) The name of a profile in /imperva/local/credentials/.aws/credentials to use for authenticating

<a id="nestedblock--asset_connection--amazon_secret"></a>
### Nested Schema for `asset_connection.amazon_secret`

Optional:

- `field_mapping` (Map of String) Field mapping for amazon secret
- `secret_asset_id` (String) Amazon secret asset id
- `secret_name` (String) Amazon secret mane


<a id="nestedblock--asset_connection--credential_fields"></a>
### Nested Schema for `asset_connection.credential_fields`

Optional:

- `credential_source` (String) HashiCorp secret asset id
- `role_arn` (String) HashiCorp secret mane


<a id="nestedblock--asset_connection--cyberark_secret"></a>
### Nested Schema for `asset_connection.cyberark_secret`

Optional:

- `field_mapping` (Map of String) Field mapping for amazon secret
- `secret_asset_id` (String) Amazon secret asset id
- `secret_name` (String) Amazon secret mane


<a id="nestedblock--asset_connection--hashicorp_secret"></a>
### Nested Schema for `asset_connection.hashicorp_secret`

Optional:

- `field_mapping` (Map of String) Field mapping for HashiCorp secret
- `path` (String) HashiCorp secret path
- `secret_asset_id` (String) HashiCorp secret asset id
- `secret_name` (String) HashiCorp secret mane


<a id="nestedblock--aws_proxy_config"></a>
### Nested Schema for `aws_proxy_config`

Optional:

- `http` (String) HTTP endpoint for aws proxy config
- `https` (String) HTTPS endpoint for aws proxy config


<a id="nestedblock--service_endpoints"></a>
### Nested Schema for `service_endpoints`

Optional:

- `logs` (String) The log endpoint for a given service

## Import

Cloud Account can be imported using the `asset_id`, e.g.:

```
$ terraform import dsf_cloud_account.example_aws_cloud_account "arn:partition:service:region:account-id"
```
