package main

import (
	"fmt"

	"github.com/Tempeny/gin_tes/src/database"

	"github.com/Tempeny/gin_tes/src/routes"
)

const (
	port = ":4200"
)

func main() {
	database.InitConnectionToDB()
	defer database.CloseDBConnection()

	fmt.Println("Hello from GO GIN App")
	routes.RunAndServe(port)
}
