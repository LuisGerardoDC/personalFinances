package usecases

import (
	"database/sql"
	"fmt"

	mssqlmodel "github.com/LuisGerardoDC/personalFinances/app/src/models/mssql"
	requestModel "github.com/LuisGerardoDC/personalFinances/app/src/models/request"
	responseModel "github.com/LuisGerardoDC/personalFinances/app/src/models/response"
)

type CreateNewBudget struct {
	DB *sql.DB
}

func (c *CreateNewBudget) CreateNewBudget(b requestModel.Budget) responseModel.Response {
	var rb responseModel.Response

	newBudget := mssqlmodel.Budget{}

	newBudget.NewBudget(b)

	err := newBudget.CreateInDB(c.DB)

	if err != nil {
		rb.Code = 500
		rb.Succes = false
		rb.Error = fmt.Sprint(err)
		return rb
	}

	rb.Budget = newBudget.ToResponseBudget()
	rb.Code = 200
	rb.Succes = false

	return rb
}
