package entities

import "errors"

var (
	ErrorBadRequest          = errors.New("server cannot process your request")
	ErrorUnAuthorized        = errors.New("your provided information might be wrong")
	ErrorNotFound            = errors.New("your request item is not found")
	ErrorConflict            = errors.New("your item already exist")
	ErrorUnprocessableEntity = errors.New("your request is correct but unable to process")
	ErrorInternalServer      = errors.New("internal server error")
)
