package mssqlmodel

import (
	"database/sql"
	"math"
	"time"

	requestModel "github.com/LuisGerardoDC/personalFinances/app/src/models/request"
	responseModel "github.com/LuisGerardoDC/personalFinances/app/src/models/response"
)

type Budget struct {
	ID              int64
	UserID          int
	Name            string
	TotalBudget     float32
	StartTime       time.Time
	EndTime         time.Time
	Records         []Record
	UsedBudget      float32
	RemainingBudget float32
}

func (b *Budget) NewBudget(rb requestModel.Budget) {
	b.RecordToMssql(rb.Assets, rb.Expenses)
	b.UserID = rb.UserID
	b.StartTime = rb.EndTime
	b.EndTime = rb.EndTime
	b.Name = rb.Name
	b.CalcBudgets()
}

func (b *Budget) CalcBudgets() {
	b.TotalBudget = 0
	b.UsedBudget = 0

	for _, record := range b.Records {
		if record.IsExpense {
			b.UsedBudget += record.Quantity
		} else {
			b.TotalBudget += record.Quantity
		}
	}

	b.RemainingBudget = b.TotalBudget - b.UsedBudget
	b.RemainingBudget = float32(math.Round(float64(b.RemainingBudget)*100) / 100)
}

func (b *Budget) RecordToMssql(assets, expences []requestModel.Record) {
	for _, asset := range assets {
		b.Records = append(b.Records, Record{
			Concept:   asset.Concept,
			Quantity:  asset.Quantity,
			Date:      asset.Date,
			IsExpense: false,
		})
	}
	for _, expence := range expences {
		b.Records = append(b.Records, Record{
			Concept:   expence.Concept,
			Quantity:  expence.Quantity,
			Date:      expence.Date,
			IsExpense: true,
		})
	}
}

func (b *Budget) CreateInDB(db *sql.DB) error {

	queryCreateBudget := "INSERT INTO budgets (UserID,Name,TotalBudget,StartTime,EndTime,UsedBudget,RemainingBudget ) VALUES (?,?,?,?,?,?,?);"
	queryCreateRecord := "INSERT INTO records (Concept,Date,Quantity,IsExpense,BudgetID ) VALUES (?,?,?,?,?);"

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	result, err := tx.Exec(queryCreateBudget,
		b.UserID,
		b.Name,
		b.TotalBudget,
		b.StartTime,
		b.EndTime,
		b.UsedBudget,
		b.RemainingBudget,
	)

	if err != nil {
		tx.Rollback()
		return err
	}
	b.ID, err = result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, record := range b.Records {
		_, err = tx.Exec(queryCreateRecord, record.Concept, record.Date, record.Quantity, record.IsExpense, b.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (b *Budget) GetByID(db *sql.DB) error {
	query := "SELECT * FROM budgets WHERE ID = @p1;"

	err := db.QueryRow(query, b.ID).Scan(
		b.UserID,
		b.Name,
		b.TotalBudget,
		b.StartTime,
		b.EndTime,
		b.UsedBudget,
		b.RemainingBudget,
	)

	return err
}

func (b *Budget) ToResponseBudget() *responseModel.Budget {
	var (
		rb         responseModel.Budget
		respRecord responseModel.Record
	)
	rb.ID = b.ID
	rb.EndTime = b.EndTime
	rb.StartTime = b.StartTime
	rb.Name = b.Name
	rb.RemainingBudget = b.RemainingBudget
	rb.UsedBudget = b.UsedBudget
	rb.TotalBudget = b.TotalBudget

	for _, record := range b.Records {
		respRecord = responseModel.Record{
			ID:       record.ID,
			Concept:  record.Concept,
			Quantity: record.Quantity,
			Date:     record.Date,
		}

		if record.IsExpense {
			rb.Expenses = append(rb.Expenses, respRecord)
		} else {
			rb.Assets = append(rb.Assets, respRecord)
		}
	}

	return &rb
}
