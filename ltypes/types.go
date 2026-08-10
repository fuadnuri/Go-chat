package ltypes

type DirectMessage struct {
	SenderUsername   string `json:"sender"`
	RecieverUsername string `json:"reciever"`
	Text             string `json:"text"`
}
