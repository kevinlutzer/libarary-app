package shared

type ApiRequest interface {
	Validate() error
}
