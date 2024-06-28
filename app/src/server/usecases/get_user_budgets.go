package usecases

import (
	"database/sql"

	responseModel "github.com/LuisGerardoDC/personalFinances/app/src/models/response"
)

type GetUserBudgets struct {
	DB *sql.DB
}

func (g *GetUserBudgets) GetBudgets(userid int64) responseModel.Response {
	var (
		response responseModel.Response
	)

	return response
}
