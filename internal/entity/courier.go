package entity

type Courier struct {
	ID          int
	Point       Point
	Coordinates Coordinates
	Orders      []Order
	Penalty     float64
	Speed       float64
}
