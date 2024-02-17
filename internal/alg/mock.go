package alg

import "github.com/vicdevcode/drivee-test/internal/entity"

func FakeData() ([]entity.Order, []entity.Courier) {
	return []entity.Order{
			{
				Start: entity.Point{X: 3, Y: 18},
				End:   entity.Point{X: 15, Y: 8},
				Price: 50,
			},
			{
				Start: entity.Point{X: 9, Y: 4},
				End:   entity.Point{X: 13, Y: 5},
				Price: 50,
			},
			{
				Start: entity.Point{X: 6, Y: 8},
				End:   entity.Point{X: 14, Y: 13},
				Price: 50,
			},
			{
				Start: entity.Point{X: 4, Y: 13},
				End:   entity.Point{X: 8, Y: 11},
				Price: 50,
			},
			{
				Start: entity.Point{X: 3, Y: 2},
				End:   entity.Point{X: 16, Y: 2},
				Price: 50,
			},
			{
				Start: entity.Point{X: 10, Y: 18},
				End:   entity.Point{X: 11, Y: 12},
				Price: 50,
			},
			{
				Start: entity.Point{X: 2, Y: 4},
				End:   entity.Point{X: 17, Y: 10},
				Price: 50,
			},
		}, []entity.Courier{
			{
				ID:     0,
				Point:  entity.Point{X: 2, Y: 8},
				Orders: nil,
			},
			{
				ID:     1,
				Point:  entity.Point{X: 13, Y: 16},
				Orders: nil,
			},
		}
}
