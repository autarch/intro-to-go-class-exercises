/*

For this exercise, you'll start with the code you just created for exercise
#7. You can paste it in below.

The goal of this exercise is to define an interface that will be shared by the
User and Group types. You will update your user code, while the exercise's
test code supplies the Group implementation.

First, rename the existing username type to Name. You will still have a
User.username field, but it will use this Name type (which will be shared by
the Group struct as well). This will require changes to your NewUser()
constructor as well as to the existing Username() and SetUsername() methods.

Then add a new method, Name(). This should return the Name for the user.

Finally, define an interface named HasName which requires a single method,
Name(), which returns a Name value. The tests will define a Group type that
also implements this method.

*/
