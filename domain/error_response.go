package domain

import "errors"

type ErrorResponse struct {
	Message string `json:"message"`
}

var (
	ErrAreaNotFound = errors.New("Area not found")
	ErrInternalServerError = errors.New("Internal server error")
	ErrBadRequest = errors.New("Bad request")
	ErrAreaFetch = errors.New("Error fetching area")
	ErrAreaSave = errors.New("Error saving area")
	ErrPlayerFetch = errors.New("Error fetching player")
	ErrMatchFetch = errors.New("Error fetching match")
	MongoDBInternalError = errors.New("MongoDB internal error")
)