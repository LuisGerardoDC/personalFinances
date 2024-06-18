package mssqlmodel

import "time"

type Record struct {
	Concept  string    `json:"concept"`
	Quantity float32   `json:"quantity"`
	Date     time.Time `json:"date"`
}
