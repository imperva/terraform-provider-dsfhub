---
layout: "dsfhub"
page_title: "dsfhub_secrets_manager Resource"
sidebar_current: "docs-dsfhub-resource-dsfhub-secrets-manager"
description: |-
Provides a dsfhub_secrets_manager resource for AWS.
---

# dsfhub_secrets_manager (Resource)

Provides a dsfhub_secrets_manager resource for AWS.

## Example usage:

```hcl
# ### DSF Provider ###
provider "dsfhub" {
	dsfhub_token = var.dsfhub_token # TF_VAR_dsfhub_token env variable
	dsfhub_host = var.dsfhub_host # TF_VAR_dsfhub_host env variable
	#insecure_ssl = false
}

# ### Resource example for AWS ###
resource "dsfhub_secret_manager" "example_aws" {
	server_type = "AWS"
	# ### required ### 
	admin_email = var.admin_email	# The email address to notify about this asset
	asset_display_name = var.asset_display_name	# User-friendly name of the asset, defined by user.
	asset_id = var.asset_id	# Asset ID
	gateway_id = var.gateway_id	# Gateway ID

	# ### optional ### 
	# application = var.application	# The Asset ID of the application asset that \"owns\" the asset.
	# asset_source = var.asset_source	# The source platform/vendor/system of the asset data. Usually the service responsible for creating that asset document
	# available_regions = var.available_regions	# A list of regions to use in discovery actions that iterate through region
	# aws_proxy_config = var.aws_proxy_config	# AWS specific proxy configuration
	# credentials_endpoint = var.credentials_endpoint	# A specific sts endpoint to use
	# criticality = var.criticality # Example Values: "1", "2", "3", "4"	# The asset's importance to the business. These values are measured on a scale from \"Most critical\" (1) to \"Least critical\" (4). Allowed values: 1, 2, 3, 4
	# jsonar_uid = var.jsonar_uid	# Unique identifier (UID) attached to the Sonar machine controlling the asset
	# location = var.location	# Current human-readable description of the physical location of the asset, or region.
	# managed_by = var.managed_by	# Email of the person who maintains the asset; can be different from the owner specified in the owned_by field. Defaults to admin_email.
	# owned_by = var.owned_by	# Email of Owner / person responsible for the asset; can be different from the person in the managed_by field. Defaults to admin_email.
	# proxy = var.proxy	# Proxy to use for AWS calls if aws_proxy_config is populated the proxy field will get populated from the http value there
	# service_endpoints = var.service_endpoints	# Specify particular endpoints for a given service in the form of <service name>: \"endpoint\"
	# used_for = var.used_for # Example Values: "Production", "Test", "Development", "Demonstration", "QA", "Staging", "Training", "Disaster Recovery"	# Designates how this asset is used / the environment that the asset is supporting.
	asset_connection {
		auth_mechanism = "profile"
		# ### required ### 
		reason = null # Example Values: "default" # reason description: "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset"
		region = null # region description: "Default AWS region for this asset"
		username = null # username description: "The name of a profile in /imperva/local/credentials/.aws/credentials to use for authenticating"
		# ### optional ### 
		# amazon_secret = null # amazon_secret description: "Configuration to integrate with AWS Secrets Manager"
		# ca_certs_path = null # ca_certs_path description: "Certificate authority certificates path; what location should the sysetm look for certificate information from. Equivalent to --capath in a curl call"
		# credential_fields = null # credential_fields description: "Document containing values to build a profile from. Filling this will create a profile using the given profile name"
		# cyberark_secret = null # cyberark_secret description: "Configuration to integrate with CyberArk Vault"
		# hashicorp_secret = null # hashicorp_secret description: "Configuration to integrate with HashiCorp Vault"
		# ssl = null # ssl description: "If true, use SSL when connecting"
	}
	asset_connection {
		auth_mechanism = "key"
		# ### required ### 
		access_id = null # access_id description: "The Access key ID of AWS secret access key used to authenticate"
		reason = null # Example Values: "default" # reason description: "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset"
		region = null # region description: "Default AWS region for this asset"
		secret_key = null # secret_key description: "The Secret access key used to authenticate"
		# ### optional ### 
		# amazon_secret = null # amazon_secret description: "Configuration to integrate with AWS Secrets Manager"
		# ca_certs_path = null # ca_certs_path description: "Certificate authority certificates path; what location should the sysetm look for certificate information from. Equivalent to --capath in a curl call"
		# cyberark_secret = null # cyberark_secret description: "Configuration to integrate with CyberArk Vault"
		# hashicorp_secret = null # hashicorp_secret description: "Configuration to integrate with HashiCorp Vault"
		# ssl = null # ssl description: "If true, use SSL when connecting"
	}
	asset_connection {
		auth_mechanism = "iam_role"
		# ### required ### 
		reason = null # Example Values: "default" # reason description: "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset"
		region = null # region description: "Default AWS region for this asset"
		# ### optional ### 
		# amazon_secret = null # amazon_secret description: "Configuration to integrate with AWS Secrets Manager"
		# cyberark_secret = null # cyberark_secret description: "Configuration to integrate with CyberArk Vault"
		# external_id = null # external_id description: "External ID to use when assuming a role"
		# hashicorp_secret = null # hashicorp_secret description: "Configuration to integrate with HashiCorp Vault"
		# ssl = null # ssl description: "If true, use SSL when connecting"
	}
	asset_connection {
		auth_mechanism = "default"
		# ### required ### 
		reason = null # Example Values: "default" # reason description: "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset"
		region = null # region description: "Default AWS region for this asset"
		# ### optional ### 
		# amazon_secret = null # amazon_secret description: "Configuration to integrate with AWS Secrets Manager"
		# cyberark_secret = null # cyberark_secret description: "Configuration to integrate with CyberArk Vault"
		# hashicorp_secret = null # hashicorp_secret description: "Configuration to integrate with HashiCorp Vault"
		# ssl = null # ssl description: "If true, use SSL when connecting"
	}
}
```


