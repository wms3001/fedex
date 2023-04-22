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
11. 邮编验证
```go
post := model.ValidatePostal{}
post.PostalCode = "42101"
post.CountryCode = "US"
post.CarrierCode = "FDXE"
post.StateOrProvinceCode = "KY"
post.ShipDate = "2023-04-22"
post.CheckForMismatch = true
res := fedex.ValidatePostal(post)
log.Println(res)
```
12. Track Multiple Piece Shipment
```go
track := model.TrackMultiplePieceShipment{}
track.MasterTrackingNumberInfo.ShipDateBegin = "2023-04-01"
track.MasterTrackingNumberInfo.ShipDateEnd = "2023-04-22"
track.MasterTrackingNumberInfo.TrackingNumberInfo.TrackingNumberUniqueId = "213424234"
track.MasterTrackingNumberInfo.TrackingNumberInfo.CarrierCode = "FDXE"
track.MasterTrackingNumberInfo.TrackingNumberInfo.TrackingNumber = "858488600850"
track.IncludeDetailedScans = true
track.AssociatedType = "STANDARD_MPS"
track.PagingDetails.PagingToken = "38903279038"
track.PagingDetails.ResultsPerPage = 56
res := fedex.TrackMultiplePieceShipment(track)
```
13. Send Notification
```go
send := model.SendNotification{}
res := fedex.SendNotification(send)
log.Println(res)
```
14. Track by Tracking Number
```go
number := model.TrackbyTrackingNumber{}
number.IncludeDetailedScans = true
numberInfo := model.TrackbyTrackingNumberInfo{}
numberInfo.ShipDateBegin = "2023-04-01"
numberInfo.ShipDateEnd = "2023-04-22"
numberInfo.TrackingNumberInfo.TrackingNumberUniqueId = "erwer"
numberInfo.TrackingNumberInfo.CarrierCode = "FDXE"
numberInfo.TrackingNumberInfo.TrackingNumber = "128667043726"
number.TrackingInfo = append(number.TrackingInfo, numberInfo)
res := fedex.TrackbyTrackingNumber(number)
log.Println(res)
```
15. 创建托运单
```go
{
  "mergeLabelDocOption": "LABELS_AND_DOCS",
  "requestedShipment": {
    "shipDatestamp": "2019-10-14",
    "totalDeclaredValue": {
      "amount": 12.45,
      "currency": "USD"
    },
    "shipper": {
      "address": {
        "streetLines": [
          "10 FedEx Parkway",
          "Suite 302"
        ],
        "city": "Beverly Hills",
        "stateOrProvinceCode": "CA",
        "postalCode": "90210",
        "countryCode": "US",
        "residential": false
      },
      "contact": {
        "personName": "John Taylor",
        "emailAddress": "sample@company.com",
        "phoneExtension": "91",
        "phoneNumber": "XXXX567890",
        "companyName": "Fedex"
      },
      "tins": [
        {
          "number": "XXX567",
          "tinType": "FEDERAL",
          "usage": "usage",
          "effectiveDate": "2000-01-23T04:56:07.000+00:00",
          "expirationDate": "2000-01-23T04:56:07.000+00:00"
        }
      ]
    },
    "soldTo": {
      "address": {
        "streetLines": [
          "10 FedEx Parkway",
          "Suite 302"
        ],
        "city": "Beverly Hills",
        "stateOrProvinceCode": "CA",
        "postalCode": "90210",
        "countryCode": "US",
        "residential": false
      },
      "contact": {
        "personName": "John Taylor",
        "emailAddress": "sample@company.com",
        "phoneExtension": "91",
        "phoneNumber": "1234567890",
        "companyName": "Fedex"
      },
      "tins": [
        {
          "number": "123567",
          "tinType": "FEDERAL",
          "usage": "usage",
          "effectiveDate": "2000-01-23T04:56:07.000+00:00",
          "expirationDate": "2000-01-23T04:56:07.000+00:00"
        }
      ],
      "accountNumber": {
        "value": "Your account number"
      }
    },
    "recipients": [
      {
        "address": {
          "streetLines": [
            "10 FedEx Parkway",
            "Suite 302"
          ],
          "city": "Beverly Hills",
          "stateOrProvinceCode": "CA",
          "postalCode": "90210",
          "countryCode": "US",
          "residential": false
        },
        "contact": {
          "personName": "John Taylor",
          "emailAddress": "sample@company.com",
          "phoneExtension": "000",
          "phoneNumber": "XXXX345671",
          "companyName": "FedEx"
        },
        "tins": [
          {
            "number": "123567",
            "tinType": "FEDERAL",
            "usage": "usage",
            "effectiveDate": "2000-01-23T04:56:07.000+00:00",
            "expirationDate": "2000-01-23T04:56:07.000+00:00"
          }
        ],
        "deliveryInstructions": "Delivery Instructions"
      }
    ],
    "recipientLocationNumber": "1234567",
    "pickupType": "USE_SCHEDULED_PICKUP",
    "serviceType": "PRIORITY_OVERNIGHT",
    "packagingType": "YOUR_PACKAGING",
    "totalWeight": 20.6,
    "origin": {
      "address": {
        "streetLines": [
          "10 FedEx Parkway",
          "Suite 302"
        ],
        "city": "Beverly Hills",
        "stateOrProvinceCode": "CA",
        "postalCode": "38127",
        "countryCode": "US",
        "residential": false
      },
      "contact": {
        "personName": "person name",
        "emailAddress": "email address",
        "phoneNumber": "phone number",
        "phoneExtension": "phone extension",
        "companyName": "company name",
        "faxNumber": "fax number"
      }
    },
    "shippingChargesPayment": {
      "paymentType": "SENDER",
      "payor": {
        "responsibleParty": {
          "address": {
            "streetLines": [
              "10 FedEx Parkway",
              "Suite 302"
            ],
            "city": "Beverly Hills",
            "stateOrProvinceCode": "CA",
            "postalCode": "90210",
            "countryCode": "US",
            "residential": false
          },
          "contact": {
            "personName": "John Taylor",
            "emailAddress": "sample@company.com",
            "phoneNumber": "XXXX567890",
            "phoneExtension": "phone extension",
            "companyName": "Fedex",
            "faxNumber": "fax number"
          },
          "accountNumber": {
            "value": "Your account number"
          }
        }
      }
    },
    "shipmentSpecialServices": {
      "specialServiceTypes": [
        "THIRD_PARTY_CONSIGNEE",
        "PROTECTION_FROM_FREEZING"
      ],
      "etdDetail": {
        "attributes": [
          "POST_SHIPMENT_UPLOAD_REQUESTED"
        ],
        "attachedDocuments": [
          {
            "documentType": "PRO_FORMA_INVOICE",
            "documentReference": "DocumentReference",
            "description": "PRO FORMA INVOICE",
            "documentId": "090927d680038c61"
          }
        ],
        "requestedDocumentTypes": [
          "VICS_BILL_OF_LADING",
          "GENERAL_AGENCY_AGREEMENT"
        ]
      },
      "returnShipmentDetail": {
        "returnEmailDetail": {
          "merchantPhoneNumber": "19012635656",
          "allowedSpecialService": [
            "SATURDAY_DELIVERY"
          ]
        },
        "rma": {
          "reason": "Wrong Size or Color"
        },
        "returnAssociationDetail": {
          "shipDatestamp": "2019-10-01",
          "trackingNumber": "123456789"
        },
        "returnType": "PRINT_RETURN_LABEL"
      },
      "deliveryOnInvoiceAcceptanceDetail": {
        "recipient": {
          "address": {
            "streetLines": [
              "23, RUE JOSEPH-DE MA",
              "Suite 302"
            ],
            "city": "Beverly Hills",
            "stateOrProvinceCode": "CA",
            "postalCode": "90210",
            "countryCode": "US",
            "residential": false
          },
          "contact": {
            "personName": "John Taylor",
            "emailAddress": "sample@company.com",
            "phoneExtension": "000",
            "phoneNumber": "1234567890",
            "companyName": "Fedex"
          },
          "tins": [
            {
              "number": "123567",
              "tinType": "FEDERAL",
              "usage": "usage",
              "effectiveDate": "2000-01-23T04:56:07.000+00:00",
              "expirationDate": "2000-01-23T04:56:07.000+00:00"
            }
          ],
          "deliveryInstructions": "Delivery Instructions"
        }
      },
      "internationalTrafficInArmsRegulationsDetail": {
        "licenseOrExemptionNumber": "9871234"
      },
      "pendingShipmentDetail": {
        "pendingShipmentType": "EMAIL",
        "processingOptions": {
          "options": [
            "ALLOW_MODIFICATIONS"
          ]
        },
        "recommendedDocumentSpecification": {
          "types": "ANTIQUE_STATEMENT_EUROPEAN_UNION"
        },
        "emailLabelDetail": {
          "recipients": [
            {
              "emailAddress": "neena@fedex.com",
              "optionsRequested": {
                "options": [
                  "PRODUCE_PAPERLESS_SHIPPING_FORMAT",
                  "SUPPRESS_ACCESS_EMAILS"
                ]
              },
              "role": "SHIPMENT_COMPLETOR",
              "locale": "en_US"
            }
          ],
          "message": "your optional message"
        },
        "attachedDocuments": [
          {
            "documentType": "PRO_FORMA_INVOICE",
            "documentReference": "DocumentReference",
            "description": "PRO FORMA INVOICE",
            "documentId": "090927d680038c61"
          }
        ],
        "expirationTimeStamp": "2020-01-01"
      },
      "holdAtLocationDetail": {
        "locationId": "YBZA",
        "locationContactAndAddress": {
          "address": {
            "streetLines": [
              "10 FedEx Parkway",
              "Suite 302"
            ],
            "city": "Beverly Hills",
            "stateOrProvinceCode": "CA",
            "postalCode": "38127",
            "countryCode": "US",
            "residential": false
          },
          "contact": {
            "personName": "person name",
            "emailAddress": "email address",
            "phoneNumber": "phone number",
            "phoneExtension": "phone extension",
            "companyName": "company name",
            "faxNumber": "fax number"
          }
        },
        "locationType": "FEDEX_ONSITE"
      },
      "shipmentCODDetail": {
        "addTransportationChargesDetail": {
          "rateType": "ACCOUNT",
          "rateLevelType": "BUNDLED_RATE",
          "chargeLevelType": "CURRENT_PACKAGE",
          "chargeType": "COD_SURCHARGE"
        },
        "codRecipient": {
          "address": {
            "streetLines": [
              "10 FedEx Parkway",
              "Suite 302"
            ],
            "city": "Beverly Hills",
            "stateOrProvinceCode": "CA",
            "postalCode": "90210",
            "countryCode": "US",
            "residential": false
          },
          "contact": {
            "personName": "John Taylor",
            "emailAddress": "sample@company.com",
            "phoneExtension": "000",
            "phoneNumber": "XXXX345671",
            "companyName": "Fedex"
          },
          "accountNumber": {
            "value": "Your account number"
          },
          "tins": [
            {
              "number": "123567",
              "tinType": "FEDERAL",
              "usage": "usage",
              "effectiveDate": "2000-01-23T04:56:07.000+00:00",
              "expirationDate": "2000-01-23T04:56:07.000+00:00"
            }
          ]
        },
        "remitToName": "remitToName",
        "codCollectionType": "ANY",
        "financialInstitutionContactAndAddress": {
          "address": {
            "streetLines": [
              "10 FedEx Parkway",
              "Suite 302"
            ],
            "city": "Beverly Hills",
            "stateOrProvinceCode": "CA",
            "postalCode": "38127",
            "countryCode": "US",
            "residential": false
          },
          "contact": {
            "personName": "person name",
            "emailAddress": "email address",
            "phoneNumber": "phone number",
            "phoneExtension": "phone extension",
            "companyName": "company name",
            "faxNumber": "fax number"
          }
        },
        "codCollectionAmount": {
          "amount": 12.45,
          "currency": "USD"
        },
        "returnReferenceIndicatorType": "INVOICE",
        "shipmentCodAmount": {
          "amount": 12.45,
          "currency": "USD"
        }
      },
      "shipmentDryIceDetail": {
        "totalWeight": {
          "units": "LB",
          "value": 10
        },
        "packageCount": 12
      },
      "internationalControlledExportDetail": {
        "licenseOrPermitExpirationDate": "2019-12-03",
        "licenseOrPermitNumber": "11",
        "entryNumber": "125",
        "foreignTradeZoneCode": "US",
        "type": "WAREHOUSE_WITHDRAWAL"
      },
      "homeDeliveryPremiumDetail": {
        "phoneNumber": {
          "areaCode": "901",
          "localNumber": "3575012",
          "extension": "200",
          "personalIdentificationNumber": "98712345"
        },
        "deliveryDate": "2019-06-26",
        "homedeliveryPremiumType": "APPOINTMENT"
      }
    },
    "emailNotificationDetail": {
      "aggregationType": "PER_PACKAGE",
      "emailNotificationRecipients": [
        {
          "name": "Joe Smith",
          "emailNotificationRecipientType": "SHIPPER",
          "emailAddress": "jsmith3@aol.com",
          "notificationFormatType": "TEXT",
          "notificationType": "EMAIL",
          "locale": "en_US",
          "notificationEventType": [
            "ON_PICKUP_DRIVER_ARRIVED",
            "ON_SHIPMENT"
          ]
        }
      ],
      "personalMessage": "your personal message here"
    },
    "expressFreightDetail": {
      "bookingConfirmationNumber": "123456789812",
      "shippersLoadAndCount": 123,
      "packingListEnclosed": true
    },
    "variableHandlingChargeDetail": {
      "rateType": "PREFERRED_CURRENCY",
      "percentValue": 12.45,
      "rateLevelType": "INDIVIDUAL_PACKAGE_RATE",
      "fixedValue": {
        "amount": 24.45,
        "currency": "USD"
      },
      "rateElementBasis": "NET_CHARGE_EXCLUDING_TAXES"
    },
    "customsClearanceDetail": {
      "regulatoryControls": "NOT_IN_FREE_CIRCULATION",
      "brokers": [
        {
          "broker": {
            "address": {
              "streetLines": [
                "10 FedEx Parkway",
                "Suite 302"
              ],
              "city": "Beverly Hills",
              "stateOrProvinceCode": "CA",
              "postalCode": "90210",
              "countryCode": "US",
              "residential": false
            },
            "contact": {
              "personName": "John Taylor",
              "emailAddress": "sample@company.com",
              "phoneNumber": "1234567890",
              "phoneExtension": 91,
              "companyName": "Fedex",
              "faxNumber": 1234567
            },
            "accountNumber": {
              "value": "Your account number"
            },
            "tins": [
              {
                "number": "number",
                "tinType": "FEDERAL",
                "usage": "usage",
                "effectiveDate": "2000-01-23T04:56:07.000+00:00",
                "expirationDate": "2000-01-23T04:56:07.000+00:00"
              }
            ],
            "deliveryInstructions": "deliveryInstructions"
          },
          "type": "IMPORT"
        }
      ],
      "commercialInvoice": {
        "originatorName": "originator Name",
        "comments": [
          "optional comments for the commercial invoice"
        ],
        "customerReferences": [
          {
            "customerReferenceType": "INVOICE_NUMBER",
            "value": "3686"
          }
        ],
        "taxesOrMiscellaneousCharge": {
          "amount": 12.45,
          "currency": "USD"
        },
        "taxesOrMiscellaneousChargeType": "COMMISSIONS",
        "freightCharge": {
          "amount": 12.45,
          "currency": "USD"
        },
        "packingCosts": {
          "amount": 12.45,
          "currency": "USD"
        },
        "handlingCosts": {
          "amount": 12.45,
          "currency": "USD"
        },
        "declarationStatement": "declarationStatement",
        "termsOfSale": "FCA",
        "specialInstructions": "specialInstructions\"",
        "shipmentPurpose": "REPAIR_AND_RETURN",
        "emailNotificationDetail": {
          "emailAddress": "neena@fedex.com",
          "type": "EMAILED",
          "recipientType": "SHIPPER"
        }
      },
      "freightOnValue": "OWN_RISK",
      "dutiesPayment": {
        "payor": {
          "responsibleParty": {
            "address": {
              "streetLines": [
                "10 FedEx Parkway",
                "Suite 302"
              ],
              "city": "Beverly Hills",
              "stateOrProvinceCode": "CA",
              "postalCode": "38127",
              "countryCode": "US",
              "residential": false
            },
            "contact": {
              "personName": "John Taylor",
              "emailAddress": "sample@company.com",
              "phoneNumber": "1234567890",
              "phoneExtension": "phone extension",
              "companyName": "Fedex",
              "faxNumber": "fax number"
            },
            "accountNumber": {
              "value": "Your account number"
            },
            "tins": [
              {
                "number": "number",
                "tinType": "FEDERAL",
                "usage": "usage",
                "effectiveDate": "2000-01-23T04:56:07.000+00:00",
                "expirationDate": "2000-01-23T04:56:07.000+00:00"
              },
              {
                "number": "number",
                "tinType": "FEDERAL",
                "usage": "usage",
                "effectiveDate": "2000-01-23T04:56:07.000+00:00",
                "expirationDate": "2000-01-23T04:56:07.000+00:00"
              }
            ]
          }
        },
        "billingDetails": {
          "billingCode": "billingCode",
          "billingType": "billingType",
          "aliasId": "aliasId",
          "accountNickname": "accountNickname",
          "accountNumber": "Your account number",
          "accountNumberCountryCode": "US"
        },
        "paymentType": "SENDER"
      },
      "commodities": [
        {
          "unitPrice": {
            "amount": 12.45,
            "currency": "USD"
          },
          "additionalMeasures": [
            {
              "quantity": 12.45,
              "units": "KG"
            }
          ],
          "numberOfPieces": 12,
          "quantity": 125,
          "quantityUnits": "Ea",
          "customsValue": {
            "amount": "1556.25",
            "currency": "USD"
          },
          "countryOfManufacture": "US",
          "cIMarksAndNumbers": "87123",
          "harmonizedCode": "0613",
          "description": "description",
          "name": "non-threaded rivets",
          "weight": {
            "units": "KG",
            "value": 68
          },
          "exportLicenseNumber": "26456",
          "exportLicenseExpirationDate": "2023-04-22T02:31:07Z",
          "partNumber": "167",
          "purpose": "BUSINESS",
          "usmcaDetail": {
            "originCriterion": "A"
          }
        }
      ],
      "isDocumentOnly": true,
      "recipientCustomsId": {
        "type": "PASSPORT",
        "value": "123"
      },
      "customsOption": {
        "description": "Description",
        "type": "COURTESY_RETURN_LABEL"
      },
      "importerOfRecord": {
        "address": {
          "streetLines": [
            "10 FedEx Parkway",
            "Suite 302"
          ],
          "city": "Beverly Hills",
          "stateOrProvinceCode": "CA",
          "postalCode": "90210",
          "countryCode": "US",
          "residential": false
        },
        "contact": {
          "personName": "John Taylor",
          "emailAddress": "sample@company.com",
          "phoneExtension": "000",
          "phoneNumber": "XXXX345671",
          "companyName": "Fedex"
        },
        "accountNumber": {
          "value": "Your account number"
        },
        "tins": [
          {
            "number": "123567",
            "tinType": "FEDERAL",
            "usage": "usage",
            "effectiveDate": "2000-01-23T04:56:07.000+00:00",
            "expirationDate": "2000-01-23T04:56:07.000+00:00"
          }
        ]
      },
      "generatedDocumentLocale": "en_US",
      "exportDetail": {
        "destinationControlDetail": {
          "endUser": "dest country user",
          "statementTypes": "DEPARTMENT_OF_COMMERCE",
          "destinationCountries": [
            "USA",
            "India"
          ]
        },
        "b13AFilingOption": "NOT_REQUIRED",
        "exportComplianceStatement": "12345678901234567",
        "permitNumber": "12345"
      },
      "totalCustomsValue": {
        "amount": 12.45,
        "currency": "USD"
      },
      "partiesToTransactionAreRelated": true,
      "declarationStatementDetail": {
        "usmcaLowValueStatementDetail": {
          "countryOfOriginLowValueDocumentRequested": true,
          "customsRole": "EXPORTER"
        }
      },
      "insuranceCharge": {
        "amount": 12.45,
        "currency": "USD"
      }
    },
    "smartPostInfoDetail": {
      "ancillaryEndorsement": "RETURN_SERVICE",
      "hubId": "5015",
      "indicia": "PRESORTED_STANDARD",
      "specialServices": "USPS_DELIVERY_CONFIRMATION"
    },
    "blockInsightVisibility": true,
    "labelSpecification": {
      "labelFormatType": "COMMON2D",
      "labelOrder": "SHIPPING_LABEL_FIRST",
      "customerSpecifiedDetail": {
        "maskedData": [
          "CUSTOMS_VALUE",
          "TOTAL_WEIGHT"
        ],
        "regulatoryLabels": [
          {
            "generationOptions": "CONTENT_ON_SHIPPING_LABEL_ONLY",
            "type": "ALCOHOL_SHIPMENT_LABEL"
          }
        ],
        "additionalLabels": [
          {
            "type": "CONSIGNEE",
            "count": 1
          }
        ],
        "docTabContent": {
          "docTabContentType": "BARCODED",
          "zone001": {
            "docTabZoneSpecifications": [
              {
                "zoneNumber": 0,
                "header": "string",
                "dataField": "string",
                "literalValue": "string",
                "justification": "RIGHT"
              }
            ]
          },
          "barcoded": {
            "symbology": "UCC128",
            "specification": {
              "zoneNumber": 0,
              "header": "string",
              "dataField": "string",
              "literalValue": "string",
              "justification": "RIGHT"
            }
          }
        }
      },
      "printedLabelOrigin": {
        "address": {
          "streetLines": [
            "10 FedEx Parkway",
            "Suite 302"
          ],
          "city": "Beverly Hills",
          "stateOrProvinceCode": "CA",
          "postalCode": "38127",
          "countryCode": "US",
          "residential": false
        },
        "contact": {
          "personName": "person name",
          "emailAddress": "email address",
          "phoneNumber": "phone number",
          "phoneExtension": "phone extension",
          "companyName": "company name",
          "faxNumber": "fax number"
        }
      },
      "labelStockType": "PAPER_85X11_TOP_HALF_LABEL",
      "labelRotation": "UPSIDE_DOWN",
      "imageType": "PDF",
      "labelPrintingOrientation": "TOP_EDGE_OF_TEXT_FIRST",
      "returnedDispositionDetail": true
    },
    "shippingDocumentSpecification": {
      "generalAgencyAgreementDetail": {
        "documentFormat": {
          "provideInstructions": true,
          "optionsRequested": {
            "options": [
              "SUPPRESS_ADDITIONAL_LANGUAGES",
              "SHIPPING_LABEL_LAST"
            ]
          },
          "stockType": "PAPER_LETTER",
          "dispositions": [
            {
              "eMailDetail": {
                "eMailRecipients": [
                  {
                    "emailAddress": "email@fedex.com",
                    "recipientType": "THIRD_PARTY"
                  }
                ],
                "locale": "en_US",
                "grouping": "NONE"
              },
              "dispositionType": "CONFIRMED"
            }
          ],
          "locale": "en_US",
          "docType": "PDF"
        }
      },
      "returnInstructionsDetail": {
        "customText": "This is additional text printed on Return instr",
        "documentFormat": {
          "provideInstructions": true,
          "optionsRequested": {
            "options": [
              "SUPPRESS_ADDITIONAL_LANGUAGES",
              "SHIPPING_LABEL_LAST"
            ]
          },
          "stockType": "PAPER_LETTER",
          "dispositions": [
            {
              "eMailDetail": {
                "eMailRecipients": [
                  {
                    "emailAddress": "email@fedex.com",
                    "recipientType": "THIRD_PARTY"
                  }
                ],
                "locale": "en_US",
                "grouping": "NONE"
              },
              "dispositionType": "CONFIRMED"
            }
          ],
          "locale": "en_US\"",
          "docType": "PNG"
        }
      },
      "op900Detail": {
        "customerImageUsages": [
          {
            "id": "IMAGE_5",
            "type": "SIGNATURE",
            "providedImageType": "SIGNATURE"
          }
        ],
        "signatureName": "Signature Name",
        "documentFormat": {
          "provideInstructions": true,
          "optionsRequested": {
            "options": [
              "SUPPRESS_ADDITIONAL_LANGUAGES",
              "SHIPPING_LABEL_LAST"
            ]
          },
          "stockType": "PAPER_LETTER",
          "dispositions": [
            {
              "eMailDetail": {
                "eMailRecipients": [
                  {
                    "emailAddress": "email@fedex.com",
                    "recipientType": "THIRD_PARTY"
                  }
                ],
                "locale": "en_US",
                "grouping": "NONE"
              },
              "dispositionType": "CONFIRMED"
            }
          ],
          "locale": "en_US",
          "docType": "PDF"
        }
      },
      "usmcaCertificationOfOriginDetail": {
        "customerImageUsages": [
          {
            "id": "IMAGE_5",
            "type": "SIGNATURE",
            "providedImageType": "SIGNATURE"
          }
        ],
        "documentFormat": {
          "provideInstructions": true,
          "optionsRequested": {
            "options": [
              "SUPPRESS_ADDITIONAL_LANGUAGES",
              "SHIPPING_LABEL_LAST"
            ]
          },
          "stockType": "PAPER_LETTER",
          "dispositions": [
            {
              "eMailDetail": {
                "eMailRecipients": [
                  {
                    "emailAddress": "email@fedex.com",
                    "recipientType": "THIRD_PARTY"
                  }
                ],
                "locale": "en_US",
                "grouping": "NONE"
              },
              "dispositionType": "CONFIRMED"
            }
          ],
          "locale": "en_US",
          "docType": "PDF"
        },
        "certifierSpecification": "EXPORTER",
        "importerSpecification": "UNKNOWN",
        "producerSpecification": "SAME_AS_EXPORTER",
        "producer": {
          "address": {
            "streetLines": [
              "10 FedEx Parkway",
              "Suite 302"
            ],
            "city": "Beverly Hills",
            "stateOrProvinceCode": "CA",
            "postalCode": "90210",
            "countryCode": "US",
            "residential": false
          },
          "contact": {
            "personName": "John Taylor",
            "emailAddress": "sample@company.com",
            "phoneExtension": "000",
            "phoneNumber": "XXXX345671",
            "companyName": "Fedex"
          },
          "accountNumber": {
            "value": "Your account number"
          },
          "tins": [
            {
              "number": "123567",
              "tinType": "FEDERAL",
              "usage": "usage",
              "effectiveDate": "2000-01-23T04:56:07.000+00:00",
              "expirationDate": "2000-01-23T04:56:07.000+00:00"
            }
          ]
        },
        "blanketPeriod": {
          "begins": "22-01-2020",
          "ends": "2-01-2020"
        },
        "certifierJobTitle": "Manager"
      },
      "usmcaCommercialInvoiceCertificationOfOriginDetail": {
        "customerImageUsages": [
          {
            "id": "IMAGE_5",
            "type": "SIGNATURE",
            "providedImageType": "SIGNATURE"
          }
        ],
        "documentFormat": {
          "provideInstructions": true,
          "optionsRequested": {
            "options": [
              "SUPPRESS_ADDITIONAL_LANGUAGES",
              "SHIPPING_LABEL_LAST"
            ]
          },
          "stockType": "PAPER_LETTER",
          "dispositions": [
            {
              "eMailDetail": {
                "eMailRecipients": [
                  {
                    "emailAddress": "email@fedex.com",
                    "recipientType": "THIRD_PARTY"
                  }
                ],
                "locale": "en_US",
                "grouping": "NONE"
              },
              "dispositionType": "CONFIRMED"
            }
          ],
          "locale": "en_US",
          "docType": "PDF"
        },
        "certifierSpecification": "EXPORTER",
        "importerSpecification": "UNKNOWN",
        "producerSpecification": "SAME_AS_EXPORTER",
        "producer": {
          "address": {
            "streetLines": [
              "10 FedEx Parkway",
              "Suite 302"
            ],
            "city": "Beverly Hills",
            "stateOrProvinceCode": "CA",
            "postalCode": "90210",
            "countryCode": "US",
            "residential": false
          },
          "contact": {
            "personName": "John Taylor",
            "emailAddress": "sample@company.com",
            "phoneExtension": "000",
            "phoneNumber": "XXXX345671",
            "companyName": "Fedex"
          },
          "accountNumber": {
            "value": "Your account number"
          },
          "tins": [
            {
              "number": "123567",
              "tinType": "FEDERAL",
              "usage": "usage",
              "effectiveDate": "2000-01-23T04:56:07.000+00:00",
              "expirationDate": "2000-01-23T04:56:07.000+00:00"
            }
          ]
        },
        "certifierJobTitle": "Manager"
      },
      "shippingDocumentTypes": [
        "RETURN_INSTRUCTIONS"
      ],
      "certificateOfOrigin": {
        "customerImageUsages": [
          {
            "id": "IMAGE_5",
            "type": "SIGNATURE",
            "providedImageType": "SIGNATURE"
          }
        ],
        "documentFormat": {
          "provideInstructions": true,
          "optionsRequested": {
            "options": [
              "SUPPRESS_ADDITIONAL_LANGUAGES",
              "SHIPPING_LABEL_LAST"
            ]
          },
          "stockType": "PAPER_LETTER",
          "dispositions": [
            {
              "eMailDetail": {
                "eMailRecipients": [
                  {
                    "emailAddress": "email@fedex.com",
                    "recipientType": "THIRD_PARTY"
                  }
                ],
                "locale": "en_US",
                "grouping": "NONE"
              },
              "dispositionType": "CONFIRMED"
            }
          ],
          "locale": "en_US",
          "docType": "PDF"
        }
      },
      "commercialInvoiceDetail": {
        "customerImageUsages": [
          {
            "id": "IMAGE_5",
            "type": "SIGNATURE",
            "providedImageType": "SIGNATURE"
          }
        ],
        "documentFormat": {
          "provideInstructions": true,
          "optionsRequested": {
            "options": [
              "SUPPRESS_ADDITIONAL_LANGUAGES",
              "SHIPPING_LABEL_LAST"
            ]
          },
          "stockType": "PAPER_LETTER",
          "dispositions": [
            {
              "eMailDetail": {
                "eMailRecipients": [
                  {
                    "emailAddress": "email@fedex.com",
                    "recipientType": "THIRD_PARTY"
                  }
                ],
                "locale": "en_US",
                "grouping": "NONE"
              },
              "dispositionType": "CONFIRMED"
            }
          ],
          "locale": "en_US",
          "docType": "PDF"
        }
      }
    },
    "rateRequestType": [
      "LIST",
      "PREFERRED"
    ],
    "preferredCurrency": "USD",
    "totalPackageCount": 25,
    "masterTrackingId": {
      "formId": "0201",
      "trackingIdType": "EXPRESS",
      "uspsApplicationId": "92",
      "trackingNumber": "49092000070120032835"
    },
    "requestedPackageLineItems": [
      {
        "sequenceNumber": "1",
        "subPackagingType": "BUCKET",
        "customerReferences": [
          {
            "customerReferenceType": "INVOICE_NUMBER",
            "value": "3686"
          }
        ],
        "declaredValue": {
          "amount": 12.45,
          "currency": "USD"
        },
        "weight": {
          "units": "KG",
          "value": 68
        },
        "dimensions": {
          "length": 100,
          "width": 50,
          "height": 30,
          "units": "CM"
        },
        "groupPackageCount": 2,
        "itemDescriptionForClearance": "description",
        "contentRecord": [
          {
            "itemNumber": "2876",
            "receivedQuantity": 256,
            "description": "Description",
            "partNumber": "456"
          }
        ],
        "itemDescription": "item description for the package",
        "variableHandlingChargeDetail": {
          "rateType": "PREFERRED_CURRENCY",
          "percentValue": 12.45,
          "rateLevelType": "INDIVIDUAL_PACKAGE_RATE",
          "fixedValue": {
            "amount": 24.45,
            "currency": "USD"
          },
          "rateElementBasis": "NET_CHARGE_EXCLUDING_TAXES"
        },
        "packageSpecialServices": {
          "specialServiceTypes": [
            "ALCOHOL",
            "NON_STANDARD_CONTAINER"
          ],
          "signatureOptionType": "ADULT",
          "priorityAlertDetail": {
            "enhancementTypes": [
              "PRIORITY_ALERT_PLUS"
            ],
            "content": [
              "string"
            ]
          },
          "signatureOptionDetail": {
            "signatureReleaseNumber": "23456"
          },
          "alcoholDetail": {
            "alcoholRecipientType": "LICENSEE",
            "shipperAgreementType": "Retailer"
          },
          "dangerousGoodsDetail": {
            "cargoAircraftOnly": false,
            "accessibility": "INACCESSIBLE",
            "options": [
              "LIMITED_QUANTITIES_COMMODITIES",
              "ORM_D"
            ]
          },
          "packageCODDetail": {
            "codCollectionAmount": {
              "amount": 12.45,
              "currency": "USD"
            }
          },
          "pieceCountVerificationBoxCount": 0,
          "batteryDetails": [
            {
              "batteryPackingType": "CONTAINED_IN_EQUIPMENT",
              "batteryRegulatoryType": "IATA_SECTION_II",
              "batteryMaterialType": "LITHIUM_METAL"
            }
          ],
          "dryIceWeight": {
            "units": "KG",
            "value": 68
          }
        },
        "trackingNumber": "756477399"
      }
    ]
  },
  "labelResponseOptions": "LABEL",
  "accountNumber": {
    "value": "Your account number"
  },
  "shipAction": "CONFIRM",
  "processingOptionType": "ALLOW_ASYNCHRONOUS",
  "oneLabelAtATime": true
}
```
```go
shipment := model.CreateShipment{}
res := fedex.CreateShipment(shipment)
log.Println(res)
```