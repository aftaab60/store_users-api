package users

import (
	"fmt"
	"github.com/aftaab60/store_users-api/datasources/mysql"
	"github.com/aftaab60/store_users-api/utils/date_utils"
	"github.com/aftaab60/store_users-api/utils/errors"
	"strings"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "no rows in result set"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created from user where id=?;"
	queryInsertUser  = "INSERT into user(first_name, last_name, email, date_created) values(?,?,?,?);"
)

var userDB = make(map[int64]*User)

func (user *User) Get() *errors.RestErr {
	stmt, err := mysql.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	if err := stmt.QueryRow(user.Id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("user id %d not found", user.Id))
		}
		return errors.NewNotFoundError(fmt.Sprintf("Error when trying to get user %d", user.Id))
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := mysql.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	user.DateCreated = date_utils.NowTimeString()
	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error saving user into database: %s", err.Error()))
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error saving user into database: %s", err.Error()))
	}
	user.Id = userId
	return nil
}
