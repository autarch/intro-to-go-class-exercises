package user

import (
	"crypto/sha1"
	"errors"
)

type EntityName string
type password [20]byte

type User struct {
	username EntityName
	password password
}

type HasName interface {
	Name() EntityName
}

func NewUser(u string, p string) (User, error) {
	if u == "" {
		return User{}, errors.New("The username must be a non-empty string.")
	}
	if p == "" {
		return User{}, errors.New("The password must be a non-empty string.")
	}

	return User{
		username: EntityName(u),
		password: password(sha1.Sum([]byte(p))),
	}, nil
}

func (u User) PasswordIsValid(p string) (bool, error) {
	if p == "" {
		return false, errors.New("The password must be a non-empty string.")
	}

	return sha1.Sum([]byte(p)) == u.password, nil
}

func (u User) Username() EntityName {
	return u.username
}

func (u *User) SetUsername(new string) {
	u.username = EntityName(new)
}

func (u User) Name() EntityName {
	return u.Username()
}
