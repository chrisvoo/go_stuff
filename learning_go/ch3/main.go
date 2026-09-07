package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	exercise1()
	exercise2()
	exercise3()
}

/*
Write a program that defines a variable named greetings of type slice of strings with the

	following values: "Hello", "Hola", "नमस्कार", "こんにちは", and "Привіт".
	Create a subslice containing the first two values;
	a second subslice with the second, third, and fourth values;
	and a third subslice with the fourth and fifth values.
	Print out all four slices.
*/
func exercise1() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	sub1 := greetings[:2]
	sub2 := greetings[1:4]
	sub3 := greetings[3:]

	/*
		greetings: [Hello Hola नमस्कार こんにちは Привіт]
		sub1: [Hello Hola]
		sub2: [Hola नमस्कार こんにちは]
		sub3: [こんにちは Привіт]
	*/

	fmt.Println("greetings:", greetings)
	fmt.Println("sub1:", sub1)
	fmt.Println("sub2:", sub2)
	fmt.Println("sub3:", sub3)
}

/*
Write a program that defines a string variable called message with the value "Hi  and " and
prints the fourth rune in it as a character, not a number.
*/
func exercise2() {
	/* Fourth rune in message: 👨 */
	message := "Hi 👨 and 🧍‍♂️"
	fmt.Printf("Fourth rune in message: %c\n", []rune(message)[3])
}

/* Write a program that defines a struct called Employee with three fields: firstName, lastName, and id. 
   The first two fields are of type string, and the last field (id) is of type int. Create three instances of this struct 
   using whatever values you’d like. Initialize the first one using the struct literal style without names, the second 
   using the struct literal style with names, and the third with a var declaration. Use dot notation to populate the 
   fields in the third struct. Print out all three structs.*/
func exercise3() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	emp1 := Employee{"John", "Doe", 1}
	emp2 := Employee{firstName: "Jane", lastName: "Smith", id: 2}
	var emp3 Employee
	emp3.firstName = "Alice"
	emp3.lastName = "Johnson"
	emp3.id = 3

	/*
	{John Doe 1}
	{Jane Smith 2}
	{Alice Johnson 3}
	*/

	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(emp3)
}

func recap() {
	array()
	slicesFunc()
	slicingSlices()
	shareMemory()
	confusing()
	fullSliceExpression()
	copySlice()
	toArray()
	mapsFunc()
	setWorkaround()
}

func array() {
	fmt.Println("Array ----------------------------")
	var x = [...]int{1, 2, 3}
	var y = [3]int{1, 2, 3}
	fmt.Println(x == y) // prints true

	x[0] = 10
	fmt.Println(x[2])   // prints 3
	fmt.Println(len(x)) // prints 3
}

func slicesFunc() {
	fmt.Println("Slices ----------------------------")
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}
	c := []int{1, 2, 3, 4, 5, 6}
	// d := []string{"a", "b", "c"}
	fmt.Println(slices.Equal(a, b)) // prints true
	fmt.Println(slices.Equal(a, c)) // prints false
	// fmt.Println(slices.Equal(a, d)) // does not compile, in call to slices.Equal, type []string of d does not match inferred type []int for

	var e []int
	fmt.Println(e, len(e), cap(e)) // prints [] 0 0
	e = append(e, 10, 11, 12, 13)  // assign result to the variable that's passed in
	fmt.Println(e)                 // prints [10, 11, 12, 13]
	fmt.Println(e, len(e), cap(e))

	f := []int{20, 30, 40}
	e = append(e, f...)
	fmt.Println(e)                 // prints [10, 11, 12, 13, 20, 30, 40]
	clear(e)                       // clears the slice, but does not free the underlying array
	fmt.Println(e)                 // prints [0, 0, 0, 0, 0, 0, 0]
	fmt.Println(e, len(e), cap(e)) // prints [0, 0, 0, 0, 0, 0, 0] 7 7
}

/*
Whenever you take a slice from another slice, the subslice’s capacity is set to the
capacity of the original slice, minus the starting offset of the subslice within the
original slice.
*/
func slicingSlices() {
	fmt.Println("Slicing Slices ----------------------------")
	base := []string{"a", "b", "c", "d"}
	copybase := base[:2]
	another := base[1:]
	yetanother := base[1:3]
	last := base[:]
	fmt.Println("base:", base)             // prints [a, b, c, d]
	fmt.Println("copybase:", copybase)     // prints [a, b]
	fmt.Println("another:", another)       // prints [b, c, d]
	fmt.Println("yetanother:", yetanother) // prints [b, c]
	fmt.Println("last:", last)             // prints [a, b, c, d]
}

func shareMemory() {
	fmt.Println("Sharing memory ----------------------------")
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	fmt.Println(cap(x), cap(y)) // prints 4 4
	y = append(y, "z")
	fmt.Println("x:", x) // prints [a, b, z, d]
	fmt.Println("y:", y) // prints [a, b, z]

}

func confusing() {
	fmt.Println("Confusing ----------------------------")
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")   // [a, b, c, d]
	y := x[:2]                          // [a, b]
	z := x[2:]                          // [c, d]
	fmt.Println(cap(x), cap(y), cap(z)) // prints 5 5 3

	y = append(y, "i", "j", "k")
	fmt.Println("x 1:", x) // prints [a b i j]
	fmt.Println("y 1:", y) // prints [a b i j k]
	fmt.Println("z 1:", z) // prints [i j]

	x = append(x, "x")
	fmt.Println("x 2:", x) // prints [a b i j x]
	fmt.Println("y 2:", y) // prints [a b i j x]
	fmt.Println("z 2:", z) // prints [i j]

	z = append(z, "y")
	fmt.Println("x 3:", x) // prints [a b i j y]
	fmt.Println("y 3:", y) // prints [a b i j y]
	fmt.Println("z 3:", z) // prints [i j y]
}

/*
Because you limited the capacity of the subslices to their lengths, appending additional

	elements onto y and z created new slices that didn’t interact with the other slices.
*/
func fullSliceExpression() {
	fmt.Println("Full Slice Expression ----------------------------")
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2:2]
	z := x[2:4:4]
	fmt.Println("x:", x) // prints [a, b, c, d], cap 5
	fmt.Println("y:", y) // prints [a, b], cap 2
	fmt.Println("z:", z) // prints [c, d], cap 2
	y = append(y, "i", "j", "k")
	fmt.Println("x:", x) // prints [a, b, c, d]
	fmt.Println("y:", y) // prints [a, b, i, j, k]
	fmt.Println("z:", z) // prints [c, d]
}

func copySlice() {
	fmt.Println("Copy Slice ----------------------------")
	x := []int{1, 2, 3, 4}
	y := make([]int, 2)
	num := copy(y, x)
	fmt.Println(y, num) // prints [1, 2] 2

	// overlapping
	x = []int{1, 2, 3, 4}
	num = copy(x[:3], x[1:])
	fmt.Println(x, num) // prints [2, 3, 4, 4] 3
}

func toArray() {
	fmt.Println("Convert Slice to Array ----------------------------")
	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice)
	smallArray := [2]int(xSlice)
	xSlice[0] = 10
	fmt.Println(xSlice)
	fmt.Println(xArray)
	fmt.Println(smallArray)
}

func mapsFunc() {
	fmt.Println("Maps ----------------------------")
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])
	fmt.Println(totalWins) // prints map[Kittens:1 Lions:3 Orcas:1]

	fmt.Println("Maps ok idiom ----------------------------")
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok) // prints 0 false

	delete(m, "hello")
	fmt.Println(m) // prints map[world:0]

	clear(m)
	fmt.Println(m) // prints map[]

	fmt.Println("Maps comparison ----------------------------")
	m = map[string]int{
		"hello": 5,
		"world": 10,
	}
	n := map[string]int{
		"world": 10,
		"hello": 5,
	}
	fmt.Println(maps.Equal(m, n))
}

/* A set is a data type that ensures there is at most one of a value */
func setWorkaround() {
	fmt.Println("Map as a set ----------------------------")
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}
}
