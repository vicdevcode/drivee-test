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
			{
				Start: entity.Coordinates{Latitude: 62.030567, Longitude: 129.732124},
				End:   entity.Coordinates{Latitude: 62.024381, Longitude: 129.732268},
				Price: 50,
			},
			{
				Start: entity.Coordinates{Latitude: 62.022144, Longitude: 129.721964},
				End:   entity.Coordinates{Latitude: 62.038718, Longitude: 129.692643},
				Price: 50,
			},
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
				Speed:       10,
				Penalty:     1,
			},
		}
}
