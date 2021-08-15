package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

//ErrorResponse error response object.
type ErrorResponse struct {
	Err        error  `json:"-"`
	StatusCode int    `json:"-"`
	StatusText string `json:"status_text"`
	Message    string `json:"message"`
}

// Error types
var (
	ErrMethodNotAllowed = &ErrorResponse{StatusCode: 405, Message: "Method not allowed"}
	ErrNotFound         = &ErrorResponse{StatusCode: 404, Message: "Resource not found"}
	ErrBadRequest       = &ErrorResponse{StatusCode: 400, Message: "Bad request"}
)

//Render is the function to render the ErrorResponse
func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

//ErrorRenderer is the function to render the Bad Request
func ErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: 400,
		StatusText: "Bad request",
		Message:    err.Error(),
	}
}

//ServerErrorRenderer is the function to render the Server Error
func ServerErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: 500,
		StatusText: "Internal server error",
		Message:    err.Error(),
	}
}

//InvalidParameterErrorRenderer is the function to render the Invalid Parameter Error
func InvalidParameterErrorRenderer(msg string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: 400,
		StatusText: "Bad request",
		Message:    msg,
	}
}
