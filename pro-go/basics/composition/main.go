package main

/*
Composition is the process by which new types are created by combining structs and
interfaces. go mod init composition
*/

import (
	"composition/store"
	"fmt"
)

func main() {
	kayak := store.NewProduct("Kayak", "Watersports", 275)
	lifejacket := &store.Product{Name: "Lifejacket", Category: "Watersports"}
	for _, p := range []*store.Product{kayak, lifejacket} {
		fmt.Println("Name:", p.Name, "Category:", p.Category, "Price:", p.Price(0.2))
	}

	boats := []*store.Boat{
		store.NewBoat("Kayak", 275, 1, false),
		store.NewBoat("Canoe", 400, 3, false),
		store.NewBoat("Tender", 650.25, 2, true),
	}
	for _, b := range boats {
		/*
			Go allows the fields of the nested type to be accessed in two ways. The first is the
			conventional approach of navigating the hierarchy of types to reach the value that is
			required (b.Product.Name).
			Go also allows nested field types to be used directly (b.Name). This is known as field promotion, where
			Go essentially flattens the types so that the Boat type behaves as though it defines the fields
			that are provided by the nested Product type.
			Go can perform promotion only if there is no field or method defined with the same name on the enclosing type
		*/
		fmt.Println("Conventional:", b.Product.Name, "Direct:", b.Name, "Price:", b.Price(0.2))
	}

	rentals := []*store.RentalBoat{
		store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
		store.NewRentalBoat("Yacht", 50000, 5, true, true, "Bob", "Alice"),
		store.NewRentalBoat("Super Yacht", 100000, 15, true, true, "Dora", "Charlie"),
	}
	for _, r := range rentals {
		fmt.Println("Rental Boat:", r.Name, "Rental Price:", r.Price(0.2), "Captain:", r.Captain)
	}

	product := store.NewProduct("Kayak", "Watersports", 279)
	deal := store.NewSpecialDeal("Weekend Special", product, 50)
	Name, price, Price := deal.GetDetails()
	fmt.Println("Name:", Name, "From deal:", deal.Name)
	fmt.Println("Price field:", price)
	fmt.Println("Price method:", Price)

	products := map[string]store.ItemForSale{
		"Kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball":  store.NewProduct("Soccer Ball", "Soccer", 19.50),
	}
	// using single-type statemnents allows to assert the right type and so the properties will be accessible.
	for key, p := range products {
		switch item := p.(type) {
		case *store.Product:
			fmt.Println("Name:", item.Name, "Category:", item.Category,
				"Price:", item.Price(0.2))
		case *store.Boat:
			fmt.Println("Name:", item.Name, "Category:", item.Category,
				"Price:", item.Price(0.2))
		default:
			fmt.Println("Key:", key, "Price:", p.Price(0.2))
		}
	}

	// An alternative solution is to define interface methods that provide access to the property values.
	for key, p := range products {
        switch item := p.(type) {
            case store.Describable:
                fmt.Println("Name:", item.GetName(), "Category:", item.GetCategory(),
                    "Price:", item.Price(0.2))
            default:
                fmt.Println("Key:", key, "Price:", p.Price(0.2))
        }
    }
}
