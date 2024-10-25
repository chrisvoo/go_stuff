package main

// There is a special alias, known as the dot import, that allows a package’s
// features to be used without using a prefix
import (
	"fmt"
	currencyFmt "packages/fmt" // package alias
	//. "packages/fmt" 		   // dot import
	"github.com/fatih/color"   // external package, the last part is the prefix
	_ "packages/data" // avoid unused error, but need just init func
	"packages/store"
	"packages/store/cart"
)

/*
 1. created the packages folder
 2. cd into it
 3. go mod init packages
*/

func main() {
	product := store.Product{
		Name:     "Kayak",
		Category: "Watersports",
	}

	anotherProduct := store.NewProduct("Kayak", "Watersports", 279)
	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Another product Price:", anotherProduct.Price())
	color.Green("Price: " + currencyFmt.ToCurrency(anotherProduct.Price()))

	cart := cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{*anotherProduct},
	}
	fmt.Println("Name:", cart.CustomerName)
	color.Green("Total: " + currencyFmt.ToCurrency(cart.GetTotal()))
}
