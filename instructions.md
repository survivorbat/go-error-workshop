# 💡 Instructions

These instructions guide you through the exercises and teach you the ins
and outs of errors in Go.
Every paragraph has its own exercise(s) attached to allow you to put
the knowledge into practice.

The code snippets serve as examples, but aren't intended to be copy-pasted directly.
For working code, please refer to the exercises.

## Error basics

### Creating errors

Take the following function.

```go
package users

import (
  "fmt"
  "errors"
  "strings"
)

func CreateUser(name string) error {
  if name == "" {
    return errors.New("name cannot be empty")
  }

  if strings.Contains(name, "root") {
    return fmt.Errorf("name %s cannot contain root", name)
  }

  // [...]

  return nil
}
```

Here you see two functions used to create errors.
To check what error was returned, a caller has to read the `.Error()` message.

```go
package main

import (
  "fmt"
  "strings"

  "my-module/users"
)

func main() {
  name := "roots"

  err := users.CreateUser(name)
  if err != nil {
    msg := err.Error()

    if msg == "name cannot be empty" {
      fmt.Println("Please ensure a name is set")
    }

    if strings.Contains(msg, "cannot contain root") {
      fmt.Println("Due to security reasons, your name cannot contain the word root")
    }
  }
}
```

These string comparisons are the only way to accomplish this.
It is not possible for the caller to compare `err` to another error object.

```go
err := users.CreateUser(name)

// Does not work
err == errors.New("name cannot be empty")

// Does not work
errors.Is(err, errors.New("name cannot be empty"))
```

That's because `errors.New` and `fmt.Errorf` don't create strings, they create
distinct error values.
Under the hood, an unexported `*errorString` is returned, and the pointers
are not equal when compared against each other.

### Exercises

Now please complete the following exercises.

- [Exercise 1 (10 mins)](./exercise-1)

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

var ErrNameEmpty = errors.New("name cannot be empty")
var ErrRootNotAllowed = errors.New("name cannot contain root")

func CreateUser(name string) error {
  if name == "" {
    return ErrNameEmpty
  }

  if strings.Contains(name, "root") {
    return ErrRootNotAllowed
  }

  // [...]

  return nil
}
```

Now, the returned error can be compared against the package-level variables.
The following comparisons are now possible.

```go
package main

import (
  "fmt"
  "errors"
  "my-module/users"
)
func main() {
  name := "roots"

  err := users.CreateUser(name)
  if err != nil {
    if errors.Is(err, users.ErrNameEmpty) {
      fmt.Println("Please ensure a name is set")
    }

    if errors.Is(err, users.ErrRootNotAllowed)  {
      fmt.Println("Due to security reasons, your name cannot contain the word root")
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
  return fmt.Errorf("%s contains root: %w", name, ErrRootNotAllowed)
}
```

This will expand the `.Error()` output to the following.

```text
roots contains root: name cannot contain root
```

This is immensely useful when debugging a large codebase.
Imagine the following code.

```go
func CreateUser(name string) error {
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
func CreateUser(name string) error {
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

import (
  "fmt"
  "errors"
)

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

- [Exercise 2 (5 mins)](./exercise-2)

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

import (
  "fmt"
  "errors"

  "my-module/users"
)

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

Custom errors don't work with `errors.Is` directly, but you can implement the `Is`
method on it to decide your own equality.

```go
package users

import "errors"

type UserNotFoundError struct {
  Name string
}

func (u *UserNotFoundError) Error() string {
  return "User " + u.Name + " not found"
}

func (u *UserNotFoundError) Is(err error) bool {
  userNotFoundErr, isType := errors.AsType[*UserNotFoundError](err)
  if !isType {
    return false
  }

  return u.Name == userNotFoundErr.Name
}
```

For unwrapping, the `Unwrap` method can be implemented.

```go
package users

type UserNotFoundError struct {
  Name string

  InnerError error
}

func (u *UserNotFoundError) Error() string {
  return "User " + u.Name + " not found"
}

func (u *UserNotFoundError) Unwrap() error {
  return u.InnerError
}
```

Which allows `errors.Is` and `errors.AsType` to look for wrapped errors in your
custom error type.

### Exercises

Now please complete the following exercises.

- [Exercise 3 (10 mins)](./exercise-3)
- [Exercise 4 (30 mins)](./exercise-4)

## Testing errors

### Asserting sentinel errors

The [stretchr/testify](https://github.com/stretchr/testify) library has many assertion
functions available, specifically for errors.
Take for example the following.

```go
package users

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestCreateUser_IsAllowedWithANormalName(t *testing.T) {
  t.Parallel()
  // Arrange
  name := "survivorbat"

  // Act
  err := CreateUser(name)

  // Assert
  assert.NoError(t, err)
}

func TestCreateUser_ReturnsErrorOnRootName(t *testing.T) {
  t.Parallel()
  // Arrange
  name := "The-root-of-the-problem"

  // Act
  err := CreateUser(name)

  // Assert
  assert.Error(t, err) // ❌ Weak assertion
}
```

This function checks whether no error was returned on a "normal" username, and
an error occurred if the caller tries to create a root-like user.
Except, it doesn't catch a bug that was accidentally introduced in the `CreateUser` function.

```go
package users

import (
  "fmt"
  "strings"
)

func CreateUser(name string) error {
  if name == "" {
    return ErrRootNotAllowed // 🪲 BUG
  }

  if strings.Contains(name, "root") {
    return fmt.Errorf("%s contains admin: %w", name, ErrNameEmpty) // 🪲 BUG
  }

  // [...]

  return nil
}
```

In this scenario, the tests don't catch that:

- The wrong sentinel error is returned
- The message mentions the wrong username

This test could be made more robust by using more specific error assertions.

```go
// Asserts the actual error and its message
assert.ErrorIs(t, err, ErrNameEmpty)
assert.ErrorContains(t, err, "contains root")
```

### Asserting error types

Like `assert.ErrorIs`, there is also an `assert.ErrorAs` assertion available.

```go
package users

import (
  "testing"

  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/require"
)

func TestFindUser_ReturnsErrorOnNotFound(t *testing.T) {
  t.Parallel()
  // Arrange
  name := "survivorbat"

  // Act
  err := FindUser(name)

  // Assert
  var actualErr *UserNotFoundError
  require.ErrorAs(t, err, &actualErr)

  assert.Equal(t, name, actualErr.Name)
}
```

Notice the use of the `require` package here.
This ensures that the test exits if `err` is not a `*UserNotFoundError`, otherwise
the `actualErr.Name` would cause a nil pointer panic.

### Exercises

Now please complete the following exercises.

- [Exercise 5 (10 mins)](./exercise-5)

## Conclusion

There is more to errors than a simple call to `errors.New`.
