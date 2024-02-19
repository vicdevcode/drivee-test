package entity

type Order struct {
	ID        int
	Start     Coordinates
	End       Coordinates
	Price     int
	CourierID *int
}
