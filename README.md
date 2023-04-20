# fedex

------

实现了fedex各个接口的请求

------

1. 安装
```go
go get github.com/wms3001/fedex
```
2. 获取access_token
```go
fedex := &Fedex{}
fedex.Url = "https://apis-sandbox.fedex.com"
fedex.Client_Id = "l7dfb2cabe001c4286bdf680d6aba072df"
fedex.Client_Secret = "4de70663756d4b6d9184acf9462789c0"
fedex.Grant_Type = "client_credentials"
fedex.Auth()
str, _ := json.Marshal(fedex)
log.Println(string(str))
```
3. 获取fedex服务站
```go
location := model.LocationRequest{}
location.Location.Address.CountryCode = "CN"
location.Location.Address.City = "ShenZhen"
res := fedex.Location(location)
log.Println(res)
```
4. 交易文件上传
```go
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
```
5. 上传图片
```go
fedex.DocUrl = "https://documentapitest.prod.fedex.com/sandbox"
uploadImage := model.UploadImage{}
path := "/home/IMAGE_2.png"
uploadImage.Document.ContentType = "image/png"
uploadImage.Document.ReferenceId = "231424"
uploadImage.Document.Name = "IMAGE_2.png"
uploadImage.Document.Meta.ImageIndex = "IMAGE_2"
uploadImage.Document.Meta.ImageType = "SIGNATURE"
uploadImage.Rules.WorkflowName = "LetterheadSignature"
res := fedex.UploadImage(uploadImage, path)
```
```go
{
    "output": {
        "meta": {
            "imageType": "SIGNATURE",
            "imageIndex": "IMAGE_1"
        },
        "documentReferenceId": "1234",
        "status": "SUCCESS"
    },
    "customerTransactionId": ""
}

```