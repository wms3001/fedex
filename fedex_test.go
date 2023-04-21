package fedex

import (
	"encoding/json"
	"github.com/wms3001/fedex/model"
	"log"
	"testing"
)

//var token string
//token = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzY29wZSI6WyJDWFMiXSwiUGF5bG9hZCI6eyJjbGllbnRJZGVudGl0eSI6eyJjbGllbnRLZXkiOiJsN2RmYjJjYWJlMDAxYzQyODZiZGY2ODBkNmFiYTA3MmRmIn0sImF1dGhlbnRpY2F0aW9uUmVhbG0iOiJDTUFDIiwiYWRkaXRpb25hbElkZW50aXR5Ijp7InRpbWVTdGFtcCI6IjE4LUFwci0yMDIzIDAyOjIyOjIxIEVTVCIsImdyYW50X3R5cGUiOiJjbGllbnRfY3JlZGVudGlhbHMiLCJhcGltb2RlIjoiU2FuZGJveCIsImN4c0lzcyI6Imh0dHBzOi8vY3hzYXV0aHNlcnZlci1zdGFnaW5nLmFwcC5wYWFzLmZlZGV4LmNvbS90b2tlbi9vYXV0aDIifSwicGVyc29uYVR5cGUiOiJEaXJlY3RJbnRlZ3JhdG9yX0IyQiJ9LCJleHAiOjE2ODE4MDYxNDEsImp0aSI6Ijk0YjBjMGZlLTRjY2EtNDUyZi1iYzBkLWM2NzM3ODBhY2I2NCJ9.jg3nycURIQcUpOwzqkoWWuuUQxFyq19yqZrv1G1vitwBg8_MDEiK___rh98sixrzdG683lyxKN9OwzYmJqMMwCBEy8S7-1Jej-2YrrOroldBc1rVgkBr6tkABoX6aMrYw1b173cLOJ8Kzgbhw5GOaSRATWIVIBN8mPZrPGcBKQIGl12E9cMlR7lmG3AdV7xHhpTLFNE7QUO1vaK8KUax_zXrh-H6UZ36AgU45R7hCFR9PDbOiKLTpaXI0kuD8Iij__XoZQOQJlgvI5D8Nj1WYkJwJR83vxnGtm3d9uHc-wjMQdWRc9PEt6aD3iaeQYTdkcGXYy_15i5TtciY5FVZ-McfDYA3LknLTNBhd2vt8zaz5-bQAwNf73t1661Wij74cjmSWYa3TIyqaaIMuZSuHaHbBoUYBg_KdV4Fgzfnvvy6yHMlSvUKLroSDKy-ib5zPaQ-iIkKsyfzjMNQs8QSCeiD6tLWDxDKRC599FCAevN0GF0EHx7yRPAqmOz4Rnqij-sjlaOMdFi50gx3PN0fdo8EbpP92yYTTgLNWrrvPnFKAPUClzRwBNItzreVacgxW5crZtl3sdk2S202ASlnvqSsnA-k1Nnlf8CkhlrSc3oKwa1rPk4Of7oNPn9VK-OHt4EcgTi3SWyjH7xb3zkUnaimkR3bgn5WZpeZT4Td9K8"

var fedex Fedex

func TestFedex_Auth(t *testing.T) {
	fedex.Url = "https://apis-sandbox.fedex.com"
	fedex.Client_Id = "l7dfb2cabe001c4286bdf680d6aba072df"
	fedex.Client_Secret = "4de70663756d4b6d9184acf9462789c0"
	fedex.Grant_Type = "client_credentials"
	fedex.Auth()
	str, _ := json.Marshal(fedex)
	log.Println(string(str))
}

func TestFedex_Location(t *testing.T) {
	location := model.LocationRequest{}
	location.Location.Address.CountryCode = "CN"
	location.Location.Address.City = "ShenZhen"
	res := fedex.Location(location)
	log.Println(res)
}

func TestFedex_GroundEndOfDayClose(t *testing.T) {
	groundEnd := model.GroundEndRequest{}
	res := fedex.GroundEndOfDayClose(groundEnd)
	log.Println(res)
}

func TestFedex_UploadDocument(t *testing.T) {
	fedex.DocUrl = "https://documentapitest.prod.fedex.com/sandbox"
	uploadDocument := model.UploadDocument{}
	path := "/home/file"
	uploadDocument.WorkflowName = "ETDPreshipment"
	uploadDocument.ContentType = "text/plain"
	uploadDocument.CarrierCode = "FDXE"
	uploadDocument.Name = "file"
	uploadDocument.Meta.ShipDocumentType = "COMMERCIAL_INVOICE"
	uploadDocument.Meta.DestinationCountryCode = "US"
	uploadDocument.Meta.FormCode = "USMCA"
	uploadDocument.Meta.OriginCountryCode = "CN"
	uploadDocument.Meta.DestinationLocationCode = "GVTKK"
	uploadDocument.Meta.OriginLocationCode = "ZHN"
	res := fedex.UploadDocument(uploadDocument, path)
	log.Println(res)
}

func TestFedex_UploadImage(t *testing.T) {
	fedex.DocUrl = "https://documentapitest.prod.fedex.com/sandbox"
	uploadImage := model.UploadImage{}
	path := "/home/IMAGE_3.png"
	uploadImage.Document.ContentType = "image/png"
	uploadImage.Document.ReferenceId = "IMAGE_3"
	uploadImage.Document.Name = "IMAGE_3.png"
	uploadImage.Document.Meta.ImageIndex = "IMAGE_3"
	uploadImage.Document.Meta.ImageType = "SIGNATURE"
	uploadImage.Rules.WorkflowName = "LetterheadSignature"
	res := fedex.UploadImage(uploadImage, path)
	log.Println(res)
}

func TestFedex_GlobalTrade(t *testing.T) {
	trade := model.GlobalTrade{}
	trade.CarrierCode = "FDXE"
	trade.OriginAddress.CountryCode = "CN"
	trade.DestinationAddress.CountryCode = "US"
	res := fedex.GlobalTrade(trade)
	log.Println(res)
}

func TestFedex_CheckPickupAvailability(t *testing.T) {
	availability := model.CheckPickupAvailability{}
	availability.AssociatedAccountNumber = "740561073"
	availability.PickupAddress.CountryCode = "CN"
	availability.PickupAddress.PostalCode = "518118"
	availability.PickupRequestType = append(availability.PickupRequestType, "SAME_DAY")
	availability.Carriers = append(availability.Carriers, "FDXE")
	availability.CountryRelationship = "INTERNATIONAL"
	availability.AssociatedAccountNumberType = "FEDEX_EXPRESS"
	availability.PickupType = "ON_CALL"
	availability.DispatchDate = "2023-04-21"
	availability.PackageReadyTime = "15:00:00"
	availability.CustomerCloseTime = "20:00:00"
	res := fedex.CheckPickupAvailability(availability)
	log.Println(res)
}

func TestFedex_CheckPickup(t *testing.T) {
	pickUp := model.CreatePickup{}
	res := fedex.CheckPickup(pickUp)
	log.Println(res)
}

func TestFedex_CancelPickup(t *testing.T) {
	cancelPickup := model.CancelPickup{}
	res := fedex.CancelPickup(cancelPickup)
	log.Println(res)
}

func TestFedex_ValidateAddress(t *testing.T) {
	validate := model.ValidateAddress{}
	address := model.ValidAddress{}
	addressDetail := model.ValidAddressDetail{}
	addressDetail.AddressVerificationId = "string"
	addressDetail.City = "IRVING"
	addressDetail.PostalCode = "75063-8659"
	addressDetail.CountryCode = "US"
	addressDetail.StateOrProvinceCode = "TX"
	addressDetail.StreetLines = append(addressDetail.StreetLines, "7372 PARKRIDGE BLVD")
	addressDetail.StreetLines = append(addressDetail.StreetLines, "APT 286")
	addressDetail.StreetLines = append(addressDetail.StreetLines, "2903 sprank")
	address.Address = addressDetail
	validate.AddressesToValidate = append(validate.AddressesToValidate, address)
	res := fedex.ValidateAddress(validate)
	log.Println(res)
}
