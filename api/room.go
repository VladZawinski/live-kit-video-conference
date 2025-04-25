package api

import (
	"encoding/json"
	"net/http"

	"github.com/live-kit-video-conference/service"
)

type CreateRoomDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerID     int    `json:"owner_id"`
}

type RoomHandler struct {
	RoomService service.RoomService
}

func NewRoomHandler(roomService service.RoomService) RoomHandler {
	return RoomHandler{RoomService: roomService}
}

func (h RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var roomDto CreateRoomDto
	if err := json.NewDecoder(r.Body).Decode(&roomDto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	room, err := h.RoomService.CreateRoom(&service.CreateRoomModel{
		Name:        roomDto.Name,
		Description: roomDto.Description,
		OwnerID:     roomDto.OwnerID,
	})
	if err != nil {
		http.Error(w, "Failed to create room", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ApiResponse{
		Data: room,
	})
}

func (h RoomHandler) ListRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := h.RoomService.ListRoom()
	if err != nil {
		http.Error(w, "Failed to list rooms", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ApiResponse{
		Data: rooms,
	})
}
