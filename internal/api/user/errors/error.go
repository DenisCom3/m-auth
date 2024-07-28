package errors

import "github.com/pkg/errors"

var (
	ErrPassNotEq = errors.New("password not equals")
)
