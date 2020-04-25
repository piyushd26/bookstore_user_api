package services

import (
	"bookstore/bookstore_user_api/src/domain/users"
	"bookstore/bookstore_user_api/utils/errors"
)

func CreateUser(user *users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestError) {
	if userId <= 0 {
		return nil, errors.BadRequestError("userId less than zero")
	}
	user := users.User{
		ID: userId,
	}
	err := user.Get()
	if err != nil {
		return nil, err
	}
	return &user, nil
}
