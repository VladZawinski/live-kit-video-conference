package service

import (
	"context"

	"github.com/live-kit-video-conference/model"
	"github.com/live-kit-video-conference/repository"
	"github.com/live-kit-video-conference/sdk"
)

type RoomService interface {
	CreateRoom(room *CreateRoomModel) (*model.Room, error)
	ListRoom() ([]*model.Room, error)
	GetRoomByID(id int) (*model.Room, error)
}

type roomService struct {
	RoomSdk sdk.RoomSdkService
	Room    repository.RoomRepository
}

func (r roomService) GetRoomByID(id int) (*model.Room, error) {
	room, err := r.Room.GetByID(id)
	if err != nil {
		return nil, err
	}
	if room == nil {
		return nil, nil
	}
	return room, nil
}

func (r roomService) CreateRoom(room *CreateRoomModel) (*model.Room, error) {
	// Create Sdk room
	sdkRoom, err := r.RoomSdk.CreateRoom(context.Background(), room.Name)
	if err != nil {
		return nil, err
	}
	roomModel := model.Room{
		Name:        room.Name,
		OwnerID:     room.OwnerID,
		Description: room.Description,
		SID:         sdkRoom.Sid,
	}
	// Create Db room
	_, err = r.Room.Create(&roomModel)
	if err != nil {
		return nil, err
	}
	return &roomModel, nil
}

func (r roomService) ListRoom() ([]*model.Room, error) {
	rooms, err := r.Room.List()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func NewRoomService(roomSdk sdk.RoomSdkService, room repository.RoomRepository) RoomService {
	return roomService{RoomSdk: roomSdk, Room: room}
}

type CreateRoomModel struct {
	Name        string
	Description string
	OwnerID     int
}
