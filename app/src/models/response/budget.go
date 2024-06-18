package responseModel

import (
	models "github.com/LuisGerardoDC/personalFinances/app/src/models/mssql"
)

type Budget struct {
	Succes string         `json:"succes,omitempty"`
	Error  string         `json:"error,omitempty"`
	Code   int            `json:"code,omitempty"`
	Budget *models.Budget `json:"budget,omitempty"`
}
