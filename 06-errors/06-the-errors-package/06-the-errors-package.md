# The Errors Package

The Go standard library provides an "errors" package that makes it easy to deal with errors.

Read the godoc for the errors.New() function https://pkg.go.dev/errors#New, but here's a simple example:

```go
var err error = errors.New("something went wrong")
```

## Assignment

Textio's software architects may have overcomplicated the requirements from the last coding assignment... oops. All we needed was a new generic error message that returns the string `no dividing by 0` when a user attempts to get us to perform the taboo.

Complete the `divide` function. Use the `errors.New()` function to return an error when `y == 0` that reads `"no dividing by 0"`.

## Hint

Remember you always need to return 2 values from the divide function! If there is an error you should be returning a `float64` along with the `errors.New()`
