/*

For this exercise, you'll start with the code you just created for exercise
#3. You can paste it in below.

You're going to create two new functions in this package. The first,
PasswordCheckAnd(), will take four arguments. The first argument is a User
struct. The second is a password string. The next two arguments are functions
themselves, and they have the following signature:

    func X(u *User) { ... }

In other words, these are functions that accept a single argument, a pointer
to a User struct, and return nothing.

The function you're writing will call the first function if the password
matches, and the second if it doesn't.

The other function you will create is called PasswordCheckFunc. This function
takes a user and returns a function with the following signature:

    func X(password string) bool { ... }

This function will take a password and return true or false, based on whether
it matches the user's password. You'll close over the passed in User struct in
the function you return.

Note 1: This is not particularly idiomatic Go, but I want you to get some
practice dealing with first class functions in Go while building on what
you've already written.

Note 2: If first class functions are making your brain explode, don't worry,
you won't have to build on this particular exercise in the future.

*/
