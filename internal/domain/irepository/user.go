package irepository

import (
	"github.com/fuadnuri/Go-chat/internal/domain/entities"
)

type IUserRepository interface {
	GetByID(id uint) (*entities.UserEntity, error)
	GetByUsername(username string) (*entities.UserEntity, error)
	CreateUser(user entities.UserEntity) (*entities.UserEntity, error)
}
