/*

For this exercise, you'll be using the "crypto/sha1" core package. See
http://golang.org/pkg/crypto/sha1/ for documentation.

Create a package named user that defines the following types:

- username (based on the string type)
- password (based on the [20]byte type)
- User, a struct containing a "username" and "password" field (you can figure
  out the appropriate types)

The package should implement the following functions:

- func NewUser(username string, password string) User { ... }
- func PasswordIsValid(user User, password string) bool { ... }

The NewUser() function should hash the password using the sha1.Sum() function
and store it in the User struct. You'll also use the sha1 package in the
PasswordIsValid() function to hash the string that you're passed.

Note that sha1.Sum() expects a byte array ([]byte), not a string, so you'll
need to convert the password to []byte when passing them to sha1.Sum().

To run the tests, run "go test -v".

    > go test -v

*/
