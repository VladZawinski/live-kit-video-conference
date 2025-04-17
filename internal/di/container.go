package di

import (
	"log"
	"os"

	"github.com/live-kit-video-conference/internal/room"
	"github.com/live-kit-video-conference/internal/token"
	lksdk "github.com/livekit/server-sdk-go/v2"
)

type Services struct {
	RoomService  room.RoomService
	TokenService token.TokenService
}

func InjectServices() *Services {
	host := "http://localhost:7880"
	apiKey := os.Getenv("LIVEKIT_API_KEY")
	apiSecret := os.Getenv("LIVEKIT_API_SECRET")
	log.Println(apiSecret)
	roomClient := lksdk.NewRoomServiceClient(host, apiKey, apiSecret)

	roomService := room.NewRoomService(roomClient)
	tokenService := token.NewTokenService(apiKey, apiSecret)

	return &Services{
		RoomService:  roomService,
		TokenService: tokenService,
	}
}
