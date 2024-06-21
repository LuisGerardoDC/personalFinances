package mssqlmodel

import (
	"database/sql"
	"time"
)

type Record struct {
	ID         int
	Concept    string
	Quantity   float32
	Date       time.Time
	IsExpensse bool
}

func (r *Record) SaveRecord(bd *sql.DB) error {
	return nil
}

func (r Record) SaveRecords(bd *sql.DB, records []Record) {

}
