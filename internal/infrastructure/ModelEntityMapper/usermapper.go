package modelentitymapper

import (
	"github.com/fuadnuri/Go-chat/internal/domain/entities"
	"github.com/fuadnuri/Go-chat/internal/infrastructure/database/models"
)

func UserModelToEntity(usermodel models.UserModel) *entities.UserEntity {
	return &entities.UserEntity{
		ID:        &usermodel.ID,
		Username:  usermodel.Username,
		Firstname: usermodel.Firstname,
		Lastname:  usermodel.Lastname,
		Email:     usermodel.Email,
		CreatedAt: usermodel.CreatedAt,
	}
}

func UserEntityToModel(UserEntity entities.UserEntity) *entities.UserEntity {
	return &entities.UserEntity{
		Username:  UserEntity.Username,
		Firstname: UserEntity.Firstname,
		Lastname:  UserEntity.Lastname,
		Email:     UserEntity.Email,
		CreatedAt: UserEntity.CreatedAt,
	}
}
