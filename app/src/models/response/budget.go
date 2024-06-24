package responseModel

import "time"

type Budget struct {
	TotalBudget     float32   `json:"totalBudget"`
	UsedBudget      float32   `json:"usedBudget"`
	RemainingBudget float32   `json:"remainigBudget"`
	Assets          []Record  `json:"assets"`
	Expenses        []Record  `json:"expenses"`
	StartTime       time.Time `json:"starttime"`
	EndTime         time.Time `json:"endtime"`
	Name            string    `json:"name"`
	ID              int64     `json:"id"`
}
