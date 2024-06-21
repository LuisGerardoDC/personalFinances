package utils

import (
	"database/sql"
	"fmt"
)

func ConnectDB() *sql.DB {

	usuario := "personalFinancesBack"
	pass := "Yolu54loki"
	host := "localhost:"
	dbName := "personal_finances"

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, dbName))
	if err != nil {
		return nil
	}
	return db
}
