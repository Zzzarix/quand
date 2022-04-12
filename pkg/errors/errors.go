package errors

import "errors"

var (
	IncorrectEmail             = errors.New("email is incorrect")
	UserWithEmailAlreadyExists = errors.New("user with current email address already exists")
	UserWithNameAlreadyExists  = errors.New("user with current user name already exists")
)
