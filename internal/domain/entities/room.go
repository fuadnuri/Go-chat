package entities

import "time"

type RoomEntity struct {
	RoomID    *uint64
	Members   []*UserEntity
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewRoom(room_name string) *RoomEntity {
	return &RoomEntity{
		RoomID:   nil,
		Members:  nil,
		RoomName: room_name,
	}
}

func (r RoomEntity) GetMembers() []*UserEntity {
	return r.Members
}
