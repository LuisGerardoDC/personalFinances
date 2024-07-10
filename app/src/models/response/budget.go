package responseModel

import "time"

type Budget struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	TotalBudget     float32   `json:"totalBudget"`
	RemainingBudget float32   `json:"remainigBudget"`
	UsedBudget      float32   `json:"usedBudget"`
	StartTime       time.Time `json:"starttime"`
	EndTime         time.Time `json:"endtime"`
	Assets          []Record  `json:"assets,omitempty"`
	Expenses        []Record  `json:"expenses,omitempty"`
}
