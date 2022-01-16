package web

import "errors"

var (
	AccountDisabled    = errors.New("account disabled")
	AlreadyApproved    = errors.New("can't change status, already approved")
	AlreadySettled     = errors.New("can't change status, already settled")
	BadRequest         = errors.New("something is not right")
	BadCredential      = errors.New("bad credentials")
	CannotProcess      = errors.New("can't process")
	DuplicateData      = errors.New("duplicate data")
	Failed             = errors.New("something when wrong when trying to create")
	Forbidden          = errors.New("forbidden")
	InvalidToken       = errors.New("invalid callback token")
	InternalServerErr  = errors.New("something gone wrong, contact administrator")
	IDNotFound         = errors.New("id not found")
	NotEnoughPoints    = errors.New("not enough points")
	NotFound           = errors.New("data not found")
	TransactionExpired = errors.New("transaction expired")
)

type ErrorResp struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors,omitempty"`
}
