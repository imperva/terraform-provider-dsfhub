package dsfhub

import (
	"errors"
	"slices"
	"strings"
)

// Config represents the configuration required for the DSF Client
type Config struct {
	// API Identifier
	DSFHUBToken string

	// API Key
	DSFHUBHost string

	// InsecureSSL
	InsecureSSL bool

	// Params including syncType
	Params map[string]string
}

var missingAPITokenMessage = "DSF HUB API Token must be provided"
var missingDSFHostMessage = "DSF HUB host/API endpoint must be provided"
var validSyncTypes = []string { "SYNC_GW_BLOCKING", "SYNC_GW_NON_BLOCKING", "DO_NOT_SYNC_GW"}
var invalidSyncTypeMessage = "Invalid sync_type. Available values: " + strings.Join(validSyncTypes, ", ")

// Client configures and returns a fully initialized DSF Client
func (c *Config) Client() (interface{}, error) {
	// Check DSFToken 
	if strings.TrimSpace(c.DSFHUBToken) == "" {
		return nil, errors.New(missingAPITokenMessage)
	}
	// Check DSFHost 
	if strings.TrimSpace(c.DSFHUBHost) == "" {
		return nil, errors.New(missingDSFHostMessage)
	}
	// Check sync_type param
	if syncType, exists := c.Params["syncType"]; exists {
		if ! IsValidSyncType(syncType) {
			return nil, errors.New(invalidSyncTypeMessage)
		}
	}

	// Create client
	client := NewClient(c)

	// Verify client credentials
	gatewaysResponse, err := client.Verify()
	client.gateways = gatewaysResponse
	if err != nil {
		return nil, err
	}

	return client, nil
}

func IsValidSyncType (sync_type string) bool {
	return slices.Contains(validSyncTypes, sync_type)
}
