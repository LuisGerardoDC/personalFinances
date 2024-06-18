package models_test

import (
	"testing"
	"time"

	"github.com/LuisGerardoDC/personalFinances/app/src/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewBudget(t *testing.T) {
	mockedStartTime := time.Now()
	mockedEndTime := time.Now().Add(time.Hour * 360)
	mockedAssets := []models.Record{
		{
			Concept:  "Test",
			Quantity: 1000.0,
			Date:     mockedStartTime,
		},
	}
	expectedBudget := models.Budget{
		StartTime:       mockedStartTime,
		EndTime:         mockedEndTime,
		Assets:          mockedAssets,
		Expenses:        nil,
		UsedBudget:      0,
		RemainingBudget: 1000.0,
	}

	gottedBudget := models.Budget{}

	gottedBudget.NewBudget(mockedAssets, mockedStartTime, mockedEndTime)

	assert.Equal(t, expectedBudget.StartTime, gottedBudget.StartTime)
	assert.Equal(t, expectedBudget.EndTime, gottedBudget.EndTime)
	assert.Equal(t, expectedBudget.Assets, gottedBudget.Assets)
	assert.Equal(t, expectedBudget.Expenses, gottedBudget.Expenses)
	assert.Equal(t, expectedBudget.UsedBudget, gottedBudget.UsedBudget)
	assert.Equal(t, expectedBudget.RemainingBudget, gottedBudget.RemainingBudget)

}

func TestCalcBudget(t *testing.T) {
	mockedStartTime := time.Now()
	mockedAssets := []models.Record{
		{
			Concept:  "Test",
			Quantity: 1000.0,
			Date:     mockedStartTime,
		},
		{
			Concept:  "Test1",
			Quantity: 2000.0,
			Date:     mockedStartTime,
		},
		{
			Concept:  "Test1",
			Quantity: 223.45,
			Date:     mockedStartTime,
		},
	}
	mockedExpenses := []models.Record{
		{
			Concept:  "Test",
			Quantity: 10,
			Date:     mockedStartTime,
		},
		{
			Concept:  "Test1",
			Quantity: 200,
			Date:     mockedStartTime,
		},
		{
			Concept:  "Test1",
			Quantity: 3000,
			Date:     mockedStartTime,
		},
	}

	ExpectedTotalBudget := float32(3223.45)
	ExpectedUsedBudget := float32(3210)
	ExpectedRemainingBudget := float32(13.45)

	mockedBudget := models.Budget{
		Assets:   mockedAssets,
		Expenses: mockedExpenses,
	}
	mockedBudget.CalcBudgets()

	assert.Equal(t, mockedBudget.RemainingBudget, ExpectedRemainingBudget)
	assert.Equal(t, mockedBudget.UsedBudget, ExpectedUsedBudget)
	assert.Equal(t, mockedBudget.TotalBudget, ExpectedTotalBudget)

}
