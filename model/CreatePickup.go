package model

import "time"

type CreatePickup struct {
	AssociatedAccountNumber struct {
		Value string `json:"value"`
	} `json:"associatedAccountNumber"`
	OriginDetail struct {
		PickupAddressType string `json:"pickupAddressType"`
		PickupLocation    struct {
			Contact struct {
				CompanyName    string `json:"companyName"`
				PersonName     string `json:"personName"`
				PhoneNumber    string `json:"phoneNumber"`
				PhoneExtension string `json:"phoneExtension"`
			} `json:"contact"`
			Address struct {
				StreetLines           []string `json:"streetLines"`
				UrbanizationCode      string   `json:"urbanizationCode"`
				City                  string   `json:"city"`
				StateOrProvinceCode   string   `json:"stateOrProvinceCode"`
				PostalCode            string   `json:"postalCode"`
				CountryCode           string   `json:"countryCode"`
				Residential           bool     `json:"residential"`
				AddressClassification string   `json:"addressClassification"`
			} `json:"address"`
			AccountNumber struct {
				Value string `json:"value"`
			} `json:"accountNumber"`
			DeliveryInstructions string `json:"deliveryInstructions"`
		} `json:"pickupLocation"`
		ReadyDateTimestamp      time.Time `json:"readyDateTimestamp"`
		CustomerCloseTime       string    `json:"customerCloseTime"`
		PickupDateType          string    `json:"pickupDateType"`
		PackageLocation         string    `json:"packageLocation"`
		BuildingPart            string    `json:"buildingPart"`
		BuildingPartDescription string    `json:"buildingPartDescription"`
		EarlyPickup             bool      `json:"earlyPickup"`
		SuppliesRequested       string    `json:"suppliesRequested"`
		GeographicalPostalCode  string    `json:"geographicalPostalCode"`
	} `json:"originDetail"`
	AssociatedAccountNumberType string `json:"associatedAccountNumberType"`
	TotalWeight                 struct {
		Units string `json:"units"`
		Value int    `json:"value"`
	} `json:"totalWeight"`
	PackageCount           int    `json:"packageCount"`
	CarrierCode            string `json:"carrierCode"`
	AccountAddressOfRecord struct {
		StreetLines           []string `json:"streetLines"`
		City                  string   `json:"city"`
		StateOrProvinceCode   string   `json:"stateOrProvinceCode"`
		PostalCode            string   `json:"postalCode"`
		CountryCode           string   `json:"countryCode"`
		Residential           bool     `json:"residential"`
		AddressClassification string   `json:"addressClassification"`
	} `json:"accountAddressOfRecord"`
	Remarks              string `json:"remarks"`
	CountryRelationships string `json:"countryRelationships"`
	PickupType           string `json:"pickupType"`
	TrackingNumber       string `json:"trackingNumber"`
	CommodityDescription string `json:"commodityDescription"`
	ExpressFreightDetail struct {
		TruckType     string `json:"truckType"`
		Service       string `json:"service"`
		TrailerLength string `json:"trailerLength"`
		BookingNumber string `json:"bookingNumber"`
		Dimensions    struct {
			Length int    `json:"length"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
			Units  string `json:"units"`
		} `json:"dimensions"`
	} `json:"expressFreightDetail"`
	OversizePackageCount     int `json:"oversizePackageCount"`
	PickupNotificationDetail struct {
		EmailDetails []struct {
			Address string `json:"address"`
			Locale  string `json:"locale"`
		} `json:"emailDetails"`
		Format      string `json:"format"`
		UserMessage string `json:"userMessage"`
	} `json:"pickupNotificationDetail"`
}
