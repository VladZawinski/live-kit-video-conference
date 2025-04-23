package service

import (
	"github.com/live-kit-video-conference/model"
	"github.com/live-kit-video-conference/repository"
	"github.com/live-kit-video-conference/sdk"
)

type UserService struct {
	User     repository.UserRepository
	TokenSdk sdk.TokenSdkService
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return UserService{User: userRepository}
}

func (s UserService) CreateUser(username string, password string) error {
	user := model.User{
		Username: username,
	}
	return s.User.Create(&user)
}

func (s UserService) GetJoinToken(username string, roomName string, isPublisher bool) (string, error) {
	return s.TokenSdk.GenerateJoinToken(username, roomName, isPublisher)
}
