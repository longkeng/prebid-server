package openrtb_ext

// ExtImp7xbid defines the contract for bidrequest.imp[i].ext.7xbid
type ExtImp7xbid struct {
	placementId  string `json:"placementid"`
	currency     string `json:"cur"`
	userAgent    string `json:"ua"`
	location     string `json:"loc"`
	topframe     int    `json:"topframe"`
	screenWidth  int    `json:"sw"`
	screenHeight int    `json:"sh"`
	cb           int    `json:"cb"`
	tpaf         int    `json:"tpaf"`
	cks          int    `json:"cks"`
	requestId    string `json:"requestid"`
}
