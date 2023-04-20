package model

type FedexError struct {
	FedexResponse
	Errors []FedexErrorDetail `json:"errors"`
}
