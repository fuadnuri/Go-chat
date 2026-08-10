package irepository

import (
	"github.com/fuadnuri/Go-chat/internal/domain/entities"
)

type IUserRepository interface {
	getUserByID(id uint64) entities.UserEntity
	getuserByUsername(username string) entities.UserEntity
	CreateUser(user entities.UserEntity) *entities.UserEntity
}
