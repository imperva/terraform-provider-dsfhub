package dsfhub

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// parseResourceAttributeReference parses a terraform field and
// determines whether it is a reference to another resource. If the field is
// a reference, return the input string and if not, return it wrapped in
// double-quotes.
func parseResourceAttributeReference(field string) string {
	var regExpr string = `dsfhub_[A-Za-z0-9_-].+\.[A-Za-z0-9_-].+` //e.g. dsfhub_cloud_account.my-cloud-account, dsfhub_cloud_account.my-cloud-account.asset_id
	var parsedField string

	isReference, _ := regexp.Match(regExpr, []byte(field))
	if isReference {
		parsedField = field
	} else {
		parsedField = fmt.Sprintf("\"%s\"", field)
	}
	return parsedField
}

// ConfigCompose can be called to concatenate multiple strings to build test configurations
func ConfigCompose(config ...string) string {
	var str strings.Builder

	for _, conf := range config {
		str.WriteString(conf)
	}

	return str.String()
}

// ignoreChangesBlock creates a lifecycle block to be added to a resource config containing the
// ignore_changes feature - an array specifying a list of attribute names that
// may change in the future, but should not affect said resource after its creation.
//
// See the following for more details
// https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle#ignore_changes
//
// Required for ignoring fields masked in the connection
// e.g. client_secret   = "*****", password = "*****"
func ignoreChangesBlock(attributes []string) string {
	if len(attributes) == 0 {
		return ""
	}
	var ignoreChangesBlock, ignoredFields string
	// builds a list like
	// ignore_changes = [ attribute1, attribute2 ]
	ignoredFields = `[ ` + strings.Join(attributes, `, `) + ` ]`
	ignoreChangesBlock = fmt.Sprintf(`
  lifecycle {
    ignore_changes = %[1]s
  }`, ignoredFields)

	return ignoreChangesBlock
}

// getGatewayId verifies that the GATEWAY_ID environment variable is set and retrieves it
func getGatewayId(t *testing.T) string {
	gatewayId := os.Getenv("GATEWAY_ID")
	if gatewayId == "" {
		t.Fatal("GATEWAY_ID environment variable must be set")
	}
	return gatewayId
}

// flipAuditPullEnabledConfig takes a config string and flips any instances of
// the "audit_pull_enabled" field
func flipAuditPullEnabledConfig(s string) string {
	regtrue := regexp.MustCompile(`(audit_pull_enabled\s*=\s*)(true)`)
	regfalse := regexp.MustCompile(`(audit_pull_enabled\s*=\s*)(false)`)
	regTRUE := regexp.MustCompile(`(audit_pull_enabled\s*=\s*)(TRUE)`)

	s = regtrue.ReplaceAllString(s, "${1}TRUE") // flip these first
	s = regfalse.ReplaceAllString(s, "${1}true")
	s = regTRUE.ReplaceAllString(s, "${1}false")

	return s
}

// createValidateImportStep returns a test step for validating import on a resource
func createValidateImportStep(resourceTypeAndName string) resource.TestStep {
	step := resource.TestStep{
		ResourceName:      resourceTypeAndName,
		ImportState:       true,
		ImportStateVerify: true,
	}
	return step
}

// createResourceAttrChecks creates an array of TestCheckFunc for a map of resources and
// fields to check for each resource.
//
// Accepts values like:
//
// [
//
//	 "resource1": [
//	     "field1": "value1",
//		 "field2": "value2",
//	 ],
//	 ...,
//
// ]
//
// If the resource is disconnected from the gateway, set disconnected to true to
// modify resource attributes such as audit_pull_enabled and gateway_service
// that determine whether the resource is disconnected from the gateway.
func createResourceAttrChecks(attrChecks map[string]map[string]string, disconnected bool) resource.TestCheckFunc {
	var checks []resource.TestCheckFunc
	for resourceName, resourceChecks := range attrChecks {
		for field, value := range resourceChecks {
			if disconnected {
				switch field {
				case "audit_pull_enabled":
					value = "false"
				case "gateway_service":
					value = ""
				}
			}
			check := resource.TestCheckResourceAttr(resourceName, field, value)
			checks = append(checks, check)
		}
	}
	return resource.ComposeTestCheckFunc(checks...)
}

// createDebugStep creates a TestStep including a debug message and
// logs the state of the resource. Can be used after a Test Step that changes state
func createDebugStep(config string, resourceTypeAndName string, debugString string) resource.TestStep {
	return resource.TestStep{
		Config: config,
		Check:  debugResourceState(resourceTypeAndName, debugString),
	}
}

// connectDisconnectGatewaySteps returns test steps to connect a resource to the gateway and then disconnect it.
//
// AWS: do not pass in audit_pull_enabled: true for the DS asset because it runs ACTION_log_group_discovery
// which attempts to find an actual log group in AWS. audit_pull_enabled: true propagates to DS asset
// during the Enable audit playbook run on the log group.
//
// Azure/GCP: audit_pull_enabled: true on the DS asset propagates to the log aggregator asset in Enable audit playbook run
// on the DS asset.
func createConnectDisconnectGatewayTestSteps(config string, initialAttrChecks map[string]map[string]string, refreshAttrChecks map[string]map[string]string, validateImport bool, resourcesToValidate []string) []resource.TestStep {
	var steps []resource.TestStep

	// connect to gateway and check attributes
	connectGatewayStep := resource.TestStep{
		Config: config,
		Check:  createResourceAttrChecks(initialAttrChecks, false),
	}
	// steps = append(steps, connectGatewayStep, createDebugStep(config, resourcesToValidate[0], "post connectGatewayStep"))
	steps = append(steps, connectGatewayStep)

	// refresh the state (required to trigger another read of context) and check attributes if specified
	if refreshAttrChecks != nil {
		refreshStep := resource.TestStep{
			RefreshState: true,
			Check:        createResourceAttrChecks(refreshAttrChecks, false),
		}
		steps = append(steps, refreshStep)
	}

	// disconnect gateway and check atttributes
	var flippedConfig = flipAuditPullEnabledConfig(config)
	log.Printf("[DEBUG] Flipped config: %s", flippedConfig)
	var flippedChecks = createResourceAttrChecks(initialAttrChecks, true)
	disconnectGatewayStep := resource.TestStep{
		Config: flippedConfig,
		Check:  flippedChecks,
	}

	steps = append(steps, disconnectGatewayStep)
	// steps = append(steps, disconnectGatewayStep, createDebugStep(flippedConfig, resourcesToValidate[0], "post disconnectGatewayStep"))

	// refresh the state (required to trigger another read of context) and check attributes if specified
	if refreshAttrChecks != nil {
		var flippedChecks = createResourceAttrChecks(initialAttrChecks, true)
		refreshStep := resource.TestStep{
			RefreshState: true,
			Check:        flippedChecks,
		}
		steps = append(steps, refreshStep)
	}

	// if specified, validate import
	if validateImport {
		for _, resourceTypeAndName := range resourcesToValidate {
			steps = append(steps, createValidateImportStep(resourceTypeAndName))
		}
	}
	return steps
}

// createAddRemoveConnectionSteps returns test steps to add one or more connections to a data source resource
// and then remove it or them.
// TODO
func createAddRemoveConnectionSteps(config string, initialAttrChecks map[string]map[string]string, refreshAttrChecks map[string]map[string]string, validateImport bool, resourcesToValidate string) []resource.TestStep {
	return nil
}

// skipTestForKnownIssue takes in a DSF Hub version and skips the test if that
// version matches the one being used to run the test
func skipTestForKnownIssue(t *testing.T, version string, details string, skipIfNoVersion bool) {
	dsfhubVersion := os.Getenv("DSFHUB_VERSION")
	if dsfhubVersion == "" && skipIfNoVersion {
		t.Skipf("No DSFHUB_VERSION set (test is not being run on a Sonar machine). Skipping test %s, details: '%s'", t.Name(), details)
	}
	if dsfhubVersion == version {
		t.Skipf("Skipping test %s for DSFHUB_VERSION '%s', details: '%s'", t.Name(), dsfhubVersion, details)
	}
}

// debugResourceState prints the state of a resource to the console and log
func debugResourceState(resourceName string, debugString string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		fmt.Printf("\nDEBUG: [%s] Resource state for %s: %#v\n", debugString, resourceName, rs.Primary.Attributes)
		log.Printf("\nDEBUG: [%s] Resource state for %s: %#v\n", debugString, resourceName, rs.Primary.Attributes)
		return nil
	}
}

// nullIfEmpty returns "null" if the input string is empty
func nullIfEmpty(s string) string {
	if s == "" {
		return "null"
	}
	return s
}

// TODO
type InitialRefreshChecks struct {
	InitialChecks map[string]map[string]string
	RefreshChecks map[string]map[string]string
}

// Returns maps of initial and refresh checks with expected field and value pairs for connect/disconnect gateway tests.
// audit_pull_enabled is checked for all resources that are passed in by default.
//
// For AWS sources, the initial checks are for the log aggregator resource after it has been created
// and the refresh checks are for the data source resource upon a state refresh, after being created
// and connected to the gateway.
//
// For Azure and GCP sources, the initial checks are for the data source resource after it has been created
// and the refresh checks are for the log aggregator resource upon a state refresh, after being created and connected to the gateway.
//
// State refreshes are required because the Connect/DisconnectGateway playbooks cause a state drift.
func buildConnectDisconnectGatewayChecks(
	initialChecks map[string]map[string]string,
	refreshChecks map[string]map[string]string,
) (map[string]map[string]string, map[string]map[string]string) {
	// Add "audit_pull_enabled": "true" by default to all resources' checks
	for _, checks := range []map[string]map[string]string{initialChecks, refreshChecks} {
		if checks == nil {
			continue
		}
		for _, attrs := range checks {
			if attrs != nil {
				attrs["audit_pull_enabled"] = "true"
			}
		}
	}
	return initialChecks, refreshChecks
}

// createBasicConfigs creates two basic configurations for a resource, with and without a connection
func createBasicConfigs(resourceName string, asset_id string, gatewayId string) (string, string) {
	configNoConnection := ConfigCompose(testAccDSFDataSourceConfig_Basic(
		resourceName,
		testAdminEmail,
		asset_id,
		gatewayId,
		testServerHostName,
		testDSServerType,
	))
	configWithConnection := ConfigCompose(testAccDSFDataSourceConfig_Basic_Connection(
		resourceName,
		testAdminEmail,
		asset_id,
		gatewayId,
		testServerHostName,
		testDSServerType,
	))
	return configNoConnection, configWithConnection
}

// Retrieves the list of "required" fields and "auth_mechanisms" for a given server type using requiredDataSourceFieldsJson
func getDSFDataSourceFields(serverType string) (map[string]interface{}, []interface{}, error) {
	var requiredFields map[string]interface{}
	err := json.Unmarshal([]byte(requiredDataSourceFieldsJson), &requiredFields)
	if err != nil {
		return nil, nil, err
	}

	serverTypes := requiredFields["ServerTypes"].(map[string]interface{})
	serverTypeFields := serverTypes[serverType].(map[string]interface{})
	authMechanisms := serverTypeFields["auth_mechanisms"].(map[string]interface{})
	required := serverTypeFields["required"].([]interface{})

	return authMechanisms, required, nil
}
