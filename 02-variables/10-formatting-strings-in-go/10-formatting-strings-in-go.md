# Formatting Strings in Go

Go follows the printf tradition from the C language. In my opinion, string formatting/interpolation in Go is currently less elegant than JavaScript and Python.
https://cplusplus.com/reference/cstdio/printf/

- `fmt.Printf` - Prints a formatted string to standard out.
  https://pkg.go.dev/fmt#Printf

- `fmt.Sprintf()` - Returns the formatted string
  https://pkg.go.dev/fmt#Sprintf

## Examples

These formatting verbs work with both fmt.Printf and fmt.Sprintf.

### Interpolate a string

```go
s := fmt.Sprintf("I am %s years old", "way too many")
// I am way too many years old
```

### Interpolate an integer in decimal form

```go
s := fmt.Sprintf("I am %d years old", 10)
// I am 10 years old
```

### Interpolate a decimal

```go
s := fmt.Sprintf("I am %f years old", 10.523)
// I am 10.523000 years old
```

```go
// The ".2" rounds the number to 2 decimal places
s := fmt.Sprintf("I am %.2f years old", 10.523)
// I am 10.52 years old
```

If you're interested in all the formatting options, feel free to take a look at the fmt package's docs here.

## Assignment

Create a new variable called `msg` on line 9. It's a string that contains the following:

```
Hi NAME, your open rate is OPENRATE percentNEWLINE
```

Where `NAME` is the given `name`, `OPENRATE` is the `openRate` rounded to the nearest "tenths" place, and `NEWLINE` is the \n escape sequence.
https://en.wikipedia.org/wiki/Newline
