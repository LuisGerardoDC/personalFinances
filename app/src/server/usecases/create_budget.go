package usecases

import (
	"database/sql"

	mssqlmodel "github.com/LuisGerardoDC/personalFinances/app/src/models/mssql"
	requestModel "github.com/LuisGerardoDC/personalFinances/app/src/models/request"
	responseModel "github.com/LuisGerardoDC/personalFinances/app/src/models/response"
)

type CreateNewBudget struct {
	DB *sql.DB
}

func (c *CreateNewBudget) CreateNewBudget(b requestModel.Budget) responseModel.Budget {
	var rb responseModel.Budget

	newBudget := mssqlmodel.Budget{}
	newBudget.NewBudget(b.Assets, b.StartTime, b.EndTime)

	rb.Budget = &newBudget
	rb.Code = 200
	rb.Succes = "True"

	return rb
}
