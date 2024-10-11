package main

// Values do not have to be provided for all fields when creating a struct value
type Product struct {
	name, category string
	price          float64
}

func (p Product) getName() string {
	return p.name
}

func (p Product) getCost(_ bool) float64 {
	return p.price
}

type ProductList []Product

type calcFunc func(float64) float64 // Function Type Alias

type Service struct {
	description    string
	durationMonths int
	monthlyFee     float64
}

func (s Service) getName() string {
	return s.description
}
func (s Service) getCost(recur bool) float64 {
	if recur {
		return s.monthlyFee * float64(s.durationMonths)
	}
	return s.monthlyFee
}
// Go simply requires that all the methods specified by the interface are defined. 
type Expense interface {
	getName() string
	getCost(annual bool) float64
}
