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
	b.StartTime = rb.StartTime
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
	var (
		getGudgetByID = "SELECT ID, UserID, Name, TotalBudget,StartTime,EndTime, UsedBudget, RemainingBudget FROM budgets WHERE ID = ?;"
		startTime     []byte
		endTime       []byte
	)
	// Gets Budget
	err := db.QueryRow(getGudgetByID, b.ID).Scan(
		&b.ID,
		&b.UserID,
		&b.Name,
		&b.TotalBudget,
		&startTime,
		&endTime,
		&b.UsedBudget,
		&b.RemainingBudget,
	)
	if err != nil {
		return err
	}
	b.StartTime, err = time.Parse("2006-01-02 15:04:05", string(startTime))
	if err != nil {
		return err
	}
	b.EndTime, err = time.Parse("2006-01-02 15:04:05", string(endTime))
	if err != nil {

		return err
	}

	return nil
}

func (b *Budget) GetRecords(db *sql.DB) error {
	var (
		getBudgetRecords = "SELECT ID, Concept, Date, Quantity, IsExpense FROM records WHERE BudgetID = ?;"
		records          = []Record{}
		readRecord       Record
		date             []byte
	)

	// Gets Budget's records
	stmt, err := db.Prepare(getBudgetRecords)
	if err != nil {
		return err
	}

	rows, err := stmt.Query(b.ID)
	if err != nil {

		return err
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
			return err
		}
		readRecord.Date, err = time.Parse("2006-01-02 15:04:05", string(date))
		if err != nil {
			return err
		}
		records = append(records, readRecord)
	}

	b.Records = records

	return nil
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
