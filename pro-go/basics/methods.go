package main

import (
	"fmt"
)

// this line disables linting for all the following statements. Just replace disable with enable to
// re-enable the rule for the rest of the file
// revive:disable:exported
func PrintNumber(number int) {
	fmt.Println(number)
}

func printPrice(product string, price float64, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}

// Variadic parameters allow a function to receive a variable number of arguments
func printSuppliers(product string, suppliers ...string) {
	if len(suppliers) == 0 {
		fmt.Println("Product:", product, "Supplier: (none)")
	} else {
		for _, supplier := range suppliers {
			fmt.Println("Product:", product, "Supplier:", supplier)
		}
	}
}

func swapValues(first, second *int) {
	fmt.Println("Before swap:", *first, *second)
	temp := *first
	*first = *second
	*second = temp
	fmt.Println("After swap:", *first, *second)
}

func swapValuesMultiResult(first, second int) (int, int) {
	return second, first
}

func calcTax(price float64) (float64, bool) {
	if price > 100 {
		return price * 0.2, true
	}
	return 0, false
}

/*
Named resuls. The empty return is necessary, it allows the current values assigned to the named results to be returned.
Within the function, the results can be used as regular variables
*/
func calcTotalPrice(products map[string]float64, minSpend float64) (total, tax float64) {
	total = minSpend
	for _, price := range products {
		if taxAmount, due := calcTax(price); due {
			total += taxAmount
			tax += taxAmount
		} else {
			total += price
		}
	}
	return
}



func calcWithTax(price float64) float64 {
	return price + (price * 0.2)
}

func calcWithoutTax(price float64) float64 {
	return price
}

func printPriceWithFunc(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

func selectCalculator(price float64) calcFunc {
	if price > 100 {
		return calcWithTax
	}
	return calcWithoutTax

	/* with anonymous functions

		if (price > 100) {
	        var withTax calcFunc = func (price float64) float64 {
	            return price + (price * 0.2)
	        }
	        return withTax
	    }
	     withoutTax := func (price float64) float64 {
	        return price
	    }
	    return withoutTax
	*/
}

// The defer keyword is used to schedule a function call that will be performed immediately before the current function
// returns. The main use for the defer keyword is to call functions that release resources, such as closing open files or HTTP connection.
// The defer keyword lets you group the statements that create, use, and release the resource together.
// Go will perform the calls scheduled with the defer keyword in the order in which they were defined (LIFO)
func calcGrandTotalPriceWithDefer(products map[string]float64) (count int, total float64) {
	fmt.Println("Function started")
	defer fmt.Println("First defer call")
	count = len(products)
	for _, price := range products {
		total += price
	}
	defer fmt.Println("Second defer call")
	fmt.Println("Function about to return")
	return
}



func printDetails(product *Product) {
	fmt.Println("Name:", product.name, "Category:", product.category, "Price", product.price)
}

// similar as printDetails above, but this time it's a Product method, since we've specified the receiver (before the method name).
// A method whose receiver is a pointer type can also be invoked through a regular value of the underlying type
func (product *Product) printDetails() {
	fmt.Println("Name:", product.name, "Category:", product.category,
		"Price", product.price)
}

func (products *ProductList) calcCategoryTotals() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}
	return totals
}