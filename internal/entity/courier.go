package entity

type Courier struct {
	ID           int
	Coordinates  Coordinates
	Orders       []Order
	Path         []int
	Penalty      float64
	Speed        float64
	FullPathTime float64
}
