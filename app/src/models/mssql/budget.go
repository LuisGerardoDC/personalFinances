package mssqlmodel

import (
	"database/sql"
	"math"
	"time"

	requestModel "github.com/LuisGerardoDC/personalFinances/app/src/models/request"
)

type Budget struct {
	ID              int
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
		if record.IsExpensse {
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
			Concept:    asset.Concept,
			Quantity:   asset.Quantity,
			Date:       asset.Date,
			IsExpensse: false,
		})
	}
	for _, expence := range expences {
		b.Records = append(b.Records, Record{
			Concept:    expence.Concept,
			Quantity:   expence.Quantity,
			Date:       expence.Date,
			IsExpensse: true,
		})
	}
}

func (b *Budget) CreateInDB(db *sql.DB) error {

	queryCreateBudget := "INSERT INTO budgets (UserID,Name,TotalBudget,StartTime,EndTime,UsedBudget,RemainingBudget ) VALUES (?,?,?,?,?,?,?);"
	queryCreateRecord := "INSERT INTO records (Concept,Date,Quantity,IsEpensse,BudgetID ) VALUES (?,?,?,?,?);"

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
	insertedID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	b.ID = int(insertedID)

	for _, record := range b.Records {
		_, err = tx.Exec(queryCreateRecord, record.Concept, record.Date, record.Quantity, record.IsExpensse, b.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
