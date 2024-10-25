package store

// Interfaces can specify only methods
type ItemForSale interface {
    Price(taxRate float64) float64
}