package dsfhub

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const endpointDsfDataSource = "/data-sources"

// CreateDSFDataSource adds a DSF data source to be monitored DSF
func (c *Client) CreateDSFDataSource(dsfDataSource ResourceWrapper) (*ResourceWrapper, error) {
	log.Printf("[INFO] Adding DSFDataSource: %s to gateway: %s\n", dsfDataSource.Data.ServerType, dsfDataSource.Data.GatewayID)

	//dsfDataSource := DSFDataSource{}
	dsfDataSourceJSON, err := json.Marshal(dsfDataSource)
	log.Printf("[DEBUG] Adding DSFDataSource dsfDataSourceJSON: %s\n", dsfDataSourceJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to JSON marshal dsfDataSource: %s", err)
	}

	resp, err := c.MakeCall(http.MethodPost, endpointDsfDataSource, dsfDataSourceJSON)
	if err != nil {
		log.Printf("[INFO] err.Error(): %s\n", err.Error())
		return nil, fmt.Errorf("error adding DSFDataSource for serverType: %s and gatewayId: %s | err: %s", dsfDataSource.Data.ServerType, dsfDataSource.Data.GatewayID, err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] Add DSFDataSource JSON response: %s\n", string(responseBody))

	// Parse the JSON
	var createDSFDataSourceResponse ResourceWrapper
	err = json.Unmarshal([]byte(responseBody), &createDSFDataSourceResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing add DSFDataSource JSON response serverType: %s and gatewayId: %s | err: %s", dsfDataSource.Data.ServerType, dsfDataSource.Data.GatewayID, err)
	}
	if createDSFDataSourceResponse.Errors != nil {
		return nil, fmt.Errorf("errors found in json response: %s", responseBody)
	}
	return &createDSFDataSourceResponse, nil
}

// ReadDSFDataSource gets the DSF data source by ID
func (c *Client) ReadDSFDataSource(dataSourceId string) (*ResourceWrapper, error) {
	log.Printf("[INFO] Getting DSFDataSource for dataSourceId: %s)\n", dataSourceId)

	reqURL := fmt.Sprintf(endpointDsfDataSource+"/%s", url.PathEscape(dataSourceId))
	resp, err := c.MakeCall(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading DSFDataSource for dataSourceId: %s | err: %s", dataSourceId, err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] ReadDSFDataSource JSON response: %s\n", string(responseBody))

	// Parse the JSON
	var readDSFDataSourceDataResponse ResourceWrapper
	err = json.Unmarshal([]byte(responseBody), &readDSFDataSourceDataResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing DSFDataSource JSON response for dataSourceId: %s | dsfDataSource: %s err: %s", dataSourceId, responseBody, err)
	}

	if readDSFDataSourceDataResponse.Errors != nil {
		return nil, fmt.Errorf("errors found in json response: %s", responseBody)
	}

	return &readDSFDataSourceDataResponse, nil
}

// UpdateDSFDataSource will update a specific data source in DSF referenced by the dataSourceId
func (c *Client) UpdateDSFDataSource(dataSourceId string, dsfDataSourceData ResourceWrapper) (*ResourceWrapper, error) {
	log.Printf("[INFO] Getting DSF data source for dataSourceId: %s)\n", dataSourceId)

	//dsfDataSource := DSFDataSource{}
	dsfDataSourceJSON, err := json.Marshal(dsfDataSourceData)
	log.Printf("[DEBUG] Updating DSFDataSource dsfDataSourceJSON: %s\n", dsfDataSourceJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to JSON marshal dsfDataSource: %s", err)
	}

	reqURL := fmt.Sprintf(endpointDsfDataSource+"/%s", url.PathEscape(dataSourceId))
	resp, err := c.MakeCall(http.MethodPut, reqURL, dsfDataSourceJSON)
	if err != nil {
		return nil, fmt.Errorf("error updating DSFDataSource for dataSourceId: %s | err: %s", dataSourceId, err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] Update DSFDataSource JSON response: %s\n", string(responseBody))

	// Parse the JSON
	var updateDSFDataSourceDataResponse ResourceWrapper
	err = json.Unmarshal([]byte(responseBody), &updateDSFDataSourceDataResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing update DSFDataSource JSON response for dataSourceId: %s | err: %s", dataSourceId, err)
	}

	if updateDSFDataSourceDataResponse.Errors != nil {
		return nil, fmt.Errorf("errors found in json response: %s", responseBody)
	}

	return &updateDSFDataSourceDataResponse, nil
}

// DeleteDSFDataSource deletes a data source in DSF
func (c *Client) DeleteDSFDataSource(dataSourceId string) (*ResourceResponse, error) {
	log.Printf("[INFO] Deleting DSFDataSource with dataSourceId: %s\n", dataSourceId)

	reqURL := fmt.Sprintf(endpointDsfDataSource+"/%s", url.PathEscape(dataSourceId))
	resp, err := c.MakeCall(http.MethodDelete, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error deleting DSFDataSource for dataSourceId: %s, %s", dataSourceId, err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] Delete DSFDataSource with JSON response: %s\n", string(responseBody))

	// Parse the JSON
	var deleteDSFDataSourceResponse ResourceResponse
	err = json.Unmarshal([]byte(responseBody), &deleteDSFDataSourceResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing delete DSFDataSource JSON response for dataSourceId: %s, %s", dataSourceId, err)
	}

	if deleteDSFDataSourceResponse.Errors != nil {
		return nil, fmt.Errorf("errors found in json response: %s", responseBody)
	}

	return &deleteDSFDataSourceResponse, nil
}

// EnableAuditDSFDataSource enables logging for a DSF data source
func (c *Client) EnableAuditDSFDataSource(dataSourceId string) (*UpdateAuditResponse, error) {
	log.Printf("[INFO] Enabling audit for dataSourceId: %v\n", dataSourceId)

	reqURL := fmt.Sprintf(endpointDsfDataSource+"/%s/operations/enable-audit-collection", url.PathEscape(dataSourceId))
	resp, err := c.MakeCall(http.MethodPost, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error enabling audit for dataSourceId: %s | err: %s\n", dataSourceId, err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] Enable audit for DSFDataSource '%v' JSON response: %s\n", dataSourceId, string(responseBody))

	// Parse the JSON
	var enableAuditResponse UpdateAuditResponse
	err = json.Unmarshal([]byte(responseBody), &enableAuditResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing enable audit DSFDataSource JSON response dataSourceId: %s | err: %s\n", dataSourceId, err)
	}
	if enableAuditResponse.Errors != nil {
		return nil, fmt.Errorf("errors found in json response: %s", responseBody)
	}
	return &enableAuditResponse, nil
}

// DisableAuditDSFDataSource enables logging for a DSF data source
func (c *Client) DisableAuditDSFDataSource(dataSourceId string) (*UpdateAuditResponse, error) {
	log.Printf("[INFO] Disabling audit for dataSourceId: %v\n", dataSourceId)

	reqURL := fmt.Sprintf(endpointDsfDataSource+"/%s/operations/disable-audit-collection", url.PathEscape(dataSourceId))
	resp, err := c.MakeCall(http.MethodPost, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error disabling audit for dataSourceId: %s | err: %s\n", dataSourceId, err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] Disable audit for DSFDataSource '%v' JSON response: %s\n", dataSourceId, string(responseBody))

	// Parse the JSON
	var disableAuditResponse UpdateAuditResponse
	err = json.Unmarshal([]byte(responseBody), &disableAuditResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing disable audit DSFDataSource JSON response dataSourceId: %s | err: %s\n", dataSourceId, err)
	}
	if disableAuditResponse.Errors != nil {
		return nil, fmt.Errorf("errors found in json response: %s", responseBody)
	}
	return &disableAuditResponse, nil
}
