package shared

type ApiResponseType string

const (
	Success ApiResponseType = "success"
	Error   ApiResponseType = "err"
)

type ApiResponse[D any] struct {
	Data D               `json:"result"`
	Type ApiResponseType `json:"type"`
}
