package entities

import "mime/multipart"

type UploadFile struct {
	UploaderUid string
	FileHeader  *multipart.FileHeader
}

const (
	StatusUploadSuccess = "upload sucessfully"
	StatusDeleteSuccess = "delete sucessfully"
)
