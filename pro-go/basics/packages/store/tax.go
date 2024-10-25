package store

const defaultTaxRate float64 = 0.2
const minThreshold = 10

var categoryMaxPrices = map[string]float64{
	"Watersports": 250,
	"Soccer":      150,
	"Chess":       50,
}

/*
initialization function, which is invoked automatically when the package is loaded.
It cannot be invoked directly. A single file can define multiple init functions, all of
which will be executed.
Each code file can have its own initialization function. When using the standard Go
compiler, the initialization functions are executed based on the alphabetic order of
the filenames. But this order is not part of the Go language specification and should 
not be relied on. Your initialization functions should be self-contained and not rely
on other init functions having been invoked previously.
*/
func init() {
	for category, price := range categoryMaxPrices {
		categoryMaxPrices[category] = price + (price * defaultTaxRate)
	}
}

type taxRate struct {
	rate, threshold float64
}

func newTaxRate(rate, threshold float64) *taxRate {
	if rate == 0 {
		rate = defaultTaxRate
	}
	if threshold < minThreshold {
		threshold = minThreshold
	}
	return &taxRate{rate, threshold}
}

func (taxRate *taxRate) calcTax(product *Product) (price float64) {
	if product.price > taxRate.threshold {
		price = product.price + (product.price * taxRate.rate)
	} else {
		price = product.price
	}
	if max, ok := categoryMaxPrices[product.Category]; ok && price > max {
		price = max
	}
	return
}
