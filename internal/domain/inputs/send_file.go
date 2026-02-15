package inputs

type SendFile struct {
	ChatID   string `json:"chatId"`
	FileURL  string `json:"urlFile"`
	FileName string `json:"fileName"`
}

func (i *SendFile) Validate() error
