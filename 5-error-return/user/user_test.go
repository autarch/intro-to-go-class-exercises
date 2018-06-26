package user

import (
	"errors"
	"testing"

	"github.com/autarch/intro-to-go-class-exercises/helpers/usertests"
	"github.com/stretchr/testify/assert"
)

func TestTypes(t *testing.T) {
	usertests.TestUserTypes(t, User{}, "username")
}

func TestNew(t *testing.T) {
	u, err := NewUser("", "password")

	assert.NotEqual(t, err, nil, "got an error when the username is empty")
	assert.Equal(
		t,
		errors.New("The username must be a non-empty string."),
		err,
		"got the expected error for an empty username",
	)
	assert.Equal(t, User{}, u, "user returned on error is empty")

	u, err = NewUser("username", "")

	assert.NotEqual(t, err, nil, "got an error when the password is empty")
	assert.Equal(
		t,
		errors.New("The password must be a non-empty string."),
		err,
		"got the expected error for an empty password",
	)
	assert.Equal(t, User{}, u, "user returned on error is empty")

	u, err = NewUser("username", "password")
	assert.Equal(t, nil, err, "got no error when username and password are both non-empty")
	assert.NotEmpty(t, u, "user returned on success is not empty")
}

func TestPasswordIsValid(t *testing.T) {
	u, err := NewUser("username", "password")

	ok, err := PasswordIsValid(u, "bad")
	assert.False(t, ok, "ok is false when password is invalid")
	assert.Equal(t, nil, err, "error is nil when password is invalid")

	ok, err = PasswordIsValid(u, "password")
	assert.True(t, ok, "ok is true when password is valid")
	assert.Equal(t, nil, err, "error is nil when password is valid")

	ok, err = PasswordIsValid(u, "")
	assert.False(t, ok, "ok is false when password is empty")
	assert.Equal(
		t,
		errors.New("The password must be a non-empty string."),
		err,
		"got the expected error for an empty password",
	)
}
