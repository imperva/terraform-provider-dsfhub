package dsfhub

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const endpointCloudAccounts = "/cloud-accounts"

// CreateCloudAccount adds a cloud account source to DSF
func (c *Client) CreateCloudAccount(cloudAccount ResourceWrapper) (*ResourceWrapper, error) {
	log.Printf("[INFO] Adding CloudAccount ServerType: %s | AssetID: %s | to gatewayID: %s\n", cloudAccount.Data.ServerType, cloudAccount.Data.AssetData.AssetID, cloudAccount.Data.GatewayID)

	//dsfDataSource := DSFDataSource{}
	cloudAccountJSON, err := json.Marshal(cloudAccount)
	log.Printf("[DEBUG] Adding CloudAccount - JSON: %s\n", cloudAccountJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to JSON marshal CloudAccount: %s\n", err)
	}

	resp, err := c.MakeCall(http.MethodPost, endpointCloudAccounts, cloudAccountJSON)
	if err != nil {
		return nil, fmt.Errorf("error adding CloudAccount of serverType: %s and gatewayID: %s | err: %s\n", cloudAccount.Data.ServerType, cloudAccount.Data.GatewayID, err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] Add CloudAccount JSON response: %s\n", string(responseBody))

	// Parse the JSON
	var createCloudAccountResponse ResourceWrapper
	err = json.Unmarshal([]byte(responseBody), &createCloudAccountResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing add CloudAccount JSON response serverType: %s and gatewayID: %s | err: %s\n", cloudAccount.Data.ServerType, cloudAccount.Data.GatewayID, err)
	}
	if createCloudAccountResponse.Errors != nil {
		return nil, fmt.Errorf("errors found in json response: %s", responseBody)
	}
	return &createCloudAccountResponse, nil
}

// ReadCloudAccount gets the CloudAccount by ID
func (c *Client) ReadCloudAccount(cloudAccountId string) (*ResourceWrapper, error) {
	log.Printf("[INFO] Getting CloudAccount for cloudAccountId: %s)\n", cloudAccountId)

	reqURL := fmt.Sprintf(endpointCloudAccounts+"/%s", url.PathEscape(cloudAccountId))
	resp, err := c.MakeCall(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading CloudAccount for cloudAccountId: %s | err: %s\n", cloudAccountId, err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] ReadCloudAcount JSON response: %s\n", string(responseBody))

	// Parse the JSON
	var readCloudAccountResponse ResourceWrapper
	err = json.Unmarshal([]byte(responseBody), &readCloudAccountResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing CloudAccount JSON response for cloudAccountId: %s | responseBody: %s err: %s\n", cloudAccountId, responseBody, err)
	}
	if readCloudAccountResponse.Errors != nil {
		return nil, fmt.Errorf("errors found in json response: %s", responseBody)
	}

	return &readCloudAccountResponse, nil
}

// UpdateCloudAccount will update a specific CloudAccount record in DSF referenced by the cloudAccountId
func (c *Client) UpdateCloudAccount(cloudAccountId string, cloudAccountIdData ResourceWrapper) (*ResourceWrapper, error) {
	log.Printf("[INFO] Updating CloudAccount with cloudAccountId: %s)\n", cloudAccountId)

	cloudAccountJSON, err := json.Marshal(cloudAccountIdData)
	log.Printf("[DEBUG] Adding CloudAccount - JSON: %s\n", cloudAccountJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to JSON marshal cloudAccount: %s", err)
	}

	reqURL := fmt.Sprintf(endpointCloudAccounts+"/%s", url.PathEscape(cloudAccountId))
	resp, err := c.MakeCall(http.MethodPut, reqURL, cloudAccountJSON)
	if err != nil {
		return nil, fmt.Errorf("error updating CloudAccount with cloudAccountId: %s | err: %s\n", cloudAccountId, err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] DSF update CloudAccount JSON response: %s\n", string(responseBody))

	// Parse the JSON
	var updateCloudAccountResponse ResourceWrapper
	err = json.Unmarshal([]byte(responseBody), &updateCloudAccountResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing update CloudAccount JSON response for cloudAccountId: %s | err: %s\n", cloudAccountId, err)
	}
	if updateCloudAccountResponse.Errors != nil {
		return nil, fmt.Errorf("errors found in json response: %s", responseBody)
	}

	return &updateCloudAccountResponse, nil
}

// DeleteCloudAccount deletes a CloudAccount in DSF
func (c *Client) DeleteCloudAccount(cloudAccountId string) (*ResourceResponse, error) {
	log.Printf("[INFO] Deleting CloudAccount with cloudAccountId: %s\n", cloudAccountId)

	reqURL := fmt.Sprintf(endpointCloudAccounts+"/%s", url.PathEscape(cloudAccountId))
	resp, err := c.MakeCall(http.MethodDelete, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error deleting CloudAccount for cloudAccountId: %s, %s\n", cloudAccountId, err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] DSF delete CloudAccount with JSON response: %s\n", string(responseBody))

	// Parse the JSON
	var deleteCloudAccountResponse ResourceResponse
	err = json.Unmarshal([]byte(responseBody), &deleteCloudAccountResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing delete CloudAccount JSON response for cloudAccountId: %s, %s\n", cloudAccountId, err)
	}
	if deleteCloudAccountResponse.Errors != nil {
		return nil, fmt.Errorf("errors found in json response: %s", responseBody)
	}

	return &deleteCloudAccountResponse, nil
}
