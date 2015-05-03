package http

type HTTPError interface {
	Error() string
	HTTPStatus() uint
}

type httpError struct {
	status  uint
	message string
}

type HTTPBadRequestError httpError

func NewHTTPBadRequestError() HTTPBadRequestError {
	return HTTPBadRequestError{status: 400, message: "Bad Request"}
}

func (e HTTPBadRequestError) HTTPStatus() uint {
	return e.status
}

func (e HTTPBadRequestError) Error() string {
	return e.message
}

type HTTPNotFoundError httpError

func NewHTTPNotFoundError() HTTPNotFoundError {
	return HTTPNotFoundError{status: 404, message: "Not Found"}
}

func (e HTTPNotFoundError) HTTPStatus() uint {
	return e.status
}

func (e HTTPNotFoundError) Error() string {
	return e.message
}

func ErrorResponse(e error) (uint, string) {
	switch h := e.(type) {
	case HTTPError:
		return h.HTTPStatus(), h.Error()
	default:
		return 500, h.Error()
	}
}
