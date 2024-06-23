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

func (a *AddRecord2Busget) CreateNewRecord(r requestModel.Record) responseModel.Record {
	var rr responseModel.Record

	newRecord := mssqlmodel.Record{}

	newRecord.RequestToMssql(r)

	err := newRecord.SaveRecord(a.DB)
	if err != nil {
		rr.Code = 500
		rr.Error = fmt.Sprint(err.Error())
		return rr
	}
	rr.Code = 200
	rr.Succes = true
	rr.RecordID = newRecord.ID
	return rr
}
