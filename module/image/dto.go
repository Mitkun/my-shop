package image

type UploadDTO struct {
	Name     string
	FileName string
	FileType string
	FileSize int
	FileData []byte
}
