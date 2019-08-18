package main

import (
	"fmt"

	"github.com/Tempeny/gin_tes/src/routes"
)

const (
	port = ":4200"
)

func main() {
	fmt.Println("Hello from GO 2")
	routes.RunAndServe(port)
}
