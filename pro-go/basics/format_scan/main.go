package main

import (
	"fmt"
	"os"
)

func main() {
	arg := "strings"
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	switch arg {
	case "strings":
		{
			//  The template is scanned for verbs, which are denoted by the percentage sign (the % character) followed by a format specifier.
			fmt.Printf("Product: %v, Price: $%4.2f\n", Kayak.Name, Kayak.Price)
			name, err := GetProductName(1)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(name)
			}

			Printfln("Value: %v", Kayak)
			Printfln("Go syntax: %#v", Kayak)
			Printfln("Type: %T", Kayak)
			Printfln("Value with fields: %+v", Kayak)

			fmt.Println(Kayak)

			number := 250
			Printfln("Binary: %b", number)
			Printfln("Decimal: %d", number)
			Printfln("Octal: %o, %O", number, number)
			Printfln("Hexadecimal: %x, %X", number, number)

			number2 := 250.5
			Printfln("Decimal without exponent: %8.2f", number2) // includes padding
			Printfln("Sign: >>%+.2f<<", number2)
			Printfln("Zeros for Padding: >>%010.2f<<", number2)
			Printfln("Right Padding: >>%-8.2f<<", number2)

			Printfln("String: %s", name)
			Printfln("Character: %c", []rune(name)[0])
			Printfln("Unicode: %U", []rune(name)[0])
		}
	case "scan":
		{
			/*
				The fmt package provides functions for scanning strings, which is the process
				of parsing strings that contain values separated by spaces.
			*/
			var name string
			var category string
			var price float64
			fmt.Print("Enter text to scan: ")
			/* The Scan function reads a string from the standard input
			   and scans it for values separated by spaces. The values parsed from the string
			   are assigned to the parameters in the order in which they are defined. */
			n, err := fmt.Scanln(&name, &category, &price)
			if err == nil {
				Printfln("Scanned %v values", n)
				Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
			} else {
				Printfln("Error: %v", err.Error())
			}

			source := "Lifejacket Watersports 48.95"
			n, err = fmt.Sscan(source, &name, &category, &price)
			if (err == nil) {
				Printfln("Scanned %v values", n)
				Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
			} else {
				Printfln("Error: %v", err.Error())
			}
		}
	}
}
