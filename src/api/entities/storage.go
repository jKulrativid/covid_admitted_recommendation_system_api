package entities

import "mime/multipart"

type UploadFile struct {
	UploaderUid string
	FileHeader  *multipart.FileHeader
}

type DeleteFileList struct {
	List []interface{} `json:"list"`
}

const (
	StatusUploadSuccess = "upload sucessfully"
	StatusDeleteSuccess = "delete sucessfully"
)
