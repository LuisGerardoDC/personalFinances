package mssqlmodel

import (
	"database/sql"
	"time"

	requestModel "github.com/LuisGerardoDC/personalFinances/app/src/models/request"
)

type Record struct {
	ID        int64
	Concept   string
	Quantity  float32
	Date      time.Time
	IsExpense bool
	BudgetID  int64
}

func (r *Record) SaveRecord(db *sql.DB) error {

	var (
		insertQuery       = "INSERT INTO records (Concept,Date,Quantity,IsExpense,BudgetID) VALUES (?,?,?,?,?);"
		getBudgetRecords  = "SELECT Quantity, IsExpense FROM records WHERE BudgetID = ?; "
		updateBudgetQuery = "UPDATE budgets SET TotalBudget = ?, UsedBudget = ?, RemainingBudget = ? WHERE ID = ?;"
		budgetRecords     = []Record{}
		budget            = Budget{}
	)

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Saves NewRecord
	result, err := tx.Exec(insertQuery,
		r.Concept,
		r.Date,
		r.Quantity,
		r.IsExpense,
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
	// Gets all Budget Records
	stmt, err := tx.Prepare(getBudgetRecords)
	if err != nil {
		tx.Rollback()
		return err
	}

	rows, err := stmt.Query(r.BudgetID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for rows.Next() {
		var gotR Record
		err = rows.Scan(&gotR.Quantity, &gotR.IsExpense)
		if err != nil {
			tx.Rollback()
			return err
		}
		budgetRecords = append(budgetRecords, gotR)
	}
	budget.ID = r.BudgetID
	budget.Records = budgetRecords

	// Calc new quantites budget
	budget.CalcBudgets()

	// Update Budget in db
	_, err = tx.Exec(updateBudgetQuery,
		budget.TotalBudget,
		budget.UsedBudget,
		budget.RemainingBudget,
		budget.ID,
	)
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
	r.IsExpense = rq.IsExpense
}

func (r Record) RemoveRecord(budgetID, RecordID int64) error {
	// Saves NewRecord
	// Gets all Budget Records
	// Calc new quantites budget
	// Update Budget in db
	return nil
}

func (r Record) GetRecordsByBudgetID(budgetID int64, db *sql.DB) ([]Record, error) {
	var (
		query      = "SELECT ID, Concept, Date, Quantity, IsExpense FROM records WHERE BudgetID = ?;"
		records    = []Record{}
		readRecord Record
		date       []byte
	)
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(budgetID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&readRecord.ID,
			&readRecord.Concept,
			&date,
			&readRecord.Quantity,
			&readRecord.IsExpense,
		)
		if err != nil {
			return nil, err
		}
		readRecord.Date, err = time.Parse("2006-01-02 15:04:05", string(date))
		if err != nil {
			return nil, err
		}
		records = append(records, readRecord)
	}

	return records, nil
}
