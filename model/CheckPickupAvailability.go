package model

type CheckPickupAvailability struct {
	PickupAddress struct {
		//StreetLines           []string `json:"streetLines"`
		//UrbanizationCode      string   `json:"urbanizationCode"`
		//City                  string   `json:"city"`
		//StateOrProvinceCode   string   `json:"stateOrProvinceCode"`
		PostalCode  string `json:"postalCode"`
		CountryCode string `json:"countryCode"`
		//Residential           bool     `json:"residential"`
		//AddressClassification string   `json:"addressClassification"`
	} `json:"pickupAddress"`
	DispatchDate                string   `json:"dispatchDate"`
	PackageReadyTime            string   `json:"packageReadyTime"`
	CustomerCloseTime           string   `json:"customerCloseTime"`
	PickupType                  string   `json:"pickupType"`
	PickupRequestType           []string `json:"pickupRequestType"`
	NumberOfBusinessDays        int      `json:"numberOfBusinessDays"`
	AssociatedAccountNumber     string   `json:"associatedAccountNumber"`
	AssociatedAccountNumberType string   `json:"associatedAccountNumberType"`
	Carriers                    []string `json:"carriers"`
	CountryRelationship         string   `json:"countryRelationship"`
}
