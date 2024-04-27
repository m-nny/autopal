package brain

import "errors"

var (
	ErrAlreadyExists = errors.New("Alredy exists")
	ErrNotFound      = errors.New("Not found")
)
