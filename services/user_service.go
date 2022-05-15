package services

import (
	"github.com/sharkx018/bookstore_users-api/domain/users"
	"github.com/sharkx018/bookstore_users-api/utils/crypto_utils"
	"github.com/sharkx018/bookstore_users-api/utils/date_utils"
	"github.com/sharkx018/bookstore_users-api/utils/errors"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	GetUser(userId int64) (*users.User, *errors.RestErr)
	CreateUser(user users.User) (*users.User, *errors.RestErr)
	UpdateUser(isPartail bool, user users.User) (*users.User, *errors.RestErr)
	DeleteUser(userId int64) *errors.RestErr
	SearchUser(status string) (users.Users, *errors.RestErr)
}

func (us *userService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	if userId <= 0 {
		return nil, errors.NewBadRequestError("invalid user id")
	}

	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil

}

func (us *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := users.Validate(&user); err != nil {
		return nil, err
	}
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *userService) UpdateUser(isPartail bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := us.GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if err := users.Validate(&user); err != nil {
		return nil, err
	}

	if isPartail {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil

}

func (us *userService) DeleteUser(userId int64) *errors.RestErr {
	user := users.User{Id: userId}
	return user.Delete()
}

func (us *userService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
