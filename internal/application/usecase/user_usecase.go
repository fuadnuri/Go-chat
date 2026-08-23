package usecase

import (
	"os/user"

	"github.com/fuadnuri/Go-chat/internal/application/dto"
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
func (uc *UserUsecase) UserLogin(user dto.AuthDTO) (*dto.AuthResponseDto,error){
	authResponse:=dto.AuthResponseDto{}
	if user,err:=uc.Repo.Login(user.Username,user.Password);err!=nil{
		return nil,err
	}
		
}