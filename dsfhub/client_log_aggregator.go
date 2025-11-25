package dsfhub

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const endpointLogAggregators = "/log-aggregators"

// CreateLogAggregator adds a log aggregator to DSF
func (c *Client) CreateLogAggregator(logAggregator ResourceWrapper) (*ResourceWrapper, error) {
	log.Printf("[INFO] Adding LogAggregator ServerType: %s | AssetID: %s | to gatewayID: %s\n", logAggregator.Data.ServerType, logAggregator.Data.AssetData.AssetID, logAggregator.Data.GatewayID)

	logAggregatorJSON, err := json.Marshal(logAggregator)
	log.Printf("[DEBUG] Adding LogAggregator - JSON: %s\n", logAggregatorJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to JSON marshal LogAggregator: %s\n", err)
	}

	resp, err := c.MakeCallWithQueryParams(http.MethodPost, endpointLogAggregators, logAggregatorJSON, c.config.Params)
	if err != nil {
		return nil, fmt.Errorf("error adding LogAggregator of serverType: %s and gatewayID: %s | err: %s\n", logAggregator.Data.ServerType, logAggregator.Data.GatewayID, err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] Add LogAggregator JSON response: %s\n", string(responseBody))

	// Parse the JSON
	var createLogAggregatorResponse ResourceWrapper
	err = json.Unmarshal([]byte(responseBody), &createLogAggregatorResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing add LogAggregator JSON response serverType: %s and gatewayID: %s | err: %s\n", logAggregator.Data.ServerType, logAggregator.Data.GatewayID, err)
	}
	if createLogAggregatorResponse.Errors != nil {
		return nil, fmt.Errorf("errors found in json response: %s", responseBody)
	}
	return &createLogAggregatorResponse, nil
}

// ReadLogAggregator gets the LogAggregator by ID
func (c *Client) ReadLogAggregator(logAggregatorId string) (*ResourceWrapper, error) {
	log.Printf("[INFO] Getting LogAggregator for logAggregatorId: %s\n", logAggregatorId)

	reqURL := fmt.Sprintf(endpointLogAggregators+"/%s", url.PathEscape(logAggregatorId))
	resp, err := c.MakeCall(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading LogAggregator for logAggregatorId: %s | err: %s\n", logAggregatorId, err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] ReadLogAggregator JSON response: %s\n", string(responseBody))

	// Parse the JSON
	var readLogAggregatorResponse ResourceWrapper
	err = json.Unmarshal([]byte(responseBody), &readLogAggregatorResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing LogAggregator JSON response for logAggregatorId: %s | responseBody: %s err: %s\n", logAggregatorId, responseBody, err)
	}
	if readLogAggregatorResponse.Errors != nil {
		return nil, fmt.Errorf("errors found in json response: %s", responseBody)
	}

	return &readLogAggregatorResponse, nil
}

// ReadLogAggregators all LogAggregators
func (c *Client) ReadLogAggregators() (*ResourcesWrapper, error) {
	log.Printf("[INFO] Getting LogAggregators\n")

	resp, err := c.MakeCall(http.MethodGet, endpointLogAggregators, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading LogAggregators | err: %s\n", err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] ReadLogAggregators JSON response: %s\n", string(responseBody))

	// Parse the JSON
	var readLogAggregatorsResponse ResourcesWrapper
	err = json.Unmarshal([]byte(responseBody), &readLogAggregatorsResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing LogAggregators JSON response: %s err: %s\n", responseBody, err)
	}
	if readLogAggregatorsResponse.Errors != nil {
		return nil, fmt.Errorf("errors found in json response: %s", responseBody)
	}

	return &readLogAggregatorsResponse, nil
}

// UpdateLogAggregator will update a specific LogAggregator record in DSF referenced by the logAggregatorId
func (c *Client) UpdateLogAggregator(logAggregatorId string, logAggregatorData ResourceWrapper) (*ResourceWrapper, error) {
	log.Printf("[INFO] Updating LogAggregator with logAggregatorId: %s)\n", logAggregatorId)

	logAggregatorJSON, err := json.Marshal(logAggregatorData)
	log.Printf("[DEBUG] Adding LogAggregator - JSON: %s\n", logAggregatorJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to JSON marshal logAggregator: %s", err)
	}

	reqURL := fmt.Sprintf(endpointLogAggregators+"/%s", url.PathEscape(logAggregatorId))
	resp, err := c.MakeCallWithQueryParams(http.MethodPut, reqURL, logAggregatorJSON, c.config.Params)
	if err != nil {
		return nil, fmt.Errorf("error updating LogAggregator with logAggregatorId: %s | err: %s\n", logAggregatorId, err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] DSF update LogAggregator JSON response: %s\n", string(responseBody))

	// Parse the JSON
	var updateLogAggregatorResponse ResourceWrapper
	err = json.Unmarshal([]byte(responseBody), &updateLogAggregatorResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing update LogAggregator JSON response for LogAggregatorId: %s | err: %s\n", logAggregatorId, err)
	}
	if updateLogAggregatorResponse.Errors != nil {
		return nil, fmt.Errorf("errors found in json response: %s", responseBody)
	}

	return &updateLogAggregatorResponse, nil
}

// DeleteLogAggregator deletes a LogAggregator in DSF
func (c *Client) DeleteLogAggregator(logAggregatorId string) (*ResourceResponse, error) {
	log.Printf("[INFO] Deleting LogAggregator with logAggregatorId: %s\n", logAggregatorId)

	reqURL := fmt.Sprintf(endpointLogAggregators+"/%s", url.PathEscape(logAggregatorId))
	resp, err := c.MakeCall(http.MethodDelete, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error deleting LogAggregator for logAggregatorId: %s, %s\n", logAggregatorId, err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] DSF delete LogAggregator with JSON response: %s\n", string(responseBody))

	// Parse the JSON
	var deleteLogAggregatorResponse ResourceResponse
	err = json.Unmarshal([]byte(responseBody), &deleteLogAggregatorResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing delete LogAggregator JSON response for logAggregatorId: %s, %s\n", logAggregatorId, err)
	}
	if deleteLogAggregatorResponse.Errors != nil {
		return nil, fmt.Errorf("errors found in json response: %s", responseBody)
	}

	return &deleteLogAggregatorResponse, nil
}
