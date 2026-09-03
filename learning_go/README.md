# Learning Go, 2nd Edition

## Types and variables

**Predeclared types** are the built-in types in the language. A Go **literal** is an explicitly specified number, character, or string.

### String literals

A **rune** literal represents a character and is surrounded by single quotes.
An **interpreted** string literal is a string literal with double quotes contain zero or more rune literals. If you need to include backslashes, double quotes, or newlines in your string, using a **raw string literal** is easier. These are delimited with backquotes (`) and can contain any character except a backquote.
The **zero value** for a string is the empty string. Strings in Go are **immutable**; you can reassign the value of a string variable, but you cannot change the value of the string that is assigned to it.

### Integers

Type integer (signed/unsigned) from one to eight bytes (int8 - uint64).   There are some aliases (listing a subset here):
* byte for uint8
* int for int32/int64 (literal default)

> [!NOTE] Which integer to use
> * If you are working with a binary file format or network protocol that has an integer of a specific size or sign, use the corresponding integer type.
> * If you are writing a library function that should work with any integer
> type, take advantage of Go’s generics support and use a generic type
> * In all other cases, just use int

### Floating types

There's float32 and float64. Like the integer types, the zero value for the floating-point types is 0.
Unless you have to be compatible with an existing format, use float64. Floating-point literals have a default type of float64, so always using float64 is the simplest option.

> [!Warning]
> Do not use them to represent money or any other value that must have an exact decimal representation. Don't also use `!=` and `==` for comparing floats

### Conversions

Go doesn’t allow automatic type promotion between variables. You must use a type conversion when variable types do not match. Even different-sized integers and floats must be converted to the same type to interact.

```go
var x int = 10
var y float64 = 30.2
var sum1 float64 = float64(x) + y
var sum2 int = x + int(y)
fmt.Println(sum1, sum2) // 40.2 40
```

You cannot treat another Go type as a boolean, **Go doesn’t allow truthiness**.  If you want to convert from another data type to boolean, you must use one of the comparison operators

### Variable declarations

```go
var x, y int = 10, 20
var x int = 10
var x = 10 // int
var x, y = 10, "hello"
var x int // zero value
// declaration list
var (
    x    int
    y        = 20
    z    int = 30
    d, e     = 40, "hello"
    f, g string
)
```

When you are within a function, you can use the := operator to replace a var declaration that uses type inference.
If you are declaring a variable at the package level, you must use var because `:=` is not legal outside of functions

```go
x := 10
x, y := 30, "hello"
```

In general:
* When initializing a variable to its zero value, use var x int to make it clear that the zero value is intended.
* When assigning an untyped constant or a literal to a variable and the default type for the constant or literal isn’t the type you want for the variable, use the long var form with the type specified.
* declare all your new variables with `var` to make it clear which variables are new, and then use the assignment operator (=) to assign values to both new and old variables.
* you should only declare variables in the package block that are effectively immutable

### Other considerations

* every declared local variable must be read. It is a compile-time error to declare a local variable and to not read its value.
* Go requires identifier names to start with a letter or underscore, and the name can contain numbers, underscores, and letters
* The smaller the scope for a variable, the shorter the name that’s used for it.