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
// AddDSFDataSource Tests
////////////////////////////////////////////////////////////////

func TestClientAddDSFDataSourceBadConnection(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddDSFDataSourceBadConnection \n")
	config := &Config{DSFHUBToken: "foo", DSFHUBHost: testInvalidDSFHUBHost}
	client := &Client{config: config, httpClient: &http.Client{Timeout: time.Millisecond * 1}}

	dsfDataSource := ResourceWrapper{
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

	createDSFDataSourceResponse, err := client.CreateDSFDataSource(dsfDataSource)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error adding DSFDataSource for serverType: %s and gatewayId: %s", testDSServerType, testGatewayId)) {
		t.Errorf("Should have received a client error, got: %s", err)
	}
	if createDSFDataSourceResponse != nil {
		t.Errorf("Should have received a nil createDSFDataSourceResponse instance")
	}
}

func TestClientAddDSFDataSourceBadJSON(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddDSFDataSourceBadJSON \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointDsfDataSource)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	dsfDataSource := ResourceWrapper{
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

	createDSFDataSourceResponse, err := client.CreateDSFDataSource(dsfDataSource)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error parsing add DSFDataSource JSON response serverType: %s and gatewayId: %s", testDSServerType, testGatewayId)) {
		t.Errorf("Should have received a JSON parse error, got: %s", err)
	}
	if createDSFDataSourceResponse != nil {
		t.Errorf("Should have received a nil createDSFDataSourceResponse instance")
	}
}

func TestClientAddDSFDataSourceInvalidDSFDataSource(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddDSFDataSourceInvalidDSFDataSource \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointDsfDataSource)

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

	dsfDataSource := ResourceWrapper{
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

	createDSFDataSourceResponse, err := client.CreateDSFDataSource(dsfDataSource)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	//log.Printf("[INFO] err.Error(): %s\n", err.Error())
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("errors found in json response")) {
		t.Errorf("Should have received an invalid request error missing parameters, got: %s", err)
	}
	if createDSFDataSourceResponse != nil {
		t.Errorf("Should have received a nil createDSFDataSourceResponse instance")
	}
}

func TestClientAddDSFDataSourceValid(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddDSFDataSourceValid \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointDsfDataSource)

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

	dsfDataSource := ResourceWrapper{
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

	createDSFDataSourceResponse, err := client.CreateDSFDataSource(dsfDataSource)
	if err != nil {
		t.Errorf("Should not have received an error: %s", err)
	}
	if createDSFDataSourceResponse == nil {
		t.Errorf("Should not have received a nil createDSFDataSourceResponse instance")
	}
	if createDSFDataSourceResponse.Data.ID == "" {
		t.Errorf("Should not have received a blank id from the data source read repsonse")
	}
}

func TestClientAddDSFDataSourceValidFromLocalSchema(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddDSFDataSourceValidFromLocalSchema \n")

	log.Printf("[INFO] Loading test data 'dataSourceTestDataJson'\n")
	var dataSourceTestData TestDataMap
	err := json.Unmarshal([]byte(dataSourceTestDataJson), &dataSourceTestData)
	if err != nil {
		log.Printf("[DEBUG] Unable to load test data 'dataSourceTestDataJson' %s\n", err)
		panic(err)
	}

	for serverType, dsfDataSource := range dataSourceTestData.ServerType {
		log.Printf("[DEBUG] serverType %v, dsfDataSource, %v\n", serverType, dsfDataSource)
	}

	DSFHUBToken := "foo"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointDsfDataSource)

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

	dsfDataSource := ResourceWrapper{
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

	createDSFDataSourceResponse, err := client.CreateDSFDataSource(dsfDataSource)
	if err != nil {
		t.Errorf("Should not have received an error: %s", err)
	}
	if createDSFDataSourceResponse == nil {
		t.Errorf("Should not have received a nil createDSFDataSourceResponse instance")
	}
	if createDSFDataSourceResponse.Data.ID == "" {
		t.Errorf("Should not have received a blank id from the data source read repsonse")
	}
}

//////////////////////////////////////////////////////////////////
//// ReadDSFDataSource Tests
//////////////////////////////////////////////////////////////////

func TestClientReadDSFDataSourceBadConnection(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientReadDSFDataSourceBadConnection \n")
	config := &Config{DSFHUBToken: "foo", DSFHUBHost: testInvalidDSFHUBHost}
	client := &Client{config: config, httpClient: &http.Client{Timeout: time.Millisecond * 1}}

	readDSFDataSourceResponse, err := client.ReadDSFDataSource(testArn)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error reading DSFDataSource for dataSourceId: %s", testArn)) {
		t.Errorf("Should have received a client error, got: %s", err)
	}
	if readDSFDataSourceResponse != nil {
		t.Errorf("Should have received a nil readDSFDataSourceResponse instance")
	}
}

func TestClientReadDSFDataSourceBadJSON(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddDSFDataSourceBadJSON \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointDsfDataSource + "/" + url.PathEscape(testArn))

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	readDSFDataSourceResponse, err := client.ReadDSFDataSource(testArn)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error parsing DSFDataSource JSON response for dataSourceId: %s", testArn)) {
		t.Errorf("Should have received a JSON parse error, got: %s", err)
	}
	if readDSFDataSourceResponse != nil {
		t.Errorf("Should have received a nil readDSFDataSourceResponse instance")
	}
}

func TestClientReadDSFDataSourceInvalidDataSourceId(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientReadDSFDataSourceInvalidDataSourceId \n")
	DSFHUBToken := "foo"
	invalidDSFDataSourceId := "abcde12345"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointDsfDataSource + "/" + invalidDSFDataSourceId)

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

	readDSFDataSourceResponse, err := client.ReadDSFDataSource(invalidDSFDataSourceId)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("errors found in json response")) {
		t.Errorf("Should have received invalid dsf data source id error, got: %s", err)
	}
	if readDSFDataSourceResponse != nil {
		t.Errorf("Should have received a nil readIncapRuleResponse instance")
	}
}

func TestClientReadDSFDataSourceValidDSFDataSourceId(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientReadDSFDataSourceValidDSFDataSourceId \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointDsfDataSource + "/" + url.PathEscape(testArn))

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

	readDSFDataSourceResponse, err := client.ReadDSFDataSource(testArn)
	if err != nil {
		t.Errorf("Should not have received an error: %s", err)
	}
	if readDSFDataSourceResponse == nil {
		t.Errorf("Should not have received a nil readDSFDataSourceResponse instance")
	}
}

//////////////////////////////////////////////////////////////////
//// UpdateDSFDataSource Tests
//////////////////////////////////////////////////////////////////

func TestClientUpdateDSFDataSourceBadConnection(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientUpdateDSFDataSourceBadConnection \n")
	config := &Config{DSFHUBToken: "foo", DSFHUBHost: testInvalidDSFHUBHost}
	client := &Client{config: config, httpClient: &http.Client{Timeout: time.Millisecond * 1}}

	dsfDataSource := ResourceWrapper{
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

	updateDSFDataSourceResponse, err := client.UpdateDSFDataSource(testArn, dsfDataSource)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error updating DSFDataSource for dataSourceId: %s", testArn)) {
		t.Errorf("Should have received an client error, got: %s", err)
	}
	if updateDSFDataSourceResponse != nil {
		t.Errorf("Should have received a nil updateDSFDataSourceResponse instance")
	}
}

func TestClientUpdateDSFDataSourceBadJSON(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientUpdateDSFDataSourceBadJSON \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointDsfDataSource + "/" + url.PathEscape(testArn))

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	dsfDataSource := ResourceWrapper{
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

	updateDSFDataSourceResponse, err := client.UpdateDSFDataSource(testArn, dsfDataSource)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error parsing update DSFDataSource JSON response for dataSourceId: %s", testArn)) {
		t.Errorf("Should have received a JSON parse error, got: %s", err)
	}
	if updateDSFDataSourceResponse != nil {
		t.Errorf("Should have received a nil updateDSFDataSourceResponse instance")
	}
}

func TestClientUpdateDSFDataSourceInvalidDSFDataSource(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientUpdateDSFDataSourceInvalidDSFDataSource \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointDsfDataSource + "/" + url.PathEscape(testArn))

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

	dsfDataSource := ResourceWrapper{
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

	updateDSFDataSourceResponse, err := client.UpdateDSFDataSource(testArn, dsfDataSource)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	log.Printf("[INFO] err.Error(): %s\n", err.Error())
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("errors found in json response")) {
		t.Errorf("Should have received an invalid request error missing parameters, got: %s", err)
	}
	if updateDSFDataSourceResponse != nil {
		t.Errorf("Should have received a nil updateDSFDataSourceResponse instance")
	}
}

func TestClientUpdateDSFDataSourceValid(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddDSFDataSourceValid \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointDsfDataSource + "/" + url.PathEscape(testArn))

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

	dsfDataSource := ResourceWrapper{
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

	updateDSFDataSourceResponse, err := client.UpdateDSFDataSource(testArn, dsfDataSource)
	if err != nil {
		t.Errorf("Should not have received an error")
	}
	if updateDSFDataSourceResponse == nil {
		t.Errorf("Should not have received a nil updateDSFDataSourceResponse instance")
	}
	if updateDSFDataSourceResponse.Data.ID == "" {
		t.Errorf("Should not have received a blank id from the data source read repsonse")
	}
}

//////////////////////////////////////////////////////////////////
//// DeleteDSFDataSource Tests
//////////////////////////////////////////////////////////////////

func TestClientDeleteDSFDataSourceBadConnection(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientDeleteDSFDataSourceBadConnection \n")
	config := &Config{DSFHUBToken: "foo", DSFHUBHost: testInvalidDSFHUBHost}
	client := &Client{config: config, httpClient: &http.Client{Timeout: time.Millisecond * 1}}

	deleteDSFDataSourceResponse, err := client.DeleteDSFDataSource(testArn)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error deleting DSFDataSource for dataSourceId: %s", testArn)) {
		t.Errorf("Should have received a client error, got: %s", err)
	}
	if deleteDSFDataSourceResponse != nil {
		t.Errorf("Should have received a nil deleteDSFDataSourceResponse instance")
	}
}

func TestClientDeleteDSFDataSourceInvalidDataSourceId(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientDeleteDSFDataSourceInvalidDataSourceId \n")
	DSFHUBToken := "foo"
	invalidDSFDataSourceId := "abcde12345"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointDsfDataSource + "/" + invalidDSFDataSourceId)

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

	deleteDSFDataSourceResponse, err := client.DeleteDSFDataSource(invalidDSFDataSourceId)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("errors found in json response")) {
		t.Errorf("Should have received invalid dsf data source id error, got: %s", err)
	}
	if deleteDSFDataSourceResponse != nil {
		t.Errorf("Should have received a nil deleteDSFDataSourceResponse instance")
	}
}

func TestClientDeleteDSFDataSourceValidDataSourceId(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientDeleteDSFDataSourceValidDataSourceId \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointDsfDataSource + "/" + url.PathEscape(testArn))

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

	deleteDSFDataSourceResponse, err := client.DeleteDSFDataSource(testArn)
	if err != nil {
		t.Errorf("Should not have received an error: %s", err)
	}
	if deleteDSFDataSourceResponse == nil {
		t.Errorf("Should not have received a nil deleteDSFDataSourceResponse instance")
	}
}
