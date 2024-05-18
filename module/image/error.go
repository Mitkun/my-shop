package image

import "errors"

var (
	ErrCannotUploadImage = errors.New("cannot upload image")
	ErrCannotFindImage   = errors.New("cannot find image")
)
