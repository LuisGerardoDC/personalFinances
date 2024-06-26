package mssqlmodel

import (
	"testing"
	"time"

	requestModel "github.com/LuisGerardoDC/personalFinances/app/src/models/request"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewBudget(t *testing.T) {
	var (
		mockedStartTime = time.Now()
		mockedEndTime   = time.Now().Add(time.Hour * 360)

		mockedAssets = []requestModel.Record{
			{
				Concept:  "asset 1",
				Quantity: 1000.0,
				Date:     mockedStartTime,
			},
		}
		mockedExpenses = []requestModel.Record{
			{
				Concept:  "Expense 1",
				Quantity: 100.0,
				Date:     mockedStartTime,
			},
		}
		mockedReqBudget = requestModel.Budget{
			Name:      "Test1",
			StartTime: mockedStartTime,
			EndTime:   mockedEndTime,
			UserID:    123456,
			Assets:    mockedAssets,
			Expenses:  mockedExpenses,
		}

		expectedBudget = Budget{
			StartTime:       mockedStartTime,
			EndTime:         mockedEndTime,
			UsedBudget:      100,
			RemainingBudget: 900.0,
			TotalBudget:     1000,
		}

		gottedBudget = Budget{}
	)
	gottedBudget.NewBudget(mockedReqBudget)

	assert.Equal(t, expectedBudget.StartTime, gottedBudget.StartTime)
	assert.Equal(t, expectedBudget.EndTime, gottedBudget.EndTime)
	assert.Equal(t, expectedBudget.UsedBudget, gottedBudget.UsedBudget)
	assert.Equal(t, expectedBudget.RemainingBudget, gottedBudget.RemainingBudget)
	assert.Equal(t, expectedBudget.TotalBudget, gottedBudget.TotalBudget)
}

func TestCalcBudget(t *testing.T) {
	var (
		mockedStartTime = time.Now()
		mockedRecords   = []Record{
			{
				Concept:   "Test",
				Quantity:  1000.0,
				Date:      mockedStartTime,
				IsExpense: false,
			},
			{
				Concept:   "Test1",
				Quantity:  2000.0,
				Date:      mockedStartTime,
				IsExpense: false,
			},
			{
				Concept:   "Test1",
				Quantity:  223.45,
				Date:      mockedStartTime,
				IsExpense: false,
			},
			{
				Concept:   "Test",
				Quantity:  10,
				Date:      mockedStartTime,
				IsExpense: true,
			},
			{
				Concept:   "Test1",
				Quantity:  200,
				Date:      mockedStartTime,
				IsExpense: true,
			},
			{
				Concept:   "Test1",
				Quantity:  3000,
				Date:      mockedStartTime,
				IsExpense: true,
			},
		}

		ExpectedTotalBudget     = float32(3223.45)
		ExpectedUsedBudget      = float32(3210)
		ExpectedRemainingBudget = float32(13.45)

		mockedBudget = Budget{
			Records: mockedRecords,
		}
	)

	mockedBudget.CalcBudgets()

	assert.Equal(t, mockedBudget.RemainingBudget, ExpectedRemainingBudget)
	assert.Equal(t, mockedBudget.UsedBudget, ExpectedUsedBudget)
	assert.Equal(t, mockedBudget.TotalBudget, ExpectedTotalBudget)
}

func TestRecordToMssql(t *testing.T) {
	var (
		mockedTime   = time.Now()
		budget       = Budget{}
		mockedAssets = []requestModel.Record{
			{
				Concept:  "asset1",
				Quantity: 1000.0,
				Date:     mockedTime,
			},
			{
				Concept:  "asset2",
				Quantity: 2000.0,
				Date:     mockedTime,
			},
			{
				Concept:  "asset3",
				Quantity: 223.45,
				Date:     mockedTime,
			},
		}
		mockedExpenses = []requestModel.Record{
			{
				Concept:  "expense1",
				Quantity: 10,
				Date:     mockedTime,
			},
			{
				Concept:  "expense2",
				Quantity: 200,
				Date:     mockedTime,
			},
			{
				Concept:  "expense3",
				Quantity: 3000,
				Date:     mockedTime,
			},
		}
	)

	budget.RecordToMssql(mockedAssets, mockedExpenses)
	assert.Equal(t, len(budget.Records), 6)
}
