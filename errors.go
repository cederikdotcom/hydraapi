package hydraapi

import "net/http"

// ErrorResponse is the standard error shape for all hydra API endpoints.
type ErrorResponse struct {
	Error  string `json:"error"`
	Code   int    `json:"code"`
	Detail string `json:"detail,omitempty"`
}

// WriteError writes a standard JSON error response.
func WriteError(w http.ResponseWriter, status int, msg string) {
	WriteJSON(w, status, ErrorResponse{
		Error: msg,
		Code:  status,
	})
}

// WriteErrorDetail writes a standard JSON error response with additional detail.
func WriteErrorDetail(w http.ResponseWriter, status int, msg, detail string) {
	WriteJSON(w, status, ErrorResponse{
		Error:  msg,
		Code:   status,
		Detail: detail,
	})
}
