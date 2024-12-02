package dsfhub

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// testAccParseResourceAttributeReference parses a terraform field and
// determines whether it is a reference to another resource. If the field is
// a reference, return the input string and if not, return it wrapped in
// double-quotes.
func testAccParseResourceAttributeReference(field string) string {
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

// ConfigCompose can be called to concatenate multiple strings to build test
// configurations
func ConfigCompose(config ...string) string {
	var str strings.Builder

	for _, conf := range config {
		str.WriteString(conf)
	}

	return str.String()
}

// ignoreChangesBlock creates a lifecycle block to be added to a resource config
// containing the ignore_changes feature - an array specifying a list of attribute
// names that may change in the future, but should not affect said resource after
// its creation.
//
// See the following for more details
// https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle#ignore_changes
func createIgnoreChangesBlock(ignoredAttributes []string) string {
	var ignoreChangesBlock, ignoredFields string

	// build list like
	// ignore_changes = [ attribute1, attribute2 ]
	ignoredFields = `[ ` + strings.Join(ignoredAttributes, `, `) + ` ]`
	ignoreChangesBlock = fmt.Sprintf(`
  lifecycle {
    ignore_changes = %[1]s
  }`, ignoredFields)

	return ignoreChangesBlock
}

// ignoreAssetConnectionBlock builds the lifecycle meta argument block to ignore
// fields masked in the connection.
// e.g. client_secret   = "*****", password = "*****"
func ignoreAssetConnectionChangesBlock() string {
	var lifecycleBlock string
	ignoredAttributes := []string{`asset_connection`}
	lifecycleBlock = createIgnoreChangesBlock(ignoredAttributes)
	return lifecycleBlock
}

// checkGatewayId checks that the GATEWAY_ID environment variable is set correctly
// for acceptance tests
func checkGatewayId(t *testing.T) string {
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

	s = regtrue.ReplaceAllString(s, "${1}TRUE") // get em outta the way
	s = regfalse.ReplaceAllString(s, "${1}true")
	s = regTRUE.ReplaceAllString(s, "${1}false")

	return s
}

// validateImportStep returns a test step for validating import on a resource
func validateImportStep(resourceTypeAndName string) resource.TestStep {
	step := resource.TestStep{
		ResourceName:      resourceTypeAndName,
		ImportState:       true,
		ImportStateVerify: true,
	}
	return step
}

// createResourceAttrChecks creates a TestCheckFunc for a map of resources and
// fields to check for each resource
// accepts values like:
// [
//
//	 "resource1": [
//	 	 "field1": "value1",
//		 "field2": "value2",
//	 ],
//	 ...,
//
// ]
func createResourceAttrChecks(attrChecks map[string]map[string]string) resource.TestCheckFunc {
	var checks []resource.TestCheckFunc
	for resourceName, resourceChecks := range attrChecks {
		for field, value := range resourceChecks {
			check := resource.TestCheckResourceAttr(resourceName, field, value)
			checks = append(checks, check)
		}
	}
	return resource.ComposeTestCheckFunc(checks...)
}

// disconnectedResourceAttrChecks updates a list of resource attribute checks
// to account for the resource being disconnected from gateway.
//
// For example, updating "audit_pull_enabled" to "false", or "gateway_service"
// to an empty string.
func disconnectedResourceAttrChecks(attrChecks map[string]map[string]string) resource.TestCheckFunc {
	var checks []resource.TestCheckFunc
	for resourceName, resourceChecks := range attrChecks {
		for field, value := range resourceChecks {
			switch field {
			case "audit_pull_enabled":
				value = "false"
			case "gateway_service":
				value = ""
			}
			check := resource.TestCheckResourceAttr(resourceName, field, value)
			checks = append(checks, check)
		}
	}
	return resource.ComposeTestCheckFunc(checks...)
}

// connectDisconnectGatewaySteps returns test steps to connect a resource to
// gateway and then disconnect it.
func connectDisconnectGatewaySteps(config string, initialAttrChecks map[string]map[string]string, refreshAttrChecks map[string]map[string]string, validate bool, resourceTypeAndName string) []resource.TestStep {
	var steps []resource.TestStep

	// connect to gateway and check attributes
	step := resource.TestStep{
		Config: config,
		Check:  createResourceAttrChecks(initialAttrChecks),
	}
	steps = append(steps, step)

	// refresh the state and check attributes if specified
	if refreshAttrChecks != nil {
		step = resource.TestStep{
			RefreshState: true,
			Check:        createResourceAttrChecks(refreshAttrChecks),
		}
		steps = append(steps, step)
	}

	// disconnect gateway and check atttributes
	var flippedConfig = flipAuditPullEnabledConfig(config)
	var flippedChecks = disconnectedResourceAttrChecks(initialAttrChecks)
	step = resource.TestStep{
		Config: flippedConfig,
		Check:  flippedChecks,
	}
	steps = append(steps, step)

	// refresh the state and check attributes if specified
	if refreshAttrChecks != nil {
		var flippedChecks = disconnectedResourceAttrChecks(initialAttrChecks)
		step = resource.TestStep{
			RefreshState: true,
			Check:        flippedChecks,
		}
		steps = append(steps, step)
	}

	// if specified, validate import
	if validate {
		steps = append(steps, validateImportStep(resourceTypeAndName))
	}
	return steps
}

func skipTestForKnownIssue(t *testing.T, version string, details string) {
	dsfhubVersion := os.Getenv("DSFHUB_VERSION")
	if dsfhubVersion == version {
		t.Skipf("Skipping test %s for DSFHUB_VERSION '%s', details: '%s'", t.Name(), dsfhubVersion, details)
	}
}
