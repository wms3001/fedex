package model

type TrackbyTrackingNumber struct {
	IncludeDetailedScans bool                        `json:"includeDetailedScans"`
	TrackingInfo         []TrackbyTrackingNumberInfo `json:"trackingInfo"`
}

type TrackbyTrackingNumberInfo struct {
	ShipDateBegin      string `json:"shipDateBegin"`
	ShipDateEnd        string `json:"shipDateEnd"`
	TrackingNumberInfo struct {
		TrackingNumber         string `json:"trackingNumber"`
		CarrierCode            string `json:"carrierCode"`
		TrackingNumberUniqueId string `json:"trackingNumberUniqueId"`
	} `json:"trackingNumberInfo"`
}
