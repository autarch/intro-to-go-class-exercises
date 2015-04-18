package user

import (
	"crypto/sha1"
	"io/ioutil"
	"reflect"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestType(t *testing.T) {
	u, err := ioutil.ReadFile("user.go")
	if err != nil {
		t.Fatal(err)
	}

	assert := assert.New(t)

	if !assert.True(regexp.MustCompile("type username string").Match(u),
		"user.go defines the username type correctly") {

		t.Fatal("cannot continue tests without username type")
	}

	if !assert.True(regexp.MustCompile(`type password \[20\]byte`).Match(u),
		"user.go defines the password type correctly") {

		t.Fatal("cannot continue tests without password type")
	}

	if !assert.True(regexp.MustCompile("type User struct {").Match(u),
		"user.go defines a User struct type") {

		t.Fatal("cannot continue tests without User struct type")
	}

	user := User{}
	userType := reflect.TypeOf(user)
	if !assert.Equal(userType.Kind(), reflect.Struct, "User", "User type's kind is Struct") {
		t.Fatal("Cannot continue unless the User type is a struct")
	}
	if !assert.Equal(userType.NumField(), 2, "User type has two fields") {
		t.Fatal("Cannot continue unless the User type has exactly two fields")
	}

	type fieldInfo struct {
		name string
		kind reflect.Kind
	}
	fields := map[string]fieldInfo{
		"username": fieldInfo{
			name: "username",
			kind: reflect.String,
		},
		"password": fieldInfo{
			name: "password",
			kind: reflect.Array,
		},
	}
	for f, info := range fields {
		field, ok := userType.FieldByName(f)
		if !assert.True(ok, "User type has a field named \""+f+"\"") {
			t.Fatal("Cannot continue unless the User type has a " + f + " field")
		}

		if !assert.Equal(field.Type.Name(), info.name,
			"User type's "+f+" field is a "+info.name) {

			t.Fatal("Cannot continue unless the User." + f + " type name is " + info.name)
		}

		if !assert.Equal(field.Type.Kind(), info.kind,
			"User type's "+f+" field is a "+info.kind.String()) {

			t.Fatal("Cannot continue unless the User." + f + " type kind is " + info.kind.String())
		}
	}

	field, _ := userType.FieldByName("password")
	if !assert.Equal(field.Type.Elem(), reflect.TypeOf(byte(0)), "password field is an array of bytes") {
		t.Fatal("Cannot continue unless the User.password field is an array of bytes")
	}
	if !assert.Equal(field.Type.Len(), 20, "password field's array length is 20 bytes") {
		t.Fatal("Cannot continue unless the User.password field's array length is 20 bytes")
	}
}

func TestFunctions(t *testing.T) {
	assert := assert.New(t)

	u := NewUser("ringo", "apple")
	if !assert.Equal(reflect.TypeOf(u).Name(), "User", "NewUser() returns a User struct") {
		t.Fatal("Cannot continue unless NewUser() returns a User struct")
	}

	assert.Equal(u.username, username("ringo"), "the returned user's username is \"ringo\"")
	assert.Equal(u.password, password(sha1.Sum([]byte("apple"))),
		"the returned user's password is the hash for \"apple\"")

	assert.True(PasswordIsValid(u, "apple"), "PasswordIsValid returns true for a matching password")
	assert.False(PasswordIsValid(u, "orange"), "PasswordIsValid returns false for a non-matching password")
}
