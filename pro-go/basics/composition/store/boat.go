package store

type Boat struct {
	*Product  // embedded type
	Capacity  int
	Motorized bool
}

func NewBoat(name string, price float64, capacity int, motorized bool) *Boat {
	return &Boat{
		NewProduct(name, "Watersports", price),
		capacity,
		motorized,
	}
}
