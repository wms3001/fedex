package fedex

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/wms3001/fedex/model"
	"log"
)

type Fedex struct {
	Url           string            `json:"url"`
	DocUrl        string            `json:"doc_url"`
	Client_Id     string            `json:"client_Id"`
	Client_Secret string            `json:"client_Secret"`
	Grant_Type    string            `json:"grant_Type""`
	Access_Token  model.AccessToken `json:"access_Token"`
}

// 获取授权
func (fedex *Fedex) Auth() {
	url := fedex.Url + "/oauth/token"
	var body map[string]string
	body = make(map[string]string)
	body["client_id"] = fedex.Client_Id
	body["client_secret"] = fedex.Client_Secret
	body["grant_type"] = fedex.Grant_Type
	res := fedex.PostForm(url, body)
	accessToken := model.AccessToken{}
	json.Unmarshal([]byte(res), &accessToken)
	fedex.Access_Token = accessToken
}

// 获取fedex服务站
func (fedex *Fedex) Location(req model.LocationRequest) string {
	url := fedex.Url + "/location/v1/locations"
	bt, _ := json.Marshal(req)
	rep := fedex.PostJson(url, string(bt))
	return rep
}

func (fedex *Fedex) GroundEndOfDayClose(req model.GroundEndRequest) string {
	url := fedex.Url + "/ship/v1/endofday"
	bt, _ := json.Marshal(req)
	rep := fedex.PostJson(url, string(bt))
	return rep
}

func (fedex *Fedex) UploadDocument(document model.UploadDocument, filePath string) string {
	url := fedex.DocUrl + "/documents/v1/etds/upload"
	var body map[string]string
	body = make(map[string]string)
	up, _ := json.Marshal(document)
	body["document"] = string(up)
	res := fedex.PostFormFile(url, body, filePath)
	return res
}

func (fedex *Fedex) UploadImage(upImage model.UploadImage, imgPath string) string {
	url := fedex.DocUrl + "/documents/v1/lhsimages/upload"
	var body map[string]string
	body = make(map[string]string)
	up, _ := json.Marshal(upImage)
	body["document"] = string(up)
	res := fedex.PostFormFile(url, body, imgPath)
	return res
}

func (fedex *Fedex) GlobalTrade(trade model.GlobalTrade) string {
	url := fedex.Url + "/globaltrade/v1/shipments/regulatorydetails/retrieve"
	req, _ := json.Marshal(trade)
	rep := fedex.PostJson(url, string(req))
	return rep
}

func (fedex *Fedex) CheckPickupAvailability(availability model.CheckPickupAvailability) string {
	url := fedex.Url + "/pickup/v1/pickups/availabilities"
	req, _ := json.Marshal(availability)
	log.Println(string(req))
	rep := fedex.PostJson(url, string(req))
	return rep
}

func (fedex *Fedex) CheckPickup(pickup model.CreatePickup) string {
	url := fedex.Url + "/pickup/v1/pickups"
	req, _ := json.Marshal(pickup)
	rep := fedex.PostJson(url, string(req))
	return rep
}

func (fedex *Fedex) CancelPickup(pickup model.CancelPickup) string {
	url := fedex.Url + "/pickup/v1/pickups/cancel"
	req, _ := json.Marshal(pickup)
	rep := fedex.PutJson(url, string(req))
	return rep
}

func (fedex *Fedex) ValidateAddress(address model.ValidateAddress) string {
	url := fedex.Url + "/address/v1/addresses/resolve"
	req, _ := json.Marshal(address)
	log.Println(string(req))
	rep := fedex.PostJson(url, string(req))
	return rep
}

func (fedex *Fedex) PostJson(url string, body string) string {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+fedex.Access_Token.Access_Token).
		SetBody(body).
		Post(url)
	if err != nil {
		response := model.FedexError{}
		response.TransactionId = "err"
		errd := model.FedexErrorDetail{}
		errd.Code = "-1"
		errd.Message = err.Error()
		response.Errors = append(response.Errors, errd)
		str, _ := json.Marshal(response)
		return string(str)
	} else {
		return string(resp.Body())
	}
}

func (fedex *Fedex) PutJson(url string, body string) string {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+fedex.Access_Token.Access_Token).
		SetBody(body).
		Put(url)
	if err != nil {
		response := model.FedexError{}
		response.TransactionId = "err"
		errd := model.FedexErrorDetail{}
		errd.Code = "-1"
		errd.Message = err.Error()
		response.Errors = append(response.Errors, errd)
		str, _ := json.Marshal(response)
		return string(str)
	} else {
		return string(resp.Body())
	}
}

func (fedex *Fedex) PostFormFile(url string, body map[string]string, filePath string) string {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetHeader("Authorization", "Bearer "+fedex.Access_Token.Access_Token).
		SetFormData(body).
		SetFile("attachment", filePath).
		Post(url)
	if err != nil {
		response := model.FedexError{}
		response.TransactionId = "err"
		errd := model.FedexErrorDetail{}
		errd.Code = "-1"
		errd.Message = err.Error()
		response.Errors = append(response.Errors, errd)
		str, _ := json.Marshal(response)
		return string(str)
	} else {
		return string(resp.Body())
	}
}

func (fedex *Fedex) PostForm(url string, body map[string]string) string {
	client := resty.New()
	resp, err := client.R().
		SetFormData(body).
		Post(url)
	if err != nil {
		response := model.FedexError{}
		response.TransactionId = "err"
		errd := model.FedexErrorDetail{}
		errd.Code = "-1"
		errd.Message = err.Error()
		response.Errors = append(response.Errors, errd)
		str, _ := json.Marshal(response)
		return string(str)
	} else {
		return string(resp.Body())
	}
}
