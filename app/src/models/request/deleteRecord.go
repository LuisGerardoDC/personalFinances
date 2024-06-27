package requestModel

type DeleteRecord struct {
	BudgetID int `json:"budgetId"`
	RecordID int `json:"recordId"`
}
