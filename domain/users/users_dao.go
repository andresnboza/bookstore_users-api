package users

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/user/bookstore_users-api/datasources/mysql/users_db"
	"github.com/user/bookstore_users-api/utils/date_utils"
	"github.com/user/bookstore_users-api/utils/errors"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRow       = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	// Preparing the statement to the user save
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close() // Very important to execute

	// Making the call to the database with the statement
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRow) {
			return errors.NewNotFoundError(fmt.Sprintf("user with id (%d) does NOT exists", user.Id))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d: %s", user.Id, err.Error()))
	}

	return nil
}

func (user *User) Save() *errors.RestErr {

	// Preparing the statement to the user save
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close() // Very important to execute

	// Setting the user.date_created
	user.DateCreated = date_utils.GetNowString()

	// Inserting the user with the valid statement
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		// Looking if we are handling a mysql error, an error that we can't have any certainty
		sqlErr, ok := saveErr.(*mysql.MySQLError)
		if !ok {
			return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", saveErr.Error()))
		}
		// Now looing for more custom errors
		switch sqlErr.Number {
		case 1062:
			return errors.NewBadRequestError(fmt.Sprintf("email (%s) already exists", user.Email))
		default:
			return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", saveErr.Error()))
		}
	}

	// Getting the id of the recently created user
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying get last inserted id of user: %s", err.Error()))
	}

	// Setting the user id to the result
	user.Id = userId
	return nil
}
