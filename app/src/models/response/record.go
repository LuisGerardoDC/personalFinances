package responseModel

import "time"

type Record struct {
	ID       int64     `json:"id"`
	Concept  string    `json:"concept"`
	Quantity float32   `json:"quantity"`
	Date     time.Time `json:"date"`
}
