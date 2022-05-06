package users

import (
	"github.com/aftaab60/store_users-api/utils/date_utils"
	"github.com/aftaab60/store_users-api/utils/errors"
)

var userDB = make(map[int64]*User)

func (user *User) Get() *errors.RestErr {
	currUser := userDB[user.Id]
	if currUser == nil {
		return errors.NewNotFoundError("user id does not exist")
	}
	user.Email = currUser.Email
	user.FirstName = currUser.FirstName
	user.LastName = currUser.LastName
	user.DateCreated = currUser.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	if userDB[user.Id] != nil {
		return errors.NewBadRequestError("User already exists")
	}
	user.DateCreated = date_utils.NowTimeString()
	userDB[user.Id] = user
	return nil
}
