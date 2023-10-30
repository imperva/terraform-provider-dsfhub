---
layout: "dsfhub"
page_title: "dsfhub_data_source Resource"
sidebar_current: "docs-dsfhub-resource-dsfhub-data-source"
description: |-
Provides a dsfhub_data_source resource for HIVE.
---

# dsfhub_data_source (Resource)

Provides a dsfhub_data_source resource for HIVE.

## Example usage:

```hcl
# ### DSF Provider ###
provider "dsfhub" {
	dsfhub_token = var.dsfhub_token # TF_VAR_dsfhub_token env variable
	dsfhub_host = var.dsfhub_host # TF_VAR_dsfhub_host env variable
	#insecure_ssl = false
}

# ### Resource example for HIVE ###
resource "dsfhub_data_source" "example_hive" {
	server_type = "HIVE"
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
	# audit_info = var.audit_info	# Normally auto-populated when enabling the audit policy, it is a sub-document in JSON format containing configuration information for audit management. See documentation for values that can be added manually depending on asset type. Editing this value does NOT enable the audit policy.
	# audit_pull_enabled = var.audit_pull_enabled	# If true, sonargateway will collect the audit logs for this system if it can.
	# content_type = var.content_type	# Content type should be set to the desired <'parent' asset 'Server Type'>, which is the one that uses this asset as a destination for logs. NOTE: The content_type field will take precedence on the lookup for parent_asset_id field when checking which server is sending logs to this asset.
	# criticality = var.criticality # Example Values: "1", "2", "3", "4"	# The asset's importance to the business. These values are measured on a scale from \"Most critical\" (1) to \"Least critical\" (4). Allowed values: 1, 2, 3, 4
	# db_engine = var.db_engine	# Specifies the version of the engine being used by the database (e.g. oracle-ee, oracle-se, oracle-se1, oracle-se2)
	# enable_audit_management = var.enable_audit_management	# If true, Sonar is responsible for setting and updating the policies
	# enable_audit_monitoring = var.enable_audit_monitoring	# If true, Sonar sends emails/alerts when the audit policies change.
	# entitlement_enabled = var.entitlement_enabled	# If true, Entitlement Management system is enabled.
	# gateway_service = var.gateway_service	# The name of the gateway pull service (if any) used to retrieve logs for this source. Usually set by the connect gateway playbook.
	# jsonar_uid = var.jsonar_uid	# Unique identifier (UID) attached to the Sonar machine controlling the asset
	# location = var.location	# Current human-readable description of the physical location of the asset, or region.
	# managed_by = var.managed_by	# Email of the person who maintains the asset; can be different from the owner specified in the owned_by field. Defaults to admin_email.
	# max_concurrent_conn = var.max_concurrent_conn	# Maximum number of concurrent connections that sensitive data management should use at once.
	# owned_by = var.owned_by	# Email of Owner / person responsible for the asset; can be different from the person in the managed_by field. Defaults to admin_email.
	# parent_asset_id = var.parent_asset_id	# The name of an asset that this asset is part of (/related to). E.g. an AWS resource will generally have an AWS account asset as its parent. Also used to connect some log aggregating asset with the sources of their logs. E.g. An AWS LOG GROUP asset can have an AWS RDS as its parent, indicating that that is the log group for that RDS.
	# region = var.region	# For cloud systems with regions, the default region or region used with this asset
	# sdm_enabled = var.sdm_enabled	# Sensitive data management (SDM) is enabled if this parameter is set to True.
	# used_for = var.used_for # Example Values: "Production", "Test", "Development", "Demonstration", "QA", "Staging", "Training", "Disaster Recovery"	# Designates how this asset is used / the environment that the asset is supporting.
	# version = var.version	# Denotes the version of the asset
	asset_connection {
		auth_mechanism = "password"
		# ### required ### 
		database_name = null # database_name description: "Specifies the name of the database (or default DB) to connect to."
		password = null # password description: "The password of the user being used to authenticate"
		reason = null # Example Values: "default", "ad-hoc-query" # reason description: "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset"
		username = null # username description: "The user to use when connecting"
		# ### optional ### 
		# amazon_secret = null # amazon_secret description: "Configuration to integrate with AWS Secrets Manager"
		# cyberark_secret = null # cyberark_secret description: "Configuration to integrate with CyberArk Vault"
		# hashicorp_secret = null # hashicorp_secret description: "Configuration to integrate with HashiCorp Vault"
		# hive_server_type = null # hive server type description: ""
		# httppath = null # httppath description: ""
		# jdbc_ssl_trust_server_certificate = false # jdbc_ssl_trust_server_certificate description: ""
		# jdbc_ssl_trust_store_location = null # jdbc_ssl_trust_store_location description: ""
		# jdbc_ssl_trust_store_password = null # jdbc_ssl_trust_store_password description: ""
		# kerberos_host_fqdn = null # kerberos_host_fqdn description: ""
		# ssl = null # ssl description: "If true, use SSL when connecting"
		# ssl_server_cert = null # ssl_server_cert description: "Path to server certificate to use during authentication"
		# thrift_transport = null # thrift transport description: ""
		# transportmode = null # transportmode description: ""
	}
	asset_connection {
		auth_mechanism = "kerberos"
		# ### required ### 
		database_name = null # database_name description: "Specifies the name of the database (or default DB) to connect to."
		reason = null # Example Values: "default", "ad-hoc-query" # reason description: "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset"
		# ### optional ### 
		# amazon_secret = null # amazon_secret description: "Configuration to integrate with AWS Secrets Manager"
		# cache_file = null # cache_file description: ""
		# cyberark_secret = null # cyberark_secret description: "Configuration to integrate with CyberArk Vault"
		# db_role = null # db_role description: ""
		# external = null # external description: ""
		# extra_kinit_parameters = null # extra_kinit_parameters description: ""
		# hashicorp_secret = null # hashicorp_secret description: "Configuration to integrate with HashiCorp Vault"
		# hive_server_type = null # hive server type description: ""
		# httppath = null # httppath description: ""
		# jdbc_ssl_trust_server_certificate = false # jdbc_ssl_trust_server_certificate description: ""
		# jdbc_ssl_trust_store_location = null # jdbc_ssl_trust_store_location description: ""
		# jdbc_ssl_trust_store_password = null # jdbc_ssl_trust_store_password description: ""
		# kerberos_host_fqdn = null # kerberos_host_fqdn description: ""
		# kerberos_kdc = null # kerberos_kdc description: ""
		# kerberos_service_kdc = null # kerberos_service_kdc description: ""
		# kerberos_service_realm = null # kerberos_service_realm description: ""
		# kerberos_spn = null # kerberos_spn description: ""
		# keytab_file = null # keytab_file description: "Specify a non-default keytab location"
		# kinit_program_path = null # kinit_program_path description: ""
		# password = null # password description: ""
		# principal = null # principal description: "The principal used to authenticate"
		# ssl = null # ssl description: "If true, use SSL when connecting"
		# ssl_server_cert = null # ssl_server_cert description: "Path to server certificate to use during authentication"
		# thrift_transport = null # thrift transport description: ""
		# transportmode = null # transportmode description: ""
		# use_keytab = null # use_keytab description: "If true, authenticate using a key tab"
	}
}
```


