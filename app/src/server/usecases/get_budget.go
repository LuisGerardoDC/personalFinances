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
	)

	budget.ID = id
	err := budget.GetByID(g.DB)
	if err != nil {
		response.Code = 500
		response.Succes = false
		response.Error = err.Error()
		return response
	}

	err = budget.GetRecords(g.DB)
	if err != nil {
		response.Code = 500
		response.Succes = false
		response.Error = err.Error()
		return response
	}

	response.Budget = budget.ToResponseBudget()
	response.Code = 200
	response.Succes = true

	return response
}
