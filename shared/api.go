package shared

type ApiRequest interface {
	Validate() error
}

type ApiResponseType string

const (
	Success ApiResponseType = "success"
	Error   ApiResponseType = "err"
)

type ApiResponse[D any] struct {
	Data D               `json:"result"`
	Type ApiResponseType `json:"type"`
}
