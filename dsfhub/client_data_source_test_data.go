package dsfhub

type TestDataMap struct {
	ServerType map[string]ResourceWrapper `json:"ServerTypes"`
}

var dataSourceTestDataJson = `{
    "ServerTypes": {
        "AEROSPIKE": {
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
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
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
        "ALIBABA APSARA MONGODB": {
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
                    "audit_type": "LOGSTORE",
                    "cluster_id": "string",
                    "cluster_member_id": "string",
                    "cluster_name": "string",
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
                                "Server Port": 0,
                                "ssl": true
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "gateway_service": "string",
                    "is_cluster": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "logs_destination_asset_id": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "parent_asset_id": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "service_name": "string",
                    "used_for": "Production",
                    "version": 1.0,
                    "virtual_hostname": "string",
                    "virtual_ip": "string"
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "ALIBABA APSARA RDS MYSQL": {
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
                                "Server Port": 0,
                                "ssl": true
                            },
                            "reason": "string"
                        }
                    ],
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
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "ALIBABA APSARA RDS POSTGRESQL": {
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
                                "Server Port": 0,
                                "ssl": true
                            },
                            "reason": "string"
                        }
                    ],
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
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "ALIBABA MAX COMPUTE": {
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
                    "criticality": 1,
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "jsonar_uid": "string",
                    "logs_destination_asset_id": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "parent_asset_id": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "used_for": "Production"
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "ALIBABA OSS": {
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
                    "criticality": 1,
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "jsonar_uid": "string",
                    "logs_destination_asset_id": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "parent_asset_id": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "used_for": "Production"
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AMBARI": {
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
        "AWS ATHENA": {
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
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
                    "proxy": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "service_endpoint": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AWS DOCUMENTDB": {
            "data": {
                "assetData": {
                    "admin_email": "string",
                    "arn": "string",
                    "asset_display_name": "string",
                    "asset_id": "string",
                    "asset_source": "string",
                    "audit_info": {
                        "policy_template_name": "<template name>"
                    },
                    "audit_pull_enabled": true,
                    "audit_type": "LOG_GROUP",
                    "cluster_id": "string",
                    "cluster_member_id": "string",
                    "cluster_name": "string",
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
                                "ca_file": "string",
                                "cert_file": "string",
                                "dns_srv": true,
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
                                "passphrase": "string",
                                "password": "string",
                                "replica_set": "string",
                                "self_signed": true,
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "credentials_endpoint": "string",
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "is_cluster": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "logs_destination_asset_id": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "parent_asset_id": "string",
                    "proxy": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "service_endpoint": "string",
                    "service_name": "string",
                    "used_for": "Production",
                    "version": 1.0,
                    "virtual_hostname": "string",
                    "virtual_ip": "string"
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AWS DOCUMENTDB CLUSTER": {
            "data": {
                "assetData": {
                    "admin_email": "string",
                    "arn": "string",
                    "asset_display_name": "string",
                    "asset_id": "string",
                    "asset_source": "string",
                    "audit_info": {
                        "policy_template_name": "<template name>"
                    },
                    "audit_pull_enabled": true,
                    "audit_type": "LOG_GROUP",
                    "cluster_id": "string",
                    "cluster_member_id": "string",
                    "cluster_name": "string",
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
                                "ca_file": "string",
                                "cert_file": "string",
                                "dns_srv": true,
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
                                "passphrase": "string",
                                "password": "string",
                                "replica_set": "string",
                                "self_signed": true,
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "credentials_endpoint": "string",
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "is_cluster": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "logs_destination_asset_id": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "parent_asset_id": "string",
                    "proxy": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "service_endpoint": "string",
                    "service_name": "string",
                    "used_for": "Production",
                    "version": 1.0,
                    "virtual_hostname": "string",
                    "virtual_ip": "string"
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AWS DYNAMODB": {
            "data": {
                "assetData": {
                    "admin_email": "string",
                    "arn": "string",
                    "asset_display_name": "string",
                    "asset_id": "string",
                    "asset_source": "string",
                    "audit_info": {
                        "policy_template_name": "<template name>"
                    },
                    "audit_pull_enabled": true,
                    "audit_type": "LOG_GROUP",
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
                                "secret_key": "string",
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
                    "proxy": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Port": "string",
                    "service_endpoints": {
                        "logs": "https://logs.us-east-1.amazonaws.com"
                    },
                    "used_for": "Production"
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AWS GLUE": {
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
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
                    "proxy": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "service_endpoint": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AWS LAKE FORMATION": {
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
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
                    "proxy": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "service_endpoint": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AWS NEPTUNE": {
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
                                "cluster_id": "string",
                                "cluster_member_id": "string",
                                "cluster_name": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "is_cluster": true,
                                "ssl": true,
                                "username": "string",
                                "virtual_hostname": "string",
                                "virtual_ip": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Port": "string",
                    "service_endpoint": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AWS NEPTUNE CLUSTER": {
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
                                "cluster_id": "string",
                                "cluster_member_id": "string",
                                "cluster_name": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "is_cluster": true,
                                "ssl": true,
                                "username": "string",
                                "virtual_hostname": "string",
                                "virtual_ip": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Port": "string",
                    "service_endpoint": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AWS RDS AURORA MYSQL": {
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
                                "driver": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "odbc_connection_string": "string",
                                "password": "string",
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
                    "proxy": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "service_endpoint": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AWS RDS AURORA MYSQL CLUSTER": {
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
                    "availability_zones": ["string"],
                    "cluster_engine": "string",
                    "cluster_id": "string",
                    "cluster_member_id": "string",
                    "cluster_name": "string",
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
                                "driver": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "odbc_connection_string": "string",
                                "password": "string",
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "credentials_endpoint": "string",
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "db_instances_display_name": ["string"],
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "enabled_logs_exports": ["string"],
                    "entitlement_enabled": true,
                    "gateway_service": "string",
                    "is_cluster": true,
                    "is_multi_zones": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "logs_destination_asset_id": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "parent_asset_id": "string",
                    "proxy": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "service_endpoint": "string",
                    "used_for": "Production",
                    "version": 1.0,
                    "virtual_hostname": "string",
                    "virtual_ip": "string"
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AWS RDS AURORA POSTGRESQL": {
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
                    "cluster_id": "string",
                    "cluster_member_id": "string",
                    "cluster_name": "string",
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
                                "autocommit": true,
                                "ca_file": "string",
                                "cert_file": "string",
                                "driver": "string",
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
                                "odbc_connection_string": "string",
                                "passphrase": "string",
                                "password": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "credentials_endpoint": "string",
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "gateway_service": "string",
                    "is_cluster": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "logs_destination_asset_id": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "parent_asset_id": "string",
                    "proxy": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "service_endpoint": "string",
                    "used_for": "Production",
                    "version": 1.0,
                    "virtual_hostname": "string",
                    "virtual_ip": "string"
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AWS RDS AURORA POSTGRESQL CLUSTER": {
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
                    "availability_zones": ["string"],
                    "cluster_engine": "string",
                    "cluster_id": "string",
                    "cluster_member_id": "string",
                    "cluster_name": "string",
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
                                "autocommit": true,
                                "ca_file": "string",
                                "cert_file": "string",
                                "driver": "string",
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
                                "odbc_connection_string": "string",
                                "passphrase": "string",
                                "password": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "credentials_endpoint": "string",
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "db_instances_display_name": ["string"],
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "enabled_logs_exports": ["string"],
                    "entitlement_enabled": true,
                    "gateway_service": "string",
                    "is_cluster": true,
                    "is_multi_zones": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "logs_destination_asset_id": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "parent_asset_id": "string",
                    "proxy": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "service_endpoint": "string",
                    "used_for": "Production",
                    "version": 1.0,
                    "virtual_hostname": "string",
                    "virtual_ip": "string"
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AWS RDS MARIADB": {
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
                                "driver": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "odbc_connection_string": "string",
                                "password": "string",
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
                    "proxy": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "service_endpoint": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AWS RDS MS SQL SERVER": {
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
                                "autocommit": true,
                                "driver": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "jdbc_ssl_trust_server_certificate": true,
                                "jdbc_ssl_trust_store_location": "string",
                                "jdbc_ssl_trust_store_password": "string",
                                "odbc_connection_string": "string",
                                "password": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
                    "log_bucket_id": "string",
                    "logs_destination_asset_id": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "parent_asset_id": "string",
                    "proxy": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "service_endpoint": "string",
                    "service_name": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AWS RDS MYSQL": {
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
                                "driver": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "odbc_connection_string": "string",
                                "password": "string",
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
                    "proxy": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "service_endpoint": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AWS RDS ORACLE": {
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
                                "autocommit": true,
                                "db_role": "string",
                                "dn": "string",
                                "driver": "string",
                                "dsn": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "net_service_name": "string",
                                "odbc_connection_string": "string",
                                "password": "string",
                                "ssl_server_cert": "string",
                                "username": "string",
                                "wallet_dir": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "credentials_endpoint": "string",
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "gateway_service": "string",
                    "host_timezone_offset": "string",
                    "jsonar_uid": "string",
                    "location": "string",
                    "logs_destination_asset_id": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "parent_asset_id": "string",
                    "proxy": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "service_endpoint": "string",
                    "service_name": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AWS RDS POSTGRESQL": {
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
                                "autocommit": true,
                                "ca_file": "string",
                                "cert_file": "string",
                                "driver": "string",
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
                                "odbc_connection_string": "string",
                                "passphrase": "string",
                                "password": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
                    "proxy": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "service_endpoint": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AWS REDSHIFT": {
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
                                "autocommit": true,
                                "aws_connection_id": "string",
                                "ca_file": "string",
                                "cert_file": "string",
                                "database_name": "string",
                                "driver": "string",
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
                                "odbc_connection_string": "string",
                                "passphrase": "string",
                                "password": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "tmp_user": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "credentials_endpoint": "string",
                    "criticality": 1,
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "gateway_service": "string",
                    "host_timezone_offset": "string",
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "parent_asset_id": "string",
                    "proxy": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "service_endpoint": "string",
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
        "AZURE COSMOSDB": {
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
        "AZURE COSMOSDB MONGO": {
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
        "AZURE COSMOSDB TABLE": {
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
        "AZURE MARIADB": {
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
                                "database_name": "string",
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
                                "passphrase": "string",
                                "password": "string",
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "gateway_service": "string",
                    "jsonar_uid": "string",
                    "location": "string",
                    "logs_destination_asset_id": "string",
                    "managed_by": "string",
                    "owned_by": "string",
                    "parent_asset_id": "string",
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AZURE MS SQL SERVER": {
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
                                "autocommit": true,
                                "driver": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "jdbc_ssl_trust_server_certificate": true,
                                "jdbc_ssl_trust_store_location": "string",
                                "jdbc_ssl_trust_store_password": "string",
                                "odbc_connection_string": "string",
                                "password": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "service_name": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AZURE MYSQL": {
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
                                "passphrase": "string",
                                "password": "string",
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
                    "owned_by": "string",
                    "parent_asset_id": "string",
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AZURE POSTGRESQL": {
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
                                "autocommit": true,
                                "ca_file": "string",
                                "cert_file": "string",
                                "driver": "string",
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
                                "odbc_connection_string": "string",
                                "passphrase": "string",
                                "password": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
        "AZURE POSTGRESQL FLEXIBLE": {
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
                                "autocommit": true,
                                "ca_file": "string",
                                "cert_file": "string",
                                "driver": "string",
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
                                "odbc_connection_string": "string",
                                "passphrase": "string",
                                "password": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
        "AZURE SQL MANAGED INSTANCE": {
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
                                "autocommit": true,
                                "driver": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "jdbc_ssl_trust_server_certificate": true,
                                "jdbc_ssl_trust_store_location": "string",
                                "jdbc_ssl_trust_store_password": "string",
                                "odbc_connection_string": "string",
                                "password": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "service_name": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "AZURE STORAGE ACCOUNT": {
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
                    "audit_type": "BLOB",
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
        "CASSANDRA": {
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
                    "audit_type": "CASSANDRA",
                    "cluster_id": "string",
                    "cluster_member_id": "string",
                    "cluster_name": "string",
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
                                "password": "string",
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "is_cluster": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "owned_by": "string",
                    "region": "string",
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "used_for": "Production",
                    "version": 1.0,
                    "virtual_hostname": "string",
                    "virtual_ip": "string"
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "CLOUDANT": {
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
                    "connections": [
                        {
                            "connectionData": {
                                "account_name": "string",
                                "amazon_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "secret_asset_id": "<AWS_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "api_key": "string",
                                "crn": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "hosts": "string",
                                "password": "string",
                                "region": "string",
                                "service_key": "string",
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
        "CLOUDANT LOCAL": {
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
                    "connections": [
                        {
                            "connectionData": {}
                        }
                    ],
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
        "CLOUDERA": {
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
                    "audit_type": "CLOUDERA",
                    "connections": [
                        {
                            "connectionData": {}
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
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
        "COCKROACHDB": {
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
                    "audit_type": "COCKROACH_V21",
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
                                "autocommit": true,
                                "ca_file": "string",
                                "cert_file": "string",
                                "driver": "string",
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
                                "odbc_connection_string": "string",
                                "passphrase": "string",
                                "password": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
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
        "COUCHBASE": {
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
                    "audit_type": "COUCHBASE_6",
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
                                "bucket": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "password": "string",
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
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
        "DATASTAX": {
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
                    "audit_type": "RSYSLOG",
                    "cluster_id": "string",
                    "cluster_member_id": "string",
                    "cluster_name": "string",
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
                                "password": "string",
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "is_cluster": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "owned_by": "string",
                    "region": "string",
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "used_for": "Production",
                    "version": 1.0,
                    "virtual_hostname": "string",
                    "virtual_ip": "string"
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "DB2": {
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
                                "autocommit": true,
                                "database_name": "string",
                                "driver": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "jdbc_ssl_trust_store_location": "string",
                                "jdbc_ssl_trust_store_password": "string",
                                "odbc_connection_string": "string",
                                "password": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
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
        "EDB POSTGRESQL": {
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
                                "autocommit": true,
                                "ca_file": "string",
                                "cert_file": "string",
                                "driver": "string",
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
                                "odbc_connection_string": "string",
                                "passphrase": "string",
                                "password": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
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
        "ELASTICSEARCH": {
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
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
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
        "ELOQUENCE": {
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
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
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
        "EMR": {
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
                    "connections": [
                        {
                            "connectionData": {}
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
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
        "GCP BIGQUERY": {
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
                    "audit_type": "BIGQUERY",
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
                    "database_name": "string",
                    "db_engine": "string",
                    "duration_threshold": 0,
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
        "GCP BIGTABLE": {
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
        "GCP MS SQL SERVER": {
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
                    "audit_type": "MSSQL",
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
                                "autocommit": true,
                                "driver": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "jdbc_ssl_trust_server_certificate": true,
                                "jdbc_ssl_trust_store_location": "string",
                                "jdbc_ssl_trust_store_password": "string",
                                "odbc_connection_string": "string",
                                "password": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "service_name": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "GCP MYSQL": {
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
                    "audit_type": "MYSQL",
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
                                "ca_file": "string",
                                "cert_file": "string",
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
                                "passphrase": "string",
                                "password": "string",
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
        "GCP POSTGRESQL": {
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
                    "audit_type": "POSTGRESQL",
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
                                "autocommit": true,
                                "ca_file": "string",
                                "cert_file": "string",
                                "driver": "string",
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
                                "odbc_connection_string": "string",
                                "passphrase": "string",
                                "password": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
        "GCP SPANNER": {
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
                    "audit_type": "SPANNER",
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
        "HBASE": {
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
                                "keytab_file": "string",
                                "kinit_program_path": "string",
                                "password": "string",
                                "principal": "string",
                                "proxy_auto_detect": "string",
                                "proxy_password": "string",
                                "proxy_port": "string",
                                "proxy_server": "string",
                                "proxy_ssl_type": "string",
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
        "HDFS": {
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
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
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
        "HIVE": {
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
                                "database_name": "string",
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
                                "hive_server_type": "string",
                                "httppath": "string",
                                "jdbc_ssl_trust_server_certificate": true,
                                "jdbc_ssl_trust_store_location": "string",
                                "jdbc_ssl_trust_store_password": "string",
                                "kerberos_host_fqdn": "string",
                                "kerberos_kdc": "string",
                                "kerberos_service_kdc": "string",
                                "kerberos_service_realm": "string",
                                "kerberos_spn": "string",
                                "keytab_file": "string",
                                "kinit_program_path": "string",
                                "password": "string",
                                "principal": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "thrift_transport": 0,
                                "transportmode": "string",
                                "use_keytab": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "content_type": "string",
                    "criticality": 1,
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
        "HORTONWORKS": {
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
                    "connections": [
                        {
                            "connectionData": {}
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
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
        "IMPALA": {
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
                                "autocommit": true,
                                "driver": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "host_name_mismatch": true,
                                "key_file": "string",
                                "odbc_connection_string": "string",
                                "passphrase": "string",
                                "password": "string",
                                "schema": "string",
                                "self_signed_cert": true,
                                "ssl": true,
                                "ssl_server_cert": "string",
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
        "INFORMIX": {
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
                                "database_name": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "password": "string",
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
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
        "IRIS": {
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
                                "autocommit": true,
                                "driver": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "odbc_connection_string": "string",
                                "password": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "gateway_service": "string",
                    "host_timezone_offset": "string",
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
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
        "KINETICA": {
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
                                "autocommit": true,
                                "driver": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "odbc_connection_string": "string",
                                "password": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
        "KNOX GATEWAY": {
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
        "MARIADB": {
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
                                "password": "string",
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
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
        "MARKLOGIC": {
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
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
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
        "MONGODB": {
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
                    "cluster_id": "string",
                    "cluster_member_id": "string",
                    "cluster_name": "string",
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
                                "ca_file": "string",
                                "cert_file": "string",
                                "dns_srv": true,
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
                                "passphrase": "string",
                                "password": "string",
                                "replica_set": "string",
                                "self_signed": true,
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "is_cluster": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "service_name": "string",
                    "used_for": "Production",
                    "version": 1.0,
                    "virtual_hostname": "string",
                    "virtual_ip": "string"
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "MONGODB ATLAS": {
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
                    "cluster_id": "string",
                    "cluster_member_id": "string",
                    "cluster_name": "string",
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
                                "secret_key": "string",
                                "ssl": true
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "gateway_service": "string",
                    "is_cluster": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "provider_url": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "used_for": "Production",
                    "version": 1.0,
                    "virtual_hostname": "string",
                    "virtual_ip": "string"
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "MS SQL SERVER": {
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
                                "autocommit": true,
                                "driver": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "jdbc_ssl_trust_server_certificate": true,
                                "jdbc_ssl_trust_store_location": "string",
                                "jdbc_ssl_trust_store_password": "string",
                                "odbc_connection_string": "string",
                                "password": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "ignore_latest_of": "string",
                    "jsonar_uid": "string",
                    "location": "string",
                    "logs_destination_asset_id": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "service_name": "string",
                    "used_for": "Production",
                    "version": 1.0,
                    "xel_directory": "string"
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "MYSQL": {
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
                                "password": "string",
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
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
        "NEO4J": {
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
                    "audit_type": "SYSLOG",
                    "connections": [
                        {
                            "connectionData": {}
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "logs_destination_asset_id": "string",
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
        "NETEZZA": {
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
                                "database_name": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "password": "string",
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "gateway_service": "string",
                    "host_timezone_offset": "string",
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
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
        "ORACLE": {
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
                    "audit_type": "SYSLOG",
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
                                "autocommit": true,
                                "cache_file": "string",
                                "db_role": "string",
                                "dn": "string",
                                "driver": "string",
                                "dsn": "string",
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
                                "keytab_file": "string",
                                "kinit_program_path": "string",
                                "net_service_name": "string",
                                "odbc_connection_string": "string",
                                "password": "string",
                                "principal": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "use_keytab": true,
                                "username": "string",
                                "wallet_dir": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "gateway_service": "string",
                    "host_timezone_offset": "string",
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "service_name": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "PERCONA MONGODB": {
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
                    "cluster_id": "string",
                    "cluster_member_id": "string",
                    "cluster_name": "string",
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
                                "ca_file": "string",
                                "cert_file": "string",
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
                                "passphrase": "string",
                                "password": "string",
                                "replica_set": "string",
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "is_cluster": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "used_for": "Production",
                    "version": 1.0,
                    "virtual_hostname": "string",
                    "virtual_ip": "string"
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "PERCONA MYSQL": {
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
                                "password": "string",
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
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
        "POSTGRESQL": {
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
                                "autocommit": true,
                                "ca_file": "string",
                                "cert_file": "string",
                                "driver": "string",
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
                                "odbc_connection_string": "string",
                                "passphrase": "string",
                                "password": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
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
        "PROGRESS OPENEDGE": {
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
                                "autocommit": true,
                                "database_name": "string",
                                "driver": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "odbc_connection_string": "string",
                                "password": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
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
        "REDIS": {
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
                    "cluster_id": "string",
                    "cluster_member_id": "string",
                    "cluster_name": "string",
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
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "is_cluster": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "region": "string",
                    "sdm_enabled": true,
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "used_for": "Production",
                    "version": 1.0,
                    "virtual_hostname": "string",
                    "virtual_ip": "string"
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "SAP HANA": {
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
                                "autocommit": true,
                                "driver": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "odbc_connection_string": "string",
                                "password": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
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
        "SCYLLADB": {
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
                                "password": "string",
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
        "SNOWFLAKE": {
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
                                "ca_file": "string",
                                "cert_file": "string",
                                "client_id": "string",
                                "client_secret": "string",
                                "content_type": "string",
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
                                "oauth_parameters": {
                                    "parameter": "value"
                                },
                                "password": "string",
                                "principal": "string",
                                "resource_id": "string",
                                "schema": "string",
                                "self_signed": true,
                                "snowflake_role": "string",
                                "ssl": true,
                                "tenant_id": "string",
                                "token": "string",
                                "token_endpoint": "string",
                                "username": "string",
                                "warehouse": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "gateway_service": "string",
                    "host_timezone_offset": "string",
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
                    "proxy": "string",
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
        "SPLUNK": {
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
                    "audit_type": "LEEF",
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
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "password": "string",
                                "secret_key": "string",
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
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
                    "region": "string",
                    "sdm_enabled": true,
                    "searches": "string",
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "used_for": "Production",
                    "version": 1.0
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "SYBASE": {
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
                                "autocommit": true,
                                "cache_file": "string",
                                "db_role": "string",
                                "driver": "string",
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
                                "kerberos_retry_count": 0,
                                "kerberos_service_kdc": "string",
                                "kerberos_service_realm": "string",
                                "kerberos_spn": "string",
                                "keytab_file": "string",
                                "kinit_program_path": "string",
                                "odbc_connection_string": "string",
                                "password": "string",
                                "principal": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "use_keytab": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "gateway_service": "string",
                    "host_timezone_offset": "string",
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
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
        "TERADATA": {
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
                    "audit_type": "TERADATA_15",
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
                                "autocommit": true,
                                "database_name": "string",
                                "driver": "string",
                                "hashicorp_secret": {
                                    "field_mapping": {
                                        "<local_field1>": "<remote_field1>",
                                        "<local_field2>": "<remote_field2>"
                                    },
                                    "path": "secret/",
                                    "secret_asset_id": "<hashicorp_asset_id>",
                                    "secret_name": "<secret_name>"
                                },
                                "odbc_connection_string": "string",
                                "password": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "gateway_service": "string",
                    "host_timezone_offset": "string",
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "max_concurrent_conn": "string",
                    "owned_by": "string",
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
        "YARN": {
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
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
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
        "YUGABYTE CQL": {
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
                    "cluster_id": "string",
                    "cluster_member_id": "string",
                    "cluster_name": "string",
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
                                "password": "string",
                                "ssl": true,
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
                    "gateway_service": "string",
                    "is_cluster": true,
                    "jsonar_uid": "string",
                    "location": "string",
                    "managed_by": "string",
                    "owned_by": "string",
                    "region": "string",
                    "Server Host Name": "string",
                    "Server Port": "string",
                    "used_for": "Production",
                    "version": 1.0,
                    "virtual_hostname": "string",
                    "virtual_ip": "string"
                },
                "asset_id": "string",
                "gateway_id": "string"
            }
        },
        "YUGABYTE SQL": {
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
                                "autocommit": true,
                                "ca_file": "string",
                                "cert_file": "string",
                                "driver": "string",
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
                                "odbc_connection_string": "string",
                                "passphrase": "string",
                                "password": "string",
                                "ssl": true,
                                "ssl_server_cert": "string",
                                "username": "string"
                            },
                            "reason": "string"
                        }
                    ],
                    "criticality": 1,
                    "database_name": "string",
                    "db_engine": "string",
                    "enable_audit_management": true,
                    "enable_audit_monitoring": true,
                    "entitlement_enabled": true,
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
        "databases": {}
    }
}`
