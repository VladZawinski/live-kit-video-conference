package sdk

import (
	"os"

	lksdk "github.com/livekit/server-sdk-go/v2"
)

type SdkService struct {
	Room  RoomSdkService
	Token TokenSdkService
}

func InjectSdkServices() *SdkService {
	host := "http://localhost:7880"
	apiKey := os.Getenv("LIVEKIT_API_KEY")
	apiSecret := os.Getenv("LIVEKIT_API_SECRET")
	roomClient := lksdk.NewRoomServiceClient(host, apiKey, apiSecret)

	roomService := NewRoomSdkService(roomClient)
	tokenService := NewTokenSdkService(apiKey, apiSecret)

	return &SdkService{
		Room:  roomService,
		Token: tokenService,
	}
}
