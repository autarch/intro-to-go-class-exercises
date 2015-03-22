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

The NewUser() function should use the sha1 package to hash the password (use
sha1.Sum()) and store it in the User struct. You'll also USE the sha1 package
in the PasswordIsValid function to hash the string that you're passed.

Note that you'll have to convert strings to byte arrays in order to pass the
password string to sha1.Sum().

To run the tests, run "go test -v".

    > go test -v

*/
