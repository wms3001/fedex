package model

type TrackMultiplePieceShipment struct {
	IncludeDetailedScans     bool                        `json:"includeDetailedScans"`
	AssociatedType           string                      `json:"associatedType"`
	MasterTrackingNumberInfo MasterTrackingNumber        `json:"masterTrackingNumberInfo"`
	PagingDetails            MasterTrackingPagingDetails `json:"pagingDetails"`
}

type MasterTrackingNumber struct {
	ShipDateEnd        string                `json:"shipDateEnd"`
	ShipDateBegin      string                `json:"shipDateBegin"`
	TrackingNumberInfo TrackingNumberMessage `json:"trackingNumberInfo"`
}

type TrackingNumberMessage struct {
	TrackingNumberUniqueId string `json:"trackingNumberUniqueId"`
	CarrierCode            string `json:"carrierCode"`
	TrackingNumber         string `json:"trackingNumber"`
}

type MasterTrackingPagingDetails struct {
	ResultsPerPage int    `json:"resultsPerPage"`
	PagingToken    string `json:"pagingToken"`
}
