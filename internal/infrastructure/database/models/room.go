package models

import (
	"gorm.io/gorm"
)

type RoomModel struct {
	gorm.Model
	RoomName string
}
