package database

import (
	"log"

	"github.com/OlegOdnoral/go_gin_app/src/config"
	"github.com/go-pg/pg"
)

var (
	dbConnection *pg.DB
)

//InitConnectionToDB set config for database connetion from config.toml file
func InitConnectionToDB() {
	if dbConfig, err := config.GetDBConfig(); err != nil {
		log.Fatalln("Invalid DB config")
	} else {
		//fmt.Printf("%+v\n", dbConfig)
		dbConnection = pg.Connect(&pg.Options{
			Addr:            dbConfig.Addr,
			User:            dbConfig.User,
			Password:        dbConfig.Password,
			Database:        dbConfig.Database,
			ApplicationName: dbConfig.ApplicationName,
		})
		//fmt.Printf("%+v\n", dbConnection.PoolStats())
	}

}

//CloseDBConnection close database connection
func CloseDBConnection() {
	dbConnection.Close()
}
