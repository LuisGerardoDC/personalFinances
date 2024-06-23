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

	/*
		1.- Guardar new Record
		2.- Obtener de DB todos los records del Budget
		3.- Calcular las cantidades del budget (totalBudget,UsedBudget,RemainingBudget)
		4.- Actualizar Budget en db

	*/
	// todo update Budget
	insertQuery := "INSERT INTO budgets (Concept,Date,Quantity,IsExpense,BudgetID) VALUES (?,?,?,?,?);"
	getBudgetRecords := "SELECT Quantity, IsExpense FROM records WHERE BudgetID = @p1; "

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	result, err := tx.Exec(insertQuery,
		r.Concept,
		r.Date,
		r.Quantity,
		r.IsExpensse,
		r.BudgetID,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	r.ID, err = result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *Record) RequestToMssql(rq requestModel.Record) {
	r.BudgetID = int64(rq.BudgetID)
	r.Concept = rq.Concept
	r.Quantity = rq.Quantity
	r.Date = rq.Date
	r.IsExpensse = rq.IsExpense
}
