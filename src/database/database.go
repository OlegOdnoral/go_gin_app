package database

import (
	"fmt"
	"log"

	"github.com/Tempeny/gin_tes/src/config"
	"github.com/go-pg/pg"
)

var (
	dbConnection *pg.DB
)

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

func PoolStats() {
	fmt.Printf("%+v\n", dbConnection.PoolStats())
}

func CloseDBConnection() {
	dbConnection.Close()
}
