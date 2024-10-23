package dsfhub

import (
	"fmt"
	"regexp"
)

func testAccParseResourceAttributeReference(field string) string {
	var expr string = `dsfhub_[A-Za-z0-9_-].+\.[A-Za-z0-9_-].+` //e.g. dsfhub_cloud_account.my-cloud-account, dsfhub_cloud_account.my-cloud-account.asset_id
	var ret string 

	isRef, _ := regexp.Match(expr, []byte(field)) 
	if isRef {
		ret = field
	} else {
		ret = fmt.Sprintf("\"%s\"", field)
		}
	return ret
}
