package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host     = "mysql_users_host"
	mysql_users_schema   = "mysql_users_host"
)

var (
	Client *sql.DB

	// os.Setenv(mysql_users_username, 	"root")
	// os.Setenv(mysql_users_password,	"admin123")
	// os.Setenv(mysql_users_host, 	    "127.0.0.1:3306")
	// os.Setenv(mysql_users_schema, 	"users_db")

	// username = os.Getenv(mysql_users_username)
	// password = os.Getenv(mysql_users_password)
	// host     = os.Getenv(mysql_users_host)
	// schema   = os.Getenv(mysql_users_schema)

	// TODO: eliminate these when useing docker
	username = "andres"
	password = "Taquito.123"
	host     = "127.0.0.1:3306"
	schema   = "users_db"
)

// This is a function that is executed when the package in imported
func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
