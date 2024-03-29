package mysql_utils

import (
	"fmt"
	"github.com/aftaab60/store_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlError, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("no record matching given id"))
		}
		return errors.NewInternalServerError("error parsing database response")
	}
	switch sqlError.Number {
	case 1062:
		return errors.NewBadRequestError("record already existing")
	}
	return errors.NewInternalServerError("error processing request")
}
