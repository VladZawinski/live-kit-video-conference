package api

import (
	"net/http"

	"github.com/live-kit-video-conference/service"
)

type RoomHandler struct {
	RoomService service.RoomService
}

func NewRoomHandler(roomService service.RoomService) RoomHandler {
	return RoomHandler{RoomService: roomService}
}

func (h RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {

}

func (h RoomHandler) ListRooms(w http.ResponseWriter, r *http.Request) {

}
