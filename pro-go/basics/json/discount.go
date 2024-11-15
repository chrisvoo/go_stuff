package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name, Category string
	price          float64
}

type DiscountedProduct struct {
	*Product
	Discount float64
}

// How a struct is encoded can be customized using struct tags. By default, the JSON Encoder includes struct fields,
// even when they have not been assigned a value (use omitempty to avoid this)
type CustomDiscountedProduct struct {
	*Product `json:"p,omitempty"`
	Discount float64 `json:"d,string"` // forced as string
	color    string  `json:"-"`        // omitted
}

// implements the Marshaler interface for pointers to the DiscountedProduct struct type
func (dp *DiscountedProduct) MarshalJSON() (jsn []byte, err error) {
	if dp.Product != nil {
		m := map[string]interface{}{
			"product": dp.Name,
			"cost":    dp.price - dp.Discount,
		}
		jsn, err = json.Marshal(m)
	}
	return
}

func Printfln(template string, values ...interface{}) {
    fmt.Printf(template + "\n", values...)
}
