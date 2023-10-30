package dsfhub

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// //////////////////////////////////////////////////////////////
// Verify Tests
// //////////////////////////////////////////////////////////////
func TestClientVerifyBadConnection(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientVerifyBadConnection \n")
	config := &Config{DSFHUBToken: "foo", DSFHUBHost: "https://invalid.host.com"}
	client := &Client{config: config, httpClient: &http.Client{Timeout: time.Millisecond * 1}}
	_, err := client.Verify()
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), "error checking token") {
		t.Errorf("Should have received a client error, got: %s", err)
	}
}

func TestClientVerifyBadJSON(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientVerifyBadJSON \n")
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != baseAPIPrefix+endpointGateways {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpointGateways, req.URL.String())
		}
		rw.Write([]byte(`{`))
	}))
	defer server.Close()
	config := &Config{DSFHUBToken: "foo", DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}
	_, err := client.Verify()
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), "error parsing gateways JSON response") {
		t.Errorf("Should have received a JSON parse error, got: %s", err)
	}
}

func TestClientVerifyInvalidDSFHUBToken(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientVerifyInvalidDSFHUBToken \n")
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != baseAPIPrefix+endpointGateways {
			t.Errorf("Should have have hit %s endpoint. Got: %s", endpointGateways, req.URL.String())
		}
		rw.Write([]byte(`{"code": 403,"message": "Forbidden"}`))
	}))
	defer server.Close()
	config := &Config{DSFHUBToken: "foo", DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}
	_, err := client.Verify()
	if err == nil {
		t.Errorf("Should have received an error")
	}
	if !strings.HasPrefix(err.Error(), "error authenticating to DSF API with token when checking gateways") {
		t.Errorf("Should have received a 403 forbidden error, got: %s", err)
	}
}

func TestClientVerifyValidAccount(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestClientVerifyValidAccount \n")
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != baseAPIPrefix+endpointGateways {
			t.Errorf("Should have have hit /%s endpoint. Got: %s", endpointGateways, req.URL.String())
		}
		rw.Write([]byte(`{ "data": [ { "applianceId": 1, "applianceType": "DSF_HUB", "id": "a1b2c3-4d5e-6f7g-8h9i-9adf5a7d8a72-172.16.1.123", "name": "ba-dsf-4.12-hub", "hostname": "172.16.1.123", "serverType": "IMPERVA WAREHOUSE", "sonar": { "jsonarUid": "a1b2c3-4d5e-6f7g-8h9i-678910" } } ] }`))
	}))
	defer server.Close()
	config := &Config{DSFHUBToken: "foo", DSFHUBHost: server.URL}
	client := &Client{config: config, httpClient: &http.Client{}}
	_, err := client.Verify()
	if err != nil {
		t.Errorf("Should not have received an error, got: %s", err)
	}
}
