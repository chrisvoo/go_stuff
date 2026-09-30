package main

import (
	"fmt"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	start := time.Now()
	exercise3Append()
	fmt.Println("Time taken:", time.Since(start))
}

/*
Create a struct named Person with three fields: FirstName and LastName of type string
and Age of type int.
Write a function called MakePerson that takes in firstName, lastName, and age and
returns a Person.
Write a second function MakePersonPointer that takes in firstName, lastName, and age and
returns a *Person.
Call both from main. Compile your program with go build -gcflags="-m".
This both compiles your code and prints out which values escape to the heap.
Are you surprised about what escapes?

The -gcflags="-m" flag reveals two major optimizations the Go compiler performs on your code: inlining and escape analysis.
To optimize performance, the Go compiler eliminates function call overhead by replacing the
function calls in main with the actual code from inside those functions.
Because MakePerson and MakePersonPointer are very short, the compiler flags that they "can inline",
and then confirms it actually did inline them inside main.
Despite the scary word "leaking," this is not a memory leak. In compiler terminology, this simply means
the data flow "leaks" from the input parameters directly into the return value. For MakePerson, the
firstName and lastName strings are passed in and placed unmodified into the returned struct (~r0 represents
the first return value).
With the other function, It realizes that p must outlive the function that created it, so it safely and
automatically moves p to the heap. The heap is a different area of memory managed by Go's garbage
collector, allowing the struct to persist until nothing is pointing to it anymore.
*/

func exercise1() {
	/*
		# ch6
		./main.go:25:6: can inline MakePerson
		./main.go:33:6: can inline MakePersonPointer
		./main.go:9:6: can inline main
		./main.go:10:12: inlining call to MakePerson
		./main.go:11:19: inlining call to MakePersonPointer
		./main.go:25:17: leaking param: firstName to result ~r0 level=0
		./main.go:25:35: leaking param: lastName to result ~r0 level=0
		./main.go:33:24: leaking param: firstName
		./main.go:33:42: leaking param: lastName
		./main.go:34:2: moved to heap: p */
	MakePerson("Chris", "Castles", 45)
	MakePersonPointer("ChrisP", "CastlesP", 46)
}

func MakePerson(firstName string, lastName string, age int) Person {
	/* The compiler builds the struct on the stack, copies that value back to the caller (main), and deletes the original. No heap allocation is needed. */
	return Person{
		firstName,
		lastName,
		age,
	}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	/* If p stayed on the stack, that memory would be destroyed when MakePersonPointer finished, leaving the caller with a pointer to dead memory. */
	p := Person{
		firstName,
		lastName,
		age,
	}

	return &p
}

/*
Write two functions. The UpdateSlice function takes in a []string and a string. It sets the last position
in the passed-in slice to the passed-in string. At the end of UpdateSlice, print the slice after making
the change. The GrowSlice function also takes in a []string and a string. It appends the string onto
the slice. At the end of GrowSlice, print the slice after making the change. Call these functions
from main. Print out the slice before each function is called and after each function is called.
Do you understand why some changes are visible in main and why some changes are not?
*/
func exercise2() {
	/*
		Before UpdateSlice:  [one two]
		[one three]
		After UpdateSlice:  [one three]
		Before GrowSlice:  [one three]
		[one three four]
		After GrowSlice:  [one three]
	*/
	mySlice := []string{"one", "two"}

	fmt.Println("Before UpdateSlice: ", mySlice)
	UpdateSlice(mySlice, "three")
	fmt.Println("After UpdateSlice: ", mySlice)

	fmt.Println("Before GrowSlice: ", mySlice)
	GrowSlice(mySlice, "four")
	fmt.Println("After GrowSlice: ", mySlice)
}

/*
When you run slice[len(slice)-1] = param, you are telling Go to follow the pointer in the copied
slice header down to the shared backing array and overwrite the data at that index. Because the slice
in main is looking at that exact same backing array, and its length hasn't changed, main instantly sees
the modified data.
*/
func UpdateSlice(slice []string, param string) {
	slice[len(slice)-1] = param
	fmt.Println(slice)
}

/*
When you run append(slice, param), the append function evaluates the underlying array. One of two
things happens:
- If there is enough capacity, it adds the new element to the existing backing array.
- If there is not enough capacity, it allocates a brand-new backing array, copies the old data over, and adds the new element.
In either case, append always returns a brand-new slice header with an updated Length (and potentially a
new Pointer and Capacity).
*/
func GrowSlice(slice []string, param string) {
	slice = append(slice, param)
	fmt.Println(slice)
}

/*
Write a program that builds a []Person with 10,000,000 entries (they could all be the same names and ages).
See how long it takes to run. Change the value of GOGC and see how that affects the time it takes
for the program to complete.
Set the environment variable GODEBUG=gctrace=1 to see when garbage collections happen and see how
changing GOGC changes the number of garbage collections. What happens if you create the slice with a
capacity of 10,000,000?
*/
func exercise3Append() {
	/*
		Because you start with a capacity of 1,000, the append function quickly runs out of room. Go has
		to allocate a new, larger backing array, copy the data over, and abandon the old array. To reach
		10,000,000 elements, Go must reallocate and abandon dozens of increasingly massive backing arrays.
		All of those abandoned arrays become "garbage", which forces the garbage collector to run repeatedly
		to free up memory.
	*/
	persons := make([]Person, 1000)
	for range 10000000 {
		persons = append(persons, Person{
			firstName: "Ciccio",
			lastName:  "Pasticcio",
			age:       45,
		})
	}

	fmt.Println(persons[10000000-1])
}

func exercise3BigLoop() {
	/*
		By allocating the exact capacity up front, Go provisions one single backing array. You iterate through
		 and populate it directly via index notation. Zero arrays are abandoned, zero garbage is created, and the garbage collector will barely run at all. This version will execute significantly faster.
	*/
	persons := make([]Person, 10000000)
	for i := 0; i < 10000000; i++ {
		persons[i] = Person{
			firstName: "Ciccio",
			lastName:  "Pasticcio",
			age:       45,
		}
	}

	fmt.Println(persons[10000000-1])
}
