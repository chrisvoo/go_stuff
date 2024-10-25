package store

type Product struct {
	Name, Category string
	price          float64
}

/*
Since there are no classes, there aren't constructors either so we use the convention of
a function named New<Type> for creating new entities
*/
func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (p *Product) Price(taxRate float64) float64 {
	return p.price + (p.price * taxRate)
}

/*
One interface can enclose another, with the effect that types must implement all the methods
defined by the enclosing and enclosed interfaces.
*/
type Describable interface {
	GetName() string
	GetCategory() string
	ItemForSale
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetCategory() string {
	return p.Category
}
