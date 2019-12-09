package bid7x

import (
	"testing"
)

func TestValidParams(t *testing.T) {
	validator, err := openrtb_ext.NewBidderParamsValidator("../..static/bidder-params")
	if err != nil {
		t.Fatalf("Failed top fetch the json-schemas. %v", err)
	}

	

}


var validParams = []string{
	`{"placementId": "12345"}`
}

var invalidParams = []string{
	`{"placementId": 1234}`
}
