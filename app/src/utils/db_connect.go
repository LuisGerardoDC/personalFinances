package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {

	usuario := "personalFinancesBack"
	pass := "Yolu@54&loki"
	host := "localhost:3306"
	dbName := "personal_finances"

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", usuario, pass, host, dbName))
	if err != nil {
		panic(err.Error())
	}
	return db
}
