package entities

import (
	"mime/multipart"
)

type Image struct {
	userId      string
	imageHeader *multipart.FileHeader
	imageData   multipart.File
}
