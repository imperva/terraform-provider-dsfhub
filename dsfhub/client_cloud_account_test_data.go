package dsfhub

var cloudAccountTestDataJson = `{
    "ServerTypes": {
        "ALIBABA": {
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
                                "amazon_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "secret_asset_id": "<AWS_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "role_name": "string",
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
                    "Server Port": 1234,
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
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
        "AZURE": {
            "data": {
                "assetData": {
                    "admin_email": "string",
                    "asset_display_name": "string",
                    "asset_id": "string",
                    "asset_source": "string",
                    "connections": [
                        {
                            "connectionData": {
                                "amazon_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "secret_asset_id": "<AWS_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "client_secret": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "key_file": "string",
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
                    "server_host_name": "string",
                    "used_for": "Production"
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "GCP": {
            "data": {
                "assetData": {
                    "admin_email": "string",
                    "asset_display_name": "string",
                    "asset_id": "string",
                    "asset_source": "string",
                    "connections": [
                        {
                            "connectionData": {
                                "amazon_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "secret_asset_id": "<AWS_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "key_file": "string",
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
                    "server_host_name": "string",
                    "used_for": "Production"
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "cloud_accounts": {}
    }
}`
