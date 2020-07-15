package response

import "errors"

var (
	ErrBadParams   = errors.New("bad params")
	ErrEmpty       = errors.New("")
	ErrServerError = errors.New("server error")
)
