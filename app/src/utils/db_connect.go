package utils

import (
	"database/sql"
	"fmt"
)

func ConnectDB() *sql.DB {

	usuario := ""
	pass := ""
	host := ""
	dbName := ""

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, dbName))
	if err != nil {
		return nil
	}
	return db
}
