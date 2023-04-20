package model

type UploadDocument struct {
	WorkflowName string `json:"workflowName"`
	CarrierCode  string `json:"carrierCode"`
	Name         string `json:"name"`
	ContentType  string `json:"contentType"`
	Meta         struct {
		ShipDocumentType        string `json:"shipDocumentType"`
		FormCode                string `json:"formCode"`
		TrackingNumber          string `json:"trackingNumber"`
		ShipmentDate            string `json:"shipmentDate"`
		OriginLocationCode      string `json:"originLocationCode"`
		OriginCountryCode       string `json:"originCountryCode"`
		DestinationLocationCode string `json:"destinationLocationCode"`
		DestinationCountryCode  string `json:"destinationCountryCode"`
	} `json:"meta"`
}
