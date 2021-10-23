package errorx

import "errors"

var (
	ErrBadInput = errors.New("bad input")
	ErrNotFound = errors.New("not found")
	ErrInternal = errors.New("internal")
)
