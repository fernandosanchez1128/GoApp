package Services

import (
	"GoApp/Model"
	"GoApp/Repository"
)

type UserInterface interface {
	GetUser(document string) (user Model.User, err error)
}

type UserService struct {
	userRepository *Repository.UserRepository
}

func NewUserService(userRepository *Repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (userService *UserService) GetUser(document string) (Model.User, error) {

	user, err := userService.userRepository.GetUserInfo(document)
	if err != nil {
		return user, err
	}
	phones, err := userService.userRepository.GetUserPhones(user.Id)
	if err != nil {
		return user, err
	}

	user.SetPhones(phones)
	return user, nil
}
