package model

type SendNotification struct {
	SenderContactName               string `json:"senderContactName"`
	SenderEMailAddress              string `json:"senderEMailAddress"`
	TrackingEventNotificationDetail struct {
		TrackingNotifications []struct {
			NotificationDetail struct {
				Localization struct {
					LanguageCode string `json:"languageCode"`
					LocaleCode   string `json:"localeCode"`
				} `json:"localization"`
				EmailDetail struct {
					EmailAddress string `json:"emailAddress"`
					Name         string `json:"name"`
				} `json:"emailDetail"`
				NotificationType string `json:"notificationType"`
			} `json:"notificationDetail"`
			Role                       string   `json:"role"`
			NotificationEventTypes     []string `json:"notificationEventTypes"`
			CurrentResultRequestedFlag bool     `json:"currentResultRequestedFlag"`
		} `json:"trackingNotifications"`
		PersonalMessage string      `json:"personalMessage"`
		SupportHTML     interface{} `json:"supportHTML"`
	} `json:"trackingEventNotificationDetail"`
	TrackingNumberInfo struct {
		TrackingNumber         string `json:"trackingNumber"`
		CarrierCode            string `json:"carrierCode"`
		TrackingNumberUniqueId string `json:"trackingNumberUniqueId"`
	} `json:"trackingNumberInfo"`
	ShipDateBegin string `json:"shipDateBegin"`
	ShipDateEnd   string `json:"shipDateEnd"`
}
