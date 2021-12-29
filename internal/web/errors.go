package web

import "errors"

var (
	AccountDisabled   = errors.New("account disabled")
	BadRequest        = errors.New("something is not right")
	BadCredential     = errors.New("bad credentials")
	DuplicateData     = errors.New("duplicate data")
	Failed            = errors.New("something when wrong when trying to create")
	Forbidden         = errors.New("forbidden")
	InternalServerErr = errors.New("something gone wrong, contact administrator")
	IDNotFound        = errors.New("id not found")
	NotFound          = errors.New("data not found")
)
