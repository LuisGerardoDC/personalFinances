package models

import "time"

type Record struct {
	Concept  string
	Quantity float32
	Date     time.Time
}
