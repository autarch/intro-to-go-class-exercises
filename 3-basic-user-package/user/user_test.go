package user

import (
	"crypto/sha1"
	"reflect"
	"testing"

	"github.com/autarch/intro-to-go-class-exercises/helpers/usertests"

	"github.com/stretchr/testify/assert"
)

func TestTypes(t *testing.T) {
	usertests.TestUserTypes(t, User{}, "username")
}

func TestFunctions(t *testing.T) {
	assert := assert.New(t)

	u := NewUser("ringo", "apple")
	if !assert.Equal("User", reflect.TypeOf(u).Name(), "NewUser() returns a User struct") {
		t.Fatal("Cannot continue unless NewUser() returns a User struct")
	}

	assert.Equal(username("ringo"), u.username, `the returned user's username is "ringo"`)
	assert.Equal(password(sha1.Sum([]byte("apple"))), u.password,
		`the returned user's password is the hash for "apple"`)

	assert.True(PasswordIsValid(u, "apple"), "PasswordIsValid returns true for a matching password")
	assert.False(PasswordIsValid(u, "orange"), "PasswordIsValid returns false for a non-matching password")
}
