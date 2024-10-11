package main

import (
	"fmt"
	"packages/store"
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

	// Adding Code Files to Packages continues...
}
