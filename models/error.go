package models

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"data"`
}

func NewError(status int, message string) Error {
	return Error{status, message}
}
