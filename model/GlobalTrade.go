package model

type GlobalTrade struct {
	ServiceType         string `json:"serviceType"`
	TotalCommodityValue struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	} `json:"totalCommodityValue"`
	OriginAddress struct {
		CountryCode         string `json:"countryCode"`
		PostalCode          string `json:"postalCode"`
		StateOrProvinceCode string `json:"stateOrProvinceCode"`
	} `json:"originAddress"`
	DestinationAddress struct {
		CountryCode         string `json:"countryCode"`
		PostalCode          string `json:"postalCode"`
		StateOrProvinceCode string `json:"stateOrProvinceCode"`
	} `json:"destinationAddress"`
	ServiceOptionType      []string `json:"serviceOptionType"`
	CustomsClearanceDetail struct {
		RegulatoryControls string `json:"regulatoryControls"`
		InsuranceCharges   struct {
			Amount   float64 `json:"amount"`
			Currency string  `json:"currency"`
		} `json:"insuranceCharges"`
		ImporterOfRecordAccountNumber struct {
			Field1 string `json:"0"`
			Field2 string `json:"1"`
			Field3 string `json:"2"`
			Field4 string `json:"3"`
			Field5 string `json:"4"`
			Field6 string `json:"5"`
			Field7 string `json:"6"`
			Field8 string `json:"7"`
			Field9 string `json:"8"`
		} `json:"importerOfRecordAccountNumber"`
		CustomsValue struct {
			Amount   float64 `json:"amount"`
			Currency string  `json:"currency"`
		} `json:"customsValue"`
		CommercialInvoice struct {
			FreightCharge struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"freightCharge"`
			Purpose string `json:"purpose"`
		} `json:"commercialInvoice"`
		Commodities []struct {
			Quantity     string `json:"quantity"`
			QuantityUOM  string `json:"quantityUOM"`
			CustomsValue struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"customsValue"`
			CountryOfManufacture string `json:"countryOfManufacture"`
			Name                 string `json:"name"`
			HarmonizedCode       string `json:"harmonizedCode"`
			Description          string `json:"description"`
			Weight               struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"weight"`
			CommodityId        int `json:"commodityId"`
			AdditionalMeasures struct {
				Uom      string `json:"uom"`
				Quantity int    `json:"quantity"`
			} `json:"additionalMeasures"`
		} `json:"commodities"`
		DocumentContent string `json:"documentContent"`
	} `json:"customsClearanceDetail"`
	ShipDate    string `json:"shipDate"`
	CarrierCode string `json:"carrierCode"`
	TotalWeight struct {
		Units string `json:"units"`
		Value int    `json:"value"`
	} `json:"totalWeight"`
	IncludeURLReferences bool   `json:"includeURLReferences"`
	ConsolidationType    string `json:"consolidationType"`
	ConsolidationRole    string `json:"consolidationRole"`
}
