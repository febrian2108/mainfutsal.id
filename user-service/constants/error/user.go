package error

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrPasswordIncorrect    = errors.New("password incorrect")
	ErrUserAlreadyExists    = errors.New("user already exists")
	ErrPasswordDoesNotMatch = errors.New("password does not match")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordIncorrect,
	ErrUserAlreadyExists,
	ErrPasswordDoesNotMatch,
}
