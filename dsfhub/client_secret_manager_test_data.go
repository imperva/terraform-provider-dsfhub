package dsfhub

var secretManagerTestDataJson = `{
    "ServerTypes": {
        "AWS": {
            "data": {
                "assetData": {
                    "admin_email": "string",
                    "arn": "string",
                    "asset_display_name": "string",
                    "asset_source": "string",
                    "available_regions": ["string"],
                    "aws_proxy_config": {
                        "http": "< proxy >",
                        "https": "< proxy >"
                    },
                    "connections": [
                        {
                            "connectionData": {
                                "access_id": "string",
                                "amazon_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "secret_asset_id": "<AWS_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "ca_certs_path": "string",
                                "credential_fields": {
                                    "credential_source": "Ec2InstanceMetadata",
                                    "role_arn": "arn:aws:iam::111777333222:role/other_role"
                                },
                                "external_id": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "region": "string",
                                "secret_key": "string",
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "credentials_endpoint": "string",
                    "criticality": 1,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "owned_by": "string",
                    "proxy": "string",
                    "service_endpoints": {
                        "logs": "https://logs.us-east-1.amazonaws.com"
                    },
                    "used_for": "Production"
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "CYBERARK": {
            "data": {
                "assetData": {
                    "admin_email": "string",
                    "asset_display_name": "string",
                    "asset_id": "string",
                    "asset_source": "string",
                    "connections": [
                        {
                            "connectionData": {
                                "ca_certs_path": "string",
                                "cert_file": "string",
                                "key_file": "string",
                                "self_signed": true,
                                "ssl": true
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "owned_by": "string",
                    "region": "string",
                    "Server Host Name": "string",
                    "Server Port": 1234,
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "HASHICORP": {
            "data": {
                "assetData": {
                    "admin_email": "string",
                    "asset_display_name": "string",
                    "asset_id": "string",
                    "asset_source": "string",
                    "connections": [
                        {
                            "connectionData": {
                                "access_id": "string",
                                "role_name": "string",
                                "secret_key": "string",
                                "self_signed": true,
                                "ssl": true
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "owned_by": "string",
                    "region": "string",
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "secrets_managers": {}
    }
}`
