package storage

import "errors"

var ErrExists = errors.New("object already exists")
var ErrNotExists = errors.New("object does not exist")
