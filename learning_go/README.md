# Learning Go, 2nd Edition

- [Learning Go, 2nd Edition](#learning-go-2nd-edition)
  - [Built-in types and variables](#built-in-types-and-variables)
    - [String literals](#string-literals)
    - [Integers](#integers)
    - [Floating types](#floating-types)
    - [Conversions](#conversions)
    - [Variable declarations](#variable-declarations)
    - [Other considerations](#other-considerations)
  - [Composite types](#composite-types)
    - [Arrays](#arrays)
    - [Slices](#slices)
      - [Slicing Slices](#slicing-slices)
    - [Strings and Runes and Bytes](#strings-and-runes-and-bytes)
    - [Maps](#maps)
    - [Structs](#structs)
  - [Blocks, Shadows, and Control Structures](#blocks-shadows-and-control-structures)
    - [Blocks and shadows](#blocks-and-shadows)
    - [Control structures](#control-structures)
      - [if/else](#ifelse)
      - [for](#for)
      - [switch](#switch)
  - [Functions](#functions)


## Built-in types and variables

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

## Composite types

### Arrays

* all elements in the array must be of the type that’s specified.
* you cannot read or write past the end of an array or use a negative index.
* Go considers the size of the array to be part of the type of the array.
* the size of an array must be a number, can't be a variable
* you can’t use a type conversion to directly convert arrays of different sizes to identical types

```go
var x [3]int // zero values
var x = [3]int{10, 20, 30} // or var x = [...]int{10, 20, 30}
var x = [12]int{1, 5: 4, 6, 10: 100, 15} // sparse array, indices with nonzero values
```

### Slices

Data structure for holding a sequence of values where its length is not part of its type. It's zero value is `nil`.

```go
var x []int // nil slice
var x = []int{} // zero length and zero capacity
var x = []int{10, 20, 30}
var x = []int{1, 5: 4, 6, 10: 100, 15}
```

Slices are not comparable, the only thing you can compare a slice with using `==` is `nil`. There are functions to compare slices in the `slices` package (Equal, EqualFunc, etc).
The **length** of a slice is the number of consecutive memory locations that have been assigned a value (built-in `len` function). Every slice also has a **capacity**, which is the number of consecutive memory locations reserved and can be larger than the length (built-in `cap` function). If you try to add additional values when the length equals the capacity, the append function uses the Go runtime to allocate a new backing array for the slice with a larger capacity. The values in the original backing array are copied to the new one, the new values are added to the end of the new backing array, and the slice is updated to refer to the new backing array. Finally, the updated slice is returned.
`make` allows you to specify the type, length, and, optionally, the capacity

```go
x := make([]int, 5) // initial length
x := make([]int, 5, 10) // initial length and capacity
```

* minimize the number of times the slice needs to grow
* If it’s possible that the slice won’t need to grow at all, use a var declaration with no assigned value to create a nil slice
* If you have some starting values, or if a slice’s values aren’t going to change, then a slice literal is a good choice
* If you have a good idea of how large your slice needs to be, but don’t know what those values will be when you are writing the program, use make:
  * If you are using a slice as a **buffer**, then specify a nonzero length.
  * If you are sure you know the exact size you want, you can specify the length and index into the slice to set the values
  * otherwise zero length and a specified capacity

#### Slicing Slices

When you take a slice from a slice, you are not making a copy of the data. Instead, you now have two variables that are sharing memory. This means that changes to an element in a slice affect all slices that share that element.

```go
	base := []string{"a", "b", "c", "d"}
	copybase := base[:2]
	another := base[1:]
	yetanother := base[1:3]
	last := base[:]
	fmt.Println("base:", base) 				// prints [a, b, c, d]
	fmt.Println("copybase:", copybase)		// prints [a, b]
	fmt.Println("another:", another)		// prints [b, c, d]
	fmt.Println("yetanother:", yetanother)	// prints [b, c]
	fmt.Println("last:", last)				// prints [a, b, c, d]
```

Whenever you take a slice from another slice, the subslice’s capacity is set to the
capacity of the original slice, minus the starting offset of the subslice within the
original slice. To avoid complicated slice situations, **you should either never use append with a subslice or make sure that append doesn’t cause an overwrite by using a full slice expression**.
If you need to create a slice that’s independent of the original, use the built-in `copy` function. The capacity of x and y doesn’t matter; it’s the length that’s important.

```go
x := []int{1, 2, 3, 4}
y := make([]int, 4)
num := copy(y, x)
fmt.Println(y, num)
```

You can covert arrays (or a subset of array) to slices. Be aware that taking a slice from an array has the same memory-sharing properties as taking a slice from a slice.

```go
xArray := [4]int{5, 6, 7, 8}
xSlice := xArray[:]
```

and viceversa:

```go
xSlice := []int{1, 2, 3, 4}
xArray := [4]int(xSlice)
smallArray := [2]int(xSlice)
xSlice[0] = 10
fmt.Println(xSlice) // [10 2 3 4]
fmt.Println(xArray) // [1 2 3 4]
fmt.Println(smallArray) // [1, 2]
```

### Strings and Runes and Bytes

Go uses a sequence of bytes to represent a string

```go
var s string = "Hello there"
var b byte = s[6]

// slice
var s string = "Hello there"
var s2 string = s[4:7] // "o t" careful, with non-english languages you don't get this
```

Conversions:

```go
var a rune    = 'x'
var s string  = string(a)
var b byte    = 'y'
var s2 string = string(b)
```

### Maps

The map type is written as `map[keyType]valueType`. The zero value for a map is nil with lenght 0. Attempting to write to a nil map variable causes a panic.

```go
totalWins := map[string]int{} // empty map literal is writable

teams := map[string][]string {
    "Orcas": []string{"Fred", "Ralph", "Bijou"},
    "Lions": []string{"Sarah", "Peter", "Billie"},
    "Kittens": []string{"Waldo", "Raul", "Ze"},
}

ages := make(map[int][]string, 10)
```

* Maps automatically grow as you add key-value pairs to them.
* If you know how many key-value pairs you plan to insert into a map, you can use make to create a map with a specific initial size.
* Passing a map to the len function tells you the number of key-value pairs in a map.
* The zero value for a map is nil.
* Maps are not comparable. You can check if they are equal to nil, but you cannot check if two maps have identical keys and values using == or differ using !=.
* The key for a map can be any comparable type. This means you cannot use a slice or a map as the key for a map.

### Structs

```go
type person struct {
    name string
    age  int
    pet  string
}

var fred person
bob := person{}
julia := person{
    "Julia",
    40,
    "cat",
}
beth := person{
    age:  30,
    name: "Beth",
}
bob.name = "Bob"
fmt.Println(bob.name)

// anonymous struct
pet := struct {
    name string
    kind string
}{
    name: "Fido",
    kind: "dog",
}
```

* A struct type that’s defined within a function can be used only within that function.
* A zero value struct has every field set to the field’s zero value.
* Structs that are entirely composed of comparable types are comparable; those with slice or map fields are not
* if two struct variables are being compared and at least one has a type that’s an anonymous struct, you can compare them without a type conversion if the fields of both structs have the same names, order, and types.

You might wonder when it’s useful to have a data type that’s associated only with a single instance. Anonymous structs are handy in two common situations. The first is when you translate external data into a struct or a struct into external data (like JSON or Protocol Buffers). This is called **unmarshaling** and **marshaling** data, respectively. Writing tests is another place where anonymous structs pop up.

## Blocks, Shadows, and Control Structures

### Blocks and shadows

Each place where a declaration occurs is called a **block**. Variables, constants, types, and functions declared outside of any functions are placed in the **package block**. Import statements are in the **file block**. built-in types, constants and function (like `make`) are in the **universe block**.
When you have a declaration with the same name as an identifier in a containing block, you **shadow** the identifier created in the outer block.

```go
/* 10
   5
   10 */
func main() {
    x := 10
    if x > 5 {
        fmt.Println(x)
        x := 5 // shadowing x
        fmt.Println(x)
    }
    fmt.Println(x)
}
```

A **shadowing variable** is a variable that has the same name as a variable in a containing block. For as long as the shadowing variable exists, you cannot access a shadowed variable.
Same thing can happen to packages if you declare a variable named like a package, and also to universe block's elements.

### Control structures

#### if/else

```go
// It lets you create variables available only where needed.
if n := rand.Intn(10); n == 0 {
    fmt.Println("That's too low")
} else if n > 5 {
    fmt.Println("That's too big:", n)
} else {
    fmt.Println("That's a good number:", n)
}
// n here is undefined
```

#### for

`for` is the only looping keyword in the language and can be used in four format. Favor a for-range loop when iterating over all the contents of an instance of one of the built-in compound types. It avoids a great deal of boilerplate code that’s required when you use an array, slice, or map with one of the other for loop styles.

```go
// Complete C-style
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// Complete C-style wiht initialization already done outside
i := 0
for ; i < 10; i++ {
    fmt.Println(i)
}

// Condition-Only for Statement: it's like a while
i := 1
for i < 100 {
  fmt.Println(i)
  i = i * 2
}

// infinite and break/continue
for {
    // things to do in the loop
    if !CONDITION {
        break
    }
}

for i := 1; i <= 100; i++ {
    if i%3 == 0 && i%5 == 0 {
        fmt.Println("FizzBuzz")
        continue
    }
    if i%3 == 0 {
        fmt.Println("Fizz")
        continue
    }
    if i%5 == 0 {
        fmt.Println("Buzz")
        continue
    }
    fmt.Println(i)
}

// for-range Statement
// only to iterate over the built-in compound types and
// user-defined types that are based on them
evenVals := []int{2, 4, 6, 8, 10, 12}
// If you don’t need to access the key, use an underscore (_)
for _, v := range evenVals {
    fmt.Println(v)
}

// or just the keys
uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
for k := range uniqueNames {
    fmt.Println(k)
}
```

#### switch

break statements aren't needed in Go for case statements, however you can use them if you want to exit prematurely, even if may signals you're doing something too complicated.

```go
for _, word := range words {
    switch size := len(word); size {
    case 1, 2, 3, 4:
        fmt.Println(word, "is a short word!")
    case 5:
        wordLen := len(word)
        fmt.Println(word, "is exactly the right length:", wordLen)
    case 6, 7, 8, 9: // nothing happens!
    default:
        fmt.Println(word, "is a long word!")
    }
}
```

## Functions
