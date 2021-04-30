package users

import (
	"github.com/user/bookstore_users-api/datasources/mysql/users_db"
	"github.com/user/bookstore_users-api/utils/date_utils"
	"github.com/user/bookstore_users-api/utils/errors"
	"github.com/user/bookstore_users-api/utils/mysql_utils"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser) // Preparing the statement to the user save
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close() // Very important to execute

	result := stmt.QueryRow(user.Id) // Making the call to the database with the statement

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser) // Preparing the statement to the user save
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close() // Very important to execute

	user.DateCreated = date_utils.GetNowString() // Setting the user.date_created

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated) // Inserting the user with the valid statement

	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId() // Getting the id of the recently created user
	if err != nil {
		return mysql_utils.ParseError(saveErr)
	}

	user.Id = userId // Setting the user id to the result
	return nil
}
