package inputs

type SendTextMessage struct {
	ChatID  string `json:"chatId"`
	Message string `json:"message"`
}

func (i *SendTextMessage) Validate() error
