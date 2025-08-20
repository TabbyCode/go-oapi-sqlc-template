package repository

import "errors"

// ErrRecordNotFound is returned when a requested record is not found in the database.
var ErrRecordNotFound = errors.New("record not found")
