package api

import (
	"net/http"

	"github.com/live-kit-video-conference/service"
)

func BuildHandlers(service service.AppService) {
	// Initialize the RoomHandler
	roomHandler := NewRoomHandler(service.Room)
	// Register the routes
	http.HandleFunc("/api/rooms", roomHandler.CreateRoom)
	http.HandleFunc("/api/rooms/list", roomHandler.ListRooms)
	authHandler := NewAuthHandler(service.User)
	http.HandleFunc("/api/auth/join", authHandler.JoinAsGuest)
	http.HandleFunc("/api/auth/token", authHandler.GetJoinToken)
}
