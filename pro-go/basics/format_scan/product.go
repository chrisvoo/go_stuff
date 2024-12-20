package main

import (
	"fmt"
)

type Product struct {
    Name, Category string
    Price float64
}

// The String method specified by the Stringer interface will be used to obtain
// a string representation of any type that defines it
func (p Product) String() string {
    return fmt.Sprintf("Product: %v, Price: $%4.2f", p.Name, p.Price)
}

var Kayak = Product {
    Name: "Kayak",
    Category: "Watersports",
    Price: 275,
}

var Products = []Product {
    { "Kayak", "Watersports", 279 },
    { "Lifejacket", "Watersports", 49.95 },
    { "Soccer Ball", "Soccer", 19.50 },
    { "Corner Flags", "Soccer", 34.95 },
    { "Stadium", "Soccer", 79500 },
    { "Thinking Cap", "Chess", 16 },
    { "Unsteady Chair", "Chess", 75 },
    { "Bling-Bling King", "Chess", 1200 },
}

func GetProductName(index int) (name string, err error) {
    if (len(Products) > index && index > 0) {
        name = fmt.Sprintf("Name of product: %v", Products[index].Name)
    } else {
        err = fmt.Errorf("error for index %v", index)
    }
    return
}

func Printfln(template string, values ...interface{}) {
    fmt.Printf(template + "\n", values...)
}