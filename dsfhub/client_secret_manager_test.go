package dsfhub

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

////////////////////////////////////////////////////////////////
// // AddDSFDataSource Tests
////////////////////////////////////////////////////////////////

func TestClientAddSecretManagerBadConnection(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddSecretManagerBadConnection \n")
	config := &Config{DSFHUBToken: "foo", DSFHUBHost: testInvalidDSFHUBHost}
	client := &Client{config: config, httpClient: &http.Client{Timeout: time.Millisecond * 1}}

	secretManager := ResourceWrapper{
		Data: ResourceData{
			AssetData: AssetData{
				AdminEmail:       testAdminEmail,
				AssetDisplayName: testAssetDisplayName,
				AssetID:          testSMAssetId,
				ServerHostName:   testServerHostName,
				ServerIP:         testServerIP,
				ServerPort:       testServerPort,
				Connections: []AssetConnection{{
					Reason: testSMConnectionReason,
					ConnectionData: ConnectionData{
						RoleName:      testSMRoleName,
						AuthMechanism: testSMAuthMechanism,
					},
				}},
			},
			ServerType: testSMServerType,
			GatewayID:  testGatewayId,
		},
	}

	createSecretManagerResponse, err := client.CreateSecretManager(secretManager)
	if err == nil {
		t.Errorf("Should have received an error")
	}

	log.Printf("[INFO] Running test err.Error() %s\n", err.Error())
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error adding SecretManager of serverType: %s and gatewayID: %s", testSMServerType, testGatewayId)) {
		t.Errorf("Should have received a client error, got: %s", err)
	}
	if createSecretManagerResponse != nil {
		t.Errorf("Should have received a nil createSecretManagerResponse instance")
	}
}

func TestClientAddSecretManagerBadJSON(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddSecretManagerBadJSON \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointSecretManagers)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	secretManager := ResourceWrapper{
		Data: ResourceData{
			AssetData: AssetData{
				AdminEmail:       testAdminEmail,
				AssetDisplayName: testAssetDisplayName,
				AssetID:          testSMAssetId,
				ServerHostName:   testServerHostName,
				ServerIP:         testServerIP,
				ServerPort:       testServerPort,
				Connections: []AssetConnection{{
					Reason: testSMConnectionReason,
					ConnectionData: ConnectionData{
						RoleName:      testSMRoleName,
						AuthMechanism: testSMAuthMechanism,
					},
				}},
			},
			ServerType: testSMServerType,
			GatewayID:  testGatewayId,
		},
	}

	createSecretManagerResponse, err := client.CreateSecretManager(secretManager)
	if err == nil {
		t.Errorf("Should have received an error")
	}

	log.Printf("[INFO] Running test err.Error() %s\n", err.Error())
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error parsing add SecretManager JSON response serverType: %s and gatewayID: %s", testSMServerType, testGatewayId)) {
		t.Errorf("Should have received a client error, got: %s", err)
	}
	if createSecretManagerResponse != nil {
		t.Errorf("Should have received a nil createSecretManagerResponse instance")
	}
}

func TestClientAddSecretManagerInvalidSecretManager(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddSecretManagerInvalidSecretManager \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointSecretManagers)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(406)
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{"errors":[{"status":400,"id":"c96afebeea4576a1","source":{"pointer":"/api/v2/secret-managers"},"title":"Bad Request","detail":"Field(s) missing or incorrect: 'serverType': must not be blank"}]}`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	secretManager := ResourceWrapper{
		Data: ResourceData{
			AssetData: AssetData{
				AdminEmail:       testAdminEmail,
				AssetDisplayName: testAssetDisplayName,
				AssetID:          testSMAssetId,
				ServerHostName:   testServerHostName,
				ServerIP:         testServerIP,
				ServerPort:       testServerPort,
				Connections: []AssetConnection{{
					Reason: testSMConnectionReason,
					ConnectionData: ConnectionData{
						RoleName:      testSMRoleName,
						AuthMechanism: testSMAuthMechanism,
					},
				}},
			},
			ServerType: testSMServerType,
			GatewayID:  testGatewayId,
		},
	}

	createSecretManagerResponse, err := client.CreateSecretManager(secretManager)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	//log.Printf("[INFO] err.Error(): %s\n", err.Error())
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("errors found in json response")) {
		t.Errorf("Should have received an invalid request error missing parameters, got: %s", err)
	}
	if createSecretManagerResponse != nil {
		t.Errorf("Should have received a nil createSecretManagerResponse instance")
	}
}

func TestClientAddSecretManagerValid(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddSecretManagerValid \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointSecretManagers)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(200)
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{"data":{"applianceId":1,"applianceType":"DSF_HUB","id":"your-host-name-here","serverType":"HASHICORP","gatewayName":"ba-dsf-4.12-gw","auditState":"NO","assetData":{"Server Host Name":"your-host-name-here2","Server IP":"1.2.3.4","Server Port":"8200","Server Type":"HASHICORP","admin_email":"email@imperva.com","asset_display_name":"your:arn:here","asset_id":"your-host-name-here","asset_source":"","jsonar_uid":"e33bfbe4-a93a-c4e5-8e9c-6e5558c2e2cd","jsonar_uid_display_name":"ba-dsf-4.12-gw","managed_by":"email@imperva.com","owned_by":"email@imperva.com","connections":[{"reason":"default2","connectionData":{"auth_mechanism":"ec2","role_name":"your-vault-role-for-ec2"}}]},"gatewayId":"e33bfbe4-a93a-c4e5-8e9c-6e5558c2e2cd","cloudAccount":{},"isMonitored":false}}`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	secretManager := ResourceWrapper{
		Data: ResourceData{
			AssetData: AssetData{
				AdminEmail:       testAdminEmail,
				AssetDisplayName: testAssetDisplayName,
				AssetID:          testSMAssetId,
				ServerHostName:   testServerHostName,
				ServerIP:         testServerIP,
				ServerPort:       testServerPort,
				Connections: []AssetConnection{{
					Reason: testSMConnectionReason,
					ConnectionData: ConnectionData{
						RoleName:      testSMRoleName,
						AuthMechanism: testSMAuthMechanism,
					},
				}},
			},
			ServerType: testSMServerType,
			GatewayID:  testGatewayId,
		},
	}

	createSecretManagerResponse, err := client.CreateSecretManager(secretManager)
	if err != nil {
		t.Errorf("Should not have received an error: %s", err)
	}
	if createSecretManagerResponse == nil {
		t.Errorf("Should not have received a nil createSecretManagerResponse instance")
	}
	if createSecretManagerResponse.Data.ID == "" {
		t.Errorf("Should not have received a blank id from the data source read repsonse")
	}
}

func TestClientAddSecretManagerValidFromLocalSchema(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddSecretManagerValidFromLocalSchema \n")

	log.Printf("[INFO] Loading test data 'secretManagerTestDataJson'\n")
	var secretManagerTestData TestDataMap
	err := json.Unmarshal([]byte(secretManagerTestDataJson), &secretManagerTestData)
	if err != nil {
		log.Printf("[DEBUG] Unable to load test data 'secretManagerTestDataJson' %s\n", err)
		panic(err)
	}

	for serverType, secretManager := range secretManagerTestData.ServerType {
		log.Printf("[DEBUG] serverType %v, secretManager, %v\n", serverType, secretManager)
	}

	DSFHUBToken := "foo"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointSecretManagers)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(200)
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{"data":{"applianceId":1,"applianceType":"DSF_HUB","id":"your-host-name-here","serverType":"HASHICORP","gatewayName":"ba-dsf-4.12-gw","auditState":"NO","assetData":{"Server Host Name":"your-host-name-here2","Server IP":"1.2.3.4","Server Port":"8200","Server Type":"HASHICORP","admin_email":"email@imperva.com","asset_display_name":"your:arn:here","asset_id":"your-host-name-here","asset_source":"","jsonar_uid":"e33bfbe4-a93a-c4e5-8e9c-6e5558c2e2cd","jsonar_uid_display_name":"ba-dsf-4.12-gw","managed_by":"email@imperva.com","owned_by":"email@imperva.com","connections":[{"reason":"default2","connectionData":{"auth_mechanism":"ec2","role_name":"your-vault-role-for-ec2"}}]},"gatewayId":"e33bfbe4-a93a-c4e5-8e9c-6e5558c2e2cd","cloudAccount":{},"isMonitored":false}}`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	secretManager := ResourceWrapper{
		Data: ResourceData{
			AssetData: AssetData{
				AdminEmail:       testAdminEmail,
				AssetDisplayName: testAssetDisplayName,
				AssetID:          testSMAssetId,
				ServerHostName:   testServerHostName,
				ServerIP:         testServerIP,
				ServerPort:       testServerPort,
				Connections: []AssetConnection{{
					Reason: testSMConnectionReason,
					ConnectionData: ConnectionData{
						RoleName:      testSMRoleName,
						AuthMechanism: testSMAuthMechanism,
					},
				}},
			},
			ServerType: testSMServerType,
			GatewayID:  testGatewayId,
		},
	}

	createSecretManagerResponse, err := client.CreateSecretManager(secretManager)
	if err != nil {
		t.Errorf("Should not have received an error: %s", err)
	}
	if createSecretManagerResponse == nil {
		t.Errorf("Should not have received a nil createSecretManagerResponse instance")
	}
	if createSecretManagerResponse.Data.ID == "" {
		t.Errorf("Should not have received a blank id from the data source read repsonse")
	}
}

////////////////////////////////////////////////////////////////////
////// ReadDSFDataSource Tests
////////////////////////////////////////////////////////////////////

func TestClientReadSecretManagerBadConnection(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientReadSecretManagerBadConnection \n")
	config := &Config{DSFHUBToken: "foo", DSFHUBHost: testInvalidDSFHUBHost}
	client := &Client{config: config, httpClient: &http.Client{Timeout: time.Millisecond * 1}}

	readSecretManagerResponse, err := client.ReadSecretManager(testSMAssetId)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error reading SecretManager for secretManagerId: %s", testSMAssetId)) {
		t.Errorf("Should have received a client error, got: %s", err)
	}
	if readSecretManagerResponse != nil {
		t.Errorf("Should have received a nil readSecretManagerResponse instance")
	}
}

func TestClientReadSecretManagerBadJSON(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientReadSecretManagerBadJSON \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointSecretManagers + "/" + url.PathEscape(testSMAssetId))

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	readSecretManagerResponse, err := client.ReadSecretManager(testSMAssetId)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error parsing SecretManager JSON response for secretManagerId: %s", testSMAssetId)) {
		t.Errorf("Should have received a JSON parse error, got: %s", err)
	}
	if readSecretManagerResponse != nil {
		t.Errorf("Should have received a nil readSecretManagerResponse instance")
	}
}

func TestClientReadSecretManagerInvalidDataSourceId(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientReadDSFDataSourceInvalidDataSourceId \n")
	DSFHUBToken := "foo"
	invalidSecretManagerId := "abcde12345"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointSecretManagers + "/" + invalidSecretManagerId)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(404)
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{"errors":[{"status":404,"id":"dc533a5fe1edd22f","source":{"pointer":"/api/v2/secret-managers/abcde12345"},"title":"Not Found","detail":"Cannot find asset with ID 'abcde12345' in 'DSF Hub'"}]}`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	readSecretManagerResponse, err := client.ReadSecretManager(invalidSecretManagerId)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("errors found in json response")) {
		t.Errorf("Should have received invalid dsf data source id error, got: %s", err)
	}
	if readSecretManagerResponse != nil {
		t.Errorf("Should have received a nil readSecretManagerResponse instance")
	}
}

func TestClientReadSecretManagerValidSecretManager(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientReadSecretManagerValidSecretManager \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointSecretManagers + "/" + url.PathEscape(testSMAssetId))

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(200)
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{"data":{"applianceId":1,"applianceType":"DSF_HUB","id":"your-host-name-here","serverType":"HASHICORP","gatewayName":"ba-dsf-4.12-gw","auditState":"NO","assetData":{"Server Host Name":"your-host-name-here2","Server IP":"1.2.3.4","Server Port":"8200","Server Type":"HASHICORP","admin_email":"email@imperva.com","asset_display_name":"your:arn:here","asset_id":"your-host-name-here","asset_source":"","jsonar_uid":"e33bfbe4-a93a-c4e5-8e9c-6e5558c2e2cd","jsonar_uid_display_name":"ba-dsf-4.12-gw","managed_by":"email@imperva.com","owned_by":"email@imperva.com","connections":[{"reason":"default2","connectionData":{"auth_mechanism":"ec2","role_name":"your-vault-role-for-ec2"}}]},"gatewayId":"e33bfbe4-a93a-c4e5-8e9c-6e5558c2e2cd","remoteSyncState":"SYNCED","cloudAccount":{},"isMonitored":false}}`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	readSecretManagerResponse, err := client.ReadSecretManager(testSMAssetId)
	if err != nil {
		t.Errorf("Should not have received an error: %s", err)
	}
	if readSecretManagerResponse == nil {
		t.Errorf("Should not have received a nil readSecretManagerResponse instance")
	}
}

////////////////////////////////////////////////////////////////////
////// UpdateDSFDataSource Tests
////////////////////////////////////////////////////////////////////

func TestClientUpdateSecretManagerBadConnection(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientUpdateSecretManagerBadConnection \n")
	config := &Config{DSFHUBToken: "foo", DSFHUBHost: testInvalidDSFHUBHost}
	client := &Client{config: config, httpClient: &http.Client{Timeout: time.Millisecond * 1}}

	secretManager := ResourceWrapper{
		Data: ResourceData{
			AssetData: AssetData{
				AdminEmail:       testAdminEmail,
				AssetDisplayName: testAssetDisplayName,
				AssetID:          testSMAssetId,
				ServerHostName:   testServerHostName,
				ServerIP:         testServerIP,
				ServerPort:       testServerPort,
				Connections: []AssetConnection{{
					Reason: testSMConnectionReason,
					ConnectionData: ConnectionData{
						RoleName:      testSMRoleName,
						AuthMechanism: testSMAuthMechanism,
					},
				}},
			},
			ServerType: testSMServerType,
			GatewayID:  testGatewayId,
		},
	}

	updateSecretManagerResponse, err := client.UpdateSecretManager(testSMAssetId, secretManager)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error updating SecretManager with secretManagerId: %s", testSMAssetId)) {
		t.Errorf("Should have received an client error, got: %s", err)
	}
	if updateSecretManagerResponse != nil {
		t.Errorf("Should have received a nil updateSecretManagerResponse instance")
	}
}

func TestClientUpdateSecretManagerBadJSON(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientUpdateSecretManagerBadJSON \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointSecretManagers + "/" + url.PathEscape(testSMAssetId))

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	secretManager := ResourceWrapper{
		Data: ResourceData{
			AssetData: AssetData{
				AdminEmail:       testAdminEmail,
				AssetDisplayName: testAssetDisplayName,
				AssetID:          testSMAssetId,
				ServerHostName:   testServerHostName,
				ServerIP:         testServerIP,
				ServerPort:       testServerPort,
				Connections: []AssetConnection{{
					Reason: testSMConnectionReason,
					ConnectionData: ConnectionData{
						RoleName:      testSMRoleName,
						AuthMechanism: testSMAuthMechanism,
					},
				}},
			},
			ServerType: testSMServerType,
			GatewayID:  testGatewayId,
		},
	}

	updateSecretManagerResponse, err := client.UpdateSecretManager(testSMAssetId, secretManager)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error parsing update SecretManager JSON response for secretManagerId: %s", testSMAssetId)) {
		t.Errorf("Should have received a JSON parse error, got: %s", err)
	}
	if updateSecretManagerResponse != nil {
		t.Errorf("Should have received a nil updateSecretManagerResponse instance")
	}
}

func TestClientUpdateSecretManagerInvalidSecretManager(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientUpdateSecretManagerInvalidSecretManager \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointSecretManagers + "/" + url.PathEscape(testSMAssetId))

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(404)
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{"errors":[{"status":400,"id":"598b7a272a5983db","source":{"pointer":"/api/v2/secret-managers/your-host-name-here"},"title":"Bad Request","detail":"Field(s) missing or incorrect: 'serverType': must not be blank"}]}`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	secretManager := ResourceWrapper{
		Data: ResourceData{
			AssetData: AssetData{
				AdminEmail:       testAdminEmail,
				AssetDisplayName: testAssetDisplayName,
				AssetID:          testSMAssetId,
				ServerHostName:   testServerHostName,
				ServerIP:         testServerIP,
				ServerPort:       testServerPort,
				Connections: []AssetConnection{{
					Reason: testSMConnectionReason,
					ConnectionData: ConnectionData{
						RoleName:      testSMRoleName,
						AuthMechanism: testSMAuthMechanism,
					},
				}},
			},
			//ServerType: testSMServerType,
			//GatewayID:  testGatewayId,
		},
	}

	updateSecretManagerResponse, err := client.UpdateSecretManager(testSMAssetId, secretManager)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	log.Printf("[INFO] err.Error(): %s\n", err.Error())
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("errors found in json response")) {
		t.Errorf("Should have received an invalid request error missing parameters, got: %s", err)
	}
	if updateSecretManagerResponse != nil {
		t.Errorf("Should have received a nil updateSecretManagerResponse instance")
	}
}

func TestClientUpdateSecretManagerValid(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientUpdateSecretManagerValid \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointSecretManagers + "/" + url.PathEscape(testSMAssetId))

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(200)
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{"data":{"applianceId":1,"applianceType":"DSF_HUB","id":"arn:aws:rds:us-east-2:123456789:db:your-db","serverType":"AWS RDS MYSQL","gatewayName":"ba-dsf-4.12-gw","auditState":"NO","assetData":{"Server Host Name":"your-db-name.abcde12345.us-east-2.rds.amazonaws.com","Server Port":"3306","Server Type":"AWS RDS MYSQL","admin_email":"test2@imperva.com","arn":"arn:aws:rds:us-east-2:123456789:db:your-db","asset_display_name":"arn:aws:rds:us-east-2:123456789:db:your-db","asset_id":"arn:aws:rds:us-east-2:123456789:db:your-db","asset_source":"","audit_pull_enabled":false,"jsonar_uid":"e33bfbe4-a93a-c4e5-8e9c-6e5558c2e2cd","jsonar_uid_display_name":"ba-dsf-4.12-gw","managed_by":"test2@imperva.com","owned_by":"test2@imperva.com","connections":[]},"gatewayId":"e33bfbe4-a93a-c4e5-8e9c-6e5558c2e2cd","cloudAccount":{},"isMonitored":false}}`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	secretManager := ResourceWrapper{
		Data: ResourceData{
			AssetData: AssetData{
				AdminEmail:       testAdminEmail,
				AssetDisplayName: testAssetDisplayName,
				AssetID:          testSMAssetId,
				ServerHostName:   testServerHostName,
				ServerIP:         testServerIP,
				ServerPort:       testServerPort,
				Connections: []AssetConnection{{
					Reason: testSMConnectionReason,
					ConnectionData: ConnectionData{
						RoleName:      testSMRoleName,
						AuthMechanism: testSMAuthMechanism,
					},
				}},
			},
			//ServerType: testSMServerType,
			//GatewayID:  testGatewayId,
		},
	}

	updateSecretManagerResponse, err := client.UpdateSecretManager(testSMAssetId, secretManager)
	if err != nil {
		t.Errorf("Should not have received an error")
	}
	if updateSecretManagerResponse == nil {
		t.Errorf("Should not have received a nil updateDSFDataSourceResponse instance")
	}
	if updateSecretManagerResponse.Data.ID == "" {
		t.Errorf("Should not have received a blank id from the data source read repsonse")
	}
}

////////////////////////////////////////////////////////////////////
////// DeleteDSFDataSource Tests
////////////////////////////////////////////////////////////////////

func TestClientDeleteSecretManagerBadConnection(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientDeleteSecretManagerBadConnection \n")
	config := &Config{DSFHUBToken: "foo", DSFHUBHost: testInvalidDSFHUBHost}
	client := &Client{config: config, httpClient: &http.Client{Timeout: time.Millisecond * 1}}

	deleteSecretManagerResponse, err := client.DeleteSecretManager(testSMAssetId)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error deleting SecretManager for secretManagerId: %s", testSMAssetId)) {
		t.Errorf("Should have received a client error, got: %s", err)
	}
	if deleteSecretManagerResponse != nil {
		t.Errorf("Should have received a nil deleteSecretManagerResponse instance")
	}
}

func TestClientDeleteSecretManagerInvalidSecretManagerId(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientDeleteSecretManagerInvalidSecretManagerId \n")
	DSFHUBToken := "foo"
	invalidSecretManagerId := "abcde12345"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointSecretManagers + "/" + invalidSecretManagerId)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(404)
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{"errors":[{"status":404,"id":"7d054cdc0f703009","source":{"pointer":"/api/v2/secret-managers/abcde12345"},"title":"Not Found","detail":"Cannot find asset with ID 'abcde12345' in 'DSF Hub'"}]}`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	deleteSecretManagerResponse, err := client.DeleteSecretManager(invalidSecretManagerId)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("errors found in json response")) {
		t.Errorf("Should have received invalid dsf data source id error, got: %s", err)
	}
	if deleteSecretManagerResponse != nil {
		t.Errorf("Should have received a nil deleteSecretManagerResponse instance")
	}
}

func TestClientDeleteSecretManagerValidSecretManagerId(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientDeleteSecretManagerValidSecretManagerId \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointSecretManagers + "/" + url.PathEscape(testSMAssetId))

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(200)
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{"data":"Asset your-host-name-here was sent for deletion using playbook b21acdf8-da53-4808-9450-e8924efb8908"}`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	deleteSecretManagerResponse, err := client.DeleteSecretManager(testSMAssetId)
	if err != nil {
		t.Errorf("Should not have received an error: %s", err)
	}
	if deleteSecretManagerResponse == nil {
		t.Errorf("Should not have received a nil deleteSecretManagerResponse instance")
	}
}
