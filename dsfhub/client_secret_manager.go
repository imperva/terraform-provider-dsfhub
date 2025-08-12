package dsfhub

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const endpointSecretManagers = "/secret-managers"

// CreateSecretManager adds a secret manager source to DSF
func (c *Client) CreateSecretManager(secretManager ResourceWrapper) (*ResourceWrapper, error) {
	log.Printf("[INFO] Adding SecretManager ServerType: %s | AssetID: %s | to gatewayID: %s\n", secretManager.Data.ServerType, secretManager.Data.AssetData.AssetID, secretManager.Data.GatewayID)

	//dsfDataSource := DSFDataSource{}
	secretManagerJSON, err := json.Marshal(secretManager)
	log.Printf("[DEBUG] Adding SecretManager - JSON: %s\n", secretManagerJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to JSON marshal SecreManager: %s\n", err)
	}

	resp, err := c.MakeCallWithQueryParams(http.MethodPost, endpointSecretManagers, secretManagerJSON, c.config.Params)
	if err != nil {
		return nil, fmt.Errorf("error adding SecretManager of serverType: %s and gatewayID: %s | err: %s\n", secretManager.Data.ServerType, secretManager.Data.GatewayID, err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] Add DSF SecretManager JSON response: %s\n", string(responseBody))

	// Parse the JSON
	var createSecretManagerResponse ResourceWrapper
	err = json.Unmarshal([]byte(responseBody), &createSecretManagerResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing add SecretManager JSON response serverType: %s and gatewayID: %s | err: %s\n", secretManager.Data.ServerType, secretManager.Data.GatewayID, err)
	}
	if createSecretManagerResponse.Errors != nil {
		return nil, fmt.Errorf("errors found in json response: %s", responseBody)
	}
	return &createSecretManagerResponse, nil
}

// ReadSecretManager gets the DSF data source by ID
func (c *Client) ReadSecretManager(secretManagerId string) (*ResourceWrapper, error) {
	log.Printf("[INFO] Getting SecretManager for secretManagerId: %s)\n", secretManagerId)

	reqURL := fmt.Sprintf(endpointSecretManagers+"/%s", url.PathEscape(secretManagerId))
	resp, err := c.MakeCall(http.MethodGet, reqURL, nil, baseAPIPrefix)
	if err != nil {
		return nil, fmt.Errorf("error reading SecretManager for secretManagerId: %s | err: %s\n", secretManagerId, err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] DSF SecretManager JSON response: %s\n", string(responseBody))

	// Parse the JSON
	var readSecretManagerResponse ResourceWrapper
	err = json.Unmarshal([]byte(responseBody), &readSecretManagerResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing SecretManager JSON response for secretManagerId: %s | secretManager: %s err: %s\n", secretManagerId, responseBody, err)
	}
	if readSecretManagerResponse.Errors != nil {
		if readSecretManagerResponse.Errors[0].Status == 404 {
			return nil, fmt.Errorf("SecretManager not found for secretManagerId: %s", secretManagerId)
		}
		return nil, fmt.Errorf("errors found in json response: %s", responseBody)
	}

	return &readSecretManagerResponse, nil
}

// ReadSecretManagers gets all secretManagers
func (c *Client) ReadSecretManagers() (*ResourcesWrapper, error) {
	log.Printf("[INFO] Getting SecretManagers\n")

	resp, err := c.MakeCall(http.MethodGet, endpointSecretManagers, nil, baseAPIPrefix)
	if err != nil {
		return nil, fmt.Errorf("error reading SecretManagers | err: %s\n", err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] DSF SecretManagers JSON response: %s\n", string(responseBody))

	// Parse the JSON
	var readSecretManagersResponse ResourcesWrapper
	err = json.Unmarshal([]byte(responseBody), &readSecretManagersResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing SecretManagers JSON response: %s err: %s\n", responseBody, err)
	}
	if readSecretManagersResponse.Errors != nil {
		return nil, fmt.Errorf("errors found in json response: %s", responseBody)
	}

	return &readSecretManagersResponse, nil
}

// UpdateSecretManager will update a specific secret-manager record in DSF referenced by the dataSourceId
func (c *Client) UpdateSecretManager(secretManagerId string, secretManager ResourceWrapper) (*ResourceWrapper, error) {
	log.Printf("[INFO] Updating SecretManager with secretManagerId: %s)\n", secretManagerId)

	secretManagerJSON, err := json.Marshal(secretManager)
	log.Printf("[DEBUG] Adding SecretManager - JSON: %s\n", secretManagerJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to JSON marshal secretManager: %s", err)
	}

	reqURL := fmt.Sprintf(endpointSecretManagers+"/%s", url.PathEscape(secretManagerId))
	resp, err := c.MakeCallWithQueryParams(http.MethodPut, reqURL, secretManagerJSON, c.config.Params)
	if err != nil {
		return nil, fmt.Errorf("error updating SecretManager with secretManagerId: %s | err: %s\n", secretManagerId, err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] DSF update SecretManager JSON response: %s\n", string(responseBody))

	// Parse the JSON
	var updateSecretManagerResponse ResourceWrapper
	err = json.Unmarshal([]byte(responseBody), &updateSecretManagerResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing update SecretManager JSON response for secretManagerId: %s | err: %s\n", secretManagerId, err)
	}
	if updateSecretManagerResponse.Errors != nil {
		return nil, fmt.Errorf("errors found in json response: %s", responseBody)
	}

	return &updateSecretManagerResponse, nil
}

// DeleteSecretManager deletes a secret-manager in DSF
func (c *Client) DeleteSecretManager(secretManagerId string) (*ResourceResponse, error) {
	log.Printf("[INFO] Deleting SecretManager with secretManagerId: %s\n", secretManagerId)

	reqURL := fmt.Sprintf(endpointSecretManagers+"/%s", url.PathEscape(secretManagerId))
	resp, err := c.MakeCall(http.MethodDelete, reqURL, nil, baseAPIPrefix)
	if err != nil {
		return nil, fmt.Errorf("error deleting SecretManager for secretManagerId: %s, %s\n", secretManagerId, err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] DSF delete SecretManager with JSON response: %s\n", string(responseBody))

	// Parse the JSON
	var deleteSecretManagerResponse ResourceResponse
	err = json.Unmarshal([]byte(responseBody), &deleteSecretManagerResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing delete SecretManager JSON response for dataSourceId: %s, %s\n", secretManagerId, err)
	}
	if deleteSecretManagerResponse.Errors != nil {
		return nil, fmt.Errorf("errors found in json response: %s", responseBody)
	}

	return &deleteSecretManagerResponse, nil
}
