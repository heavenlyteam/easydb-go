package easydb

import "errors"

var ErrEmptyToken = errors.New("an error occurred. provided token is empty")
var ErrEmptyDB = errors.New("an error occurred. provided database name is empty")
