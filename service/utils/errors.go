package utils

import "errors"

var ErrUnauthorized = errors.New("unauthorized operation")
var ErrUsernameNotValid = errors.New("username not valid")
var ErrMissingUsername = errors.New("missing username")