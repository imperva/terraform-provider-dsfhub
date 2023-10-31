---
layout: "dsfhub"
page_title: "Example dsfhub_secrets_manager Resource for HASHICORP"
description: |-
Provides a dsfhub_secrets_manager resource for HASHICORP.
---

# dsfhub_secrets_manager (Resource)

Provides a dsfhub_secrets_manager resource for HASHICORP.

## Example usage:

```hcl
# ### DSF Provider ###
provider "dsfhub" {
	dsfhub_token = var.dsfhub_token # TF_VAR_dsfhub_token env variable
	dsfhub_host = var.dsfhub_host # TF_VAR_dsfhub_host env variable
	#insecure_ssl = false
}

# ### Resource example for HASHICORP ###
resource "dsfhub_secret_manager" "example_hashicorp" {
	server_type = "HASHICORP"
	# ### required ### 
	admin_email = var.admin_email	# The email address to notify about this asset
	asset_display_name = var.asset_display_name	# User-friendly name of the asset, defined by user.
	asset_id = var.asset_id	# Asset ID
	gateway_id = var.gateway_id	# Gateway ID
	server_host_name = var.server_host_name	# Hostname (or IP if name is unknown)
	server_ip = var.server_ip	# IP address of the service where this asset is located. If no IP is available populate this field with other information that would identify the system e.g. hostname or AWS ARN, etc.
	server_port = var.server_port	# Port used by the source server

	# ### optional ### 
	# application = var.application	# The Asset ID of the application asset that \"owns\" the asset.
	# asset_source = var.asset_source	# The source platform/vendor/system of the asset data. Usually the service responsible for creating that asset document
	# criticality = var.criticality # Example Values: "1", "2", "3", "4"	# The asset's importance to the business. These values are measured on a scale from \"Most critical\" (1) to \"Least critical\" (4). Allowed values: 1, 2, 3, 4
	# jsonar_uid = var.jsonar_uid	# Unique identifier (UID) attached to the Sonar machine controlling the asset
	# location = var.location	# Current human-readable description of the physical location of the asset, or region.
	# managed_by = var.managed_by	# Email of the person who maintains the asset; can be different from the owner specified in the owned_by field. Defaults to admin_email.
	# owned_by = var.owned_by	# Email of Owner / person responsible for the asset; can be different from the person in the managed_by field. Defaults to admin_email.
	# region = var.region	# For cloud systems with regions, the default region or region used with this asset
	# used_for = var.used_for # Example Values: "Production", "Test", "Development", "Demonstration", "QA", "Staging", "Training", "Disaster Recovery"	# Designates how this asset is used / the environment that the asset is supporting.
	# version = var.version	# Denotes the version of the asset
	asset_connection {
		auth_mechanism = "root_token"
		# ### required ### 
		reason = null # Example Values: "default" # reason description: "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset"
		secret_key = null # secret_key description: "Vault token"
		# ### optional ### 
		# credential_expiry = null # credential_expiry description: ""
		# protocol = null # protocol description: ""
		# self_signed = null # self_signed description: "Accept self signed certificates"
		# ssl = null # ssl description: "If true, use SSL when connecting"
		# store_aws_credentials = null # store_aws_credentials description: ""
		# v2_key_engine = null # v2_key_engine description: "Use a KV2 secret engine"
	}
	asset_connection {
		auth_mechanism = "ec2"
		# ### required ### 
		reason = null # Example Values: "default" # reason description: "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset"
		role_name = null # role_name description: "What Hashicorp role is used in the login step"
		# ### optional ### 
		# credential_expiry = null # credential_expiry description: ""
		# nonce = null # nonce description: ""
		# protocol = null # protocol description: ""
		# self_signed = null # self_signed description: "Accept self signed certificates"
		# ssl = null # ssl description: "If true, use SSL when connecting"
		# store_aws_credentials = null # store_aws_credentials description: ""
		# v2_key_engine = null # v2_key_engine description: "Use a KV2 secret engine"
	}
	asset_connection {
		auth_mechanism = "iam_role"
		# ### required ### 
		access_id = null # access_id description: "AWS access_ID"
		aws_iam_server_id = null # aws_iam_server_id description: "e.g. vault.example.com"
		reason = null # Example Values: "default" # reason description: "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset"
		role_name = null # role_name description: "What Hashicorp role is used in the login step"
		secret_key = null # secret_key description: "AWS secret_key"
		# ### optional ### 
		# credential_expiry = null # credential_expiry description: ""
		# protocol = null # protocol description: ""
		# self_signed = null # self_signed description: "Accept self signed certificates"
		# ssl = null # ssl description: "If true, use SSL when connecting"
		# store_aws_credentials = null # store_aws_credentials description: ""
		# v2_key_engine = null # v2_key_engine description: "Use a KV2 secret engine"
	}
	asset_connection {
		auth_mechanism = "app_role"
		# ### required ### 
		reason = null # Example Values: "default" # reason description: "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset"
		role_name = null # role_name description: "Role to use for authentication"
		secret_key = null # secret_key description: "The Secret access key used to authenticate"
		# ### optional ### 
		# credential_expiry = null # credential_expiry description: ""
		# protocol = null # protocol description: ""
		# self_signed = null # self_signed description: "Accept self signed certificates"
		# ssl = null # ssl description: "If true, use SSL when connecting"
		# store_aws_credentials = null # store_aws_credentials description: ""
		# v2_key_engine = null # v2_key_engine description: "Use a KV2 secret engine"
	}
}
```


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
