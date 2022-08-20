package mysql_utils

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/sharkx018/bookstore_utils-go/rest_errors"
	"strings"
)

const (
	noRowPresent = "no rows in result set"
)

func ParseError(err error) *rest_errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), noRowPresent) {
			return rest_errors.NewNotFoundError("no records matching the given id")
		}
		return rest_errors.NewInternalServerError("error parsing database response", errors.New("database_error"))
	}

	switch sqlErr.Number {
	case 1062:
		return rest_errors.NewBadRequestError("invalid data")
	}
	return rest_errors.NewInternalServerError("error processing request", errors.New("database_error"))
}
