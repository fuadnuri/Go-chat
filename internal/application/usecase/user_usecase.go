package usecase

import (
	"github.com/fuadnuri/Go-chat/internal/domain/entities"
	"github.com/fuadnuri/Go-chat/internal/domain/irepository"
)

type UserUsecase struct {
	Repo irepository.IUserRepository
}

func NewUserUsecase(repo irepository.IUserRepository) *UserUsecase {
	return &UserUsecase{
		Repo: repo,
	}
}

func (uc *UserUsecase) GetUserByID(id uint) (*entities.UserEntity, error) {
	if user, err := uc.Repo.GetByID(id); err != nil {
		return nil, err

	} else {

		return user, nil
	}
}
func (uc *UserUsecase) UserLogin(username, password string) bool {
	_, err := uc.Repo.Login(username, password)
	return err != nil

}
