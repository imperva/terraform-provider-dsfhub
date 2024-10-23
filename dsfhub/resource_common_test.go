package dsfhub

import (
	"fmt"
	"regexp"
)

// handle reference to other resources
func testAccParseResourceReference(field string) string {
	var ret string 
	isRef, _ := regexp.Match("([A-Za-z0-9-_]+).([A-Za-z0-9-_]+).asset_id", []byte(field)) //e.g. dsfhub_cloud_account.my-cloud-account.asset_id
	if isRef {
		ret = field
	} else {
		ret = fmt.Sprintf("\"%s\"", field)
		}
	return ret
}
