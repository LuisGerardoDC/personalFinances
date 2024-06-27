package usecases

import (
	"database/sql"
	"fmt"

	mssqlmodel "github.com/LuisGerardoDC/personalFinances/app/src/models/mssql"
	requestModel "github.com/LuisGerardoDC/personalFinances/app/src/models/request"
	responseModel "github.com/LuisGerardoDC/personalFinances/app/src/models/response"
)

type RemoveRecord struct {
	DB *sql.DB
}

func (d *RemoveRecord) DeleteRecord(req requestModel.DeleteRecord) responseModel.Response {
	var (
		response responseModel.Response
		budget   = mssqlmodel.Budget{ID: int64(req.BudgetID)}
	)

	err := mssqlmodel.Record{}.RemoveRecord(int64(req.BudgetID), int64(req.RecordID), d.DB)
	if err != nil {
		response.Code = 500
		response.Error = fmt.Sprint(err.Error())
		return response
	}
	err = budget.GetByID(d.DB)
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
