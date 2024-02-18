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
			{
				Start: entity.Point{X: 20, Y: 18},
				End:   entity.Point{X: 15, Y: 30},
				Price: 50,
			},
			{
				Start: entity.Point{X: 24, Y: 14},
				End:   entity.Point{X: 23, Y: 25},
				Price: 50,
			},
			{
				Start: entity.Point{X: 1, Y: 20},
				End:   entity.Point{X: 1, Y: 12},
				Price: 50,
			},
			{
				Start: entity.Point{X: 2, Y: 1},
				End:   entity.Point{X: 5, Y: 6},
				Price: 50,
			},
			{
				Start: entity.Point{X: 13, Y: 22},
				End:   entity.Point{X: 12, Y: 15},
				Price: 50,
			},
			{
				Start: entity.Point{X: 10, Y: 12},
				End:   entity.Point{X: 16, Y: 20},
				Price: 50,
			},
			{
				Start: entity.Point{X: 12, Y: 23},
				End:   entity.Point{X: 13, Y: 8},
				Price: 50,
			},
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
			{
				Start: entity.Point{X: 20, Y: 18},
				End:   entity.Point{X: 15, Y: 30},
				Price: 50,
			},
			{
				Start: entity.Point{X: 24, Y: 14},
				End:   entity.Point{X: 23, Y: 25},
				Price: 50,
			},
			{
				Start: entity.Point{X: 1, Y: 20},
				End:   entity.Point{X: 1, Y: 12},
				Price: 50,
			},
			{
				Start: entity.Point{X: 2, Y: 1},
				End:   entity.Point{X: 5, Y: 6},
				Price: 50,
			},
			{
				Start: entity.Point{X: 13, Y: 22},
				End:   entity.Point{X: 12, Y: 15},
				Price: 50,
			},
			{
				Start: entity.Point{X: 10, Y: 12},
				End:   entity.Point{X: 16, Y: 20},
				Price: 50,
			},
			{
				Start: entity.Point{X: 12, Y: 23},
				End:   entity.Point{X: 13, Y: 8},
				Price: 50,
			},
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
			{
				Start: entity.Point{X: 20, Y: 18},
				End:   entity.Point{X: 15, Y: 30},
				Price: 50,
			},
			{
				Start: entity.Point{X: 24, Y: 14},
				End:   entity.Point{X: 23, Y: 25},
				Price: 50,
			},
		}, []entity.Courier{
			{
				ID:      0,
				Point:   entity.Point{X: 2, Y: 8},
				Orders:  nil,
				Penalty: 0.9,
			},
			{
				ID:      1,
				Point:   entity.Point{X: 13, Y: 16},
				Orders:  nil,
				Penalty: 0.9,
			},
			{
				ID:      2,
				Point:   entity.Point{X: 11, Y: 4},
				Orders:  nil,
				Penalty: 0.7,
			},
			{
				ID:      3,
				Point:   entity.Point{X: 2, Y: 16},
				Orders:  nil,
				Penalty: 0.95,
			},
		}
}
