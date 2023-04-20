package model

type GroundEndRequest struct {
	CloseReqType  string `json:"closeReqType"`
	AccountNumber struct {
		Value string `json:"value"`
	} `json:"accountNumber"`
	GroundServiceCategory      string `json:"groundServiceCategory"`
	CloseDate                  string `json:"closeDate"`
	CloseDocumentSpecification struct {
		CloseDocumentTypes []string `json:"closeDocumentTypes"`
	} `json:"closeDocumentSpecification"`
}
