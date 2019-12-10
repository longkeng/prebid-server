package bid7x

import (
	"testing"

	"github.com/prebid/prebid-server/adapters/adapterstest"
)

func TestJsonSamples(t *testing.T) {
	//TODO: FIX ME
	adapterstest.RunJSONBidderTest(t, "7xbidtest", New7xBidBidder("http://bidder.7xbid.com/api/v1/prebid/banner"))
}
