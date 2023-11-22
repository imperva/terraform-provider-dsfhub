package dsfhub

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"log"
	"net/http"
)

const contentTypeApplicationJson = "application/json"
const endpointGateways = "/gateways"

// Client represents an internal client that brokers calls to the DSF API
type Client struct {
	config          *Config
	httpClient      *http.Client
	providerVersion string
	gateways        *GatewaysResponse
}

// GatewaysResponse contains account id
type GatewaysResponse struct {
	Data []struct {
		ApplianceId   int    `json:"applianceId"`
		ApplianceType string `json:"applianceType"`
		ID            string `json:"id"`
		Name          string `json:"name"`
		Hostname      string `json:"hostname"`
		ServerType    string `json:"serverType"`
		Sonar         struct {
			JsonarUid string `json:"jsonarUid"`
		} `json:"sonar"`
	} `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UpdateAuditResponse struct {
	Data   string     `json:"data"`
	Errors []APIError `json:"errors,omitempty"`
}

type ResourceWrapper struct {
	Data   ResourceData `json:"data"`
	Errors []APIError   `json:"errors,omitempty"`
}

type APIError struct {
	Status int    `json:"status,omitempty"`
	Id     string `json:"id,omitempty"`
	Source struct {
		Pointer string `json:"pointer,omitempty"`
	} `json:"source,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
}

type ResourceData struct {
	ApplianceID     int       `json:"applianceId,omitempty"`
	ApplianceType   string    `json:"applianceType,omitempty"`
	AssetData       AssetData `json:"assetData"`
	AuditState      string    `json:"auditState,omitempty"`
	GatewayID       string    `json:"gatewayId"`
	GatewayName     string    `json:"gatewayName,omitempty"`
	ID              string    `json:"id,omitempty,omitempty"`
	IsMonitored     bool      `json:"isMonitored,omitempty"`
	ParentAssetID   string    `json:"parentAssetId,omitempty"`
	RemoteSyncState string    `json:"remoteSyncState,omitempty"`
	ServerType      string    `json:"serverType"`
}

type AssetData struct {
	AdminEmail             string            `json:"admin_email"`
	Application            string            `json:"application,omitempty"`
	Archive                bool              `json:"archive,omitempty"`
	Arn                    string            `json:"arn,omitempty"`
	AssetDisplayName       string            `json:"asset_display_name"`
	AssetID                string            `json:"asset_id,omitempty"`
	AssetSource            string            `json:"Asset Source,omitempty"`
	AuditDataType          string            `json:"audit_data_type,omitempty"`
	AuditInfo              *AuditInfo        `json:"audit_info,omitempty"`
	AuditPullEnabled       bool              `json:"audit_pull_enabled,omitempty"`
	AuditType              string            `json:"audit_type,omitempty"`
	AvailabilityZones      string            `json:"availability_zones,omitempty"`
	AvailableRegions       string            `json:"available_regions,omitempty"`
	AwsProxyConfig         *AwsProxyConfig   `json:"aws_proxy_config,omitempty"`
	CaCertsPath            string            `json:"ca_certs_path,omitempty"`
	CaFile                 string            `json:"ca_file,omitempty"`
	ClusterEngine          string            `json:"cluster_engine,omitempty"`
	ClusterID              string            `json:"cluster_id,omitempty"`
	ClusterMemberID        string            `json:"cluster_member_id,omitempty"`
	ClusterName            string            `json:"cluster_name,omitempty"`
	Connections            []AssetConnection `json:"connections,omitempty"`
	ConsumerGroup          string            `json:"consumer_group,omitempty"`
	ConsumerGroupWorkers   string            `json:"consumer_group_workers,omitempty"`
	ContentType            string            `json:"content_type,omitempty"`
	CredentialsEndpoint    string            `json:"credentials_endpoint,omitempty"`
	Criticality            int               `json:"criticality,omitempty"`
	DatabaseName           string            `json:"database_name,omitempty"`
	DbEngine               string            `json:"db_engine,omitempty"`
	DbInstancesDisplayName string            `json:"db_instances_display_name,omitempty"`
	DurationThreshold      int               `json:"duration_threshold,omitempty"`
	EnableAuditManagement  bool              `json:"enable_audit_management,omitempty"`
	EnableAuditMonitoring  bool              `json:"enable_audit_monitoring,omitempty"`
	EnabledLogsExports     string            `json:"enabled_logs_exports,omitempty"`
	//Endpoint               string            `json:"endpoint,omitempty"`
	EntitlementEnabled     bool        `json:"entitlement_enabled,omitempty"`
	GatewayService         string      `json:"gateway_service,omitempty"`
	HostTimezoneOffset     string      `json:"host_timezone_offset,omitempty"`
	IgnoreLatestOf         string      `json:"ignore_latest_of,omitempty"`
	IsCluster              bool        `json:"is_cluster,omitempty"`
	IsMultiZones           bool        `json:"is_multi_zones,omitempty"`
	JsonarUID              string      `json:"jsonar_uid,omitempty"`
	JsonarUIDDisplayName   string      `json:"jsonar_uid_display_name,omitempty"`
	Location               string      `json:"location,omitempty"`
	LogBucketID            string      `json:"log_bucket_id,omitempty"`
	LogsDestinationAssetID string      `json:"logs_destination_asset_id,omitempty"`
	Logstore               string      `json:"logstore,omitempty"`
	ManagedBy              string      `json:"managed_by,omitempty"`
	MaxConcurrentConn      string      `json:"max_concurrent_conn,omitempty"`
	OwnedBy                string      `json:"owned_by,omitempty"`
	ParentAssetID          string      `json:"parent_asset_id,omitempty"`
	Project                string      `json:"project,omitempty"`
	ADProvider             string      `json:"provider,omitempty"`
	ProviderUrl            string      `json:"provider_url,omitempty"`
	Proxy                  string      `json:"proxy,omitempty"`
	PubsubSubscription     string      `json:"pubsub_subscription,omitempty"`
	PullType               string      `json:"pull_type,omitempty"`
	Region                 string      `json:"region,omitempty"`
	SdmEnabled             bool        `json:"sdm_enabled,omitempty"`
	Searches               string      `json:"searches,omitempty"`
	ServerHostName         string      `json:"Server Host Name,omitempty"`
	ServerIP               string      `json:"Server IP,omitempty"`
	ServerPort             interface{} `json:"Server Port,omitempty"`
	//ServerType             string         `json:"serverType,omitempty"`
	ServiceEndpoint  string            `json:"service_endpoint,omitempty"`
	ServiceEndpoints *ServiceEndpoints `json:"service_endpoints,omitempty"`
	ServiceName      string            `json:"Service Name,omitempty"`
	SmtpTimeout      string            `json:"smtp_timeout,omitempty"`
	Ssl              bool              `json:"SSL,omitempty"`
	SubscriptionID   string            `json:"subscription_id,omitempty"`
	UsedFor          string            `json:"used_for,omitempty"`
	Version          float64           `json:"version,omitempty"`
	VirtualHostname  string            `json:"virtual_hostname,omitempty"`
	VirtualIp        string            `json:"virtual_ip,omitempty,omitempty"`
	XelDirectory     string            `json:"xel_directory,omitempty,omitempty"`
}

type AwsProxyConfig struct {
	HTTP  string `json:"http,omitempty"`
	HTTPS string `json:"https,omitempty"`
}

type AuditInfo struct {
	PolicyTemplateName string `json:"policy_template_name,omitempty"`
}

type AssetConnection struct {
	Reason         string         `json:"reason,omitempty"`
	AuthMechanism  string         `json:"auth_mechanism,omitempty"`
	RoleName       string         `json:"role_name,omitempty"`
	ConnectionData ConnectionData `json:"connectionData"`
}

type ServiceEndpoints struct {
	Logs string `json:"logs,omitempty"`
}

type ConnectionData struct {
	AccessID                      string            `json:"access_id,omitempty"`
	AccessKey                     string            `json:"access_key,omitempty"`
	AccessMethod                  string            `json:"access_method,omitempty"`
	AccountName                   string            `json:"account_name,omitempty"`
	AmazonSecret                  *Secret           `json:"amazon_secret,omitempty"`
	ApiKey                        string            `json:"api_key,omitempty"`
	ApplicationID                 string            `json:"application_id,omitempty"`
	AuthMechanism                 string            `json:"auth_mechanism,omitempty"`
	Autocommit                    bool              `json:"autocommit,omitempty"`
	AwsConnectionID               string            `json:"aws_connection_id,omitempty"`
	AwsIamServerID                string            `json:"aws_iam_server_id,omitempty"`
	AzureStorageAccount           string            `json:"azure_storage_account,omitempty"`
	AzureStorageContainer         string            `json:"azure_storage_container,omitempty"`
	AzureStorageSecretKey         string            `json:"azure_storage_secret_key,omitempty"`
	BaseDn                        string            `json:"base_dn,omitempty"`
	Bucket                        string            `json:"bucket,omitempty"`
	CaCertsPath                   string            `json:"ca_certs_path,omitempty"`
	CaFile                        string            `json:"ca_file,omitempty"`
	CacheFile                     string            `json:"cache_file,omitempty"`
	CertFile                      string            `json:"cert_file,omitempty"`
	ClientID                      string            `json:"client_id,omitempty"`
	ClientSecret                  string            `json:"client_secret,omitempty"`
	ClusterID                     string            `json:"cluster_id,omitempty"`
	ClusterMemberID               string            `json:"cluster_member_id,omitempty"`
	ClusterName                   string            `json:"cluster_name,omitempty"`
	ContentType                   string            `json:"content_type,omitempty"`
	CredentialExpiry              string            `json:"credential_expiry,omitempty"`
	CredentialFields              *CredentialFields `json:"credential_fields,omitempty"`
	Crn                           string            `json:"crn,omitempty"`
	CyberarkSecret                *Secret           `json:"cyberark_secret,omitempty"`
	DatabaseName                  string            `json:"database_name,omitempty"`
	DbRole                        string            `json:"db_role,omitempty"`
	DirectoryID                   string            `json:"directory_id,omitempty"`
	Dn                            string            `json:"dn,omitempty"`
	DnsSrv                        bool              `json:"dns_srv,omitempty"`
	Driver                        string            `json:"driver,omitempty"`
	Dsn                           string            `json:"DSN,omitempty"`
	EventhubAccessKey             string            `json:"eventhub_access_key,omitempty"`
	EventhubAccessPolicy          string            `json:"eventhub_access_policy,omitempty"`
	EventhubName                  string            `json:"eventhub_name,omitempty"`
	EventhubNamespace             string            `json:"eventhub_namespace,omitempty"`
	External                      bool              `json:"external,omitempty"`
	ExternalID                    string            `json:"external_id,omitempty"`
	ExtraKinitParameters          string            `json:"extra_kinit_parameters,omitempty"`
	Format                        string            `json:"format,omitempty"`
	HashicorpSecret               *Secret           `json:"hashicorp_secret,omitempty"`
	HiveServerType                string            `json:"hive_server_type,omitempty"`
	HostNameMismatch              bool              `json:"host_name_mismatch,omitempty"`
	Hosts                         string            `json:"hosts,omitempty"`
	Httppath                      string            `json:"httppath,omitempty"`
	IsCluster                     bool              `json:"is_cluster,omitempty"`
	JdbcSslTrustServerCertificate bool              `json:"jdbc_ssl_trust_server_certificate,omitempty"`
	JdbcSslTrustStoreLocation     string            `json:"jdbc_ssl_trust_store_location,omitempty"`
	JdbcSslTrustStorePassword     string            `json:"jdbc_ssl_trust_store_password,omitempty"`
	KerberosHostFqdn              string            `json:"kerberos_host_fqdn,omitempty"`
	KerberosKdc                   string            `json:"kerberos_kdc,omitempty"`
	KerberosRetryCount            int               `json:"kerberos_retry_count,omitempty"`
	KerberosServiceKdc            string            `json:"kerberos_service_kdc,omitempty"`
	KerberosServiceRealm          string            `json:"kerberos_service_realm,omitempty"`
	KerberosSpn                   string            `json:"kerberos_spn,omitempty"`
	KeyFile                       string            `json:"key_file,omitempty"`
	KeytabFile                    string            `json:"keytab_file,omitempty"`
	KinitProgramPath              string            `json:"kinit_program_path,omitempty"`
	NetServiceName                string            `json:"net_service_name,omitempty"`
	Nonce                         string            `json:"nonce,omitempty"`
	Ntlm                          bool              `json:"ntlm,omitempty"`
	OauthParameters               *OauthParameters  `json:"oauth_parameters,omitempty"`
	OdbcConnectionString          string            `json:"odbc_connection_string,omitempty"`
	PageSize                      string            `json:"page_size,omitempty"`
	Passphrase                    string            `json:"passphrase,omitempty"`
	Password                      string            `json:"password,omitempty"`
	Port                          string            `json:"port,omitempty"`
	Principal                     string            `json:"principal,omitempty"`
	Protocol                      string            `json:"protocol,omitempty"`
	ProxyAutoDetect               string            `json:"proxy_auto_detect,omitempty"`
	ProxyPassword                 string            `json:"proxy_password,omitempty"`
	ProxyPort                     string            `json:"proxy_port,omitempty"`
	ProxyServer                   string            `json:"proxy_server,omitempty"`
	ProxySslType                  string            `json:"proxy_ssl_type,omitempty"`
	Query                         string            `json:"query,omitempty"`
	RedirectUri                   string            `json:"redirect_uri,omitempty"`
	Region                        string            `json:"region,omitempty"`
	ReplicaSet                    string            `json:"replica_set,omitempty"`
	ResourceID                    string            `json:"resource_id,omitempty"`
	RoleName                      string            `json:"role_name,omitempty"`
	Schema                        string            `json:"schema,omitempty"`
	SecretKey                     string            `json:"secret_key,omitempty"`
	SecureConnection              bool              `json:"secure_connection,omitempty"`
	SelfSigned                    bool              `json:"self_signed,omitempty"`
	SelfSignedCert                bool              `json:"self_signed_cert,omitempty"`
	ServerIp                      string            `json:"server_ip,omitempty"`
	ServerPort                    int               `json:"server_port,omitempty"`
	ServiceKey                    string            `json:"service_key,omitempty"`
	SnowflakeRole                 string            `json:"snowflake_role,omitempty"`
	Ssl                           bool              `json:"ssl,omitempty"`
	SslServerCert                 string            `json:"ssl_server_cert,omitempty"`
	StoreAwsCredentials           bool              `json:"store_aws_credentials,omitempty"`
	SubscriptionID                string            `json:"subscription_id,omitempty"`
	TenantID                      string            `json:"tenant_id,omitempty"`
	ThriftTransport               int               `json:"thrift_transport,omitempty"`
	TmpUser                       bool              `json:"tmp_user,omitempty"`
	Token                         string            `json:"token,omitempty"`
	TokenEndpoint                 string            `json:"token_endpoint,omitempty"`
	Transportmode                 string            `json:"transportmode,omitempty"`
	Url                           string            `json:"url,omitempty"`
	UseKeytab                     bool              `json:"use_keytab,omitempty"`
	Username                      string            `json:"username,omitempty"`
	V2KeyEngine                   bool              `json:"v2_key_engine,omitempty"`
	VirtualHostname               string            `json:"virtual_hostname,omitempty"`
	VirtualIp                     string            `json:"virtual_ip,omitempty"`
	WalletDir                     string            `json:"wallet_dir,omitempty"`
	Warehouse                     string            `json:"warehouse,omitempty"`
}

type Secret struct {
	FieldMapping  map[string]string `json:"field_mapping,omitempty"`
	Path          string            `json:"path,omitempty,omitempty"`
	SecretAssetID string            `json:"secret_asset_id,omitempty"`
	SecretName    string            `json:"secret_name,omitempty"`
}

type CredentialFields struct {
	CredentialSource string `json:"credential_source,omitempty"`
	RoleArn          string `json:"role_arn,omitempty"`
}

type OauthParameters struct {
	Parameter string `json:"parameter,omitempty"`
}

type ResourceResponse struct {
	Data   string     `json:"data"`
	Errors []APIError `json:"errors,omitempty"`
}

// NewClient creates a new client with the provided configuration
func NewClient(config *Config) *Client {
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	if config.InsecureSSL {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	client := &http.Client{Transport: customTransport}
	return &Client{config: config, httpClient: client, providerVersion: "1.0.0"}
}

// Verify checks the API credentials
func (c *Client) Verify() (*GatewaysResponse, error) {
	log.Println("[INFO] Checking API token against DSF Host /gateways endpoint")

	resp, err := c.MakeCall(http.MethodGet, endpointGateways, nil)
	if err != nil {
		// tls: failed to verify certificate
		return nil, fmt.Errorf("error checking token: %s", err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Parse the JSON
	var gatewaysResponse GatewaysResponse
	err = json.Unmarshal([]byte(responseBody), &gatewaysResponse)
	log.Printf("[DEBUG] gatewaysResponse: %s\n", responseBody)
	if err != nil {
		return nil, fmt.Errorf("error parsing gateways JSON response: %s", err)
	}
	if gatewaysResponse.Code == 403 {
		return nil, fmt.Errorf("error authenticating to DSF API with token when checking gateways")
	} else {
		log.Printf("[INFO] Successfully authenticated to DSF API\n")
	}
	// resp.StatusCode
	// Dump JSON
	return &gatewaysResponse, nil
}

func (c *Client) MakeCall(method string, action string, data []byte) (*http.Response, error) {
	reqURL := c.config.DSFHUBHost + baseAPIPrefix + action
	req, err := PrepareJsonRequest(method, reqURL, data)
	if err != nil {
		return nil, fmt.Errorf("error preparing request: %s", err)
	}

	SetHeaders(c, req)

	return c.httpClient.Do(req)
}

func (c *Client) MakeCallWithQueryParams(method string, action string, data []byte, params map[string]string) (*http.Response, error) {
	reqURL := c.config.DSFHUBHost + baseAPIPrefix + action
	req, err := PrepareJsonRequest(method, reqURL, data)
	if err != nil {
		return nil, fmt.Errorf("error preparing request: %s", err)
	}
	q := req.URL.Query()
	for name, value := range params {
		q.Add(name, value)
	}
	req.URL.RawQuery = q.Encode()

	SetHeaders(c, req)

	return c.httpClient.Do(req)
}

func PrepareJsonRequest(method string, url string, data []byte) (*http.Request, error) {
	if data == nil {
		return http.NewRequest(method, url, nil)
	}
	return http.NewRequest(method, url, bytes.NewReader(data))
}

func SetHeaders(c *Client, req *http.Request) {
	req.Header.Set("Content-Type", contentTypeApplicationJson)
	req.Header.Set("Authorization", "Bearer "+c.config.DSFHUBToken)
	req.Header.Set("Accept", contentTypeApplicationJson)
}

func PositiveHash(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	return 0
}
