package api

type ApiResponse struct {
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	StatusCode int         `json:"code"`
}
