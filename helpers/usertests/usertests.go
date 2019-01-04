package usertests

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserTypes(t *testing.T, u interface{}, nameTypeName string) {
	src, err := ioutil.ReadFile("user.go")
	if err != nil {
		t.Fatal(err)
	}

	assert := assert.New(t)

	if !assert.True(regexp.MustCompile(fmt.Sprintf(`type\s+%s\s+string`, nameTypeName)).Match(src),
		"user.go defines the %s type correctly", nameTypeName) {

		t.Fatalf("cannot continue tests without %s type", nameTypeName)
	}

	if !assert.True(regexp.MustCompile(`type\s+password\s+\[20\]byte`).Match(src),
		"user.go defines the password type correctly") {

		t.Fatal("cannot continue tests without password type")
	}

	if !assert.True(regexp.MustCompile(`type\s+User\s+struct\s+{`).Match(src),
		"user.go defines a User struct type") {

		t.Fatal("cannot continue tests without User struct type")
	}

	userType := reflect.TypeOf(u)
	if !assert.Equal(reflect.Struct, userType.Kind(), "User", "User type's kind is Struct") {
		t.Fatal("Cannot continue unless the User type is a struct")
	}
	if !assert.Equal(2, userType.NumField(), "User type has two fields") {
		t.Fatal("Cannot continue unless the User type has exactly two fields")
	}

	type fieldInfo struct {
		name string
		kind reflect.Kind
	}
	fields := map[string]fieldInfo{
		"username": fieldInfo{
			name: nameTypeName,
			kind: reflect.String,
		},
		"password": fieldInfo{
			name: "password",
			kind: reflect.Array,
		},
	}
	for f, info := range fields {
		field, ok := userType.FieldByName(f)
		if !assert.True(ok, `User type has a field named "%s"`, f) {
			t.Fatalf("Cannot continue unless the User type has a %s field", f)
		}

		if !assert.Equal(info.name, field.Type.Name(),
			"User type's %s field is a %s", f, info.name) {

			t.Fatalf("Cannot continue unless the User.%s type name is a %s", f, info.name)
		}

		if !assert.Equal(info.kind, field.Type.Kind(),
			"User type's %s field is a %s", f, info.kind.String()) {

			t.Fatalf("Cannot continue unless the User.%s type kind is %s", f, info.kind.String())
		}
	}

	field, _ := userType.FieldByName("password")
	if !assert.Equal(reflect.TypeOf(byte(0)), field.Type.Elem(), "password field is an array of bytes") {
		t.Fatal("Cannot continue unless the User.password field is an array of bytes")
	}
	if !assert.Equal(field.Type.Len(), 20, "password field's array length is 20 bytes") {
		t.Fatal("Cannot continue unless the User.password field's array length is 20 bytes")
	}
}
