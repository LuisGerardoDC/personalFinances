package mssqlmodel

import (
	"database/sql"
	"time"

	requestModel "github.com/LuisGerardoDC/personalFinances/app/src/models/request"
)

type Record struct {
	ID         int64
	Concept    string
	Quantity   float32
	Date       time.Time
	IsExpensse bool
	BudgetID   int64
}

func (r *Record) SaveRecord(db *sql.DB) error {

	insertQuery := "INSERT INTO budgets (Concept,Date,Quantity,IsExpense,BudgetID) VALUES (?,?,?,?,?)"

	result, err := db.Exec(insertQuery,
		r.Concept,
		r.Date,
		r.Quantity,
		r.IsExpensse,
		r.BudgetID,
	)

	if err != nil {
		return err
	}

	r.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (r *Record) RequestToMssql(rq requestModel.Record) {
	r.BudgetID = int64(rq.BudgetID)
	r.Concept = rq.Concept
	r.Quantity = rq.Quantity
	r.Date = rq.Date
	r.IsExpensse = rq.IsExpense
}
