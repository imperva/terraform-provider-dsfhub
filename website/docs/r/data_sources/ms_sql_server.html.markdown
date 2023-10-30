---
layout: "dsfhub"
page_title: "dsfhub_data_source Resource"
sidebar_current: "docs-dsfhub-resource-dsfhub-data-source"
description: |-
Provides a dsfhub_data_source resource for MS SQL SERVER.
---

# dsfhub_data_source (Resource)

Provides a dsfhub_data_source resource for MS SQL SERVER.

## Example usage:

```hcl
# ### DSF Provider ###
provider "dsfhub" {
	dsfhub_token = var.dsfhub_token # TF_VAR_dsfhub_token env variable
	dsfhub_host = var.dsfhub_host # TF_VAR_dsfhub_host env variable
	#insecure_ssl = false
}

# ### Resource example for MS SQL SERVER ###
resource "dsfhub_data_source" "example_ms_sql_server" {
	server_type = "MS SQL SERVER"
	# ### required ### 
	admin_email = var.admin_email	# The email address to notify about this asset
	asset_display_name = var.asset_display_name	# User-friendly name of the asset, defined by user.
	asset_id = var.asset_id	# Asset ID
	gateway_id = var.gateway_id	# Gateway ID
	server_host_name = var.server_host_name	# Hostname (or IP if name is unknown)
	server_ip = var.server_ip	# IP address of the service where this asset is located. If no IP is available populate this field with other information that would identify the system e.g. hostname or AWS ARN, etc.

	# ### optional ### 
	# application = var.application	# The Asset ID of the application asset that \"owns\" the asset.
	# archive = var.archive	# If True archive files in the asset after being processed by sonargd. Defaults to True if field isn't present
	# asset_source = var.asset_source	# The source platform/vendor/system of the asset data. Usually the service responsible for creating that asset document
	# audit_info = var.audit_info	# Normally auto-populated when enabling the audit policy, it is a sub-document in JSON format containing configuration information for audit management. See documentation for values that can be added manually depending on asset type. Editing this value does NOT enable the audit policy.
	# audit_pull_enabled = var.audit_pull_enabled	# If true, sonargateway will collect the audit logs for this system if it can.
	# audit_type = var.audit_type # Example Values: "SSH"	# Used to indicate what mechanism should be used to fetch logs on systems supporting multiple ways to get logs, see asset specific documentation for details
	# criticality = var.criticality # Example Values: "1", "2", "3", "4"	# The asset's importance to the business. These values are measured on a scale from \"Most critical\" (1) to \"Least critical\" (4). Allowed values: 1, 2, 3, 4
	# database_name = var.database_name	# Specifies the name of the database (or default DB) to connect to.
	# db_engine = var.db_engine	# Specifies the version of the engine being used by the database (e.g. oracle-ee, oracle-se, oracle-se1, oracle-se2)
	# enable_audit_management = var.enable_audit_management	# If true, Sonar is responsible for setting and updating the policies
	# enable_audit_monitoring = var.enable_audit_monitoring	# If true, Sonar sends emails/alerts when the audit policies change.
	# entitlement_enabled = var.entitlement_enabled	# If true, Entitlement Management system is enabled.
	# ignore_latest_of = var.ignore_latest_of	# A regex defining a group. From all the files with the same group, the latest one will be ignored, so that it isn't archived until server is done writing
	# jsonar_uid = var.jsonar_uid	# Unique identifier (UID) attached to the Sonar machine controlling the asset
	# location = var.location	# Current human-readable description of the physical location of the asset, or region.
	# logs_destination_asset_id = var.logs_destination_asset_id	# The asset name of the log aggregator that stores this asset's logs.
	# managed_by = var.managed_by	# Email of the person who maintains the asset; can be different from the owner specified in the owned_by field. Defaults to admin_email.
	# max_concurrent_conn = var.max_concurrent_conn	# Maximum number of concurrent connections that sensitive data management should use at once.
	# owned_by = var.owned_by	# Email of Owner / person responsible for the asset; can be different from the person in the managed_by field. Defaults to admin_email.
	# sdm_enabled = var.sdm_enabled	# Sensitive data management (SDM) is enabled if this parameter is set to True.
	# server_port = "1433"	# Port used when connecting to this asset
	# service_name = "MS SQL SERVER"	# 
	# used_for = var.used_for # Example Values: "Production", "Test", "Development", "Demonstration", "QA", "Staging", "Training", "Disaster Recovery"	# Designates how this asset is used / the environment that the asset is supporting.
	# version = var.version	# Denotes the version of the asset
	# xel_directory = var.xel_directory	# Absolute path of the XEL files location
	asset_connection {
		auth_mechanism = "password"
		# ### required ### 
		password = null # password description: "The password of the user being used to authenticate"
		reason = null # Example Values: "default", "sonargateway", "SDM", "audit management", "ad-hoc-query" # reason description: "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset"
		username = null # username description: "The user to use when connecting"
		# ### optional ### 
		# amazon_secret = null # amazon_secret description: "Configuration to integrate with AWS Secrets Manager"
		# autocommit = null # autocommit description: "If true, Commit automatically don't wait for transaction to be explicitly committed"
		# cyberark_secret = null # cyberark_secret description: "Configuration to integrate with CyberArk Vault"
		# driver = null # driver description: "A path to a non-default driver location. If populated this driver will be used rather than the default"
		# hashicorp_secret = null # hashicorp_secret description: "Configuration to integrate with HashiCorp Vault"
		# jdbc_ssl_trust_server_certificate = false # jdbc_ssl_trust_server_certificate description: ""
		# jdbc_ssl_trust_store_location = null # jdbc_ssl_trust_store_location description: ""
		# jdbc_ssl_trust_store_password = null # jdbc_ssl_trust_store_password description: ""
		# odbc_connection_string = null # odbc_connection_string description: "Additional ODBC connection string parameters. This string will get added to the connection string"
		# ssl = null # ssl description: "If true, use SSL when connecting"
		# ssl_server_cert = null # ssl_server_cert description: "Path to server certificate to use during authentication"
	}
}
```


