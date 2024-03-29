package users

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/sharkx018/bookstore_users-api/datasources/mysql/users_db"
	"github.com/sharkx018/bookstore_utils-go/logger"
	"github.com/sharkx018/bookstore_utils-go/rest_errors"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?,?,?,?,?,?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?"
	queryFindUserByStatus       = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?;"
)

const (
	StatusActive = "active"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() rest_errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement.", err)
		return rest_errors.NewInternalServerError("Error when trying to get the user", errors.New("database error"))
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return rest_errors.NewInternalServerError("Error when trying to get the user", errors.New("database error"))
		//return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Save() rest_errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return rest_errors.NewInternalServerError("Error when trying to save the user", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		logger.Error("error when trying to save user", saveErr)
		return rest_errors.NewInternalServerError("Error when trying to save the user", errors.New("database error"))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert user after creating the new user", err)
		return rest_errors.NewInternalServerError("Error when trying to save the user", errors.New("database error"))
		//return mysql_utils.ParseError(err)
		//return errors.NewInternalServerError(fmt.Sprintf("error when trying to save the user: %s", err.Error()))
	}

	user.Id = userId

	return nil

}

func (user *User) Update() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return rest_errors.NewInternalServerError("Error when trying to update the user", errors.New("database error"))
		//return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	updatedResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if saveErr != nil {
		logger.Error("error when trying to update user", saveErr)
		return rest_errors.NewInternalServerError("Error when trying to udpate the user", errors.New("database error"))
		//return mysql_utils.ParseError(saveErr)
	}

	_, err = updatedResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert user after creating the new user", err)
		return rest_errors.NewInternalServerError("Error when trying to update the user", errors.New("database error"))
		//return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Delete() rest_errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return rest_errors.NewInternalServerError("Error when trying to delete the user", errors.New("database error"))
		//return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Id); err != nil {
		logger.Error("error when trying to update user", err)
		return rest_errors.NewInternalServerError("Error when trying to delete the user", errors.New("database error"))
		//return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, rest_errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return nil, rest_errors.NewInternalServerError("Error when trying to find by status the user", errors.New("database error"))
		//return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find users by status", err)
		return nil, rest_errors.NewInternalServerError("Error when trying to find by status the user", errors.New("database error"))
		//return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when trying to scan user row into user struct", err)
			return nil, rest_errors.NewInternalServerError("Error when trying to get the user", errors.New("database error"))
			//return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}

func (user *User) GetByEmailAndPassword() rest_errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare get user by email and password statement.", err)
		return rest_errors.NewInternalServerError("Error when trying to find the user", errors.New("database error"))
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {

		if err == sql.ErrNoRows {
			logger.Error("error when trying to get user by email and password statement", err)
			return rest_errors.NewNotFoundError("No user found")
		}
		logger.Error("error when trying to get user by email and password statement", err)
		return rest_errors.NewInternalServerError("Error when trying to find the user", errors.New("database error"))
		//return mysql_utils.ParseError(err)
	}

	return nil
}
