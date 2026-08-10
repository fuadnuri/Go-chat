package entities

type Message struct {
	ID        uint64
	Sender    uint64
	Reciever  uint64
	Content   string
	CreatedAt string
}

func NewMessage(id, sender, Reciever uint64, content string) *Message {
	return &Message{ID: id, Sender: sender, Reciever: Reciever, Content: content}
}

func (msg Message) GetSender() *uint64 {
	return &msg.Sender
}
func (msg Message) GetReciever() *uint64 {
	return &msg.Reciever
}

func (msg Message) GetContent() *string {
	return &msg.Content
}
func (msg Message) GetDateCreated() string {
	return msg.CreatedAt
}

func (msg *Message) SetContent(message string) {
	msg.Content = message
}
