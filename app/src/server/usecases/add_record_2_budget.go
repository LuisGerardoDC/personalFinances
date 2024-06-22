package usecases

import (
	"database/sql"

	requestModel "github.com/LuisGerardoDC/personalFinances/app/src/models/request"
	responseModel "github.com/LuisGerardoDC/personalFinances/app/src/models/response"
)

type AddRecord2Busget struct {
	DB *sql.DB
}

func (a *AddRecord2Busget) CreateNewRecord(r requestModel.Record) responseModel.Record {
	var rr responseModel.Record

	return rr
}
