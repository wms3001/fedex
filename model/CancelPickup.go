package model

type CancelPickup struct {
	AssociatedAccountNumber struct {
		Value string `json:"value"`
	} `json:"associatedAccountNumber"`
	PickupConfirmationCode string `json:"pickupConfirmationCode"`
	Remarks                string `json:"remarks"`
	CarrierCode            string `json:"carrierCode"`
	AccountAddressOfRecord struct {
		StreetLines           []string `json:"streetLines"`
		UrbanizationCode      string   `json:"urbanizationCode"`
		City                  string   `json:"city"`
		StateOrProvinceCode   string   `json:"stateOrProvinceCode"`
		PostalCode            string   `json:"postalCode"`
		CountryCode           string   `json:"countryCode"`
		Residential           bool     `json:"residential"`
		AddressClassification string   `json:"addressClassification"`
	} `json:"accountAddressOfRecord"`
	ScheduledDate string `json:"scheduledDate"`
	Location      string `json:"location"`
}
