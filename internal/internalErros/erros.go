package internalerros

import "errors"

var ErrInternal error = errors.New("Internal Server Error")
