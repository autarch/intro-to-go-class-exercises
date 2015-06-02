/*

The goal for this exercise is to create an interface that represents HTTP
errors. This interface will also implement the built-in Error interface.

Call this interface httpError. It should require the Error interface (Error()
string) as well as a method called HTTPStatus() which returns a uint value (an
HTTP status code).

Then create several types which implement this interface:

* HTTPBadRequestError (400)
* HTTPNotFoundError (404)

Since each of these types consists of the same two elements (a status code and
a message), it will be easiest to do this by creating a single struct type and
then create the types above using that as the underlying type.

Create constructors named NewHTTPBadRequestError() and NewHTTPNotFoundError()
for those two types. The message for each error can be anything you want (but
don't leave it empty).

Each of those types should implement the httpError interface specified above.

Finally, create a function named ErrorResponse() that accepts *any* error type
and returns a uint and a string. The uint is an HTTP status and the string the
error message. If the error is *not* an httpError, return 500 for the status
and the error's Error() as the message. Otherwise return the status and
message from the httpError.

You should use either type assertions or the type switch construct to
implement the ErrorResponse() function.

*/
