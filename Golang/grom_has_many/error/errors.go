package error

import "errors"

var (
	BadRequestError     = errors.New("bad request")
	InternalServerError = errors.New("internal server error")
)
