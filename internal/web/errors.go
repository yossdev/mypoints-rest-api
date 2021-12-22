package web

import "errors"

var (
	BadRequest        = errors.New("something is not right")
	BadCredential     = errors.New("bad credential")
	Forbidden         = errors.New("forbidden")
	InternalServerErr = errors.New("something gone wrong, contact administrator")
	NotFound          = errors.New("data not found")
	IDNotFound        = errors.New("id not found")
	DuplicateData     = errors.New("duplicate data")
)
