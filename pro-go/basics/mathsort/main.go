package main

import (
	"math"
	"math/rand"
	"sort"
)

func main() {
	val1 := 279.00
	val2 := 48.95
	Printfln("Abs: %v", math.Abs(val1))
	Printfln("Ceil: %v", math.Ceil(val2))
	Printfln("Copysign: %v", math.Copysign(val1, -5))
	Printfln("Floor: %v", math.Floor(val2))
	Printfln("Max: %v", math.Max(val1, val2))
	Printfln("Min: %v", math.Min(val1, val2))
	Printfln("Mod: %v", math.Mod(val1, val2))
	Printfln("Pow: %v", math.Pow(val1, 2))
	Printfln("Round: %v", math.Round(val2))
	Printfln("RoundToEven: %v", math.RoundToEven(val2))

	for i := 0; i < 5; i++ {
		Printfln("Value %v : %v", i, IntRange(10, 20))
	}

	var names = []string{"Alice", "Bob", "Charlie", "Dora", "Edith"}
	rand.Shuffle(len(names), func(first, second int) {
		names[first], names[second] = names[second], names[first]
	})
	for i, name := range names {
		Printfln("Index %v: Name: %v", i, name)
	}

	// sort the elements in place, rather than creating a new slice.
	ints := []int{9, 4, 2, -1, 10}
	Printfln("Ints: %v", ints)
	sort.Ints(ints)
	Printfln("Ints Sorted: %v", ints)
	floats := []float64{279, 48.95, 19.50}
	Printfln("Floats: %v", floats)
	sort.Float64s(floats)
	Printfln("Floats Sorted: %v", floats)
	strings := []string{"Kayak", "Lifejacket", "Stadium"}
	Printfln("Strings: %v", strings)
	if !sort.StringsAreSorted(strings) {
		sort.Strings(strings)
		Printfln("Strings Sorted: %v", strings)
	} else {
		Printfln("Strings Already Sorted: %v", strings)
	}

	// If you want to create a new, sorted slice, then you must use the built-in make and copy functions
	ints = []int { 9, 4, 2, -1, 10}
    sortedInts := make([]int, len(ints))
    copy(sortedInts, ints)
    sort.Ints(sortedInts)
    Printfln("Ints: %v", ints)
    Printfln("Ints Sorted: %v", sortedInts)
	// The return value is the index to insert x if x is not present (it could be len(a)).
	// The slice must be sorted in ascending order.
	indexOf4:= sort.SearchInts(sortedInts, 4)
    indexOf3 := sort.SearchInts(sortedInts, 3)
	Printfln("Index of 4: %v (present: %v)", indexOf4, sortedInts[indexOf4] == 4)
    Printfln("Index of 3: %v (present: %v)", indexOf3, sortedInts[indexOf3] == 3)

	products := []Product {
        { "Kayak", 279} ,
        { "Lifejacket", 49.95 },
        { "Soccer Ball",  19.50 },
    }
    ProductSlices(products)
    for _, p := range products {
        Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
    }

	Printfln("Are sorted? %v", ProductSlicesAreSorted(products))

	ProductSlicesByName(products)
    for _, p := range products {
        Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
    }

	SortWith(products, func (p1, p2 Product) bool {
        return p1.Name < p2.Name
    })
    for _, p := range products {
        Printfln("Name: %v, Price: %.2f",  p.Name, p.Price)
    }
}
