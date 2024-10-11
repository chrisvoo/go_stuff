// Package store provides types and methods
// commonly required for online sales
package store

// Product describes an item for sale
type Product struct {
	Name, Category string // Name and type of the product, caital letter
	price          float64 // lower letter, not directly visible outside the package
}

func NewProduct(name, category string, price float64) *Product {
    return &Product{ name, category, price }
}
func (p *Product) Price() float64 {
    return p.price
}
func (p *Product) SetPrice(newPrice float64)  {
    p.price = newPrice
}