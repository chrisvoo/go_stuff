package main

import (
	"fmt"
	"math/rand"
)

func main() {
	exercise1()
	exercise2()
	exercise3()
}

/* Write a for loop that puts 100 random numbers between 0 and 100 into an int slice. */
func exercise1() []int {
	numKeeper := make([]int, 100)

	for i := range numKeeper {
		// rand.Intn(101) generates 0 through 100 inclusive
		numKeeper[i] = rand.Intn(101)
	}

	fmt.Println(numKeeper)
	return numKeeper
}

func exercise2() {
	nums := exercise1()
	for _, v := range nums {
		byTwo := v % 2
		byThree := v % 3
		if byTwo == 0 {
			fmt.Println(v, "Two!")
			if byThree == 0 {
				fmt.Println(v, "Six!")
			}
		} else if byThree == 0 {
			fmt.Println(v, "Three!")
		} else {
			fmt.Println(v, "Never mind")
		}
	}
}

/* Start a new program. In main, declare an int variable called total. Write a for loop that uses a 
 variable named i to iterate from 0 (inclusive) to 10 (exclusive). The body of the for loop should be as follows:
 total := total + ifmt.Println(total)
 After the for loop, print out the value of total.
 What is printed out? What is the likely bug in this code? */
func exercise3() {
	var total int
	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Println(total)
	}
	fmt.Println(total)
	/* It print s 0, 1, 3, 6, 10, 15, 21, 28, 36, 45 and then 0. The bug is that the total variable is
	   being shadowed by the total variable declared in the for loop.
	   The total variable in the for loop is a new variable that is only accessible within the scope of the 
	   for loop. The total variable declared outside of the for loop is not being updated.
	   Use the = operator instead of := to update the total variable. */
}

func recap() {
	printMap()
	forRangeCopy()
	outerLoop()
	switchMe()
	exitLoopFromSwitch()
	goToExample()
}

func printMap() {
	fmt.Println("for-range with maps -----------------------")
	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	/*  Output may vary because the order of iteration over maps is not specified and can change from
	one iteration to the next. Example output:
	Loop 0
	b 2
	a 1
	c 3
	Loop 1
	a 1
	c 3
	b 2
	Loop 2
	c 3
	b 2
	a 1
	*/
	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}

	// always output maps with their keys in ascending sorted order
	fmt.Println(m)
}

/*
each time the for-range loop iterates over your compound type,

	it copies the value from the compound type to the value variable
*/
func forRangeCopy() {
	fmt.Println("for-range with slices, copy -----------------------")
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		v *= 2
	}
	fmt.Println(evenVals)
}

// labels can be used to break or continue an outer loop from within an inner loop
func outerLoop() {
	fmt.Println("Outer loop -----------------------")
	samples := []string{"hello", "apple_π!"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}
}

func switchMe() {
	fmt.Println("Switch statement -----------------------")
	words := []string{
		"a", "cow", "smile", "gopher",
		"octopus", "anthropologist",
	}

	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}
}

func exitLoopFromSwitch() {
	fmt.Println("Exit loop from switch -----------------------")
loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 2, 4, 6:
			fmt.Println(i, "is even")
		case 3:
			fmt.Println(i, "is divisible by 3 but not 2")
		case 7:
			fmt.Println("exit the loop!")
			break loop // necessary to break the outer loop, not just the switch statement
		default:
			fmt.Println(i, "is boring")
		}
	}
}

func goToExample() {
	fmt.Println("goto example -----------------------")
	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("do something when the loop completes normally")
done:
	fmt.Println("do complicated stuff no matter why we left the loop")
	fmt.Println(a)
}
