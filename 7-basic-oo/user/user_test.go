package user

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMethods(t *testing.T) {
	assert := assert.New(t)

	u, err := NewUser("ringo", "apple")
	if !assert.Equal("User", reflect.TypeOf(u).Name(), "NewUser() returns a User struct") {
		t.Fatal("Cannot continue unless NewUser() returns a User struct")
	}

	assert.Equal("ringo", u.Username(), "Username method returns username")

	u.SetUsername("shiina")
	assert.Equal("shiina", u.Username(), "Username method returns new username after call to SetUsername")

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
