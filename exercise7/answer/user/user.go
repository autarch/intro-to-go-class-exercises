package user

import (
	"crypto/sha1"
	"errors"
)

type username string
type password [20]byte

type User struct {
	username username
	password password
}

func NewUser(u string, p string) (User, error) {
	if u == "" {
		return User{}, errors.New("The username must be a non-empty string.")
	}
	if p == "" {
		return User{}, errors.New("The password must be a non-empty string.")
	}

	return User{
		username: username(u),
		password: password(sha1.Sum([]byte(p))),
	}, nil
}

func (u User) PasswordIsValid(p string) (bool, error) {
	if p == "" {
		return false, errors.New("The password must be a non-empty string.")
	}

	return sha1.Sum([]byte(p)) == u.password, nil
}

func (u User) Username() username {
	return u.username
}

func (u *User) SetUsername(new string) {
	u.username = username(new)
}
