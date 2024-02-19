package entity

type Order struct {
	Start     Coordinates
	End       Coordinates
	Price     int
	CourierID *int
}
