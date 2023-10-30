---
layout: "dsfhub"
page_title: "dsfhub_data_source Resource"
sidebar_current: "docs-dsfhub-resource-dsfhub-data-source"
description: |-
Provides a dsfhub_data_source resource for AZURE MYSQL.
---

# dsfhub_data_source (Resource)

Provides a dsfhub_data_source resource for AZURE MYSQL.

## Example usage:

```hcl
# ### DSF Provider ###
provider "dsfhub" {
	dsfhub_token = var.dsfhub_token # TF_VAR_dsfhub_token env variable
	dsfhub_host = var.dsfhub_host # TF_VAR_dsfhub_host env variable
	#insecure_ssl = false
}

# ### Resource example for AZURE MYSQL ###
resource "dsfhub_data_source" "example_azure_mysql" {
	server_type = "AZURE MYSQL"
	# ### required ### 
	admin_email = var.admin_email	# The email address to notify about this asset
	asset_display_name = var.asset_display_name	# User-friendly name of the asset, defined by user.
	asset_id = var.asset_id	# Asset ID
	gateway_id = var.gateway_id	# Gateway ID
	location = var.location	# Current human-readable description of the physical location of the asset, or region.
	server_host_name = var.server_host_name	# Hostname (or IP if name is unknown)

	# ### optional ### 
	# application = var.application	# The Asset ID of the application asset that \"owns\" the asset.
	# asset_source = var.asset_source	# The source platform/vendor/system of the asset data. Usually the service responsible for creating that asset document
	# audit_info = var.audit_info	# Normally auto-populated when enabling the audit policy, it is a sub-document in JSON format containing configuration information for audit management. See documentation for values that can be added manually depending on asset type. Editing this value does NOT enable the audit policy.
	# audit_pull_enabled = var.audit_pull_enabled	# If true, sonargateway will collect the audit logs for this system if it can.
	# criticality = var.criticality # Example Values: "1", "2", "3", "4"	# The asset's importance to the business. These values are measured on a scale from \"Most critical\" (1) to \"Least critical\" (4). Allowed values: 1, 2, 3, 4
	# database_name = var.database_name	# Specifies the name of the database (or default DB) to connect to.
	# db_engine = var.db_engine	# Specifies the version of the engine being used by the database (e.g. oracle-ee, oracle-se, oracle-se1, oracle-se2)
	# enable_audit_management = var.enable_audit_management	# If true, Sonar is responsible for setting and updating the policies
	# enable_audit_monitoring = var.enable_audit_monitoring	# If true, Sonar sends emails/alerts when the audit policies change.
	# entitlement_enabled = var.entitlement_enabled	# If true, Entitlement Management system is enabled.
	# gateway_service = var.gateway_service	# The name of the gateway pull service (if any) used to retrieve logs for this source. Usually set by the connect gateway playbook.
	# jsonar_uid = var.jsonar_uid	# Unique identifier (UID) attached to the Sonar machine controlling the asset
	# logs_destination_asset_id = var.logs_destination_asset_id	# The asset name of the log aggregator that stores this asset's logs.
	# managed_by = var.managed_by	# Email of the person who maintains the asset; can be different from the owner specified in the owned_by field. Defaults to admin_email.
	# owned_by = var.owned_by	# Email of Owner / person responsible for the asset; can be different from the person in the managed_by field. Defaults to admin_email.
	# parent_asset_id = var.parent_asset_id	# The name of an asset that this asset is part of (/related to). E.g. an AWS resource will generally have an AWS account asset as its parent. Also used to connect some log aggregating asset with the sources of their logs. E.g. An AWS LOG GROUP asset can have an AWS RDS as its parent, indicating that that is the log group for that RDS.
	# server_port = "3306"	# 
	# subscription_id = var.subscription_id	# This is the Azure account subscription ID. You can find this number under the Subscriptions page on the Azure portal
	# used_for = var.used_for # Example Values: "Production", "Test", "Development", "Demonstration", "QA", "Staging", "Training", "Disaster Recovery"	# Designates how this asset is used / the environment that the asset is supporting.
	# version = var.version	# Denotes the version of the asset
	asset_connection {
		auth_mechanism = "password"
		# ### required ### 
		password = null # password description: "The password of the user being used to authenticate"
		reason = null # Example Values: "default", "sonargateway", "audit management", "ad-hoc-query" # reason description: "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset"
		username = null # username description: "The user to use when connecting"
		# ### optional ### 
		# amazon_secret = null # amazon_secret description: "Configuration to integrate with AWS Secrets Manager"
		# cyberark_secret = null # cyberark_secret description: "Configuration to integrate with CyberArk Vault"
		# hashicorp_secret = null # hashicorp_secret description: "Configuration to integrate with HashiCorp Vault"
		# key_file = null # key_file description: "Private key file name, as specified in Azure SSL documentation"
		# passphrase = null # passphrase description: "Passphrase for the private key."
		# ssl = null # ssl description: "If true, use SSL when connecting"
	}
}
```


