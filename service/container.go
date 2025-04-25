package service

import (
	"github.com/live-kit-video-conference/repository"
	"github.com/live-kit-video-conference/sdk"
)

type AppService struct {
	Room RoomService
	User UserService
}

func InjectAppServices(sdk sdk.SdkService, repository repository.AppRepository) *AppService {
	roomService := NewRoomService(sdk.Room, repository.Room)
	userService := NewUserService(repository.User)

	return &AppService{
		Room: roomService,
		User: userService,
	}
}
