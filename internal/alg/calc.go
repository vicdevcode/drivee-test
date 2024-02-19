package alg

import (
	"math"

	"github.com/vicdevcode/drivee-test/internal/entity"
)

// Измеряем расстояние между 3 точками (C -> A -> B) либо (B -> A -> B).
// Почему B в начале? Потому что B является концом заказа,
// где окажется в итоге курьер
func calculateDistance(First entity.Coordinates, Start entity.Coordinates, End entity.Coordinates, Penalty float64) float64 {
	fromFirstToStart := getDistanceFromLatLonInKm(First.Latitude, First.Longitude, Start.Latitude, Start.Longitude)
	fromStartToEnd := getDistanceFromLatLonInKm(Start.Latitude, Start.Longitude, End.Latitude, End.Longitude)

	return (fromFirstToStart + fromStartToEnd) * Penalty
}

func deg2rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func getDistanceFromLatLonInKm(lat1, lon1, lat2, lon2 float64) float64 {
	R := 6371.0
	dLat := deg2rad(lat2 - lat1)
	dLon := deg2rad(lon2 - lon1)
	a := math.Pow(math.Sin(dLat/2), 2) + (math.Cos(deg2rad(lat1)) * math.Cos(deg2rad(lat2)) * math.Pow(math.Sin(dLon), 2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func calculateTime(V float64, S float64) float64 {
	return S / V
}
