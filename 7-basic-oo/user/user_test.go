package user

import (
	"errors"
	"reflect"
	"testing"

	"github.com/autarch/intro-to-go-class-exercises/helpers/usertests"
	"github.com/stretchr/testify/assert"
)

func TestTypes(t *testing.T) {
	usertests.TestUserTypes(t, User{}, "username")
}

/* Copied from 5-error-return */
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

func TestMethods(t *testing.T) {
	assert := assert.New(t)

	u, err := NewUser("ringo", "apple")
	if !assert.Equal("User", reflect.TypeOf(u).Name(), "NewUser() returns a User struct") {
		t.Fatal("Cannot continue unless NewUser() returns a User struct")
	}

	assert.Equal(username("ringo"), u.Username(), "Username method returns the correct username value")

	u.SetUsername("shiina")
	assert.Equal(username("shiina"), u.Username(), "Username method returns new username after call to SetUsername")

	ok, err := u.PasswordIsValid("apple")
	assert.True(ok, "PasswordIsValid returns true for a matching password")
	assert.Empty(err, "PasswordIsValid does not return an error for a valid password")

	ok, err = u.PasswordIsValid("orange")
	assert.False(ok, "PasswordIsValid returns false for a non-matching password")
	assert.Empty(err, "PasswordIsValid does not return an error for an invalid password")

	ok, err = u.PasswordIsValid("")
	assert.False(ok, "PasswordIsValid returns false for empty password")
	assert.NotEmpty(err, "PasswordIsValid return an error for an empty password")
}
