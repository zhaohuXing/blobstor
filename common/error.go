package common

import (
	"errors"
)

var ErrInternalError = errors.New("internal_error")
var ErrNotMatchCodeError = errors.New("not_match_code")
var ErrUserExistError = errors.New("user_exist")
var ErrUserNotExistError = errors.New("user_not_exist")
var ErrInvalidArgumentError = errors.New("invalid_argument")
