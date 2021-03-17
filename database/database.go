package database

import (
	"database/sql"
	"log"
)

var DbConnection *sql.DB

// SetupDatabase
func SetupDatabase() {
	var err error
	DbConnection, err = sql.Open("mysql", "root:nimda@tcp(127.0.0.1:3306)/inventorydb")
	if err != nil {
		log.Fatal(err)
	}

	/*
		// For connection pooling
			DbConnection.SetMaxOpenConns(3)
			DbConnection.SetMaxIdleConns(3)
			DbConnection.SetConnMaxLifetime(60 * time.Second)
	*/
}
