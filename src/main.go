package main

import (
	"fmt"

	"github.com/go-redis/redis"

	"github.com/OlegOdnoral/go_gin_app/src/database"

	"github.com/OlegOdnoral/go_gin_app/src/routes"

	"github.com/OlegOdnoral/go_gin_app/src/dav"
)

const (
	port = ":8080"
)

func main() {
	database.InitConnectionToDB()
	defer database.CloseDBConnection()
	testRedis()

	dav.GetFsInfo()
	routes.RunAndServe(port)

}

func testRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	writeErr := client.Set("key", "value", 0).Err()
	if writeErr != nil {
		panic(writeErr)
	}

	val, readErr := client.Get("key").Result()
	if readErr != nil {
		panic(readErr)
	}
	fmt.Println("key", val)
}
