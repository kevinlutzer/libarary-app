package shared

type ErrorType string

const (
	AlreadyExists    ErrorType = "AlreadyExists"
	NotFound         ErrorType = "NotFound"
	Internal         ErrorType = "Internal"
	InvalidArguments ErrorType = "InvalidArguments"
)

type AppError struct {
	Type ErrorType `json:"errorType"`
	Msg  string    `json:"msg"`
}

func (r *AppError) Error() string {
	return r.Msg
}

func NewError(errorType ErrorType, msg string) error {
	return &AppError{
		Type: errorType,
		Msg:  msg,
	}
}
