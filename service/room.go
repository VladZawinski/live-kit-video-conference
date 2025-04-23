package service

import (
	"github.com/live-kit-video-conference/repository"
	"github.com/live-kit-video-conference/sdk"
)

type RoomService struct {
	RoomSdk sdk.RoomSdkService
	Room    repository.RoomRepository
}

func NewRoomService(roomSdk sdk.RoomSdkService, room repository.RoomRepository) RoomService {
	return RoomService{RoomSdk: roomSdk, Room: room}
}
