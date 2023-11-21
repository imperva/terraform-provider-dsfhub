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
// AddDSFLogAggregator Tests
////////////////////////////////////////////////////////////////

func TestClientAddAddLogAggregatorBadConnection(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddLogAggregatorBadConnection \n")
	config := &Config{DSFHUBToken: "foo", DSFHUBHost: testInvalidDSFHUBHost}
	client := &Client{config: config, httpClient: &http.Client{Timeout: time.Millisecond * 1}}

	logAggregator := ResourceWrapper{
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

	createLogAggregatorResponse, err := client.CreateLogAggregator(logAggregator)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error adding LogAggregator of serverType: %s and gatewayID: %s", testDSServerType, testGatewayId)) {
		t.Errorf("Should have received a client error, got: %s", err)
	}
	if createLogAggregatorResponse != nil {
		t.Errorf("Should have received a nil createLogAggregatorResponse instance")
	}
}

func TestClientAddLogAggregatorBadJSON(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddLogAggregatorBadJSON \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointLogAggregators)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	logAggregator := ResourceWrapper{
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

	createLogAggregatorResponse, err := client.CreateLogAggregator(logAggregator)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error parsing add LogAggregator JSON response serverType: %s and gatewayID: %s", testDSServerType, testGatewayId)) {
		t.Errorf("Should have received a JSON parse error, got: %s", err)
	}
	if createLogAggregatorResponse != nil {
		t.Errorf("Should have received a nil createLogAggregatorResponse instance")
	}
}

func TestClientAddLogAggregatorInvalidLogAggregator(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddLogAggregatorInvalidLogAggregator \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointLogAggregators)

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

	logAggregator := ResourceWrapper{
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

	createLogAggregatorResponse, err := client.CreateLogAggregator(logAggregator)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	//log.Printf("[INFO] err.Error(): %s\n", err.Error())
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("errors found in json response")) {
		t.Errorf("Should have received an invalid request error missing parameters, got: %s", err)
	}
	if createLogAggregatorResponse != nil {
		t.Errorf("Should have received a nil createLogAggregatorResponse instance")
	}
}

func TestClientAddLogAggregatorValid(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddLogAggregatorValid \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointLogAggregators)

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

	logAggregator := ResourceWrapper{
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

	createLogAggregatorResponse, err := client.CreateLogAggregator(logAggregator)
	if err != nil {
		t.Errorf("Should not have received an error: %s", err)
	}
	if createLogAggregatorResponse == nil {
		t.Errorf("Should not have received a nil createLogAggregatorResponse instance")
	}
	if createLogAggregatorResponse.Data.ID == "" {
		t.Errorf("Should not have received a blank id from the data source read repsonse")
	}
}

func TestClientAddLogAggregatorValidFromLocalSchema(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddLogAggregatorValidFromLocalSchema \n")

	log.Printf("[INFO] Loading test data 'logAggregatorTestDataJson'\n")
	var logAggregatorTestData TestDataMap
	err := json.Unmarshal([]byte(logAggregatorTestDataJson), &logAggregatorTestData)
	if err != nil {
		log.Printf("[DEBUG] Unable to load test data 'logAggregatorTestDataJson' %s\n", err)
		panic(err)
	}

	for serverType, secretManager := range logAggregatorTestData.ServerType {
		log.Printf("[DEBUG] serverType %v, secretManager, %v\n", serverType, secretManager)
	}

	DSFHUBToken := "foo"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointLogAggregators)

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

	logAggregator := ResourceWrapper{
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

	createLogAggregatorResponse, err := client.CreateLogAggregator(logAggregator)
	if err != nil {
		t.Errorf("Should not have received an error: %s", err)
	}
	if createLogAggregatorResponse == nil {
		t.Errorf("Should not have received a nil createLogAggregatorResponse instance")
	}
	if createLogAggregatorResponse.Data.ID == "" {
		t.Errorf("Should not have received a blank id from the data source read repsonse")
	}
}

//////////////////////////////////////////////////////////////////
//// ReadLogAggregator Tests
//////////////////////////////////////////////////////////////////

func TestClientReadLogAggregatorBadConnection(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientReadLogAggregatorBadConnection \n")
	config := &Config{DSFHUBToken: "foo", DSFHUBHost: testInvalidDSFHUBHost}
	client := &Client{config: config, httpClient: &http.Client{Timeout: time.Millisecond * 1}}

	readLogAggregatorResponse, err := client.ReadLogAggregator(testArn)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error reading LogAggregator for logAggregatorId: %s", testArn)) {
		t.Errorf("Should have received a client error, got: %s", err)
	}
	if readLogAggregatorResponse != nil {
		t.Errorf("Should have received a nil readLogAggregatorResponse instance")
	}
}

func TestClientReadLogAggregatorBadJSON(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddLogAggregatorBadJSON \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointLogAggregators + "/" + url.PathEscape(testArn))

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	readLogAggregatorResponse, err := client.ReadLogAggregator(testArn)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error parsing LogAggregator JSON response for logAggregatorId: %s", testArn)) {
		t.Errorf("Should have received a JSON parse error, got: %s", err)
	}
	if readLogAggregatorResponse != nil {
		t.Errorf("Should have received a nil readLogAggregatorResponse instance")
	}
}

func TestClientReadLogAggregatorInvalidLogAggregatorId(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientReadLogAggregatorInvalidLogAggregatorId \n")
	DSFHUBToken := "foo"
	invalidLogAggregatorId := "abcde12345"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointLogAggregators + "/" + invalidLogAggregatorId)

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

	readLogAggregatorResponse, err := client.ReadLogAggregator(invalidLogAggregatorId)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("errors found in json response")) {
		t.Errorf("Should have received invalid dsf data source id error, got: %s", err)
	}
	if readLogAggregatorResponse != nil {
		t.Errorf("Should have received a nil readIncapRuleResponse instance")
	}
}

func TestClientReadLogAggregatorValidLogAggregatorId(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientReadLogAggregatorValidLogAggregatorId \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointLogAggregators + "/" + url.PathEscape(testArn))

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

	readLogAggregatorResponse, err := client.ReadLogAggregator(testArn)
	if err != nil {
		t.Errorf("Should not have received an error: %s", err)
	}
	if readLogAggregatorResponse == nil {
		t.Errorf("Should not have received a nil readLogAggregatorResponse instance")
	}
}

//////////////////////////////////////////////////////////////////
//// UpdateLogAggregator Tests
//////////////////////////////////////////////////////////////////

func TestClientUpdateLogAggregatorBadConnection(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientUpdateLogAggregatorBadConnection \n")
	config := &Config{DSFHUBToken: "foo", DSFHUBHost: testInvalidDSFHUBHost}
	client := &Client{config: config, httpClient: &http.Client{Timeout: time.Millisecond * 1}}

	logAggregator := ResourceWrapper{
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

	updateLogAggregatorResponse, err := client.UpdateLogAggregator(testArn, logAggregator)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error updating LogAggregator with logAggregatorId: %s", testArn)) {
		t.Errorf("Should have received an client error, got: %s", err)
	}
	if updateLogAggregatorResponse != nil {
		t.Errorf("Should have received a nil updateLogAggregatorResponse instance")
	}
}

func TestClientUpdateLogAggregatorBadJSON(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientUpdateLogAggregatorBadJSON \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointLogAggregators + "/" + url.PathEscape(testArn))

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != endpoint {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpoint, req.URL.String())
		}
		rw.Write([]byte(`{`))
	}))
	defer server.Close()

	config := &Config{DSFHUBToken: DSFHUBToken, DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}

	logAggregator := ResourceWrapper{
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

	updateLogAggregatorResponse, err := client.UpdateLogAggregator(testArn, logAggregator)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error parsing update LogAggregator JSON response for LogAggregatorId: %s", testArn)) {
		t.Errorf("Should have received a JSON parse error, got: %s", err)
	}
	if updateLogAggregatorResponse != nil {
		t.Errorf("Should have received a nil updateLogAggregatorResponse instance")
	}
}

func TestClientUpdateLogAggregatorInvalidLogAggregator(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientUpdateLogAggregatorInvalidLogAggregator \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointLogAggregators + "/" + url.PathEscape(testArn))

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

	logAggregator := ResourceWrapper{
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

	updateLogAggregatorResponse, err := client.UpdateLogAggregator(testArn, logAggregator)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	log.Printf("[INFO] err.Error(): %s\n", err.Error())
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("errors found in json response")) {
		t.Errorf("Should have received an invalid request error missing parameters, got: %s", err)
	}
	if updateLogAggregatorResponse != nil {
		t.Errorf("Should have received a nil updateLogAggregatorResponse instance")
	}
}

func TestClientUpdateLogAggregatorValid(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientAddLogAggregatorValid \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointLogAggregators + "/" + url.PathEscape(testArn))

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

	logAggregator := ResourceWrapper{
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

	updateLogAggregatorResponse, err := client.UpdateLogAggregator(testArn, logAggregator)
	if err != nil {
		t.Errorf("Should not have received an error")
	}
	if updateLogAggregatorResponse == nil {
		t.Errorf("Should not have received a nil updateLogAggregatorResponse instance")
	}
	if updateLogAggregatorResponse.Data.ID == "" {
		t.Errorf("Should not have received a blank id from the data source read repsonse")
	}
}

//////////////////////////////////////////////////////////////////
//// DeleteLogAggregator Tests
//////////////////////////////////////////////////////////////////

func TestClientDeleteLogAggregatorBadConnection(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientDeleteLogAggregatorBadConnection \n")
	config := &Config{DSFHUBToken: "foo", DSFHUBHost: testInvalidDSFHUBHost}
	client := &Client{config: config, httpClient: &http.Client{Timeout: time.Millisecond * 1}}

	deleteLogAggregatorResponse, err := client.DeleteLogAggregator(testArn)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("error deleting LogAggregator for logAggregatorId: %s", testArn)) {
		t.Errorf("Should have received a client error, got: %s", err)
	}
	if deleteLogAggregatorResponse != nil {
		t.Errorf("Should have received a nil deleteLogAggregatorResponse instance")
	}
}

func TestClientDeleteLogAggregatorInvalidLogAggregatorId(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientDeleteLogAggregatorInvalidLogAggregatorId \n")
	DSFHUBToken := "foo"
	invalidLogAggregatorId := "abcde12345"
	endpoint := fmt.Sprintf(baseAPIPrefix + endpointLogAggregators + "/" + invalidLogAggregatorId)

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

	deleteLogAggregatorResponse, err := client.DeleteLogAggregator(invalidLogAggregatorId)
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("errors found in json response")) {
		t.Errorf("Should have received invalid dsf data source id error, got: %s", err)
	}
	if deleteLogAggregatorResponse != nil {
		t.Errorf("Should have received a nil deleteLogAggregatorResponse instance")
	}
}

func TestClientDeleteLogAggregatorValidLogAggregatorId(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientDeleteLogAggregatorValidLogAggregatorId \n")
	DSFHUBToken := "foo"
	endpoint := fmt.Sprint(baseAPIPrefix + endpointLogAggregators + "/" + url.PathEscape(testArn))

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

	deleteLogAggregatorResponse, err := client.DeleteLogAggregator(testArn)
	if err != nil {
		t.Errorf("Should not have received an error: %s", err)
	}
	if deleteLogAggregatorResponse == nil {
		t.Errorf("Should not have received a nil deleteLogAggregatorResponse instance")
	}
}
