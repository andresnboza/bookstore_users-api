package main

import (
	"fmt"

	"github.com/user/bookstore_users-api/app"
)

func main() {
	fmt.Println("Server working on: http://localhost:8080")
	app.StartApplication()
}
