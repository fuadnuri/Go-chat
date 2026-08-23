package repository

import (
	"fmt"

	"github.com/fuadnuri/Go-chat/internal/domain/entities"
	modelentitymapper "github.com/fuadnuri/Go-chat/internal/infrastructure/ModelEntityMapper"
	"github.com/fuadnuri/Go-chat/internal/infrastructure/database/models"
	"gorm.io/gorm"
)

type UserRepositororyImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositororyImpl {
	return &UserRepositororyImpl{
		DB: db,
	}
}

func (repo *UserRepositororyImpl) CreateUser(user models.UserModel) error {
	result := repo.DB.Create(&user)
	if result.Error != nil {
		return fmt.Errorf("error creating user: %s", result.Error.Error())
	}
	return nil

}

func (repo *UserRepositororyImpl) GetByID(id uint) (*models.UserModel, error) {
	var user models.UserModel
	result := repo.DB.Find(&user, id)
	if result.Error != nil {
		return nil, fmt.Errorf("user couldn't be found! \n	%s", result.Error.Error())
	}
	return &user, nil

}

func (repo *UserRepositororyImpl) GetByUsername(username string) (*models.UserModel, error) {
	var user models.UserModel
	result := repo.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf(
			"couldn't find user related to the email : %s",
			result.Error.Error(),
		)
	}
	return &user, nil
}

func (repo *UserRepositororyImpl) Delete(userID uint) error {
	result := repo.DB.Delete(map[string]uint{"ID": userID})
	if result.Error != nil {
		return fmt.Errorf("error: %s", result.Error.Error())
	}
	return nil
}

// user login
func (repo *UserRepositororyImpl) Login(username, password string) (*entities.UserEntity, error) {
	var userModel models.UserModel
	if err := repo.DB.Where("username=?", username).Find(&userModel); err != nil {
		return nil, fmt.Errorf("user not found: %s", err)
	}

	if userModel.PasswordHash != password {
		return nil, fmt.Errorf("incorrect password!")
	}
	user := modelentitymapper.UserModelToEntity(userModel)

	return user, nil
}
