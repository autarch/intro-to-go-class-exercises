package http

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type myError struct {
	status  uint
	message string
}

type HTTPForbiddenError myError

func NewHTTPForbiddenError() HTTPForbiddenError {
	return HTTPForbiddenError{status: 403, message: "Forbidden"}
}

func (e HTTPForbiddenError) HTTPStatus() uint {
	return e.status
}

func (e HTTPForbiddenError) Error() string {
	return e.message
}

func TestErrors(t *testing.T) {
	fe := NewHTTPForbiddenError()
	// So gross! We construct a nil pointer to something that does the
	// interface we want, then make the Type for that. Calling Elem() on a
	// pointer type returns the type to which it is a pointer, which in this
	// case is the HTTPError interface. My eyes are bleeding!
	heface := reflect.TypeOf((*HTTPError)(nil)).Elem()
	assert.True(
		t,
		reflect.TypeOf(fe).Implements(heface),
		"HTTPError interface requires Error() and HTTPStatus() methods",
	)

	testError(t, fe, 403)

	br := NewHTTPBadRequestError()
	testError(t, br, 400)

	nf := NewHTTPNotFoundError()
	testError(t, nf, 404)

}

func testError(t *testing.T, he HTTPError, status uint) {
	assert.Equal(t, status, he.HTTPStatus(), "HTTPStatus() returns %d", status)
	assert.NotEqual(t, "", he.Error(), "Error() is not empty")

	s, m := ErrorResponse(he)
	assert.Equal(t, status, s, "ErrorResponse() returns %d for error", status)
	assert.Equal(t, he.Error(), m, "ErrorResponse() returns %s for error", he.Error())
}
