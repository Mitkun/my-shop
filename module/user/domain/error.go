package userdomain

import "errors"

var (
	ErrEmailHasExisted = errors.New("email has been existed")
)
