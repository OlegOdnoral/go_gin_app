package main

import (
	"fmt"

	"github.com/Tempeny/gin_tes/src/config"

	"github.com/Tempeny/gin_tes/src/database"

	"github.com/Tempeny/gin_tes/src/routes"
)

const (
	port = ":4200"
)

func main() {
	database.InitConnectionToDB()
	defer database.CloseDBConnection()

	config.GetConfig()

	fmt.Println("Hello from GO 2")
	routes.RunAndServe(port)
}
