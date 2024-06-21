package mssqlmodel

import (
	"database/sql"
	"math"
	"time"
)

type Budget struct {
	ID              int
	UserID          int
	Name            string
	TotalBudget     float32
	StartTime       time.Time
	EndTime         time.Time
	Assets          []Record
	Expenses        []Record
	UsedBudget      float32
	RemainingBudget float32
}

func (b *Budget) NewBudget(assets []Record, startTime, endTime time.Time) {
	b.Assets = assets
	b.StartTime = startTime
	b.EndTime = endTime
	b.CalcBudgets()
}

func (b *Budget) CalcBudgets() {
	b.TotalBudget = 0
	b.UsedBudget = 0
	for _, asset := range b.Assets {
		b.TotalBudget += asset.Quantity
	}
	for _, expense := range b.Expenses {
		b.UsedBudget += expense.Quantity
	}

	b.RemainingBudget = b.TotalBudget - b.UsedBudget
	b.RemainingBudget = float32(math.Round(float64(b.RemainingBudget)*100) / 100)
}

func (b *Budget) CreateInDB(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO budgets (UserID,Name,TotalBudget,StartTime,EndTime,UsedBudget,RemainingBudget ) OUTPUT inserted.ID VALUES (@p1,@p2,@p3,@p4,@p5,@p6,@p7)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		b.UserID,
		b.Name,
		b.TotalBudget,
		b.StartTime,
		b.EndTime,
		b.UsedBudget,
		b.RemainingBudget,
	).Scan(b.ID)

	return nil
}
