package model

import "time"

type CreateShipment struct {
	MergeLabelDocOption string `json:"mergeLabelDocOption"`
	RequestedShipment   struct {
		ShipDatestamp      string `json:"shipDatestamp"`
		TotalDeclaredValue struct {
			Amount   float64 `json:"amount"`
			Currency string  `json:"currency"`
		} `json:"totalDeclaredValue"`
		Shipper struct {
			Address struct {
				StreetLines         []string `json:"streetLines"`
				City                string   `json:"city"`
				StateOrProvinceCode string   `json:"stateOrProvinceCode"`
				PostalCode          string   `json:"postalCode"`
				CountryCode         string   `json:"countryCode"`
				Residential         bool     `json:"residential"`
			} `json:"address"`
			Contact struct {
				PersonName     string `json:"personName"`
				EmailAddress   string `json:"emailAddress"`
				PhoneExtension string `json:"phoneExtension"`
				PhoneNumber    string `json:"phoneNumber"`
				CompanyName    string `json:"companyName"`
			} `json:"contact"`
			Tins []struct {
				Number         string    `json:"number"`
				TinType        string    `json:"tinType"`
				Usage          string    `json:"usage"`
				EffectiveDate  time.Time `json:"effectiveDate"`
				ExpirationDate time.Time `json:"expirationDate"`
			} `json:"tins"`
		} `json:"shipper"`
		SoldTo struct {
			Address struct {
				StreetLines         []string `json:"streetLines"`
				City                string   `json:"city"`
				StateOrProvinceCode string   `json:"stateOrProvinceCode"`
				PostalCode          string   `json:"postalCode"`
				CountryCode         string   `json:"countryCode"`
				Residential         bool     `json:"residential"`
			} `json:"address"`
			Contact struct {
				PersonName     string `json:"personName"`
				EmailAddress   string `json:"emailAddress"`
				PhoneExtension string `json:"phoneExtension"`
				PhoneNumber    string `json:"phoneNumber"`
				CompanyName    string `json:"companyName"`
			} `json:"contact"`
			Tins []struct {
				Number         string    `json:"number"`
				TinType        string    `json:"tinType"`
				Usage          string    `json:"usage"`
				EffectiveDate  time.Time `json:"effectiveDate"`
				ExpirationDate time.Time `json:"expirationDate"`
			} `json:"tins"`
			AccountNumber struct {
				Value string `json:"value"`
			} `json:"accountNumber"`
		} `json:"soldTo"`
		Recipients []struct {
			Address struct {
				StreetLines         []string `json:"streetLines"`
				City                string   `json:"city"`
				StateOrProvinceCode string   `json:"stateOrProvinceCode"`
				PostalCode          string   `json:"postalCode"`
				CountryCode         string   `json:"countryCode"`
				Residential         bool     `json:"residential"`
			} `json:"address"`
			Contact struct {
				PersonName     string `json:"personName"`
				EmailAddress   string `json:"emailAddress"`
				PhoneExtension string `json:"phoneExtension"`
				PhoneNumber    string `json:"phoneNumber"`
				CompanyName    string `json:"companyName"`
			} `json:"contact"`
			Tins []struct {
				Number         string    `json:"number"`
				TinType        string    `json:"tinType"`
				Usage          string    `json:"usage"`
				EffectiveDate  time.Time `json:"effectiveDate"`
				ExpirationDate time.Time `json:"expirationDate"`
			} `json:"tins"`
			DeliveryInstructions string `json:"deliveryInstructions"`
		} `json:"recipients"`
		RecipientLocationNumber string  `json:"recipientLocationNumber"`
		PickupType              string  `json:"pickupType"`
		ServiceType             string  `json:"serviceType"`
		PackagingType           string  `json:"packagingType"`
		TotalWeight             float64 `json:"totalWeight"`
		Origin                  struct {
			Address struct {
				StreetLines         []string `json:"streetLines"`
				City                string   `json:"city"`
				StateOrProvinceCode string   `json:"stateOrProvinceCode"`
				PostalCode          string   `json:"postalCode"`
				CountryCode         string   `json:"countryCode"`
				Residential         bool     `json:"residential"`
			} `json:"address"`
			Contact struct {
				PersonName     string `json:"personName"`
				EmailAddress   string `json:"emailAddress"`
				PhoneNumber    string `json:"phoneNumber"`
				PhoneExtension string `json:"phoneExtension"`
				CompanyName    string `json:"companyName"`
				FaxNumber      string `json:"faxNumber"`
			} `json:"contact"`
		} `json:"origin"`
		ShippingChargesPayment struct {
			PaymentType string `json:"paymentType"`
			Payor       struct {
				ResponsibleParty struct {
					Address struct {
						StreetLines         []string `json:"streetLines"`
						City                string   `json:"city"`
						StateOrProvinceCode string   `json:"stateOrProvinceCode"`
						PostalCode          string   `json:"postalCode"`
						CountryCode         string   `json:"countryCode"`
						Residential         bool     `json:"residential"`
					} `json:"address"`
					Contact struct {
						PersonName     string `json:"personName"`
						EmailAddress   string `json:"emailAddress"`
						PhoneNumber    string `json:"phoneNumber"`
						PhoneExtension string `json:"phoneExtension"`
						CompanyName    string `json:"companyName"`
						FaxNumber      string `json:"faxNumber"`
					} `json:"contact"`
					AccountNumber struct {
						Value string `json:"value"`
					} `json:"accountNumber"`
				} `json:"responsibleParty"`
			} `json:"payor"`
		} `json:"shippingChargesPayment"`
		ShipmentSpecialServices struct {
			SpecialServiceTypes []string `json:"specialServiceTypes"`
			EtdDetail           struct {
				Attributes        []string `json:"attributes"`
				AttachedDocuments []struct {
					DocumentType      string `json:"documentType"`
					DocumentReference string `json:"documentReference"`
					Description       string `json:"description"`
					DocumentId        string `json:"documentId"`
				} `json:"attachedDocuments"`
				RequestedDocumentTypes []string `json:"requestedDocumentTypes"`
			} `json:"etdDetail"`
			ReturnShipmentDetail struct {
				ReturnEmailDetail struct {
					MerchantPhoneNumber   string   `json:"merchantPhoneNumber"`
					AllowedSpecialService []string `json:"allowedSpecialService"`
				} `json:"returnEmailDetail"`
				Rma struct {
					Reason string `json:"reason"`
				} `json:"rma"`
				ReturnAssociationDetail struct {
					ShipDatestamp  string `json:"shipDatestamp"`
					TrackingNumber string `json:"trackingNumber"`
				} `json:"returnAssociationDetail"`
				ReturnType string `json:"returnType"`
			} `json:"returnShipmentDetail"`
			DeliveryOnInvoiceAcceptanceDetail struct {
				Recipient struct {
					Address struct {
						StreetLines         []string `json:"streetLines"`
						City                string   `json:"city"`
						StateOrProvinceCode string   `json:"stateOrProvinceCode"`
						PostalCode          string   `json:"postalCode"`
						CountryCode         string   `json:"countryCode"`
						Residential         bool     `json:"residential"`
					} `json:"address"`
					Contact struct {
						PersonName     string `json:"personName"`
						EmailAddress   string `json:"emailAddress"`
						PhoneExtension string `json:"phoneExtension"`
						PhoneNumber    string `json:"phoneNumber"`
						CompanyName    string `json:"companyName"`
					} `json:"contact"`
					Tins []struct {
						Number         string    `json:"number"`
						TinType        string    `json:"tinType"`
						Usage          string    `json:"usage"`
						EffectiveDate  time.Time `json:"effectiveDate"`
						ExpirationDate time.Time `json:"expirationDate"`
					} `json:"tins"`
					DeliveryInstructions string `json:"deliveryInstructions"`
				} `json:"recipient"`
			} `json:"deliveryOnInvoiceAcceptanceDetail"`
			InternationalTrafficInArmsRegulationsDetail struct {
				LicenseOrExemptionNumber string `json:"licenseOrExemptionNumber"`
			} `json:"internationalTrafficInArmsRegulationsDetail"`
			PendingShipmentDetail struct {
				PendingShipmentType string `json:"pendingShipmentType"`
				ProcessingOptions   struct {
					Options []string `json:"options"`
				} `json:"processingOptions"`
				RecommendedDocumentSpecification struct {
					Types string `json:"types"`
				} `json:"recommendedDocumentSpecification"`
				EmailLabelDetail struct {
					Recipients []struct {
						EmailAddress     string `json:"emailAddress"`
						OptionsRequested struct {
							Options []string `json:"options"`
						} `json:"optionsRequested"`
						Role   string `json:"role"`
						Locale string `json:"locale"`
					} `json:"recipients"`
					Message string `json:"message"`
				} `json:"emailLabelDetail"`
				AttachedDocuments []struct {
					DocumentType      string `json:"documentType"`
					DocumentReference string `json:"documentReference"`
					Description       string `json:"description"`
					DocumentId        string `json:"documentId"`
				} `json:"attachedDocuments"`
				ExpirationTimeStamp string `json:"expirationTimeStamp"`
			} `json:"pendingShipmentDetail"`
			HoldAtLocationDetail struct {
				LocationId                string `json:"locationId"`
				LocationContactAndAddress struct {
					Address struct {
						StreetLines         []string `json:"streetLines"`
						City                string   `json:"city"`
						StateOrProvinceCode string   `json:"stateOrProvinceCode"`
						PostalCode          string   `json:"postalCode"`
						CountryCode         string   `json:"countryCode"`
						Residential         bool     `json:"residential"`
					} `json:"address"`
					Contact struct {
						PersonName     string `json:"personName"`
						EmailAddress   string `json:"emailAddress"`
						PhoneNumber    string `json:"phoneNumber"`
						PhoneExtension string `json:"phoneExtension"`
						CompanyName    string `json:"companyName"`
						FaxNumber      string `json:"faxNumber"`
					} `json:"contact"`
				} `json:"locationContactAndAddress"`
				LocationType string `json:"locationType"`
			} `json:"holdAtLocationDetail"`
			ShipmentCODDetail struct {
				AddTransportationChargesDetail struct {
					RateType        string `json:"rateType"`
					RateLevelType   string `json:"rateLevelType"`
					ChargeLevelType string `json:"chargeLevelType"`
					ChargeType      string `json:"chargeType"`
				} `json:"addTransportationChargesDetail"`
				CodRecipient struct {
					Address struct {
						StreetLines         []string `json:"streetLines"`
						City                string   `json:"city"`
						StateOrProvinceCode string   `json:"stateOrProvinceCode"`
						PostalCode          string   `json:"postalCode"`
						CountryCode         string   `json:"countryCode"`
						Residential         bool     `json:"residential"`
					} `json:"address"`
					Contact struct {
						PersonName     string `json:"personName"`
						EmailAddress   string `json:"emailAddress"`
						PhoneExtension string `json:"phoneExtension"`
						PhoneNumber    string `json:"phoneNumber"`
						CompanyName    string `json:"companyName"`
					} `json:"contact"`
					AccountNumber struct {
						Value string `json:"value"`
					} `json:"accountNumber"`
					Tins []struct {
						Number         string    `json:"number"`
						TinType        string    `json:"tinType"`
						Usage          string    `json:"usage"`
						EffectiveDate  time.Time `json:"effectiveDate"`
						ExpirationDate time.Time `json:"expirationDate"`
					} `json:"tins"`
				} `json:"codRecipient"`
				RemitToName                           string `json:"remitToName"`
				CodCollectionType                     string `json:"codCollectionType"`
				FinancialInstitutionContactAndAddress struct {
					Address struct {
						StreetLines         []string `json:"streetLines"`
						City                string   `json:"city"`
						StateOrProvinceCode string   `json:"stateOrProvinceCode"`
						PostalCode          string   `json:"postalCode"`
						CountryCode         string   `json:"countryCode"`
						Residential         bool     `json:"residential"`
					} `json:"address"`
					Contact struct {
						PersonName     string `json:"personName"`
						EmailAddress   string `json:"emailAddress"`
						PhoneNumber    string `json:"phoneNumber"`
						PhoneExtension string `json:"phoneExtension"`
						CompanyName    string `json:"companyName"`
						FaxNumber      string `json:"faxNumber"`
					} `json:"contact"`
				} `json:"financialInstitutionContactAndAddress"`
				CodCollectionAmount struct {
					Amount   float64 `json:"amount"`
					Currency string  `json:"currency"`
				} `json:"codCollectionAmount"`
				ReturnReferenceIndicatorType string `json:"returnReferenceIndicatorType"`
				ShipmentCodAmount            struct {
					Amount   float64 `json:"amount"`
					Currency string  `json:"currency"`
				} `json:"shipmentCodAmount"`
			} `json:"shipmentCODDetail"`
			ShipmentDryIceDetail struct {
				TotalWeight struct {
					Units string `json:"units"`
					Value int    `json:"value"`
				} `json:"totalWeight"`
				PackageCount int `json:"packageCount"`
			} `json:"shipmentDryIceDetail"`
			InternationalControlledExportDetail struct {
				LicenseOrPermitExpirationDate string `json:"licenseOrPermitExpirationDate"`
				LicenseOrPermitNumber         string `json:"licenseOrPermitNumber"`
				EntryNumber                   string `json:"entryNumber"`
				ForeignTradeZoneCode          string `json:"foreignTradeZoneCode"`
				Type                          string `json:"type"`
			} `json:"internationalControlledExportDetail"`
			HomeDeliveryPremiumDetail struct {
				PhoneNumber struct {
					AreaCode                     string `json:"areaCode"`
					LocalNumber                  string `json:"localNumber"`
					Extension                    string `json:"extension"`
					PersonalIdentificationNumber string `json:"personalIdentificationNumber"`
				} `json:"phoneNumber"`
				DeliveryDate            string `json:"deliveryDate"`
				HomedeliveryPremiumType string `json:"homedeliveryPremiumType"`
			} `json:"homeDeliveryPremiumDetail"`
		} `json:"shipmentSpecialServices"`
		EmailNotificationDetail struct {
			AggregationType             string `json:"aggregationType"`
			EmailNotificationRecipients []struct {
				Name                           string   `json:"name"`
				EmailNotificationRecipientType string   `json:"emailNotificationRecipientType"`
				EmailAddress                   string   `json:"emailAddress"`
				NotificationFormatType         string   `json:"notificationFormatType"`
				NotificationType               string   `json:"notificationType"`
				Locale                         string   `json:"locale"`
				NotificationEventType          []string `json:"notificationEventType"`
			} `json:"emailNotificationRecipients"`
			PersonalMessage string `json:"personalMessage"`
		} `json:"emailNotificationDetail"`
		ExpressFreightDetail struct {
			BookingConfirmationNumber string `json:"bookingConfirmationNumber"`
			ShippersLoadAndCount      int    `json:"shippersLoadAndCount"`
			PackingListEnclosed       bool   `json:"packingListEnclosed"`
		} `json:"expressFreightDetail"`
		VariableHandlingChargeDetail struct {
			RateType      string  `json:"rateType"`
			PercentValue  float64 `json:"percentValue"`
			RateLevelType string  `json:"rateLevelType"`
			FixedValue    struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"fixedValue"`
			RateElementBasis string `json:"rateElementBasis"`
		} `json:"variableHandlingChargeDetail"`
		CustomsClearanceDetail struct {
			RegulatoryControls string `json:"regulatoryControls"`
			Brokers            []struct {
				Broker struct {
					Address struct {
						StreetLines         []string `json:"streetLines"`
						City                string   `json:"city"`
						StateOrProvinceCode string   `json:"stateOrProvinceCode"`
						PostalCode          string   `json:"postalCode"`
						CountryCode         string   `json:"countryCode"`
						Residential         bool     `json:"residential"`
					} `json:"address"`
					Contact struct {
						PersonName     string `json:"personName"`
						EmailAddress   string `json:"emailAddress"`
						PhoneNumber    string `json:"phoneNumber"`
						PhoneExtension int    `json:"phoneExtension"`
						CompanyName    string `json:"companyName"`
						FaxNumber      int    `json:"faxNumber"`
					} `json:"contact"`
					AccountNumber struct {
						Value string `json:"value"`
					} `json:"accountNumber"`
					Tins []struct {
						Number         string    `json:"number"`
						TinType        string    `json:"tinType"`
						Usage          string    `json:"usage"`
						EffectiveDate  time.Time `json:"effectiveDate"`
						ExpirationDate time.Time `json:"expirationDate"`
					} `json:"tins"`
					DeliveryInstructions string `json:"deliveryInstructions"`
				} `json:"broker"`
				Type string `json:"type"`
			} `json:"brokers"`
			CommercialInvoice struct {
				OriginatorName     string   `json:"originatorName"`
				Comments           []string `json:"comments"`
				CustomerReferences []struct {
					CustomerReferenceType string `json:"customerReferenceType"`
					Value                 string `json:"value"`
				} `json:"customerReferences"`
				TaxesOrMiscellaneousCharge struct {
					Amount   float64 `json:"amount"`
					Currency string  `json:"currency"`
				} `json:"taxesOrMiscellaneousCharge"`
				TaxesOrMiscellaneousChargeType string `json:"taxesOrMiscellaneousChargeType"`
				FreightCharge                  struct {
					Amount   float64 `json:"amount"`
					Currency string  `json:"currency"`
				} `json:"freightCharge"`
				PackingCosts struct {
					Amount   float64 `json:"amount"`
					Currency string  `json:"currency"`
				} `json:"packingCosts"`
				HandlingCosts struct {
					Amount   float64 `json:"amount"`
					Currency string  `json:"currency"`
				} `json:"handlingCosts"`
				DeclarationStatement    string `json:"declarationStatement"`
				TermsOfSale             string `json:"termsOfSale"`
				SpecialInstructions     string `json:"specialInstructions"`
				ShipmentPurpose         string `json:"shipmentPurpose"`
				EmailNotificationDetail struct {
					EmailAddress  string `json:"emailAddress"`
					Type          string `json:"type"`
					RecipientType string `json:"recipientType"`
				} `json:"emailNotificationDetail"`
			} `json:"commercialInvoice"`
			FreightOnValue string `json:"freightOnValue"`
			DutiesPayment  struct {
				Payor struct {
					ResponsibleParty struct {
						Address struct {
							StreetLines         []string `json:"streetLines"`
							City                string   `json:"city"`
							StateOrProvinceCode string   `json:"stateOrProvinceCode"`
							PostalCode          string   `json:"postalCode"`
							CountryCode         string   `json:"countryCode"`
							Residential         bool     `json:"residential"`
						} `json:"address"`
						Contact struct {
							PersonName     string `json:"personName"`
							EmailAddress   string `json:"emailAddress"`
							PhoneNumber    string `json:"phoneNumber"`
							PhoneExtension string `json:"phoneExtension"`
							CompanyName    string `json:"companyName"`
							FaxNumber      string `json:"faxNumber"`
						} `json:"contact"`
						AccountNumber struct {
							Value string `json:"value"`
						} `json:"accountNumber"`
						Tins []struct {
							Number         string    `json:"number"`
							TinType        string    `json:"tinType"`
							Usage          string    `json:"usage"`
							EffectiveDate  time.Time `json:"effectiveDate"`
							ExpirationDate time.Time `json:"expirationDate"`
						} `json:"tins"`
					} `json:"responsibleParty"`
				} `json:"payor"`
				BillingDetails struct {
					BillingCode              string `json:"billingCode"`
					BillingType              string `json:"billingType"`
					AliasId                  string `json:"aliasId"`
					AccountNickname          string `json:"accountNickname"`
					AccountNumber            string `json:"accountNumber"`
					AccountNumberCountryCode string `json:"accountNumberCountryCode"`
				} `json:"billingDetails"`
				PaymentType string `json:"paymentType"`
			} `json:"dutiesPayment"`
			Commodities []struct {
				UnitPrice struct {
					Amount   float64 `json:"amount"`
					Currency string  `json:"currency"`
				} `json:"unitPrice"`
				AdditionalMeasures []struct {
					Quantity float64 `json:"quantity"`
					Units    string  `json:"units"`
				} `json:"additionalMeasures"`
				NumberOfPieces int    `json:"numberOfPieces"`
				Quantity       int    `json:"quantity"`
				QuantityUnits  string `json:"quantityUnits"`
				CustomsValue   struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"customsValue"`
				CountryOfManufacture string `json:"countryOfManufacture"`
				CIMarksAndNumbers    string `json:"cIMarksAndNumbers"`
				HarmonizedCode       string `json:"harmonizedCode"`
				Description          string `json:"description"`
				Name                 string `json:"name"`
				Weight               struct {
					Units string `json:"units"`
					Value int    `json:"value"`
				} `json:"weight"`
				ExportLicenseNumber         string    `json:"exportLicenseNumber"`
				ExportLicenseExpirationDate time.Time `json:"exportLicenseExpirationDate"`
				PartNumber                  string    `json:"partNumber"`
				Purpose                     string    `json:"purpose"`
				UsmcaDetail                 struct {
					OriginCriterion string `json:"originCriterion"`
				} `json:"usmcaDetail"`
			} `json:"commodities"`
			IsDocumentOnly     bool `json:"isDocumentOnly"`
			RecipientCustomsId struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"recipientCustomsId"`
			CustomsOption struct {
				Description string `json:"description"`
				Type        string `json:"type"`
			} `json:"customsOption"`
			ImporterOfRecord struct {
				Address struct {
					StreetLines         []string `json:"streetLines"`
					City                string   `json:"city"`
					StateOrProvinceCode string   `json:"stateOrProvinceCode"`
					PostalCode          string   `json:"postalCode"`
					CountryCode         string   `json:"countryCode"`
					Residential         bool     `json:"residential"`
				} `json:"address"`
				Contact struct {
					PersonName     string `json:"personName"`
					EmailAddress   string `json:"emailAddress"`
					PhoneExtension string `json:"phoneExtension"`
					PhoneNumber    string `json:"phoneNumber"`
					CompanyName    string `json:"companyName"`
				} `json:"contact"`
				AccountNumber struct {
					Value string `json:"value"`
				} `json:"accountNumber"`
				Tins []struct {
					Number         string    `json:"number"`
					TinType        string    `json:"tinType"`
					Usage          string    `json:"usage"`
					EffectiveDate  time.Time `json:"effectiveDate"`
					ExpirationDate time.Time `json:"expirationDate"`
				} `json:"tins"`
			} `json:"importerOfRecord"`
			GeneratedDocumentLocale string `json:"generatedDocumentLocale"`
			ExportDetail            struct {
				DestinationControlDetail struct {
					EndUser              string   `json:"endUser"`
					StatementTypes       string   `json:"statementTypes"`
					DestinationCountries []string `json:"destinationCountries"`
				} `json:"destinationControlDetail"`
				B13AFilingOption          string `json:"b13AFilingOption"`
				ExportComplianceStatement string `json:"exportComplianceStatement"`
				PermitNumber              string `json:"permitNumber"`
			} `json:"exportDetail"`
			TotalCustomsValue struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"totalCustomsValue"`
			PartiesToTransactionAreRelated bool `json:"partiesToTransactionAreRelated"`
			DeclarationStatementDetail     struct {
				UsmcaLowValueStatementDetail struct {
					CountryOfOriginLowValueDocumentRequested bool   `json:"countryOfOriginLowValueDocumentRequested"`
					CustomsRole                              string `json:"customsRole"`
				} `json:"usmcaLowValueStatementDetail"`
			} `json:"declarationStatementDetail"`
			InsuranceCharge struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"insuranceCharge"`
		} `json:"customsClearanceDetail"`
		SmartPostInfoDetail struct {
			AncillaryEndorsement string `json:"ancillaryEndorsement"`
			HubId                string `json:"hubId"`
			Indicia              string `json:"indicia"`
			SpecialServices      string `json:"specialServices"`
		} `json:"smartPostInfoDetail"`
		BlockInsightVisibility bool `json:"blockInsightVisibility"`
		LabelSpecification     struct {
			LabelFormatType         string `json:"labelFormatType"`
			LabelOrder              string `json:"labelOrder"`
			CustomerSpecifiedDetail struct {
				MaskedData       []string `json:"maskedData"`
				RegulatoryLabels []struct {
					GenerationOptions string `json:"generationOptions"`
					Type              string `json:"type"`
				} `json:"regulatoryLabels"`
				AdditionalLabels []struct {
					Type  string `json:"type"`
					Count int    `json:"count"`
				} `json:"additionalLabels"`
				DocTabContent struct {
					DocTabContentType string `json:"docTabContentType"`
					Zone001           struct {
						DocTabZoneSpecifications []struct {
							ZoneNumber    int    `json:"zoneNumber"`
							Header        string `json:"header"`
							DataField     string `json:"dataField"`
							LiteralValue  string `json:"literalValue"`
							Justification string `json:"justification"`
						} `json:"docTabZoneSpecifications"`
					} `json:"zone001"`
					Barcoded struct {
						Symbology     string `json:"symbology"`
						Specification struct {
							ZoneNumber    int    `json:"zoneNumber"`
							Header        string `json:"header"`
							DataField     string `json:"dataField"`
							LiteralValue  string `json:"literalValue"`
							Justification string `json:"justification"`
						} `json:"specification"`
					} `json:"barcoded"`
				} `json:"docTabContent"`
			} `json:"customerSpecifiedDetail"`
			PrintedLabelOrigin struct {
				Address struct {
					StreetLines         []string `json:"streetLines"`
					City                string   `json:"city"`
					StateOrProvinceCode string   `json:"stateOrProvinceCode"`
					PostalCode          string   `json:"postalCode"`
					CountryCode         string   `json:"countryCode"`
					Residential         bool     `json:"residential"`
				} `json:"address"`
				Contact struct {
					PersonName     string `json:"personName"`
					EmailAddress   string `json:"emailAddress"`
					PhoneNumber    string `json:"phoneNumber"`
					PhoneExtension string `json:"phoneExtension"`
					CompanyName    string `json:"companyName"`
					FaxNumber      string `json:"faxNumber"`
				} `json:"contact"`
			} `json:"printedLabelOrigin"`
			LabelStockType            string `json:"labelStockType"`
			LabelRotation             string `json:"labelRotation"`
			ImageType                 string `json:"imageType"`
			LabelPrintingOrientation  string `json:"labelPrintingOrientation"`
			ReturnedDispositionDetail bool   `json:"returnedDispositionDetail"`
		} `json:"labelSpecification"`
		ShippingDocumentSpecification struct {
			GeneralAgencyAgreementDetail struct {
				DocumentFormat struct {
					ProvideInstructions bool `json:"provideInstructions"`
					OptionsRequested    struct {
						Options []string `json:"options"`
					} `json:"optionsRequested"`
					StockType    string `json:"stockType"`
					Dispositions []struct {
						EMailDetail struct {
							EMailRecipients []struct {
								EmailAddress  string `json:"emailAddress"`
								RecipientType string `json:"recipientType"`
							} `json:"eMailRecipients"`
							Locale   string `json:"locale"`
							Grouping string `json:"grouping"`
						} `json:"eMailDetail"`
						DispositionType string `json:"dispositionType"`
					} `json:"dispositions"`
					Locale  string `json:"locale"`
					DocType string `json:"docType"`
				} `json:"documentFormat"`
			} `json:"generalAgencyAgreementDetail"`
			ReturnInstructionsDetail struct {
				CustomText     string `json:"customText"`
				DocumentFormat struct {
					ProvideInstructions bool `json:"provideInstructions"`
					OptionsRequested    struct {
						Options []string `json:"options"`
					} `json:"optionsRequested"`
					StockType    string `json:"stockType"`
					Dispositions []struct {
						EMailDetail struct {
							EMailRecipients []struct {
								EmailAddress  string `json:"emailAddress"`
								RecipientType string `json:"recipientType"`
							} `json:"eMailRecipients"`
							Locale   string `json:"locale"`
							Grouping string `json:"grouping"`
						} `json:"eMailDetail"`
						DispositionType string `json:"dispositionType"`
					} `json:"dispositions"`
					Locale  string `json:"locale"`
					DocType string `json:"docType"`
				} `json:"documentFormat"`
			} `json:"returnInstructionsDetail"`
			Op900Detail struct {
				CustomerImageUsages []struct {
					Id                string `json:"id"`
					Type              string `json:"type"`
					ProvidedImageType string `json:"providedImageType"`
				} `json:"customerImageUsages"`
				SignatureName  string `json:"signatureName"`
				DocumentFormat struct {
					ProvideInstructions bool `json:"provideInstructions"`
					OptionsRequested    struct {
						Options []string `json:"options"`
					} `json:"optionsRequested"`
					StockType    string `json:"stockType"`
					Dispositions []struct {
						EMailDetail struct {
							EMailRecipients []struct {
								EmailAddress  string `json:"emailAddress"`
								RecipientType string `json:"recipientType"`
							} `json:"eMailRecipients"`
							Locale   string `json:"locale"`
							Grouping string `json:"grouping"`
						} `json:"eMailDetail"`
						DispositionType string `json:"dispositionType"`
					} `json:"dispositions"`
					Locale  string `json:"locale"`
					DocType string `json:"docType"`
				} `json:"documentFormat"`
			} `json:"op900Detail"`
			UsmcaCertificationOfOriginDetail struct {
				CustomerImageUsages []struct {
					Id                string `json:"id"`
					Type              string `json:"type"`
					ProvidedImageType string `json:"providedImageType"`
				} `json:"customerImageUsages"`
				DocumentFormat struct {
					ProvideInstructions bool `json:"provideInstructions"`
					OptionsRequested    struct {
						Options []string `json:"options"`
					} `json:"optionsRequested"`
					StockType    string `json:"stockType"`
					Dispositions []struct {
						EMailDetail struct {
							EMailRecipients []struct {
								EmailAddress  string `json:"emailAddress"`
								RecipientType string `json:"recipientType"`
							} `json:"eMailRecipients"`
							Locale   string `json:"locale"`
							Grouping string `json:"grouping"`
						} `json:"eMailDetail"`
						DispositionType string `json:"dispositionType"`
					} `json:"dispositions"`
					Locale  string `json:"locale"`
					DocType string `json:"docType"`
				} `json:"documentFormat"`
				CertifierSpecification string `json:"certifierSpecification"`
				ImporterSpecification  string `json:"importerSpecification"`
				ProducerSpecification  string `json:"producerSpecification"`
				Producer               struct {
					Address struct {
						StreetLines         []string `json:"streetLines"`
						City                string   `json:"city"`
						StateOrProvinceCode string   `json:"stateOrProvinceCode"`
						PostalCode          string   `json:"postalCode"`
						CountryCode         string   `json:"countryCode"`
						Residential         bool     `json:"residential"`
					} `json:"address"`
					Contact struct {
						PersonName     string `json:"personName"`
						EmailAddress   string `json:"emailAddress"`
						PhoneExtension string `json:"phoneExtension"`
						PhoneNumber    string `json:"phoneNumber"`
						CompanyName    string `json:"companyName"`
					} `json:"contact"`
					AccountNumber struct {
						Value string `json:"value"`
					} `json:"accountNumber"`
					Tins []struct {
						Number         string    `json:"number"`
						TinType        string    `json:"tinType"`
						Usage          string    `json:"usage"`
						EffectiveDate  time.Time `json:"effectiveDate"`
						ExpirationDate time.Time `json:"expirationDate"`
					} `json:"tins"`
				} `json:"producer"`
				BlanketPeriod struct {
					Begins string `json:"begins"`
					Ends   string `json:"ends"`
				} `json:"blanketPeriod"`
				CertifierJobTitle string `json:"certifierJobTitle"`
			} `json:"usmcaCertificationOfOriginDetail"`
			UsmcaCommercialInvoiceCertificationOfOriginDetail struct {
				CustomerImageUsages []struct {
					Id                string `json:"id"`
					Type              string `json:"type"`
					ProvidedImageType string `json:"providedImageType"`
				} `json:"customerImageUsages"`
				DocumentFormat struct {
					ProvideInstructions bool `json:"provideInstructions"`
					OptionsRequested    struct {
						Options []string `json:"options"`
					} `json:"optionsRequested"`
					StockType    string `json:"stockType"`
					Dispositions []struct {
						EMailDetail struct {
							EMailRecipients []struct {
								EmailAddress  string `json:"emailAddress"`
								RecipientType string `json:"recipientType"`
							} `json:"eMailRecipients"`
							Locale   string `json:"locale"`
							Grouping string `json:"grouping"`
						} `json:"eMailDetail"`
						DispositionType string `json:"dispositionType"`
					} `json:"dispositions"`
					Locale  string `json:"locale"`
					DocType string `json:"docType"`
				} `json:"documentFormat"`
				CertifierSpecification string `json:"certifierSpecification"`
				ImporterSpecification  string `json:"importerSpecification"`
				ProducerSpecification  string `json:"producerSpecification"`
				Producer               struct {
					Address struct {
						StreetLines         []string `json:"streetLines"`
						City                string   `json:"city"`
						StateOrProvinceCode string   `json:"stateOrProvinceCode"`
						PostalCode          string   `json:"postalCode"`
						CountryCode         string   `json:"countryCode"`
						Residential         bool     `json:"residential"`
					} `json:"address"`
					Contact struct {
						PersonName     string `json:"personName"`
						EmailAddress   string `json:"emailAddress"`
						PhoneExtension string `json:"phoneExtension"`
						PhoneNumber    string `json:"phoneNumber"`
						CompanyName    string `json:"companyName"`
					} `json:"contact"`
					AccountNumber struct {
						Value string `json:"value"`
					} `json:"accountNumber"`
					Tins []struct {
						Number         string    `json:"number"`
						TinType        string    `json:"tinType"`
						Usage          string    `json:"usage"`
						EffectiveDate  time.Time `json:"effectiveDate"`
						ExpirationDate time.Time `json:"expirationDate"`
					} `json:"tins"`
				} `json:"producer"`
				CertifierJobTitle string `json:"certifierJobTitle"`
			} `json:"usmcaCommercialInvoiceCertificationOfOriginDetail"`
			ShippingDocumentTypes []string `json:"shippingDocumentTypes"`
			CertificateOfOrigin   struct {
				CustomerImageUsages []struct {
					Id                string `json:"id"`
					Type              string `json:"type"`
					ProvidedImageType string `json:"providedImageType"`
				} `json:"customerImageUsages"`
				DocumentFormat struct {
					ProvideInstructions bool `json:"provideInstructions"`
					OptionsRequested    struct {
						Options []string `json:"options"`
					} `json:"optionsRequested"`
					StockType    string `json:"stockType"`
					Dispositions []struct {
						EMailDetail struct {
							EMailRecipients []struct {
								EmailAddress  string `json:"emailAddress"`
								RecipientType string `json:"recipientType"`
							} `json:"eMailRecipients"`
							Locale   string `json:"locale"`
							Grouping string `json:"grouping"`
						} `json:"eMailDetail"`
						DispositionType string `json:"dispositionType"`
					} `json:"dispositions"`
					Locale  string `json:"locale"`
					DocType string `json:"docType"`
				} `json:"documentFormat"`
			} `json:"certificateOfOrigin"`
			CommercialInvoiceDetail struct {
				CustomerImageUsages []struct {
					Id                string `json:"id"`
					Type              string `json:"type"`
					ProvidedImageType string `json:"providedImageType"`
				} `json:"customerImageUsages"`
				DocumentFormat struct {
					ProvideInstructions bool `json:"provideInstructions"`
					OptionsRequested    struct {
						Options []string `json:"options"`
					} `json:"optionsRequested"`
					StockType    string `json:"stockType"`
					Dispositions []struct {
						EMailDetail struct {
							EMailRecipients []struct {
								EmailAddress  string `json:"emailAddress"`
								RecipientType string `json:"recipientType"`
							} `json:"eMailRecipients"`
							Locale   string `json:"locale"`
							Grouping string `json:"grouping"`
						} `json:"eMailDetail"`
						DispositionType string `json:"dispositionType"`
					} `json:"dispositions"`
					Locale  string `json:"locale"`
					DocType string `json:"docType"`
				} `json:"documentFormat"`
			} `json:"commercialInvoiceDetail"`
		} `json:"shippingDocumentSpecification"`
		RateRequestType   []string `json:"rateRequestType"`
		PreferredCurrency string   `json:"preferredCurrency"`
		TotalPackageCount int      `json:"totalPackageCount"`
		MasterTrackingId  struct {
			FormId            string `json:"formId"`
			TrackingIdType    string `json:"trackingIdType"`
			UspsApplicationId string `json:"uspsApplicationId"`
			TrackingNumber    string `json:"trackingNumber"`
		} `json:"masterTrackingId"`
		RequestedPackageLineItems []struct {
			SequenceNumber     string `json:"sequenceNumber"`
			SubPackagingType   string `json:"subPackagingType"`
			CustomerReferences []struct {
				CustomerReferenceType string `json:"customerReferenceType"`
				Value                 string `json:"value"`
			} `json:"customerReferences"`
			DeclaredValue struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"declaredValue"`
			Weight struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"weight"`
			Dimensions struct {
				Length int    `json:"length"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
				Units  string `json:"units"`
			} `json:"dimensions"`
			GroupPackageCount           int    `json:"groupPackageCount"`
			ItemDescriptionForClearance string `json:"itemDescriptionForClearance"`
			ContentRecord               []struct {
				ItemNumber       string `json:"itemNumber"`
				ReceivedQuantity int    `json:"receivedQuantity"`
				Description      string `json:"description"`
				PartNumber       string `json:"partNumber"`
			} `json:"contentRecord"`
			ItemDescription              string `json:"itemDescription"`
			VariableHandlingChargeDetail struct {
				RateType      string  `json:"rateType"`
				PercentValue  float64 `json:"percentValue"`
				RateLevelType string  `json:"rateLevelType"`
				FixedValue    struct {
					Amount   float64 `json:"amount"`
					Currency string  `json:"currency"`
				} `json:"fixedValue"`
				RateElementBasis string `json:"rateElementBasis"`
			} `json:"variableHandlingChargeDetail"`
			PackageSpecialServices struct {
				SpecialServiceTypes []string `json:"specialServiceTypes"`
				SignatureOptionType string   `json:"signatureOptionType"`
				PriorityAlertDetail struct {
					EnhancementTypes []string `json:"enhancementTypes"`
					Content          []string `json:"content"`
				} `json:"priorityAlertDetail"`
				SignatureOptionDetail struct {
					SignatureReleaseNumber string `json:"signatureReleaseNumber"`
				} `json:"signatureOptionDetail"`
				AlcoholDetail struct {
					AlcoholRecipientType string `json:"alcoholRecipientType"`
					ShipperAgreementType string `json:"shipperAgreementType"`
				} `json:"alcoholDetail"`
				DangerousGoodsDetail struct {
					CargoAircraftOnly bool     `json:"cargoAircraftOnly"`
					Accessibility     string   `json:"accessibility"`
					Options           []string `json:"options"`
				} `json:"dangerousGoodsDetail"`
				PackageCODDetail struct {
					CodCollectionAmount struct {
						Amount   float64 `json:"amount"`
						Currency string  `json:"currency"`
					} `json:"codCollectionAmount"`
				} `json:"packageCODDetail"`
				PieceCountVerificationBoxCount int `json:"pieceCountVerificationBoxCount"`
				BatteryDetails                 []struct {
					BatteryPackingType    string `json:"batteryPackingType"`
					BatteryRegulatoryType string `json:"batteryRegulatoryType"`
					BatteryMaterialType   string `json:"batteryMaterialType"`
				} `json:"batteryDetails"`
				DryIceWeight struct {
					Units string `json:"units"`
					Value int    `json:"value"`
				} `json:"dryIceWeight"`
			} `json:"packageSpecialServices"`
			TrackingNumber string `json:"trackingNumber"`
		} `json:"requestedPackageLineItems"`
	} `json:"requestedShipment"`
	LabelResponseOptions string `json:"labelResponseOptions"`
	AccountNumber        struct {
		Value string `json:"value"`
	} `json:"accountNumber"`
	ShipAction           string `json:"shipAction"`
	ProcessingOptionType string `json:"processingOptionType"`
	OneLabelAtATime      bool   `json:"oneLabelAtATime"`
}
