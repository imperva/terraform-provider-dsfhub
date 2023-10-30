---
layout: "dsfhub"
page_title: "dsfhub_data_source Resource"
sidebar_current: "docs-dsfhub-resource-dsfhub-data-source"
description: |-
Provides a dsfhub_data_source resource for AWS DOCUMENTDB CLUSTER.
---

# dsfhub_data_source (Resource)

Provides a dsfhub_data_source resource for AWS DOCUMENTDB CLUSTER.

## Example usage:

```hcl
# ### DSF Provider ###
provider "dsfhub" {
	dsfhub_token = var.dsfhub_token # TF_VAR_dsfhub_token env variable
	dsfhub_host = var.dsfhub_host # TF_VAR_dsfhub_host env variable
	#insecure_ssl = false
}

# ### Resource example for AWS DOCUMENTDB CLUSTER ###
resource "dsfhub_data_source" "example_aws_documentdb_cluster" {
	server_type = "AWS DOCUMENTDB CLUSTER"
	# ### required ### 
	admin_email = var.admin_email	# The email address to notify about this asset
	asset_display_name = var.asset_display_name	# User-friendly name of the asset, defined by user.
	asset_id = var.asset_id	# Asset ID
	gateway_id = var.gateway_id	# Gateway ID
	server_host_name = var.server_host_name	# Hostname (or IP if name is unknown)
	server_ip = var.server_ip	# IP address of the service where this asset is located. If no IP is available populate this field with other information that would identify the system e.g. hostname or AWS ARN, etc.

	# ### optional ### 
	# application = var.application	# The Asset ID of the application asset that \"owns\" the asset.
	# asset_source = var.asset_source	# The source platform/vendor/system of the asset data. Usually the service responsible for creating that asset document
	# audit_info = var.audit_info	# Normally auto-populated when enabling the audit policy, it is a sub-document in JSON format containing configuration information for audit management. See documentation for values that can be added manually depending on asset type. Editing this value does NOT enable the audit policy.
	# audit_pull_enabled = var.audit_pull_enabled	# If true, sonargateway will collect the audit logs for this system if it can.
	# audit_type = var.audit_type # Example Values: "LOG_GROUP"	# Used to indicate what mechanism should be used to fetch logs on systems supporting multiple ways to get logs, see asset specific documentation for details
	# cluster_id = var.cluster_id	# 
	# cluster_member_id = var.cluster_member_id	# 
	# cluster_name = var.cluster_name	# 
	# credentials_endpoint = var.credentials_endpoint	# A specific sts endpoint to use
	# criticality = var.criticality # Example Values: "1", "2", "3", "4"	# The asset's importance to the business. These values are measured on a scale from \"Most critical\" (1) to \"Least critical\" (4). Allowed values: 1, 2, 3, 4
	# database_name = "admin"	# Specifies the name of the database (or default DB) to connect to.
	# db_engine = var.db_engine	# Specifies the version of the engine being used by the database (e.g. oracle-ee, oracle-se, oracle-se1, oracle-se2)
	# enable_audit_management = var.enable_audit_management	# If true, Sonar is responsible for setting and updating the policies
	# enable_audit_monitoring = var.enable_audit_monitoring	# If true, Sonar sends emails/alerts when the audit policies change.
	# entitlement_enabled = var.entitlement_enabled	# If true, Entitlement Management system is enabled.
	# is_cluster = var.is_cluster	# 
	# jsonar_uid = var.jsonar_uid	# Unique identifier (UID) attached to the Sonar machine controlling the asset
	# location = var.location	# Current human-readable description of the physical location of the asset, or region.
	# logs_destination_asset_id = var.logs_destination_asset_id	# The asset name of the log aggregator that stores this asset's logs.
	# managed_by = var.managed_by	# Email of the person who maintains the asset; can be different from the owner specified in the owned_by field. Defaults to admin_email.
	# max_concurrent_conn = var.max_concurrent_conn	# Maximum number of concurrent connections that sensitive data management should use at once.
	# owned_by = var.owned_by	# Email of Owner / person responsible for the asset; can be different from the person in the managed_by field. Defaults to admin_email.
	# parent_asset_id = var.parent_asset_id	# The name of an asset that this asset is part of (/related to). E.g. an AWS resource will generally have an AWS account asset as its parent. Also used to connect some log aggregating asset with the sources of their logs. E.g. An AWS LOG GROUP asset can have an AWS RDS as its parent, indicating that that is the log group for that RDS.
	# proxy = var.proxy	# Proxy to use for AWS calls if aws_proxy_config is populated the proxy field will get populated from the http value there
	# region = var.region	# For cloud systems with regions, the default region or region used with this asset
	# sdm_enabled = var.sdm_enabled	# Sensitive data management (SDM) is enabled if this parameter is set to True.
	# server_port = "27017"	# 
	# service_name = var.service_name	# Authentication database, usually admin
	# service_endpoint = var.service_endpoint	# Specify a particular endpoint for a given service
	# used_for = var.used_for # Example Values: "Production", "Test", "Development", "Demonstration", "QA", "Staging", "Training", "Disaster Recovery"	# Designates how this asset is used / the environment that the asset is supporting.
	# version = var.version	# Denotes the version of the asset
	# virtual_hostname = var.virtual_hostname	# 
	# virtual_ip = var.virtual_ip	# 
	asset_connection {
		auth_mechanism = "password"
		# ### required ### 
		password = null # password description: "The password of the user being used to authenticate"
		reason = null # Example Values: "default", "sonargateway", "SDM", "audit management", "ad-hoc-query" # reason description: "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset"
		username = null # username description: ""
		# ### optional ### 
		# amazon_secret = null # amazon_secret description: "Configuration to integrate with AWS Secrets Manager"
		# ca_file = null # ca_file description: "Use the specified certificate file to verify the peer. The file may contain multiple CA certificates."
		# cert_file = null # cert_file description: "Use the specified client certificate file when getting a file with HTTPS, FTPS or another SSL-based protocol."
		# cyberark_secret = null # cyberark_secret description: "Configuration to integrate with CyberArk Vault"
		# dns_srv = null # dns srv description: ""
		# hashicorp_secret = null # hashicorp_secret description: "Configuration to integrate with HashiCorp Vault"
		# key_file = null # key_file description: "Private key file name. Allows you to provide your private key in this separate file."
		# passphrase = null # passphrase description: "Passphrase for the private key."
		# replica_set = null # replica_set description: ""
		# self_signed = null # self_signed description: "Accept self-signed certificates"
		# ssl = null # ssl description: "If true, use SSL when connecting"
	}
	asset_connection {
		auth_mechanism = "key_file"
		# ### required ### 
		key_file = null # key_file description: "Private key file name. Allows you to provide your private key in this separate file."
		reason = null # Example Values: "default", "sonargateway", "SDM", "audit management", "ad-hoc-query" # reason description: "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset"
		username = null # username description: ""
		# ### optional ### 
		# amazon_secret = null # amazon_secret description: "Configuration to integrate with AWS Secrets Manager"
		# ca_file = null # ca_file description: "Use the specified certificate file to verify the peer. The file may contain multiple CA certificates."
		# cert_file = null # cert_file description: "Use the specified client certificate file when getting a file with HTTPS, FTPS or another SSL-based protocol."
		# cyberark_secret = null # cyberark_secret description: "Configuration to integrate with CyberArk Vault"
		# dns_srv = null # dns srv description: ""
		# hashicorp_secret = null # hashicorp_secret description: "Configuration to integrate with HashiCorp Vault"
		# passphrase = null # passphrase description: "Passphrase for the private key."
		# replica_set = null # replica_set description: ""
		# self_signed = null # self_signed description: "Accept self-signed certificates"
		# ssl = null # ssl description: "If true, use SSL when connecting"
	}
}
```


