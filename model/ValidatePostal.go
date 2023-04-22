package model

type ValidatePostal struct {
	CarrierCode         string `json:"carrierCode"`
	CountryCode         string `json:"countryCode"`
	StateOrProvinceCode string `json:"stateOrProvinceCode"`
	PostalCode          string `json:"postalCode"`
	ShipDate            string `json:"shipDate"`
	RoutingCode         string `json:"routingCode"`
	CheckForMismatch    bool   `json:"checkForMismatch"`
}
