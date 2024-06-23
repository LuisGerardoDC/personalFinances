package requestModel

import (
	"time"
)

type Budget struct {
	Assets    []Record  `json:"assets"`
	Expenses  []Record  `json:"expenses"`
	StartTime time.Time `json:"starttime"`
	EndTime   time.Time `json:"endtime"`
	Name      string    `json:"name"`
	UserID    int       `json:"userId"`
}
