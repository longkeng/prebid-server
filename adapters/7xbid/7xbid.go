package bid7x

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mxmCherry/openrtb"
	"github.com/prebid/prebid-server/adapters"
	"github.com/prebid/prebid-server/errortypes"
	"github.com/prebid/prebid-server/openrtb_ext"
)

type bid7xAdapter struct {
	endpoint string
}

// MakeRequests create the object for 7xBid request

func (a *bid7xAdapter) MakeRequests(request *openrtb.BidRequest, reqInfo *adapters.ExtraRequestInfo) ([]*adapters.RequestData, []error) {
	var errs []error
	var adapterRequests []*adapters.RequestData

	adapterReq, errors := a.makeRequest(request)

	if adapterReq != nil {
		adapterRequests = append(adapterRequests, adapterReq)
	}

	errs = append(errs, errors...)

	return adapterRequests, errors
}

// Update the request object to include custome value
func (a *bid7xAdapter) makeRequest(request *openrtb.BidRequest) (*adapters.RequestData, []error) {
	var errs []error

	// Make a copy as we don't want to change the original request
	reqCopy := *request
	if err := preprocess(&reqCopy); err != nil {
		errs = append(errs, err)
	}

	reqJSON, err := json.Marshal(reqCopy)
	if err != nil {
		errs = append(errs, err)
		return nil, errs
	}

	headers := http.Header{}
	headers.Add("Content-Type", "application/json;charset=utf-8")

	return &adapters.RequestData{
		Method:  "POST",
		Uri:     a.endpoint,
		Body:    reqJSON,
		Headers: headers,
	}, errs
}

// Mutate the request to get it ready to send to 7xbid
func preprocess(request *openrtb.BidRequest) error {
	var imp = &request.Imp[0]
	var bidderExt adapters.ExtImpBidder
	if err := json.Unmarshal(imp.Ext, &bidderExt); err != nil {
		return &errortypes.BadInput{
			Message: fmt.Sprintf("Missing bidder ext: %s", err.Error()),
		}
	}

	var bid7xExt openrtb_ext.ExtImp7xbid
	if err := json.Unmarshal(bidderExt.Bidder, &bid7xExt); err != nil {
		return &errortypes.BadInput{
			Message: fmt.Sprintf("Cannot Resolve placementId: %s", err.Error()),
		}
	}

	if len(bid7xExt.PlacementId) < 0 {
		return &errortypes.BadInput{
			Message: "Invalid/Missing placementId",
		}
	}

	return nil
}

// MakeBids make the bids for the bid response
func (a *bid7xAdapter) MakeBids(internalRequest *openrtb.BidRequest, externalRequest *adapters.RequestData, response *adapters.ResponseData) (*adapters.BidderResponse, []error) {
	if response.StatusCode == http.StatusNoContent {
		return nil, nil
	}
	if response.StatusCode == http.StatusBadRequest {
		return nil, []error{
			&errortypes.BadInput{
				Message: fmt.Sprintf("Unexpected status code: %d. Run with request.debug = 1 for more info", response.StatusCode),
			}}
	}
	if response.StatusCode != http.StatusOK {
		return nil, []error{
			&errortypes.BadServerResponse{
				Message: fmt.Sprintf("Unexpected status code: %d. Run request.debug = 1 for more info", response.StatusCode),
			}}
	}

	var bidResp openrtb.BidResponse

	if err := json.Unmarshal(response.Body, &bidResp); err != nil {
		return nil, []error{err}
	}

	bidResponse := adapters.NewBidderResponseWithBidsCapacity(1)

	for _, sb := range bidResp.SeatBid {
		for i := range sb.Bid {
			bid := sb.Bid[i]
			bidResponse.Bids = append(bidResponse.Bids, &adapters.TypedBid{
				Bid:     &bid,
				BidType: getMediaType(bid.ImpID, internalRequest.Imp),
			})
		}
	}
	return bidResponse, nil
}

func getMediaType(impID string, imps []openrtb.Imp) openrtb_ext.BidType {
	bidType := openrtb_ext.BidTypeBanner

	for _, imp := range imps {
		if imp.ID == impID {
			if imp.Video != nil {
				bidType = openrtb_ext.BidTypeVideo
				break
			} else if imp.Native != nil {
				bidType = openrtb_ext.BidTypeNative
				break
			} else {
				bidType = openrtb_ext.BidTypeBanner
				break
			}
		}
	}

	return bidType
}

func New7xBidBidder(endpoint string) *bid7xAdapter {
	return &bid7xAdapter{
		endpoint: endpoint,
	}
}
