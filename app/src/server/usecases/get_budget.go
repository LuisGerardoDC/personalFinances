package usecases

import (
	"database/sql"

	mssqlmodel "github.com/LuisGerardoDC/personalFinances/app/src/models/mssql"
	responseModel "github.com/LuisGerardoDC/personalFinances/app/src/models/response"
)

type GetBudget struct {
	DB *sql.DB
}

func (g *GetBudget) GetBudget(id int64) responseModel.Response {
	var (
		response responseModel.Response
		budget   mssqlmodel.Budget
		records  []mssqlmodel.Record
	)

	budget.ID = id
	err := budget.GetByID(g.DB)
	if err != nil {
		return g.returnError(err)
	}

	records, err = mssqlmodel.Record{}.GetRecordsByBudgetID(budget.ID, g.DB)
	if err != nil {
		return g.returnError(err)
	}

	budget.Records = records
	response.Budget = budget.ToResponseBudget()
	response.Code = 200
	response.Succes = true

	return response
}

func (g *GetBudget) returnError(err error) responseModel.Response {
	return responseModel.Response{
		Error:  err.Error(),
		Succes: false,
		Code:   500,
	}
}
