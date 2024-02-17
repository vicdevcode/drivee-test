package entity

type Courier struct {
	ID     int
	Point  Point
	Orders []Order
}
