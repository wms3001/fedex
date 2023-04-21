package model

type ValidAddress struct {
	Address ValidAddressDetail `json:"address"`
}

type ValidAddressDetail struct {
	StreetLines           []string `json:"streetLines"`
	City                  string   `json:"city"`
	StateOrProvinceCode   string   `json:"stateOrProvinceCode"`
	PostalCode            string   `json:"postalCode"`
	CountryCode           string   `json:"countryCode"`
	AddressVerificationId string   `json:"addressVerificationId"`
}
