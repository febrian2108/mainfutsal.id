package error

import "errors"

const (
	Success = "success"
	Error   = "error"
)

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrSQLError            = errors.New("databasse server failed to execute query")
	ErrTooManyRequests     = errors.New("too many requests")
	ErrUnautorized         = errors.New("unauthorized")
	ErrInvalidToken        = errors.New("invalid token")
	ErrForbidden           = errors.New("forbidden")
)

var GeneralErrors = []error{
	ErrInternalServerError,
	ErrSQLError,
	ErrTooManyRequests,
	ErrUnautorized,
	ErrInvalidToken,
	ErrForbidden,
}
