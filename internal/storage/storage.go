package storage

import "errors"

var (
	ErrURLNotfound = errors.New("url not found")
	ErrURLExists   = errors.New("url exists")
)
