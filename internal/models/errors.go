package models

import "errors"

var (
	ErrShortyNotFound = errors.New("shorty not found")
	ErrShortyExists   = errors.New("shorty already exist")
)
