package dsfhub

import (
	"errors"
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

// Client configures and returns a fully initialized DSF Client
func (c *Config) Client() (interface{}, error) {
	// Check DSFToken env var
	if strings.TrimSpace(c.DSFHUBToken) == "" {
		return nil, errors.New(missingAPITokenMessage)
	}
	// Check DSFHost env var
	if strings.TrimSpace(c.DSFHUBHost) == "" {
		return nil, errors.New(missingDSFHostMessage)
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
