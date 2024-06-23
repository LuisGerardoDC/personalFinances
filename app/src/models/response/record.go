package responseModel

type Record struct {
	Succes   bool   `json:"succes,omitempty"`
	Error    string `json:"error,omitempty"`
	Code     int    `json:"code,omitempty"`
	RecordID int64  `json:"id,omitempty"`
}
