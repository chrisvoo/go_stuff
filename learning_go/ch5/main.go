package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	bytes, err := fileLen("main.go")
	if err != nil {
		fmt.Printf("Got an error: %v\n", err)
		return
	}
	fmt.Printf("Total bytes: %d\n", bytes)

	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}

/*
The simple calculator program doesn’t handle one error case: division by zero.
Change the function signature for the math operations to return both an int and an
error. In the div function, if the divisor is 0, return errors.New("division by zero")
for the error. In all other cases, return nil. Adjust the main function to check for
this error.
*/
func exercise1() {
	calculator()
}

/*
Write a function called fileLen that has an input parameter of type string and
returns an int and an error. The function takes in a filename and returns the number
of bytes in the file. If there is an error reading the file, return the error.
Use defer to make sure the file is closed properly.
*/
func fileLen(filename string) (int, error) {
	info, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return int(info.Size()), nil
}

/*
Write a function called prefixer that has an input parameter of type string and
returns a function that has an input parameter of type string and returns a string.
The returned function should prefix its input with the string passed into prefixer.
Use the following main function to test prefixer:
*/
func prefixer(input string) func(i string) string {
	return func(i string) string {
		return input + " " + i
	}
}

/*
map[2:hello 3:goodbye]
[2 4 6]
*/
func showMods() {
	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	fmt.Println(m)

	s := []int{1, 2, 3}
	modSlice(s)
	fmt.Println(s)
}

func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func recap() {
	useDivAndRemainder()
	calculator()
	closure()
	cat()
	showMods()
}

func cat() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}

/*
20
30
20
*/
func closure() {
	a := 20
	f := func() {
		fmt.Println(a)
		/* Using := instead of = inside the closure creates a new a that ceases to exist
		   when the closure exits. With =, the closure would modify the outer a */
		a := 30
		fmt.Println(a)
	}
	f()
	fmt.Println(a)
}

func useDivAndRemainder() {
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
}

func divAndRemainder(num, denom float64) (float64, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, int(num) % int(denom), nil
}

func add(i int, j int) (int, error) { return i + j, nil }

func sub(i int, j int) (int, error) { return i - j, nil }

func mul(i int, j int) (int, error) { return i * j, nil }

func div(i int, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return i / j, nil
}

func calculator() {
	type opFuncType func(int, int) (int, error)

	var opMap = map[string]opFuncType{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "/", "0"},
		{"2", "%", "2"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
}
