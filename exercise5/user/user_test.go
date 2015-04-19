package user

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	u, err := NewUser("", "password")

	assert.NotEqual(t, err, nil, "got an error when the username is empty")
	assert.Equal(
		t,
		errors.New("The username must be a non-empty string."),
		err,
		"got the expected error for an empty username",
	)
	assert.Equal(t, User{}, user, "user returned on error is empty")

	u, err = NewUser("username", "")

	assert.NotEqual(t, err, nil, "got an error when the password is empty")
	assert.Equal(
		t,
		errors.New("The password must be a non-empty string."),
		err,
		"got the expected error for an empty password",
	)
	assert.Equal(t, User{}, user, "user returned on error is empty")

	u, err = NewUser("username", "password")
	assert.Equal(t, nil, err, "got no error when username and password are both non-empty")
	assert.NotEmpty(t, u, "user returned on success is not empty")
}
