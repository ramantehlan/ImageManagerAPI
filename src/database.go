package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// ConnectDB is to connect to the database
func ConnectDB() *sql.DB {
	Log("debug", "Connecting to Database")

	// Open up our database connection.
	DB, err := sql.Open("mysql", viper.GetString("database.user")+":"+viper.GetString("database.pass")+"@tcp("+viper.GetString("database.addr")+")/"+viper.GetString("database.name"))
	ErrorHandler(err)

	// Ping to check if Credentials are correct
	err = DB.Ping()
	ErrorHandler(err)
	if err == nil {
		Log("info", "Successfully Connected to Database: "+viper.GetString("database.name")+" Address: "+viper.GetString("database.addr"))
	}

	return DB
}
