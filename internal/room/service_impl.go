package room

import (
	"context"

	"github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go/v2"
)

type roomServiceImpl struct {
	client *lksdk.RoomServiceClient
}

func NewRoomService(client *lksdk.RoomServiceClient) RoomService {
	return &roomServiceImpl{client: client}
}

func (s *roomServiceImpl) CreateRoom(ctx context.Context, name string) (*livekit.Room, error) {
	return s.client.CreateRoom(ctx, &livekit.CreateRoomRequest{
		Name: name,
	})
}

func (s *roomServiceImpl) ListRooms(ctx context.Context) ([]*livekit.Room, error) {
	resp, err := s.client.ListRooms(ctx, &livekit.ListRoomsRequest{})
	if err != nil {
		return nil, err
	}
	return resp.Rooms, nil
}

func (s *roomServiceImpl) DeleteRoom(ctx context.Context, name string) error {
	_, err := s.client.DeleteRoom(ctx, &livekit.DeleteRoomRequest{
		Room: name,
	})
	return err
}

func (s *roomServiceImpl) RoomExists(ctx context.Context, name string) (bool, error) {
	rooms, err := s.ListRooms(ctx)
	if err != nil {
		return false, err
	}
	for _, room := range rooms {
		if room.Name == name {
			return true, nil
		}
	}
	return false, nil
}
