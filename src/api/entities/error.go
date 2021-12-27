package entities

import "errors"

var (
	ErrorBadRequest     = errors.New("server cannot process your request")
	ErrorNotAuthorized  = errors.New("your provided information might be wrong")
	ErrorNotFound       = errors.New("your request item is not found")
	ErrorConflict       = errors.New("your item already exist")
	ErrorInternalServer = errors.New("internal server error")
)
