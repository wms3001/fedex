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
6. 全球交易api
```go
trade := model.GlobalTrade{}
trade.CarrierCode = "FDXE"
trade.OriginAddress.CountryCode = "CN"
trade.DestinationAddress.CountryCode = "US"
res := fedex.GlobalTrade(trade)
```
7. 检查取件是否可用
```go
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
```
8. 创建取件任务
```go
{
  "associatedAccountNumber": {
    "value": "Your account number"
  },
  "originDetail": {
    "pickupAddressType": "ACCOUNT",
    "pickupLocation": {
      "contact": {
        "companyName": "Fedex",
        "personName": "John Taylor",
        "phoneNumber": "7194446666",
        "phoneExtension": "phone extension"
      },
      "address": {
        "streetLines": [
          "123 Ship Street",
          "Suite 302"
        ],
        "urbanizationCode": "URB FAIR OAKS",
        "city": "Memphis",
        "stateOrProvinceCode": "TN",
        "postalCode": "38017",
        "countryCode": "US",
        "residential": false,
        "addressClassification": "MIXED"
      },
      "accountNumber": {
        "value": "XXX289837"
      },
      "deliveryInstructions": "deliveryInstructions"
    },
    "readyDateTimestamp": "2020-04-21T11:00:00Z",
    "customerCloseTime": "18:00:00",
    "pickupDateType": "SAME_DAY",
    "packageLocation": "FRONT",
    "buildingPart": "APARTMENT",
    "buildingPartDescription": "111",
    "earlyPickup": false,
    "suppliesRequested": "Supplies requested by customer",
    "geographicalPostalCode": "geographicalPostalCode"
  },
  "associatedAccountNumberType": "FEDEX_GROUND",
  "totalWeight": {
    "units": "KG",
    "value": 20
  },
  "packageCount": 5,
  "carrierCode": "FDXE",
  "accountAddressOfRecord": {
    "streetLines": [
      "123 Ship Street"
    ],
    "city": "Memphis",
    "stateOrProvinceCode": "TN",
    "postalCode": "38017",
    "countryCode": "US",
    "residential": false,
    "addressClassification": "MIXED"
  },
  "remarks": "Please ring bell at loading dock.",
  "countryRelationships": "DOMESTIC",
  "pickupType": "ON_CALL, PACKAGE_RETURN_PROGRAM, REGULAR_STOP.",
  "trackingNumber": "795803657326",
  "commodityDescription": "This field contains CommodityDescription",
  "expressFreightDetail": {
    "truckType": "DROP_TRAILER_AGREEMENT",
    "service": "FEDEX_1_DAY_FREIGHT",
    "trailerLength": "TRAILER_28_FT",
    "bookingNumber": "1234AGTT",
    "dimensions": {
      "length": 20,
      "width": 15,
      "height": 12,
      "units": "CM"
    }
  },
  "oversizePackageCount": 2,
  "pickupNotificationDetail": {
    "emailDetails": [
      {
        "address": "sample@gmail.com",
        "locale": "en_US"
      }
    ],
    "format": "HTML",
    "userMessage": "This is the user message"
  }
}
```
```go
pickUp := model.CreatePickup{}
res := fedex.CheckPickup(pickUp)
log.Println(res)
```
9. 取消取件任务
request 
```go
{
  "associatedAccountNumber": {
    "value": "Your account number"
  },
  "pickupConfirmationCode": "7",
  "remarks": "This is my remarks",
  "carrierCode": "FDXE",
  "accountAddressOfRecord": {
    "streetLines": [
      "123 Ship Street"
    ],
    "urbanizationCode": "URB FAIR OAKS",
    "city": "Memphis",
    "stateOrProvinceCode": "ON",
    "postalCode": "38017",
    "countryCode": "US",
    "residential": false,
    "addressClassification": "MIXED"
  },
  "scheduledDate": "2019-10-15",
  "location": "LOSA"
}
```
```go
cancelPickup := model.CancelPickup{}
res := fedex.CancelPickup(cancelPickup)
log.Println(res)
```
10. 地址验证
```go
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
```