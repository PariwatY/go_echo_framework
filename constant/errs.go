package constant

import "errors"

var (
	ErrDuplicate = errors.New("duplicate data")
	ErrNotFound  = errors.New("not found")
)
