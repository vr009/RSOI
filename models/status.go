package models

type StatusCode int

const (
	Okay StatusCode = iota
	NotFound
	InternalError
	BadRequest
	Created
	NoContent
)
