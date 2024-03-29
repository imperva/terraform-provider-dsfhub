---
subcategory: "Provider Reference"
layout: "dsfhub"
page_title: "DSFHUB Secret Manager Resource"
description: |-
  Provides a dsfhub_secret_manager resource.
---

# dsfhub_secret_manager (Resource)

Provides a secret manager resource.  A secret manager is a service used to store configuration information such as 
database passwords, API keys, or TLS certificates needed by an application at runtime.

The `dsfhub_secret_manager` resource contains the configuration parameters necessary to onboard a secret manager.  
When configuring connections for assets (including data sources, cloud accounts, log aggregators and secret managers) 
in the DSF HUB, you have the option of using a secret manager as the source for fields such as user credentials or 
other secrets. If you use this option, certain configuration fields for a connection can be populated from an onboarded 
secret manager rather than from the configured fields in the asset itself.  Documentation for the underlying API used in 
this resource can be found at [Secret Managers API Definition page](https://docs.imperva.com/bundle/v4.13-sonar-user-guide/page/84552.htm).

## Example Usage

```hcl
# Example generic variable reference:
variable "admin_email" {
  default = "your@email.com"
}
variable "gateway_id" {
  default = "7a4af7cf-4292-89d9-46ec-183756ksdjd"
}

# Example dsfhub_secret_manager specific variables for HASHICORP
variable "secret_manager_hashicorp_asset_display_name" {
  default = "arn:partition:service:region:account-id"
}
variable "secret_manager_hashicorp_asset_id" {
  default = "arn:partition:service:region:account-id"
}
variable "secret_manager_hashicorp_server_host_name" {
  default = "your-data-source-asset-id-here"
}
variable "secret_manager_hashicorp_server_ip" {
  default = "1.2.3.4"
}
variable "secret_manager_hashicorp_server_port" {
  default = 1234
}
variable "secret_manager_hashicorp_access_id" {
  default = "ABCDE12345ABCDE12345"
}
variable "secret_manager_hashicorp_aws_iam_server_id" {
  default = "your.vault.example.com"
}
variable "secret_manager_hashicorp_role_name" {
  default = "your-role-name-here"
}
variable "secret_manager_hashicorp_secret_key" {
  default = "A1b2C3d4/H5i6J7k8/A1b2C3d4/H5i6J7k8"
}

variable "hashicorp_field_mapping" {
  description = "Object listing the mappings between DSF fields and their corresponding locations with the hashicorp secret manager."
  type = map(string)
  default = {
    username = "username_field"
    password = "password_field"
  }
}

variable "hashicorp_secret_name" {
  description = "The secret name created in the KV secrets engine."
  type = string
}

variable "hashicorp_secret_path" {
  description = "The path to where you enabled a KV secrets engine on your hashcorp server."
  type = string
}


# Example dsfhub_secret_manager usage for HASHICORP
resource "dsfhub_secret_manager" "example_hashicorp" {
  server_type = "HASHICORP"
  admin_email = var.admin_email	# The email address to notify about this asset
  asset_display_name = var.secret_manager_hashicorp_asset_display_name # User-friendly name of the asset, defined by user.
  asset_id = var.secret_manager_hashicorp_asset_id # The unique identifier or resource name of the asset. For AWS, use arn, for Azure, use subscription ID, for GCP, use project ID
  gateway_id = var.gateway_id # The jsonarUid unique identifier of the agentless gateway. Example: '7a4af7cf-4292-89d9-46ec-183756ksdjd'
  server_host_name = var.secret_manager_hashicorp_server_host_name# Hostname (or IP if name is unknown)
  server_ip = var.secret_manager_hashicorp_server_ip	# IP address of the service where this asset is located. If no IP is available populate this field with other information that would identify the system e.g. hostname or AWS ARN, etc.
  server_port = var.secret_manager_hashicorp_server_port	# Port used by the source server
  asset_connection {
    auth_mechanism = "iam_role"
    access_id = var.secret_manager_hashicorp_access_id # The Access key ID of AWS secret access key used to authenticate
    aws_iam_server_id = var.secret_manager_hashicorp_aws_iam_server_id
    reason = "default" # Used to differentiate connections if multiple connections exist for this asset"
    role_name = var.secret_manager_hashicorp_role_name # AWS role name to use for authentication
    secret_key = var.secret_manager_hashicorp_secret_key # The AWS secret access key used to authenticate
  }
}

resource "dsfhub_data_source" "example_sybase_kerberos" {
  server_type = "SYBASE"

  admin_email        = var.admin_email
  asset_display_name = "example-sybase-kerberos-asset"
  asset_id           = "${var.sybase_server_host_name}:SYBASE::${var.sybase_server_port}"
  database_name      = "example-sybase-db-name"
  gateway_id         = var.gateway_id
  server_host_name   = "example-server-host-name"
  server_ip          = "example-server-host-ip"
  server_port        = "5000"
  version            = "16"

  audit_pull_enabled = true

  asset_connection {
    auth_mechanism = "kerberos"

    reason       = "default"
    external     = false
    kerberos_kdc = var.kerberos_kdc
    kerberos_spn = var.kerberos_spn
    password     = "dummy_password_val"
    principal    = "dummy_principal_val"

    hashicorp_secret {
      secret_asset_id = dsfhub_secret_manager.example_hashicorp.asset_id
      path            = var.hashicorp_secret_path
      secret_name     = var.hashicorp_secret_name

      field_mapping = var.hashicorp_field_mapping
    }
  }
}
```

## Secret Manager Types:
<ul>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/secrets_managers/aws.md">Amazon Account</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/secrets_managers/cyberark.md">CyberArk</a></li>
	<li><a href="https://github.com/imperva/terraform-provider-dsfhub/tree/main/examples/secrets_managers/hashicorp.md">HashiCorp</a></li>
</ul>


## Argument Reference

### Required

- `admin_email` (String) The email address to notify about this asset
- `gateway_id` (String) The jsonarUid unique identifier of the agentless gateway. Example: '7a4af7cf-4292-89d9-46ec-183756ksdjd'
- `server_host_name` (String) Hostname (or IP if name is unknown)
- `server_type` (String) The type of server or data service to be created as a data source. The list of available data sources is documented at: https://docs.imperva.com/bundle/v4.11-sonar-user-guide/page/84552.htm

### Optional

- `application` (String) The Asset ID of the application asset that "owns" the asset.
- `asset_connection` (Block Set) N/A (see [below for nested schema](#nestedblock--asset_connection))
- `asset_display_name` (String) User-friendly name of the asset, defined by user.
- `asset_id` (String) (String) The unique identifier or resource name of the asset.
- `asset_source` (String) The source platform/vendor/system of the asset data. Usually the service responsible for creating that asset document
- `available_regions` (String) A list of regions to use in discovery actions that iterate through region
- `aws_proxy_config` (Block Set) AWS specific proxy configuration (see [below for nested schema](#nestedblock--aws_proxy_config))
- `credentials_endpoint` (String) A specific sts endpoint to use
- `criticality` (Number) The asset's importance to the business. These values are measured on a scale from "Most critical" (1) to "Least critical" (4). Allowed values: 1, 2, 3, 4
- `jsonar_uid` (String) Unique identifier (UID) attached to the Sonar machine controlling the asset
- `jsonar_uid_display_name` (String) Unique identifier (UID) attached to the Sonar machine controlling the asset
- `location` (String) Current human-readable description of the physical location of the asset, or region.
- `managed_by` (String) Email of the person who maintains the asset; can be different from the owner specified in the owned_by field. Defaults to admin_email.
- `owned_by` (String) Email of Owner / person responsible for the asset; can be different from the person in the managed_by field. Defaults to admin_email.
- `proxy` (String) Proxy to use for AWS calls if aws_proxy_config is populated the proxy field will get populated from the http value there
- `region` (String) For cloud systems with regions, the default region or region used with this asset
- `server_ip` (String) IP address of the service where this asset is located. If no IP is available populate this field with other information that would identify the system e.g. hostname or AWS ARN, etc.
- `server_port` (String) Port used by the source server
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

Optional:

- `access_id` (String) The Access key ID of AWS secret access key used to authenticate
- `amazon_secret` (Block Set, Max: 1) Configuration to integrate with AWS Secrets Manager (see [below for nested schema](#nestedblock--asset_connection--amazon_secret))
- `aws_iam_server_id` (String) e.g. vault.example.com
- `ca_certs_path` (String) Certificate authority certificates path; what location should the sysetm look for certificate information from. Equivalent to --capath in a curl call
- `cert_file` (String)
- `credential_expiry` (String)
- `credential_fields` (Block Set) Document containing values to build a profile from. Filling this will create a profile using the given profile name (see [below for nested schema](#nestedblock--asset_connection--credential_fields))
- `cyberark_secret` (Block Set, Max: 1) Configuration to integrate with AWS Secrets Manager (see [below for nested schema](#nestedblock--asset_connection--cyberark_secret))
- `external_id` (String) External ID to use when assuming a role
- `hashicorp_secret` (Block Set) Configuration to integrate with HashiCorp Vault (see [below for nested schema](#nestedblock--asset_connection--hashicorp_secret))
- `key_file` (String)
- `nonce` (String)
- `protocol` (String)
- `query` (String)
- `region` (String) Default AWS region for this asset
- `role_name` (String) Role to use for authentication
- `secret_key` (String) The Secret access key used to authenticate
- `self_signed` (String)
- `ssl` (Boolean) If true, use SSL when connecting
- `store_aws_credentials` (String)
- `username` (String) The name of a profile in /imperva/local/credentials/.aws/credentials to use for authenticating
- `v2_key_engine` (String) Use a KV2 secret engine

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

Secret Manager can be imported using the `asset_id`, e.g.:

```
$ terraform import dsf_secret_manager.example_secret_manager_aws_log_group "arn:partition:service:region:account-id"
```
