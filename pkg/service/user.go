package service

import "TechnicalAssignment/pkg/constants"

type User struct {
	Username string
}

func (u *User) GetUsername() string {
	if u.Username == "" {
		return constants.Unregistered
	}
	return u.Username
}
