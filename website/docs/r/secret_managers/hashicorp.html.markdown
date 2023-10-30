---
layout: "dsfhub"
page_title: "dsfhub_secrets_manager Resource"
sidebar_current: "docs-dsfhub-resource-dsfhub-secrets-manager"
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


