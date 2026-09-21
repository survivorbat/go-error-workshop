# 💡 Instructions

These instructions guide you through the exercises and teach you the ins
and outs of errors in Golang.

## Error basics

### Creating errors

Take the following function.

```go
package users

import (
  "errors"
  "strings"
)

func CreateUser(name string) error {
  if name == "" {
    return errors.New("name can not be empty")
  }

  if strings.Contains(name, "root") {
    return fmt.Errorf("name %s can not contain root", name)
  }

  // [...]

  return nil
}
```

Here you see two functions used to create errors.
To check what error was returned, a caller has to read the `.Error()` message.

```go
package main

func main() {
  name := roots

  err := users.CreateUser("roots")
  if err != nil {
    msg := err.Error()

    if msg == "name can not be empty" {
      fmt.Println("Please ensure a name is set")
    }

    if strings.Contains(msg, "can not contain root") {
      fmt.Println("Due to security reasons, your name can not contain the word root")
    }
  }
}
```

These string comparisons the only way to do this.
It is not possible for the caller to compare `err` to another error object.

```go
err := uses.CreateUser()

// Does not work
err == errors.New("name can not be empty")

// Does not work
errors.Is(err, errors.New("name can not be empty"))
```

That's because `errors.New` and `fmt.Errorf` don't create strings, they create
pointers to new error objects.

### Exercises

Now please complete the following exercises.

- [Exercise 1](./exercise-1)

## Improving errors

### Sentinel errors

In the exercise you may have noticed that a lot of string comparisons
had to be done to figure out how to translate an error to the user.
On top of that, if the messages of the error somehow were to change, your code
would no longer work.

One improvement that can be made is to use **Sentinel errors**.
See the example below.

```go
package users

import (
  "errors"
  "strings"
)

var ErrNameEmpty = errors.New("name can not be empty")
var ErrNoRoot = errors.New("name can not contain root")

func CreateUser(name string) error {
  if name == "" {
    return ErrNameEmpty
  }

  if strings.Contains(name, "root") {
    return ErrNoRoot
  }

  // [...]

  return nil
}
```

Now, the returned error can be compared against the package-level variables.
The following comparisons are now possible.

```go
package main

func main() {
  name := roots

  err := users.CreateUser("roots")
  if err != nil {
    if errors.Is(err, users.ErrNameEmpty) {
      fmt.Println("Please ensure a name is set")
    }

    if errors.Is(err, users.ErrNoRoot)  {
      fmt.Println("Due to security reasons, your name can not contain the word root")
    }
  }
}
```

Making your error handling more robust and less string-based.
But wait, now the `name` variable is no longer present in the error message...
This means the `Error()` no longer contains this extra information.

### Wrapping errors

When we return errors, we have the ability to add additional information
through something called error wrapping.

```go
if strings.Contains(name, "root") {
  return fmt.Errorf("%s contains root: %w", ErrNoRoot)
}
```

This will expand the `.Error()` output to the following.

```text
roots contains root: name can not contain root`.
```

This is immensely useful when debugging a large codebase.
Imagine the following code.

```go
func createUser(name string) error {
  err := createDBUser(name)
  if err != nil {
    return err
  }

  err = createDBRoleFor(name)
  if err != nil {
    return err
  }

  err = createDBTeamFor(name)
  if err != nil {
    return err
  }

  return nil
}
```

Imagine seeing the following error in your logs.

```text
error="database table does not exist"
```

Which one of the three calls caused this error?
What table doesn't exist?
With error wrapping, more information can be added to this message.

```go
func createUser(name string) error {
  err := createDBUser(name)
  if err != nil {
    return fmt.Errorf("failed to create user %s: %w", name, err)
  }

  err = createDBRoleFor(name)
  if err != nil {
    return fmt.Errorf("failed to create role for %s: %w", name, err)
  }

  err = createDBTeamFor(name)
  if err != nil {
    return fmt.Errorf("failed to create team for %s: %w", name, err)
  }

  return nil
}
```

Applying error wrapping throughout your application adds more context
to the error as it works itself up the stack.
Eventually reaching your logging layer with a paper trail of what went wrong
and what parameters were involved.

Wrapped errors can be also unwrapped using the `errors.Unwrap` method.
The comparison methods `errors.Is` and `errors.AsType` automatically unwrap
errors while searching for an error in the tree.

```go
package main


func main() {
  a := errors.New("A")
  b := fmt.Errorf("B: %w", a)
  c := fmt.Errorf("C: %w", b)
  d := fmt.Errorf("D: %w", c)

  // Evaluates to true
  errors.Is(d, a)

  // Evaluates to false
  d == a
}
```

### Error joining

Errors can be joined together with the use of `errors.Join`.

```go
package users

var ErrUserNotFound = errors.New("user not found")

func FindUser(name string) (*User, error) {
  user, err := db.FindUser(name)
  if err != nil {
    return nil, errors.Join(err, ErrUserNotFound)
  }

  return user, nil
}
```

This is the equivalent of wrapping errors twice.

```go
package users

var ErrUserNotFound = errors.New("user not found")

func FindUser(name string) (*User, error) {
  user, err := db.FindUser(name)
  if err != nil {
    return nil, fmt.Errorf("%w %w", err, ErrUserNotFound)
  }

  return user, nil
}
```

This allows you to couple your domain errors with library errors, allowing
you to compare against your own errors elsewhere in the application instead of
having to import the library's error in other packages.

### Exercises

Now please complete the following exercises.

- [Exercise 2](./exercise-2)

## Typing errors

### Custom error types

So far we've been mostly seeing errors created from strings.
But technically, anything can be an error.
See what the error type actually is in the standard library below.

```go
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
 Error() string
}
```

This means that you're able to define your own error types easily by
just implementing the `Error() string` method on it.

```go
package users

type UserNotFoundError struct {
  Name string
}

func (u *UserNotFoundError) Error() string {
  return "User " + u.Name + " not found"
}
```

A caller would then extract the error using the `errors.AsType` method.

```go
package main

func main () {
  err := users.FindUser("Joey")

  if err != nil {
    userNotFoundErr, ok := errors.AsType[*users.UserNotFoundError](err)
    if ok {
      fmt.Println("Unable to find user with the name " + userNotFoundErr.Name)
      return
    }
  }

  // [...]
}
```

This makes it trivially easy to extract data from an error, and remove the need
for any string comparisons.

### Is and Unwrap

### Exercises

Now please complete the following exercises.

- [Exercise 3](./exercise-3)
