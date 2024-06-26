package requestModel

type DeleteRecord struct {
	BudgetID int    `json:"budgetId"`
	RecordID string `json:"recordId"`
}
