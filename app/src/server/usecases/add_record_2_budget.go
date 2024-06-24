package usecases

import (
	"database/sql"
	"fmt"

	mssqlmodel "github.com/LuisGerardoDC/personalFinances/app/src/models/mssql"
	requestModel "github.com/LuisGerardoDC/personalFinances/app/src/models/request"
	responseModel "github.com/LuisGerardoDC/personalFinances/app/src/models/response"
)

type AddRecord2Busget struct {
	DB *sql.DB
}

func (a *AddRecord2Busget) CreateNewRecord(r requestModel.Record) responseModel.Response {
	var (
		rr        responseModel.Response
		newRecord = mssqlmodel.Record{}
		budget    = mssqlmodel.Budget{ID: int64(r.BudgetID)}
	)

	newRecord.RequestToMssql(r)

	err := newRecord.SaveRecord(a.DB)
	if err != nil {
		rr.Code = 500
		rr.Error = fmt.Sprint(err.Error())
		return rr
	}

	err = budget.GetByID(a.DB)
	if err != nil {
		rr.Code = 500
		rr.Error = fmt.Sprint(err.Error())
		return rr
	}

	budgetRecords, err := mssqlmodel.Record{}.GetRecordsByBudgetID(budget.ID, a.DB)
	if err != nil {
		rr.Code = 500
		rr.Error = fmt.Sprint(err.Error())
		return rr
	}

	budget.Records = budgetRecords

	rr.Code = 200
	rr.Succes = true
	rr.Budget = budget.ToResponseBudget()

	return rr
}
