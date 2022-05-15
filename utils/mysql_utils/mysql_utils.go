package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/sharkx018/bookstore_users-api/utils/errors"
	"strings"
)

const (
	noRowPresent = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), noRowPresent) {
			return errors.NewNotFoundError("no records matching the given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error processing request")
}
