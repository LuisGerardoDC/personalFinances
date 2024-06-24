package requestModel

import (
	"time"
)

type Record struct {
	BudgetID  int       `json:"budgetId"`
	Concept   string    `json:"concept"`
	Quantity  float32   `json:"quantity"`
	Date      time.Time `json:"date"`
	IsExpense bool      `json:"isExpense"`
}
