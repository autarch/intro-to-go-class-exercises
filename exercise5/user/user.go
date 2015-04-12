/*

For this exercise, you'll start with the code you just created for exercise
#3 (not #4). You can paste it in below.

For this exercise, you're going to alter the functions you created in exercise
#3 so that they return errors.

Change the NewUser() function so that it returns two values, a User and an
error. You should check that the username and password passed to the function
are not empty (!= ""). If they *are* empty, return an empty User (User{}) and
an error. The errors you should use are:

* "The username must be a non-empty string."
* "The password must be a non-empty string."

Change the PasswordIsValid() function to also return two values, a boolean
indicating whether there was a match and an error. If the password you get is
empty, then return an false and an error with the message "The password must
be a non-empty string." If the password is non-empty but doesn't match, this
is not an error, just return false and a nil error.

Use the errors.New() function create the errors.

*/
