package entities

import (
	"mime/multipart"
)

type Image struct {
	UserId      string
	ImageHeader *multipart.FileHeader
	ImageData   multipart.File
}
