package domain

type InstanceSettings struct {
	Wid             string `json:"wid"`
	CountryInstance string `json:"countryInstance"`
	TypeAccount     string `json:"typeAccount"`

	WebhookURL      string `json:"webhookUrl"`
	WebhookURLToken string `json:"webhookUrlToken"`

	DelaySendMessagesMilliseconds int `json:"delaySendMessagesMilliseconds"`

	MarkIncomingMessagesReaded        YesNo `json:"markIncomingMessagesReaded"`
	MarkIncomingMessagesReadedOnReply YesNo `json:"markIncomingMessagesReadedOnReply"`

	SharedSession string `json:"sharedSession"`

	OutgoingWebhook           YesNo `json:"outgoingWebhook"`
	OutgoingMessageWebhook    YesNo `json:"outgoingMessageWebhook"`
	OutgoingAPIMessageWebhook YesNo `json:"outgoingAPIMessageWebhook"`
	IncomingWebhook           YesNo `json:"incomingWebhook"`
	DeviceWebhook             YesNo `json:"deviceWebhook"`

	StatusInstanceWebhook string `json:"statusInstanceWebhook"`
	StateWebhook          YesNo  `json:"stateWebhook"`

	EnableMessagesHistory string `json:"enableMessagesHistory"`
	KeepOnlineStatus      YesNo  `json:"keepOnlineStatus"`

	PollMessageWebhook    YesNo `json:"pollMessageWebhook"`
	IncomingBlockWebhook  YesNo `json:"incomingBlockWebhook"`
	IncomingCallWebhook   YesNo `json:"incomingCallWebhook"`
	EditedMessageWebhook  YesNo `json:"editedMessageWebhook"`
	DeletedMessageWebhook YesNo `json:"deletedMessageWebhook"`
}
