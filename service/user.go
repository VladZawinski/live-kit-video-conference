package service

import (
	"github.com/live-kit-video-conference/model"
	"github.com/live-kit-video-conference/repository"
	"github.com/live-kit-video-conference/sdk"
)

type UserService interface {
	CreateUser(username string, password string) error
	UserExists(username string) (bool, error)
	GetUserByUsername(username string) (*model.User, error)
	GetJoinToken(username string, roomName string, isPublisher bool) (string, error)
}

type userService struct {
	User     repository.UserRepository
	TokenSdk sdk.TokenSdkService
}

func (s userService) GetUserByUsername(username string) (*model.User, error) {
	user, err := s.User.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return user, nil
}

func (s userService) UserExists(username string) (bool, error) {
	user, err := s.User.GetByUsername(username)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, nil
	}
	return true, nil
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return userService{User: userRepository}
}

func (s userService) CreateUser(username string, password string) error {
	user := model.User{
		Username: username,
	}
	return s.User.Create(&user)
}

func (s userService) GetJoinToken(username string, roomName string, isPublisher bool) (string, error) {
	return s.TokenSdk.GenerateJoinToken(username, roomName, isPublisher)
}
