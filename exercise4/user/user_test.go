package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordCheckAnd(t *testing.T) {
	user := NewUser("bob", "pw1")

	var pName username
	pass := func(u *User) {
		pName = u.username
	}

	var fName username
	fail := func(u *User) {
		fName = u.username
	}

	PasswordCheckAnd(user, "pw1", pass, fail)
	assert.Equal(t, pName, "bob", "bob user was passed on successful password check")
	assert.Equal(t, fName, "", "the fail function was not called on a successful password")

	pName = ""
	PasswordCheckAnd(user, "bad pw", pass, fail)
	assert.Equal(t, pName, "", "the success function was not called on a failed password check")
	assert.Equal(t, fName, "bob", "bob user was passed on failed password check")
}

func TestPasswordCheckFunc(t *testing.T) {
	user := NewUser("bob", "pw1")

	check := PasswordCheckFunc(user)

	assert.True(t, check("pw1"), "PasswordCheckFunc's returned func returns true on valid password")
	assert.True(t, !check("bad pw"), "PasswordCheckFunc's returned func returns false on invalid password")
}
