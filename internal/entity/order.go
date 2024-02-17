package entity

type Order struct {
	Start     Point
	End       Point
	Price     int
	CourierID *int
}
