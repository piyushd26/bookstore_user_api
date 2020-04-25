package users

import (
	"bookstore/bookstore_user_api/utils/errors"
	"fmt"
)

var usersDb = make(map[int64]*User)

func (user *User) Save() *errors.RestError {
	currentUser := usersDb[user.ID]
	if currentUser != nil {
		return errors.UserAlreadyPresent("User Already present")
	}
	usersDb[user.ID] = user
	return nil
}

func (user *User) Get() *errors.RestError {
	result := usersDb[user.ID]
	if result == nil {
		fmt.Println("user not found")
		return errors.UserNotFoundError("User not found")
	}
	// data from the db
	user.ID = result.ID
	user.DateCreated = result.DateCreated
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	return nil
}
