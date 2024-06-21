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

func (c *CreateNewBudget) CreateNewBudget(b requestModel.Budget) responseModel.Budget {
	var rb responseModel.Budget

	newBudget := mssqlmodel.Budget{}

	// todo hacer conversion de reqAsset a mssqlAsset
	newBudget.NewBudget(b.Assets, b.StartTime, b.EndTime)

	err := newBudget.CreateInDB(c.DB)

	if err != nil {
		rb.Code = 500
		rb.Succes = "False"
		rb.Error = fmt.Sprint(err)
		return rb
	}

	rb.Budget = &newBudget
	rb.Code = 200
	rb.Succes = "True"

	return rb
}
