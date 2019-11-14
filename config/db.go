package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//DB is database struct
var DB *sql.DB

//InitDB initialize database
func InitDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}
}
