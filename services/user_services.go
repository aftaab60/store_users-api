package services

import (
	"fmt"
	"github.com/aftaab60/store_users-api/domain/users"
	"github.com/aftaab60/store_users-api/utils/crypto_utils"
	"github.com/aftaab60/store_users-api/utils/date_utils"
	"github.com/aftaab60/store_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.DateCreated = date_utils.NowTimeString()
	user.Password = crypto_utils.GetHash(user.Password)

	if saveErr := user.Save(); saveErr != nil {
		return nil, saveErr
	}
	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	user := &users.User{Id: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	currUser, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	fmt.Sprint("validation error")
	if isPartial {
		if user.FirstName != "" {
			currUser.FirstName = user.FirstName
		}
		if user.LastName != "" {
			currUser.LastName = user.LastName
		}
		if user.Email != "" {
			currUser.Email = user.Email
		}
	} else {
		currUser.FirstName = user.FirstName
		currUser.LastName = user.LastName
		currUser.Email = user.Email
	}

	if err := currUser.Validate(); err != nil {
		return nil, err
	}
	if err := currUser.Update(); err != nil {
		return nil, err
	}
	return currUser, nil
}
