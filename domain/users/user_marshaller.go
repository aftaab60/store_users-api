package users

import "encoding/json"

type PublicUser struct {
	Id          int64  `json:"id"`
	DateCreated string `json:"dateCreated"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	DateCreated string `json:"dateCreated"`
}

func (user *User) Marshall() *PrivateUser {
	var privateUser PrivateUser
	userJson, _ := json.Marshal(user)
	if err := json.Unmarshal(userJson, &privateUser); err != nil {
		return nil
	}
	return &privateUser
}
