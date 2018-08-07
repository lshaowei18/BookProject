package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type App struct {
	DB     *sql.DB
	Config Config
}

type Book struct {
	ID     int
	Name   string
	Author string
}

func main() {
	//Create new app
	app := App{}

	//Get config details
	app.GetConfig()

	//Connect to database
	var err error
	app.DB, err = sql.Open("postgres", app.Config.sqlInfo)
	CheckErr(err)

	//Ping to ensure that connection is alive
	err = app.DB.Ping()
	CheckErr(err)
}

func CheckErr(err error) {
	if err == sql.ErrNoRows {
		log.Fatal("No rows returned!")
	} else if err != nil {
		log.Fatal(err)
	}
}
