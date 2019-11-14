package config

import (
	"database/sql"
	"log"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
)

//DB is database struct
var DB *sql.DB

//InitDB initialize database
func InitDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("cloudsqlpostgres", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}
}
