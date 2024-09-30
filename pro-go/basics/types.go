package main

// Values do not have to be provided for all fields when creating a struct value
type Product struct {
	name, category string
	price          float64
}

type ProductList []Product

type calcFunc func(float64) float64 // Function Type Alias