package usecases

import (
	"database/sql"

	mssqlmodel "github.com/LuisGerardoDC/personalFinances/app/src/models/mssql"
	responseModel "github.com/LuisGerardoDC/personalFinances/app/src/models/response"
)

type GetUserBudgets struct {
	DB *sql.DB
}

func (g *GetUserBudgets) GetBudgets(userid int64) responseModel.Response {
	var (
		response         responseModel.Response
		queriesBudgetIDs = "SELECT ID FROM budgets WHERE UserID = ?"
		budgets          = []responseModel.Budget{}
	)
	rows, err := g.DB.Query(queriesBudgetIDs, userid)
	if err != nil {
		return returnError(err)
	}

	for rows.Next() {
		budget := mssqlmodel.Budget{}
		err = rows.Scan(
			&budget.ID,
		)
		if err != nil {
			return returnError(err)
		}

		err = budget.GetByID(g.DB)
		if err != nil {
			return returnError(err)
		}
		budgets = append(budgets, *budget.ToResponseBudget())
	}
	response.Code = 200
	response.Succes = true
	response.Budgets = &budgets

	return response
}

func returnError(err error) responseModel.Response {
	return responseModel.Response{
		Succes: false,
		Error:  err.Error(),
		Code:   500,
	}
}
