package mappers

import (
	"time"

	"github.com/fuadnuri/Go-chat/internal/domain/entities"
	"github.com/fuadnuri/Go-chat/internal/infrastructure/database/models"
)

func UserEntityToModel(user entities.UserEntity) *models.UserModel {
	return &models.UserModel{
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func UserModelToEntity(userModel models.UserModel) *entities.UserEntity {
	return &entities.UserEntity{
		ID:        &userModel.ID,
		Username:  userModel.Username,
		Firstname: userModel.Firstname,
		Lastname:  userModel.Lastname,
		Email:     userModel.Email,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
	}
}
