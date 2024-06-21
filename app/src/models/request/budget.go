package requestModel

import (
	"time"

	models "github.com/LuisGerardoDC/personalFinances/app/src/models/mssql"
)

type Budget struct {
	Assets    []models.Record `json:"assets"`
	StartTime time.Time       `json:"starttime"`
	EndTime   time.Time       `json:"endtime"`
	Name      string          `json:"name"`
}
