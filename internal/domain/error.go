package domain

import "errors"

var ErrEmailExists = errors.New("Email exist")
var ErrInvalidCredentials = errors.New("ErrInvalidCredentials")
