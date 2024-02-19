package alg

import "github.com/vicdevcode/drivee-test/internal/entity"

func FakeData() ([]entity.Order, []entity.Courier) {
	return []entity.Order{
			{
				Start: entity.Coordinates{Latitude: 62.026187, Longitude: 129.708867},
				End:   entity.Coordinates{Latitude: 62.037566, Longitude: 129.723141},
				Price: 50,
			},
			{
				Start: entity.Coordinates{Latitude: 62.028504, Longitude: 129.732519},
				End:   entity.Coordinates{Latitude: 62.030221, Longitude: 129.753908},
				Price: 50,
			},
			{
				Start: entity.Coordinates{Latitude: 62.049915, Longitude: 129.740550},
				End:   entity.Coordinates{Latitude: 62.024507, Longitude: 129.739338},
				Price: 50,
			},
			// {
			// 	Start: entity.Coordinates{Latitude: 4, Longitude: 13},
			// 	End:   entity.Coordinates{Latitude: 8, Longitude: 11},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 3, Longitude: 2},
			// 	End:   entity.Coordinates{Latitude: 16, Longitude: 2},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 10, Longitude: 18},
			// 	End:   entity.Coordinates{Latitude: 11, Longitude: 12},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 2, Longitude: 4},
			// 	End:   entity.Coordinates{Latitude: 17, Longitude: 10},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 20, Longitude: 18},
			// 	End:   entity.Coordinates{Latitude: 15, Longitude: 30},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 24, Longitude: 14},
			// 	End:   entity.Coordinates{Latitude: 23, Longitude: 25},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 1, Longitude: 20},
			// 	End:   entity.Coordinates{Latitude: 1, Longitude: 12},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 2, Longitude: 1},
			// 	End:   entity.Coordinates{Latitude: 5, Longitude: 6},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 13, Longitude: 22},
			// 	End:   entity.Coordinates{Latitude: 12, Longitude: 15},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 10, Longitude: 12},
			// 	End:   entity.Coordinates{Latitude: 16, Longitude: 20},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 12, Longitude: 23},
			// 	End:   entity.Coordinates{Latitude: 13, Longitude: 8},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 3, Longitude: 18},
			// 	End:   entity.Coordinates{Latitude: 15, Longitude: 8},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 9, Longitude: 4},
			// 	End:   entity.Coordinates{Latitude: 13, Longitude: 5},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 6, Longitude: 8},
			// 	End:   entity.Coordinates{Latitude: 14, Longitude: 13},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 4, Longitude: 13},
			// 	End:   entity.Coordinates{Latitude: 8, Longitude: 11},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 3, Longitude: 2},
			// 	End:   entity.Coordinates{Latitude: 16, Longitude: 2},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 10, Longitude: 18},
			// 	End:   entity.Coordinates{Latitude: 11, Longitude: 12},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 2, Longitude: 4},
			// 	End:   entity.Coordinates{Latitude: 17, Longitude: 10},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 20, Longitude: 18},
			// 	End:   entity.Coordinates{Latitude: 15, Longitude: 30},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 24, Longitude: 14},
			// 	End:   entity.Coordinates{Latitude: 23, Longitude: 25},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 1, Longitude: 20},
			// 	End:   entity.Coordinates{Latitude: 1, Longitude: 12},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 2, Longitude: 1},
			// 	End:   entity.Coordinates{Latitude: 5, Longitude: 6},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 13, Longitude: 22},
			// 	End:   entity.Coordinates{Latitude: 12, Longitude: 15},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 10, Longitude: 12},
			// 	End:   entity.Coordinates{Latitude: 16, Longitude: 20},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 12, Longitude: 23},
			// 	End:   entity.Coordinates{Latitude: 13, Longitude: 8},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 3, Longitude: 18},
			// 	End:   entity.Coordinates{Latitude: 15, Longitude: 8},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 9, Longitude: 4},
			// 	End:   entity.Coordinates{Latitude: 13, Longitude: 5},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 6, Longitude: 8},
			// 	End:   entity.Coordinates{Latitude: 14, Longitude: 13},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 4, Longitude: 13},
			// 	End:   entity.Coordinates{Latitude: 8, Longitude: 11},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 3, Longitude: 2},
			// 	End:   entity.Coordinates{Latitude: 16, Longitude: 2},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 10, Longitude: 18},
			// 	End:   entity.Coordinates{Latitude: 11, Longitude: 12},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 2, Longitude: 4},
			// 	End:   entity.Coordinates{Latitude: 17, Longitude: 10},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 20, Longitude: 18},
			// 	End:   entity.Coordinates{Latitude: 15, Longitude: 30},
			// 	Price: 50,
			// },
			// {
			// 	Start: entity.Coordinates{Latitude: 24, Longitude: 14},
			// 	End:   entity.Coordinates{Latitude: 23, Longitude: 25},
			// 	Price: 50,
			// },
		}, []entity.Courier{
			{
				ID:          0,
				Coordinates: entity.Coordinates{Latitude: 62.021967, Longitude: 129.719790},
				Orders:      nil,
				Speed:       30,
				Penalty:     1,
			},
			{
				ID:          1,
				Coordinates: entity.Coordinates{Latitude: 62.027141, Longitude: 129.735053},
				Orders:      nil,
				Speed:       30,
				Penalty:     1,
			},
			// {
			// 	ID:          2,
			// 	Coordinates: entity.Coordinates{Latitude: 11, Longitude: 4},
			// 	Orders:      nil,
			// 	Speed:       30,
			// 	Penalty:     0.7,
			// },
			// {
			// 	ID:          3,
			// 	Coordinates: entity.Coordinates{Latitude: 2, Longitude: 16},
			// 	Orders:      nil,
			// 	Speed:       30,
			// 	Penalty:     0.95,
			// },
			// {
			// 	ID:          4,
			// 	Coordinates: entity.Coordinates{Latitude: 20, Longitude: 20},
			// 	Orders:      nil,
			// 	Speed:       10,
			// 	Penalty:     0.6,
			// },
		}
}
