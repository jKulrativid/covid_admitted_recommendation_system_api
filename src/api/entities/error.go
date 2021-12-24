package entities

import "errors"

var (
	ErrorNotFound       = errors.New("your request item is not found")
	ErrorInternalServer = errors.New("internal server error")
	ErrorConflice       = errors.New("your item already exist")
)
