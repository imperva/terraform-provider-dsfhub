package dsfhub

import (
	"fmt"
	"regexp"
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
