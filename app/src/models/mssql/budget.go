package mssqlmodel

import (
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
