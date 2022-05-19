package users

import (
	"github.com/aftaab60/store_users-api/datasources/mysql"
	"github.com/aftaab60/store_users-api/utils/errors"
	"github.com/aftaab60/store_users-api/utils/mysql_utils"
)

const (
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created from user where id=?;"
	queryInsertUser = "INSERT into user(first_name, last_name, email, date_created, password) values(?,?,?,?,?);"
	queryUpdateUser = "UPDATE user SET first_name=?, last_name=?, email=? where id=?;"
)

var userDB = make(map[int64]*User)

func (user *User) Get() *errors.RestErr {
	stmt, err := mysql.Client.Prepare(queryGetUser)
	defer stmt.Close()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	if getErr := stmt.QueryRow(user.Id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := mysql.Client.Prepare(queryInsertUser)
	defer stmt.Close()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	result, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := mysql.Client.Prepare(queryUpdateUser)
	defer stmt.Close()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	if _, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}
