package dsfhub

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

////////////////////////////////////////////////////////////////
// AddCloudAccount Tests
////////////////////////////////////////////////////////////////

func TestClientAddCloudAccountBadConnection(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddCloudAccountBadConnection \n")
	config := &Config{DSFHUBToken: "foo", DSFHUBHost: testInvalidDSFHUBHost}
	client := &Client{config: config, httpClient: &http.Client{Timeout: time.Millisecond * 1}}

	cloudAccount := ResourceWrapper{
		Data: ResourceData{
			ServerType: testDSServerType,
			GatewayID:  testGatewayId,
			AssetData: AssetData{
				AdminEmail:       testAdminEmail,
				Arn:              testArn,
				AssetDisplayName: testAssetDisplayName,
				ServerHostName:   testServerHostName,
			},
		},
	}

	createCloudAccountResponse, err := client.CreateCloudAccount(cloudAccount)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error adding CloudAccount of serverType: %s and gatewayID: %s", testDSServerType, testGatewayId)) {
		t.Errorf("Should have received a client error, got: %s", err)
	}
	if createCloudAccountResponse != nil {
		t.Errorf("Should have received a nil createCloudAccountResponse instance")
	}
}

func TestClientAddCloudAccountBadJSON(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddCloudAccountBadJSON \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointCloudAccounts)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	cloudAccount := ResourceWrapper{
		Data: ResourceData{
			ServerType: testDSServerType,
			GatewayID:  testGatewayId,
			AssetData: AssetData{
				AdminEmail:       testAdminEmail,
				Arn:              testArn,
				AssetDisplayName: testAssetDisplayName,
				ServerHostName:   testServerHostName,
			},
		},
	}

	createCloudAccountResponse, err := client.CreateCloudAccount(cloudAccount)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error parsing add CloudAccount JSON response serverType: %s and gatewayID: %s", testDSServerType, testGatewayId)) {
		t.Errorf("Should have received a JSON parse error, got: %s", err)
	}
	if createCloudAccountResponse != nil {
		t.Errorf("Should have received a nil createCloudAccountResponse instance")
	}
}

func TestClientAddCloudAccountInvalidCloudAccount(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddCloudAccountInvalidCloudAccount \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointCloudAccounts)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(406)
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{"errors":[{"status":400,"id":"374e7034949cb208","source":{"pointer":"/api/v2/data-sources"},"title":"Bad Request","detail":"Field(s) missing or incorrect: 'serverType': must not be blank"}]}`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	cloudAccount := ResourceWrapper{
		Data: ResourceData{
			//ServerType: testDSServerType,
			//GatewayID:  testGatewayId,
			AssetData: AssetData{
				AdminEmail:       testAdminEmail,
				Arn:              testArn,
				AssetDisplayName: testAssetDisplayName,
				ServerHostName:   testServerHostName,
			},
		},
	}

	createCloudAccountResponse, err := client.CreateCloudAccount(cloudAccount)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	//log.Printf("[INFO] err.Error(): %s\n", err.Error())
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("errors found in json response")) {
		t.Errorf("Should have received an invalid request error missing parameters, got: %s", err)
	}
	if createCloudAccountResponse != nil {
		t.Errorf("Should have received a nil createCloudAccountResponse instance")
	}
}

func TestClientAddCloudAccountValid(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddCloudAccountValid \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointCloudAccounts)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(200)
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{"data":{"applianceId":1,"applianceType":"DSF_HUB","id":"arn:aws:rds:us-east-2:123456789:db:your-db","serverType":"AWS RDS MYSQL","gatewayName":"ba-dsf-4.12-gw","auditState":"NO","assetData":{"Server Host Name":"your-db-name.abcde12345.us-east-2.rds.amazonaws.com","Server Port":"3306","Server Type":"AWS RDS MYSQL","admin_email":"test@imperva.com","arn":"arn:aws:rds:us-east-2:123456789:db:your-db","asset_display_name":"arn:aws:rds:us-east-2:123456789:db:your-db","asset_id":"arn:aws:rds:us-east-2:123456789:db:your-db","asset_source":"","audit_pull_enabled":false,"jsonar_uid":"e33bfbe4-a93a-c4e5-8e9c-6e5558c2e2cd","jsonar_uid_display_name":"ba-dsf-4.12-gw","managed_by":"test@imperva.com","owned_by":"test@imperva.com","connections":[]},"gatewayId":"e33bfbe4-a93a-c4e5-8e9c-6e5558c2e2cd","cloudAccount":{},"isMonitored":false}}`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	cloudAccount := ResourceWrapper{
		Data: ResourceData{
			ServerType: testDSServerType,
			GatewayID:  testGatewayId,
			AssetData: AssetData{
				AdminEmail:       testAdminEmail,
				Arn:              testArn,
				AssetDisplayName: testAssetDisplayName,
				ServerHostName:   testServerHostName,
			},
		},
	}

	createCloudAccountResponse, err := client.CreateCloudAccount(cloudAccount)
	if err != nil {
		t.Errorf("Should not have received an error: %s", err)
	}
	if createCloudAccountResponse == nil {
		t.Errorf("Should not have received a nil createCloudAccountResponse instance")
	}
	if createCloudAccountResponse.Data.ID == "" {
		t.Errorf("Should not have received a blank id from the data source read repsonse")
	}
}

func TestClientAddCloudAccountValidFromLocalSchema(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddCloudAccountValidFromLocalSchema \n")

	log.Printf("[INFO] Loading test data 'cloudAccountTestDataJson'\n")
	var cloudAccountTestData TestDataMap
	err := json.Unmarshal([]byte(cloudAccountTestDataJson), &cloudAccountTestData)
	if err != nil {
		log.Printf("[DEBUG] Unable to load test data 'cloudAccountTestDataJson' %s\n", err)
		panic(err)
	}

	for serverType, secretManager := range cloudAccountTestData.ServerType {
		log.Printf("[DEBUG] serverType %v, secretManager %v, ServerPort '%s' %v\n", serverType, secretManager, secretManager.Data.AssetData.ServerPort, reflect.TypeOf(secretManager.Data.AssetData.ServerPort))
	}

	DSFHUBToken := "foo"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointCloudAccounts)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(200)
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{"data":{"applianceId":1,"applianceType":"DSF_HUB","id":"arn:aws:rds:us-east-2:123456789:db:your-db","serverType":"AWS RDS MYSQL","gatewayName":"ba-dsf-4.12-gw","auditState":"NO","assetData":{"Server Host Name":"your-db-name.abcde12345.us-east-2.rds.amazonaws.com","Server Port":"3306","Server Type":"AWS RDS MYSQL","admin_email":"test@imperva.com","arn":"arn:aws:rds:us-east-2:123456789:db:your-db","asset_display_name":"arn:aws:rds:us-east-2:123456789:db:your-db","asset_id":"arn:aws:rds:us-east-2:123456789:db:your-db","asset_source":"","audit_pull_enabled":false,"jsonar_uid":"e33bfbe4-a93a-c4e5-8e9c-6e5558c2e2cd","jsonar_uid_display_name":"ba-dsf-4.12-gw","managed_by":"test@imperva.com","owned_by":"test@imperva.com","connections":[]},"gatewayId":"e33bfbe4-a93a-c4e5-8e9c-6e5558c2e2cd","cloudAccount":{},"isMonitored":false}}`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	cloudAccount := ResourceWrapper{
		Data: ResourceData{
			ServerType: testDSServerType,
			GatewayID:  testGatewayId,
			AssetData: AssetData{
				AdminEmail:       testAdminEmail,
				Arn:              testArn,
				AssetDisplayName: testAssetDisplayName,
				ServerHostName:   testServerHostName,
			},
		},
	}

	createCloudAccountResponse, err := client.CreateCloudAccount(cloudAccount)
	if err != nil {
		t.Errorf("Should not have received an error: %s", err)
	}
	if createCloudAccountResponse == nil {
		t.Errorf("Should not have received a nil createCloudAccountResponse instance")
	}
	if createCloudAccountResponse.Data.ID == "" {
		t.Errorf("Should not have received a blank id from the data source read repsonse")
	}
}

//////////////////////////////////////////////////////////////////
//// ReadCloudAccount Tests
//////////////////////////////////////////////////////////////////

func TestClientReadCloudAccountBadConnection(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientReadCloudAccountBadConnection \n")
	config := &Config{DSFHUBToken: "foo", DSFHUBHost: testInvalidDSFHUBHost}
	client := &Client{config: config, httpClient: &http.Client{Timeout: time.Millisecond * 1}}

	readCloudAccountResponse, err := client.ReadCloudAccount(testArn)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error reading CloudAccount for cloudAccountId: %s", testArn)) {
		t.Errorf("Should have received a client error, got: %s", err)
	}
	if readCloudAccountResponse != nil {
		t.Errorf("Should have received a nil readCloudAccountResponse instance")
	}
}

func TestClientReadCloudAccountBadJSON(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddCloudAccountBadJSON \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointCloudAccounts + "/" + url.QueryEscape(testArn))

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	readCloudAccountResponse, err := client.ReadCloudAccount(testArn)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error parsing CloudAccount JSON response for cloudAccountId: %s", testArn)) {
		t.Errorf("Should have received a JSON parse error, got: %s", err)
	}
	if readCloudAccountResponse != nil {
		t.Errorf("Should have received a nil readCloudAccountResponse instance")
	}
}

func TestClientReadCloudAccountInvalidCloudAccountId(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientReadCloudAccountInvalidCloudAccountId \n")
	DSFHUBToken := "foo"
	invalidCloudAccountId := "abcde12345"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointCloudAccounts + "/" + invalidCloudAccountId)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(404)
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{"errors":[{"status":404,"id":"958b1e5ab42656f0","source":{"pointer":"/api/v2/data-sources/abcde12345"},"title":"Not Found","detail":"Cannot find asset with ID 'abcde12345' in 'DSF Hub'"}]}`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	readCloudAccountResponse, err := client.ReadCloudAccount(invalidCloudAccountId)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("errors found in json response")) {
		t.Errorf("Should have received invalid dsf data source id error, got: %s", err)
	}
	if readCloudAccountResponse != nil {
		t.Errorf("Should have received a nil readIncapRuleResponse instance")
	}
}

func TestClientReadCloudAccountValidCloudAccountId(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientReadCloudAccountValidCloudAccountId \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointCloudAccounts + "/" + url.QueryEscape(testArn))

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(200)
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{ "data": { "assetData": { "asset_display_name": "arn:aws:rds:us-east-2:123456789:db:your-db", "arn": "arn:aws:rds:us-east-2:123456789:db:your-db", "Server Host Name": "your-db-name.abcde12345.us-east-2.rds.amazonaws.com", "admin_email": "test@imperva.com", "connections": [] }, "serverType": "AWS RDS MYSQL", "gatewayId": "e33bfbe4-a93a-c4e5-8e9c-6e5558c2e2cd" } }`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	readCloudAccountResponse, err := client.ReadCloudAccount(testArn)
	if err != nil {
		t.Errorf("Should not have received an error: %s", err)
	}
	if readCloudAccountResponse == nil {
		t.Errorf("Should not have received a nil readCloudAccountResponse instance")
	}
}

//////////////////////////////////////////////////////////////////
//// UpdateCloudAccount Tests
//////////////////////////////////////////////////////////////////

func TestClientUpdateCloudAccountBadConnection(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientUpdateCloudAccountBadConnection \n")
	config := &Config{DSFHUBToken: "foo", DSFHUBHost: testInvalidDSFHUBHost}
	client := &Client{config: config, httpClient: &http.Client{Timeout: time.Millisecond * 1}}

	cloudAccount := ResourceWrapper{
		Data: ResourceData{
			ServerType: testDSServerType,
			GatewayID:  testGatewayId,
			AssetData: AssetData{
				AdminEmail:       "test-" + testAdminEmail,
				Arn:              testArn,
				AssetDisplayName: testAssetDisplayName,
				ServerHostName:   testServerHostName,
			},
		},
	}

	updateCloudAccountResponse, err := client.UpdateCloudAccount(testArn, cloudAccount)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error updating CloudAccount with cloudAccountId: %s", testArn)) {
		t.Errorf("Should have received an client error, got: %s", err)
	}
	if updateCloudAccountResponse != nil {
		t.Errorf("Should have received a nil updateCloudAccountResponse instance")
	}
}

func TestClientUpdateCloudAccountBadJSON(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientUpdateCloudAccountBadJSON \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointCloudAccounts + "/" + url.QueryEscape(testArn))

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	cloudAccount := ResourceWrapper{
		Data: ResourceData{
			ServerType: testDSServerType,
			GatewayID:  testGatewayId,
			AssetData: AssetData{
				AdminEmail:       testAdminEmail,
				Arn:              testArn,
				AssetDisplayName: testAssetDisplayName,
				ServerHostName:   testServerHostName,
			},
		},
	}

	updateCloudAccountResponse, err := client.UpdateCloudAccount(testArn, cloudAccount)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error parsing update CloudAccount JSON response for cloudAccountId: %s", testArn)) {
		t.Errorf("Should have received a JSON parse error, got: %s", err)
	}
	if updateCloudAccountResponse != nil {
		t.Errorf("Should have received a nil updateCloudAccountResponse instance")
	}
}

func TestClientUpdateCloudAccountInvalidCloudAccount(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientUpdateCloudAccountInvalidCloudAccount \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointCloudAccounts + "/" + url.QueryEscape(testArn))

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(404)
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{"errors":[{"status":400,"id":"c744ceed09c8c3da","source":{"pointer":"/api/v2/data-sources/arn:aws:rds:us-east-2:123456789:db:your-db"},"title":"Bad Request","detail":"Field(s) missing or incorrect: 'serverType': must not be blank"}]}`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	cloudAccount := ResourceWrapper{
		Data: ResourceData{
			//ServerType: testDSServerType,
			//GatewayID:  testGatewayId,
			AssetData: AssetData{
				AdminEmail:       testAdminEmail,
				Arn:              testArn,
				AssetDisplayName: testAssetDisplayName,
				ServerHostName:   testServerHostName,
			},
		},
	}

	updateCloudAccountResponse, err := client.UpdateCloudAccount(testArn, cloudAccount)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	log.Printf("[INFO] err.Error(): %s\n", err.Error())
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("errors found in json response")) {
		t.Errorf("Should have received an invalid request error missing parameters, got: %s", err)
	}
	if updateCloudAccountResponse != nil {
		t.Errorf("Should have received a nil updateCloudAccountResponse instance")
	}
}

func TestClientUpdateCloudAccountValid(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddCloudAccountValid \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointCloudAccounts + "/" + url.QueryEscape(testArn))

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

	cloudAccount := ResourceWrapper{
		Data: ResourceData{
			ServerType: testDSServerType,
			GatewayID:  testGatewayId,
			AssetData: AssetData{
				AdminEmail:       testAdminEmail,
				Arn:              testArn,
				AssetDisplayName: testAssetDisplayName,
				ServerHostName:   testServerHostName,
			},
		},
	}

	updateCloudAccountResponse, err := client.UpdateCloudAccount(testArn, cloudAccount)
	if err != nil {
		t.Errorf("Should not have received an error")
	}
	if updateCloudAccountResponse == nil {
		t.Errorf("Should not have received a nil updateCloudAccountResponse instance")
	}
	if updateCloudAccountResponse.Data.ID == "" {
		t.Errorf("Should not have received a blank id from the data source read repsonse")
	}
}

//////////////////////////////////////////////////////////////////
//// DeleteCloudAccount Tests
//////////////////////////////////////////////////////////////////

func TestClientDeleteCloudAccountBadConnection(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientDeleteCloudAccountBadConnection \n")
	config := &Config{DSFHUBToken: "foo", DSFHUBHost: testInvalidDSFHUBHost}
	client := &Client{config: config, httpClient: &http.Client{Timeout: time.Millisecond * 1}}

	deleteCloudAccountResponse, err := client.DeleteCloudAccount(testArn)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error deleting CloudAccount for cloudAccountId: %s", testArn)) {
		t.Errorf("Should have received a client error, got: %s", err)
	}
	if deleteCloudAccountResponse != nil {
		t.Errorf("Should have received a nil deleteCloudAccountResponse instance")
	}
}

func TestClientDeleteCloudAccountInvalidCloudAccountId(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientDeleteCloudAccountInvalidCloudAccountId \n")
	DSFHUBToken := "foo"
	invalidCloudAccountId := "abcde12345"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointCloudAccounts + "/" + invalidCloudAccountId)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(404)
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{"errors":[{"status":400,"id":"1edd8d35f53490df","source":{"pointer":"/api/v2/data-sources/arn:aws:rds:us-east-2:123456789:db:your-d"},"title":"Bad Request","detail":"'assetId' in URL is different than in request body"}]}`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	deleteCloudAccountResponse, err := client.DeleteCloudAccount(invalidCloudAccountId)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("errors found in json response")) {
		t.Errorf("Should have received invalid dsf data source id error, got: %s", err)
	}
	if deleteCloudAccountResponse != nil {
		t.Errorf("Should have received a nil deleteCloudAccountResponse instance")
	}
}

func TestClientDeleteCloudAccountValidCloudAccountId(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientDeleteCloudAccountValidCloudAccountId \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointCloudAccounts + "/" + url.QueryEscape(testArn))

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(200)
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{"data":"Asset arn:aws:rds:us-east-2:123456789:db:your-db was sent for deletion using playbook cc7136e8-8ac8-40cb-afc8-d846dc2c34f8"}`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	deleteCloudAccountResponse, err := client.DeleteCloudAccount(testArn)
	if err != nil {
		t.Errorf("Should not have received an error: %s", err)
	}
	if deleteCloudAccountResponse == nil {
		t.Errorf("Should not have received a nil deleteCloudAccountResponse instance")
	}
}
