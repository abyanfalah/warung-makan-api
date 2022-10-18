package utils

import "mime/multipart"

func IsImage(file *multipart.FileHeader) bool {
	contentType := file.Header.Values("Content-Type")
	theType := contentType[0]
	isImage := theType == "image/jpeg" || theType == "image/jpg" || theType == "image/png"
	return isImage
}
