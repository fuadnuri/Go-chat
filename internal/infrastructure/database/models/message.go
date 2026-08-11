package models

type Message struct {
	ID       uint64 `gorm:"primaryKey"`
	Sender   uint64
	Reciever uint64
	Text     string
}
