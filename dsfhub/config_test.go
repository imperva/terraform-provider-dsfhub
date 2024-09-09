package dsfhub

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMissingToken(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestMissingToken \n")
	config := Config{}
	client, err := config.Client()
	if err == nil {
		t.Errorf("Should have received an error, got a client: %q", client)
	}
	if err.Error() != missingAPITokenMessage {
		t.Errorf("Should have received missing API token message, got: %s", err)
	}
}

func TestMissingDSFHUBHost(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestMissingDSFHUBHost \n")
	config := Config{DSFHUBToken: "foo", DSFHUBHost: ""}
	client, err := config.Client()
	if err == nil {
		t.Errorf("Should have received an error, got a client: %q", client)
	}
	if err.Error() != missingDSFHostMessage {
		t.Errorf("Should have received missing base URL message, got: %s", err)
	}
}

func TestInvalidDSFHUBToken(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestInvalidDSFHUBToken \n")
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != baseAPIPrefix+endpointGateways {
			t.Errorf("Should have have hit /gateways endpoint. Got: %s", req.URL.String())
		}
		rw.Write([]byte(`{"code": 403,"message": "Forbidden"}`))
	}))
	defer server.Close()

	config := Config{DSFHUBToken: "foo", DSFHUBHost: server.URL}
	client, err := config.Client()
	if err == nil {
		t.Errorf("Should have received an error, got a client: %q", client)
	}
	if !strings.HasPrefix(err.Error(), "error authenticating to DSF API with token when checking gateways") {
		t.Errorf("Should have received DSF service error, got: %s", err)
	}
}

func TestValidDSFHUBToken(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestValidDSFHUBToken \n")
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != "/dsf/api/v2/gateways" {
			t.Errorf("Should have have hit /gateways endpoint. Got: %s", req.URL.String())
		}
		rw.Write([]byte(`{ "data": [ { "applianceId": 1, "applianceType": "DSF_HUB", "id": "a1b2c3-4d5e-6f7g-8h9i-9adf5a7d8a72-172.16.1.123", "name": "ba-dsf-4.12-hub", "hostname": "172.16.1.123", "serverType": "IMPERVA WAREHOUSE", "sonar": { "jsonarUid": "a1b2c3-4d5e-6f7g-8h9i-678910" } } ] }`))
	}))
	defer server.Close()

	config := Config{DSFHUBToken: "good", DSFHUBHost: server.URL}
	client, err := config.Client()
	if err != nil {
		t.Errorf("Should not have received an error, got: %s", err)
	}
	if client == nil {
		t.Error("Client should not be nil")
	}
}

func TestInvalidSyncType(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestInvalidSyncType \n")

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != "/dsf/api/v2/gateways" {
			t.Errorf("Should have have hit /gateways endpoint. Got: %s", req.URL.String())
		}
		rw.Write([]byte(`{ "data": [ { "applianceId": 1, "applianceType": "DSF_HUB", "id": "a1b2c3-4d5e-6f7g-8h9i-9adf5a7d8a72-172.16.1.123", "name": "ba-dsf-4.12-hub", "hostname": "172.16.1.123", "serverType": "IMPERVA WAREHOUSE", "sonar": { "jsonarUid": "a1b2c3-4d5e-6f7g-8h9i-678910" } } ] }`))
	}))
	defer server.Close()

	
	invalidSyncType := "BAD_SYNC_TYPE"
	log.Printf("[INFO] Configuring client with sync_type: '%v'\n", invalidSyncType)
	log.Printf("[DEBUG] Test server URL %v \n", server.URL)

	config := Config{DSFHUBToken: "good", DSFHUBHost: server.URL, Params: map[string]string{"syncType": invalidSyncType}}
	client, err := config.Client()
	if err == nil {
		t.Errorf("Should have received an error, got a client: %q", client)
	}
	if err.Error() != invalidSyncTypeMessage {
		t.Errorf("Should have invalid sync_type message, got: %s", err)
	}
}
