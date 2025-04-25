package api

import (
	"encoding/json"
	"net/http"

	"github.com/live-kit-video-conference/service"
)

type JoinAsGuestRequest struct {
	Username string `json:"username"`
}

type GetJoinTokenRequest struct {
	Username string `json:"username"`
	RoomId   int64  `json:"roomId"`
}

type AuthHandler struct {
	UserService service.UserService
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

}
