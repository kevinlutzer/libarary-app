package shared

import "fmt"

type ErrorType string

const (
	AlreadyExists    ErrorType = "AlreadyExists"
	NotFound         ErrorType = "NotFound"
	Internal         ErrorType = "Internal"
	InvalidArguments ErrorType = "InvalidArguments"
)

type AppError struct {
	Type ErrorType
	Msg  string
}

func (r *AppError) Error() string {
	return fmt.Sprintf("%s: %s", r.Type, r.Msg)
}

func NewError(errorType ErrorType, msg string) error {
	return &AppError{
		Type: errorType,
		Msg:  msg,
	}
}
