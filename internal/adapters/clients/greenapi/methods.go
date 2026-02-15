package greenapi

type Methods string

const (
	SendMessageMethod      Methods = "SendMessage"
	GetSettingsMethod      Methods = "GetSettings"
	GetStateInstanceMethod Methods = "GetStateInstance"
	SendFileByUrlMethod    Methods = "SendFileByUrl"
)

func (m Methods) String() string {
	return string(m)
}
