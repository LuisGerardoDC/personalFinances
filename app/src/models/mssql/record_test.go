package mssqlmodel

import (
	"testing"
	"time"

	requestModel "github.com/LuisGerardoDC/personalFinances/app/src/models/request"
	"github.com/stretchr/testify/assert"
)

func TestRequestToMssql(t *testing.T) {
	var (
		mockedTime          = time.Now()
		mockedRequestRecord = requestModel.Record{
			BudgetID:  123,
			Concept:   "Record 1",
			Quantity:  5628.52,
			Date:      mockedTime,
			IsExpense: true,
		}
		mockedRecord = Record{}
	)

	mockedRecord.RequestToMssql(mockedRequestRecord)

	assert.Equal(t, mockedRecord.BudgetID, int64(123))
	assert.Equal(t, mockedRecord.Concept, "Record 1")
	assert.Equal(t, mockedRecord.Quantity, float32(5628.52))
	assert.Equal(t, mockedRecord.Date, mockedTime)
	assert.Equal(t, mockedRecord.IsExpense, true)

}
