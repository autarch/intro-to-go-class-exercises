package user

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	user, err := NewUser("", "password")

	assert.NotEqual(t, err, nil, "got an error when the username is empty")
	assert.Equal(
		t,
		err,
		errors.New("The username must be a non-empty string."),
		"got the expected error for an empty username",
	)
	assert.Equal(t, user, User{}, "user returned on error is empty")

	user, err = NewUser("username", "")

	assert.NotEqual(t, err, nil, "got an error when the password is empty")
	assert.Equal(
		t,
		err,
		errors.New("The password must be a non-empty string."),
		"got the expected error for an empty password",
	)
	assert.Equal(t, user, User{}, "user returned on error is empty")

	user, err = NewUser("username", "password")
	assert.Equal(t, err, nil, "got no error when username and password are both non-empty")
	assert.NotEmpty(t, user, "user returned on success is not empty")
}
