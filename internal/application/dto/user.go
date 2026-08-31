package dto

import "time"

type UserRequest struct {
	Firstname string
	Lastname  string
	Username  string
	DoB       string
	Email     string
}

type UserResponse struct {
	Firstname string
	Lastname  string
	Username  string
	DoB       string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
