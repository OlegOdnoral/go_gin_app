package database

import "github.com/go-pg/pg"

var (
	dbConnection *pg.DB
)

func InitConnectionToDB() {
	dbConnection = pg.Connect(&pg.Options{})
}

func CloseDBConnection() {
	dbConnection.Close()
}
