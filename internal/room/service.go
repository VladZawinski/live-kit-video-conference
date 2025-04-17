package room

import (
	"context"

	"github.com/livekit/protocol/livekit"
)

type RoomService interface {
	CreateRoom(ctx context.Context, name string) (*livekit.Room, error)
	ListRooms(ctx context.Context) ([]*livekit.Room, error)
	DeleteRoom(ctx context.Context, name string) error
	RoomExists(ctx context.Context, name string) (bool, error)
}
