package main

import "fmt"

func main() {
	/* Expected output:
	First: 20 20
	Second: 20 20
	Third: 0 -2147483648 0
	*/
	first()
	second()
	third()
}

/*
Write a program that declares an integer variable called i with the value 20.

	Assign i to a floating-point variable named f. Print out i and f.
*/
func first() {
	i := 20
	var j float32 = float32(i)
	fmt.Println("First:", i, j)
}

/*
Write a program that declares a constant called value that can be assigned to both an integer

	and a floating-point variable.  Assign it to an integer called i and a floating-point variable
	called f. Print out i and f.
*/
func second() {
	const value = 20
	i := value
	var f float32 = value
	fmt.Println("Second:", i, f)
}

/*
Write a program with three variables, one named b of type byte, one named smallI of type int32,
and one named bigI of type uint64. Assign each variable the maximum legal value for its type;
then add 1 to each variable. Print out their values.
*/
func third() {
	b := byte(255)
	smallI := int32(2147483647)
	bigI := uint64(18446744073709551615)
	b++
	smallI++
	bigI++
	fmt.Println("Third:", b, smallI, bigI)

	/*
		Setup — max values for each type:

		┌──────────┬────────────────┬────────────────────────────┬───────────────────────┐
		│ Variable │      Type      │         Max value          │ Binary representation │
		├──────────┼────────────────┼────────────────────────────┼───────────────────────┤
		│ b        │ byte (= uint8) │ 255                        │ 11111111              │
		├──────────┼────────────────┼────────────────────────────┼───────────────────────┤
		│ smallI   │ int32          │ 2,147,483,647              │ 0111...1 (31 ones)    │
		├──────────┼────────────────┼────────────────────────────┼───────────────────────┤
		│ bigI     │ uint64         │ 18,446,744,073,709,551,615 │ 1111...1 (64 ones)    │
		└──────────┴────────────────┴────────────────────────────┴───────────────────────┘

		After ++ on each:

		- b (255 → 0): All 8 bits are 1. Adding 1 produces a carry that falls off the end — result is 00000000.
		- smallI (2147483647 → -2147483648): In two's complement, the max positive int32 is 0111...1. Adding 1 flips it to 1000...0, which is the most negative value (sign bit becomes 1).
		- bigI (max uint64 → 0): Same as b — all 64 bits are 1, adding 1 drops all carries, leaving 64 zeros.
	*/
}
