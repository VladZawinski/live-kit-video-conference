package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/live-kit-video-conference/service"
)

type JoinAsGuestRequest struct {
	Username string `json:"username"`
}

type GetJoinTokenRequest struct {
	Username string `json:"username"`
	RoomID   int    `json:"room_id"`
}

type AuthHandler struct {
	UserService service.UserService
	RoomService service.RoomService
}

func NewAuthHandler(userService service.UserService) AuthHandler {
	return AuthHandler{
		UserService: userService,
	}
}

func (h AuthHandler) JoinAsGuest(w http.ResponseWriter, r *http.Request) {
	var req JoinAsGuestRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	userExists, _ := h.UserService.UserExists(req.Username)
	if userExists {
		json.NewEncoder(w).Encode(ApiResponse{
			StatusCode: 200,
			Message:    "Join Successfully",
			Data:       nil,
		})
		return
	}
	err := h.UserService.CreateUser(req.Username, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(ApiResponse{
		StatusCode: 200,
		Message:    "Join Successfully",
		Data:       nil,
	})
}

func (h AuthHandler) GetJoinToken(w http.ResponseWriter, r *http.Request) {
	var req GetJoinTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	room, err := h.RoomService.GetRoomByID(req.RoomID)
	if err != nil {
		log.Println(err)
	}
	log.Printf("room: %v", room)
	// user, _ := h.UserService.GetUserByUsername(req.Username)
	// log.Println(user)
	// token, err := h.UserService.GetJoinToken(req.Username, room.Name, room.OwnerID == user.ID)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusUnauthorized)
	// 	return
	// }
	json.NewEncoder(w).Encode(ApiResponse{
		StatusCode: 200,
		Message:    "Get Join Token Successfully",
		Data:       "token",
	})
}
