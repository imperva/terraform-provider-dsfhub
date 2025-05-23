package dsfhub

var logAggregatorTestDataJson = `{
    "ServerTypes": {
        "ALIBABA LOGSTORE": {
            "data": {
                "assetData": {
                    "admin_email": "string",
                    "asset_display_name": "string",
                    "asset_id": "string",
                    "asset_source": "string",
                    "audit_data_type": "string",
                    "audit_pull_enabled": true,
                    "audit_type": "LOGSTORE",
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
                                "ssl": true
                            },
                            "reason": "string"
                        }
                    ],
                    "consumer_group_workers": "string",
                    "content_type": "string",
                    "criticality": 1,
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "endpoint": "string",
                    "entitlement_enabled": true,
                    "jsonar_uid": "string",
                    "logstore": "string",
                    "managed_by": "string",
                    "owned_by": "string",
                    "parent_asset_id": "string",
                    "project": "string",
                    "pull_type": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Port": 1234,
                    "used_for": "Production"
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AWS KINESIS": {
            "data": {
                "assetData": {
                    "admin_email": "string",
                    "arn": "string",
                    "asset_display_name": "string",
                    "asset_source": "string",
                    "audit_info": {
                        "policy_template_name": "<template name>"
                    },
                    "audit_pull_enabled": true,
                    "audit_type": "KINESIS",
                    "aws_proxy_config": {
                        "http": "< proxy >",
                        "https": "< proxy >"
                    },
                    "ca_certs_path": "string",
                    "ca_file": "string",
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
                                "region": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "content_type": "string",
                    "credentials_endpoint": "string",
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "gateway_service": "string",
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "parent_asset_id": "string",
                    "proxy": "string",
                    "sdm_enabled": true,
                    "service_endpoint": "string",
                    "ssl": true,
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AWS LOG GROUP": {
            "data": {
                "assetData": {
                    "admin_email": "string",
                    "arn": "string",
                    "asset_display_name": "string",
                    "asset_source": "string",
                    "audit_info": {
                        "policy_template_name": "<template name>"
                    },
                    "audit_pull_enabled": true,
                    "audit_type": "LOG_GROUP",
                    "aws_proxy_config": {
                        "http": "< proxy >",
                        "https": "< proxy >"
                    },
                    "ca_certs_path": "string",
                    "ca_file": "string",
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
                                "region": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "content_type": "string",
                    "credentials_endpoint": "string",
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "gateway_service": "string",
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "parent_asset_id": "string",
                    "proxy": "string",
                    "sdm_enabled": true,
                    "service_endpoint": "string",
                    "ssl": true,
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AWS S3": {
            "data": {
                "assetData": {
                    "admin_email": "string",
                    "arn": "string",
                    "asset_display_name": "string",
                    "asset_source": "string",
                    "audit_info": {
                        "policy_template_name": "<template name>"
                    },
                    "audit_pull_enabled": true,
                    "audit_type": "LOG_GROUP",
                    "available_bucket_account_ids": ["string"],
                    "available_regions": ["string"],
                    "aws_proxy_config": {
                        "http": "< proxy >",
                        "https": "< proxy >"
                    },
                    "bucket_account_id": "string",
                    "ca_certs_path": "string",
                    "ca_file": "string",
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
                                }
                            },
                            "reason": "string"
                        }
                    ],
                    "content_type": "string",
                    "credentials_endpoint": "string",
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "gateway_service": "string",
                    "jsonar_uid": "string",
                    "location": "string",
                    "logs_destination_asset_id": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "parent_asset_id": "string",
                    "provider": "string",
                    "proxy": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "service_endpoint": "string",
                    "ssl": true,
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AZURE EVENTHUB": {
            "data": {
                "assetData": {
                    "admin_email": "string",
                    "asset_display_name": "string",
                    "asset_id": "string",
                    "asset_source": "string",
                    "audit_info": {
                        "policy_template_name": "<template name>"
                    },
                    "audit_pull_enabled": true,
                    "audit_type": "COSMOS_TABLE",
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
                                "ssl": true
                            },
                            "reason": "string"
                        }
                    ],
                    "consumer_group": "string",
                    "content_type": "string",
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "gateway_service": "string",
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "owned_by": "string",
                    "parent_asset_id": "string",
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
        "GCP CLOUD STORAGE BUCKET": {
            "data": {
                "assetData": {
                    "admin_email": "string",
                    "asset_display_name": "string",
                    "asset_id": "string",
                    "asset_source": "string",
                    "audit_info": {
                        "policy_template_name": "<template name>"
                    },
                    "audit_pull_enabled": true,
                    "audit_type": "BUCKET",
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
                    "content_type": "string",
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "gateway_service": "string",
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "parent_asset_id": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "GCP PUBSUB": {
            "data": {
                "assetData": {
                    "admin_email": "string",
                    "asset_display_name": "string",
                    "asset_id": "string",
                    "asset_source": "string",
                    "audit_info": {
                        "policy_template_name": "<template name>"
                    },
                    "audit_pull_enabled": true,
                    "audit_type": "BIGTABLE",
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
                    "content_type": "string",
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "gateway_service": "string",
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "parent_asset_id": "string",
                    "proxy": "string",
                    "pubsub_subscription": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "SSH": {
            "data": {
                "assetData": {
                    "admin_email": "string",
                    "asset_display_name": "string",
                    "asset_id": "string",
                    "asset_source": "string",
                    "audit_info": {
                        "policy_template_name": "<template name>"
                    },
                    "audit_pull_enabled": true,
                    "audit_type": "SSH",
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
                                "cache_file": "string",
                                "db_role": "string",
                                "external": true,
                                "extra_kinit_parameters": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "kerberos_kdc": "string",
                                "kerberos_service_kdc": "string",
                                "kerberos_service_realm": "string",
                                "kerberos_spn": "string",
                                "key_file": "string",
                                "keytab_file": "string",
                                "kinit_program_path": "string",
                                "passphrase": "string",
                                "password": "string",
                                "principal": "string",
                                "secret_key": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "use_keytab": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "content_type": "string",
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "gateway_service": "string",
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "owned_by": "string",
                    "parent_asset_id": "string",
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
        "log_aggregators": {}
    }
}`
