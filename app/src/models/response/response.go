package responseModel

type Response struct {
	Succes  bool      `json:"succes,omitempty"`
	Error   string    `json:"error,omitempty"`
	Code    int       `json:"code,omitempty"`
	Budget  *Budget   `json:"budget,omitempty"`
	Record  *Record   `json:"record,omitempty"`
	Budgets *[]Budget `json:"budgets,omitempty"`
}
