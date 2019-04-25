package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// DB is the instance of connection
var DB *sql.DB

// ConnectDB is to connect to the database
func ConnectDB() {
	Log("info", "Connecting to Database")

	// Open up our database connection.
	DB, err := sql.Open("mysql", viper.GetString("database.user")+":"+viper.GetString("database.pass")+"@tcp("+viper.GetString("database.addr")+")/"+viper.GetString("database.name"))

	// if there is an error opening the connection, handle it
	if err != nil {
		Log("error", "Failed connection to Database")
		panic(err.Error())
	}

	// Ping to check if Credentials are correct
	err = DB.Ping()
	if err != nil {
		Log("error", "Failed connection to Database")
		panic(err.Error())
	} else {
		Log("info", "Successfully Connected to Database: "+viper.GetString("database.name")+" Address: "+viper.GetString("database.addr"))
	}

}
