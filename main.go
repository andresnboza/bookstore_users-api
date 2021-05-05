package main

import (
	"fmt"

	"bookstore_users-api/app"
)

func main() {
	fmt.Println("Server working on: http://localhost:8081")
	app.StartApplication()
}
