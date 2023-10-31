# ### Resource example for CYBERARK with default auth_mechanism ###
resource "dsfhub_secret_manager" "cyberark_default" {
	server_type = "CYBERARK"
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
	# criticality = var.criticality # Example Values: "1", "2", "3", "4"	# The asset's importance to the business. These values are measured on a scale from \"Most critical\" (1) to \"Least critical\" (4). Allowed values: 1, 2, 3, 4
	# jsonar_uid = var.jsonar_uid	# Unique identifier (UID) attached to the Sonar machine controlling the asset
	# location = var.location	# Current human-readable description of the physical location of the asset, or region.
	# managed_by = var.managed_by	# Email of the person who maintains the asset; can be different from the owner specified in the owned_by field. Defaults to admin_email.
	# owned_by = var.owned_by	# Email of Owner / person responsible for the asset; can be different from the person in the managed_by field. Defaults to admin_email.
	# region = var.region	# For cloud systems with regions, the default region or region used with this asset
	# server_port = "16114"	# 
	# used_for = var.used_for # Example Values: "Production", "Test", "Development", "Demonstration", "QA", "Staging", "Training", "Disaster Recovery"	# Designates how this asset is used / the environment that the asset is supporting.
	# version = var.version	# Denotes the version of the asset
	asset_connection {
		auth_mechanism = "default"
		# ### required ### 
		query = null # query description: ""
		reason = null # Example Values: "default" # reason description: "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset"
		# ### optional ### 
		# ca_certs_path = null # ca_certs_path description: "Certificate authority certificates path; what location should the sysetm look for certificate information from. Equivalent to --capath in a curl call"
		# cert_file = null # cert_file description: ""
		# key_file = null # key_file description: ""
		# protocol = null # protocol description: ""
		# self_signed = null # self_signed description: ""
		# ssl = null # ssl description: "If true, use SSL when connecting"
	}
}
