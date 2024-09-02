package main

// By default, the package reference is assigned the name of the package
import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
)

// this line disables linting for all the following statements. Just replace disable with enable to
// re-enable the rule for the rest of the file
// revive:disable:exported
func PrintNumber(number int) {
	fmt.Println(number)
}

func main() {

	fmt.Println("\nVariables -------------------")
	const price, tax float32 = 275, 27.50
	const quantity, inStock = 2, true // The untyped constant feature allows automatic conversion
	// Untyped constants will be converted only if the value can be represented in the target type
	fmt.Println("Total:", quantity*(price+tax))
	fmt.Println("In stock: ", inStock)

	fmt.Println("\nPointers -------------------")
	// in Go, pointers are not just memory addresses but, rather, memory addresses that may store a specific
	// type of value. A runtime error will occur if you follow a pointer that has not been assigned a value
	first := 100
	second := first
	first++
	fmt.Println("First:", first) // copies the current value of first
	fmt.Println("Second:", second)

	// address op
	var third *int = &first // third is a pointer to a mem. location
	first++
	fmt.Println("First:", first)
	/*
		The asterisk tells Go to follow the pointer and get the value at the memory location.
		This is known as dereferencing the pointer.
	*/
	fmt.Println("Second:", *third)

	// pointing at pointers
	one := 100
	two := &one
	three := &two
	fmt.Println(one)
	fmt.Println(*two)
	fmt.Println(**three) // following chains of pointers

	// pointing to an element of a list. Without the pointer, we'll be copying a value and the output
	// would be different.
	names := [3]string{"Alice", "Charlie", "Bob"}
	secondPosition := &names[1]
	fmt.Println(*secondPosition)
	sort.Strings(names[:])
	fmt.Println(*secondPosition)

	fmt.Println("\nOperations -------------------")
	var floatVal = math.MaxFloat64
	fmt.Println("Infinity?", math.IsInf((floatVal*2), 0))
	a, b := 275.00, 27.40
	sum := a + b
	fmt.Println("Sum:", sum)

	negResult := -3 % 2
	absResult := math.Abs(float64(negResult))
	negResult++
	fmt.Println(negResult, absResult)

	// Explicit conversions can be used only when the value can be represented in the target type.
	// Care must be taken when choosing the values to convert because explicit conversions can cause a
	// loss of precision in numeric values or cause overflows
	kayak := 275
	soccerBall := 19.50
	total := float64(kayak) + soccerBall
	fmt.Println(total)

	val1 := "true"
	val2 := "false"
	val3 := "not true"
	bool1, b1err := strconv.ParseBool(val1)
	bool2, b2err := strconv.ParseBool(val2)
	fmt.Println("Bool 1", bool1, b1err)
	fmt.Println("Bool 2", bool2, b2err)

	// if statements can define an initialization statement, and this allows a conversion function
	// to be called and its results to be inspected in a single statement
	if bool3, b3err := strconv.ParseBool(val3); b3err == nil {
		fmt.Println("Parsed value:", bool3)
	} else {
		fmt.Println("Cannot parse", val3)
	}

	// Specifying a size of 8 when calling the ParseInt function allows me to perform an explicit
	// conversion to the int8 type without the possibility of overflow.
	val4 := "100"
	int4, int1err := strconv.ParseInt(val4, 0, 8)
	if int1err == nil {
		smallInt := int8(int4)
		fmt.Println("Parsed value:", smallInt)
	} else {
		fmt.Println("Cannot parse", val4, int1err)
	}

	// Atoi function handles the parsing and explicit conversion in a single step
	val5 := "100"
	int5, int5err := strconv.Atoi(val5)
	if int5err == nil {
		var intResult int = int5
		fmt.Println("Parsed value: " + strconv.FormatInt(int64(intResult), 10))
	} else {
		fmt.Println("Cannot parse", val5, int1err)
	}

	val := 275
	base10String := strconv.Itoa(val)
	fmt.Println("Base 10: " + base10String)

	val6 := 49.95
	Fstring := strconv.FormatFloat(val6, 'f', 2, 64)  // ±ddd.ddd without an exponent
	Estring := strconv.FormatFloat(val6, 'e', -1, 64) // ±ddd.ddde±dd
	fmt.Println("Format F: " + Fstring)
	fmt.Println("Format E: " + Estring)

	fmt.Println("\nFlow control -------------------")
	/* Each clause in an if statement has its own scope, which means that variables can be accessed only within the
	   clause in which they are defined */
	kayakPrice := 275.00
	if kayakPrice > 500 {
		scopedVar := 500
		fmt.Println("Price is greater than", scopedVar)
	} else if kayakPrice < 100 {
		scopedVar := "Price is less than 100"
		fmt.Println(scopedVar)
	} else {
		scopedVar := false
		fmt.Println("Matched: ", scopedVar)
	}

	counter := 0
	for {
		fmt.Println("Counter:", counter)
		counter++
		if counter > 3 {
			break
		}
	}

	counter = 0 // similar to while loops in other languages
	for counter <= 3 {
		fmt.Println("Counter:", counter)
		counter++
	}

	// No types are specified in the shorthand syntax
	for i := 0; i < 5; i++ {
		PrintNumber(i)
		fmt.Println("Value:", rand.Int())
	}

	// do...while loops
	for counter := 0; true; counter++ {
		if counter == 1 {
			continue
		}

		fmt.Println("Counter:", counter)
		if counter > 3 {
			break
		}
	}

	// Enumerating Sequences
	product := "Kayak"
	for index, character := range product {
		// switch val := counter / 2; val {
		switch character {
		case 'K', 'k':
			if character == 'k' {
				fmt.Println("Lowercase k at position", index)
				break
				// use `fallthrough` to continue to the next statement
			}
			fmt.Println("Uppercase K at position", index)
		case 'y':
			fmt.Println("y at position", index)
		default:
			fmt.Println("Character", string(character), "at position", index)
		}
	}

	products := []string{"Kayak", "Lifejacket", "Soccer Ball"}
	for index, element := range products {
		fmt.Println("Index:", index, "Element:", element)
	}

	// labels
	theCounter := 0
target:
	fmt.Println("Counter", theCounter)
	theCounter++
	if theCounter < 5 {
		goto target
	}

	fmt.Println("\nArrays, Slices, and Maps -------------------")
	/*
		- to store a fixed number of values, se an array
		- to enumerate an array/slice/map, use a for loop with the range keyword
		- to store a variable number of values, use a slice
		- to compare slices. use the reflect package
		- to store key-value pairs, use a map
	*/

	// The length and element type of an array cannot be changed, and the array length must be specified as a constant
	var theNames [3]string // populated with empty strings
	theNames[0] = "Kayak"
	theNames[1] = "Lifejacket"
	theNames[2] = "Paddle"
	otherArray := &theNames
	theNames[0] = "Canoe"
	fmt.Println("names:", theNames)
	for _, value := range *otherArray {
		fmt.Printf("Value: %s, ", value)
	}
	fmt.Println("comparison:", theNames == *otherArray)

	// slices
	// texts := []string {"Kayak", "Lifejacket", "Paddle"} // literal syntax
	texts := make([]string, 3)
	texts[0] = "Kayak"
	texts[1] = "Lifejacket"
	texts[2] = "Paddle"
	texts = append(texts, "Hat")
	appendedTexts := append(texts, "Gloves")
	// In Go, slices are references to arrays. When you modify the elements of a slice, you modify the underlying array,
	// which can be shared by multiple slices. In this case, both texts and appendedTexts shared the same underlying array
	// at the time texts[0] was modified, so the change is visible in both slices.
	texts[0] = "Canoe"
	fmt.Println("sliced texts:", texts)
	fmt.Println("sliced appendedTexts:", appendedTexts)

	/* The reason modifying things[0] does not modify the first element of appendedThings is because, after
	   the append operation, appendedThings is backed by a new array. The things slice continues to reference
	   the original array, so changes to things do not impact appendedThings, which is now independent in terms
	   of its underlying storage.
	   slices have a length and a capacity. The length of a slice is how many values it can currently contain,
	   while the number of elements that can be stored in the underlying array before the slice must be resized
	   and a new array created. The capacity will always be at least the length but can be larger if additional
	   capacity has been allocated with the make function. */
	things := []string{"Kayak", "Lifejacket", "Paddle"}
	moreThings := []string{"Hat", "Gloves"}
	appendedThings := append(things, moreThings...)
	things[0] = "Canoe"
	fmt.Printf("things: %v, cap: %d, length: %d\n", things, cap(things), len(things))
	fmt.Printf("appendedThings: %v, cap: %d, length: %d\n", appendedThings, cap(appendedThings), len(appendedThings))
	fmt.Printf("Slice from existing array: %v\n", appendedThings[:3]) // [:] include all the elements

	// The copy function can be used to duplicate an existing slice, selecting some or all the elements but ensuring
	// that the new slice is backed by its own array. The destination slice is not resized, even when there is
	// capacity available in the existing backing array, which means that you must ensure there is sufficient
	// length to accommodate the number of elements you want to copy.
	// If the destination slice is smaller than the source slice, then copying continues until all the elements in the
	// destination slice have been replaced
	items := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allITems := items[1:] // "Lifejacket", "Paddle", "Hat"
	someItems := make([]string, 2)
	copy(someItems, allITems)
	fmt.Println("someItems:", someItems)
	fmt.Println("allITems", allITems)

	newItems := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNewITems := newItems[1:] // "Lifejacket", "Paddle", "Hat"
	someNewItems := make([]string, 2)
	someNewItems = append(someNewItems, allNewITems...)
	fmt.Println("someNewItems:", someNewItems)
	fmt.Println("allNewITems", allNewITems)

	copyNewItems := newItems
	fmt.Println("Equal:", reflect.DeepEqual(newItems, copyNewItems))

	// Getting the Array Underlying a Slice
	arrayPtr := (*[5]string)(appendedThings)
	array := *arrayPtr
	fmt.Println(array)

	// maps
	/* literal syntax
		products := map[string]float64 {
	        "Kayak" : 279,
	        "Lifejacket": 48.95,
	    }
	*/
	allProducts := make(map[string]float64, 10)
	allProducts["Kayak"] = 279
	allProducts["Lifejacket"] = 48.95
	fmt.Println("Map size:", len(allProducts))
	fmt.Println("Price:", allProducts["Kayak"])
	fmt.Println("Price:", allProducts["Hat"])
	fmt.Println("Price:", allProducts["Hsat"]) // returns 0!

	// maps return the zero value for the value type when reads are performed for which there is no key.
	// This can make it difficult to differentiate between a stored value that happens to be the zero value and a
	// nonexistent key. To solve this problem, maps produce two values when reading a value. “comma ok” technique
	value, ok := allProducts["Hsat"]
	if ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value: Hsat")
	}

	delete(allProducts, "Hat")

	if value, ok := allProducts["Hat"]; ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value")
	}

	// there are no guarantees that the contents of a map will be enumerated in any specific order. If you want to
	// get the values in a map in order, then the best approach is to enumerate the map and create a slice containing
	// the keys, sort the slice, and then enumerate the slice to read the values from the map

	// € is composed by three bytes, that's why we use rune
	var thePrice []rune = []rune("€48.95")
	var currency string = string(thePrice[0])
	var amountString string = string(thePrice[1:])
	amount, parseErr := strconv.ParseFloat(amountString, 64)
	fmt.Println("Length:", len(thePrice))
	fmt.Println("Currency:", currency)
	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}

	fmt.Println("\nDefining and Using Functions -------------------")
}
