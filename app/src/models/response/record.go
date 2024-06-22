package responseModel

type Record struct {
	Succes string `json:"succes,omitempty"`
	Error  string `json:"error,omitempty"`
	Code   int    `json:"code,omitempty"`
}
