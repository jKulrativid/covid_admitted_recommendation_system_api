package entities

import "errors"

var (
	// default errors by status codes
	ErrorBadRequest          = errors.New("server cannot process your request due to bad request")
	ErrorUnAuthorized        = errors.New("your provided information might be wrong")
	ErrorNotFound            = errors.New("your request item is not found")
	ErrorConflict            = errors.New("your item already exist")
	ErrorUnprocessableEntity = errors.New("your request is correct but unable to process")
	ErrorInternalServer      = errors.New("internal server error")

	// validate error
	ErrorInvalidForm = errors.New("invalid form of data")

	// jwt token error
	ErrorInvalildToken = errors.New("token invalid")
	ErrorExpiredToken  = errors.New("token expired")
)
