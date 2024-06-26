package usecases

import (
	"database/sql"

	requestModel "github.com/LuisGerardoDC/personalFinances/app/src/models/request"
	responseModel "github.com/LuisGerardoDC/personalFinances/app/src/models/response"
)

type RemoveRecord struct {
	DB *sql.DB
}

func (d *RemoveRecord) DeleteRecord(reqRecord requestModel.DeleteRecord) responseModel.Response {
	var (
		response responseModel.Response
	)

	return response
}
