package user

import (
	"crypto/sha1"
)

type username string
type password [20]byte

type User struct {
	username username
	password password
}

func NewUser(u string, p string) User {
	return User{
		username: username(u),
		password: password(sha1.Sum([]byte(p))),
	}
}

func PasswordIsValid(u User, p string) bool {
	return sha1.Sum([]byte(p)) == u.password
}
