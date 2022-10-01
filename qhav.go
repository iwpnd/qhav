package qhav

import (
	"math"
)

const earthRadiusKm = 6371.0088

const conversionToKM = 1.0
const conversionToMeters = 1000.0
const conversionToMiles = 0.621371192

type Unit func() float64

// InKilometers to return distance in kilometers
func InKilometers() Unit {
	return func() float64 {
		return earthRadiusKm * conversionToKM
	}
}

// InMeters to return distance in meters
func InMeters() Unit {
	return func() float64 {
		return earthRadiusKm * conversionToMeters
	}
}

// InMiles to return distance in miles
func InMiles() Unit {
	return func() float64 {
		return earthRadiusKm * conversionToMiles
	}
}

func degreeToRad(p []float64) []float64 {
	points := make([]float64, len(p))

	for i, v := range p {
		points[i] = v * math.Pi / 180
	}

	return points
}

// Haversine to calculate haversine distance between to points
func Haversine(p1, p2 []float64, unit ...Unit) float64 {
	radius := earthRadiusKm

	if len(unit) != 0 {
		for _, u := range unit {
			radius = u()
		}
	}

	dp1 := degreeToRad(p1)
	dp2 := degreeToRad(p2)

	lat1, lng1 := dp1[1], dp1[0]
	lat2, lng2 := dp2[1], dp2[0]

	dlat := lat1 - lat2
	dlng := lng1 - lng2

	d := math.Pow(math.Sin(dlat*0.5), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlng*0.5), 2)

	return 2 * radius * math.Asin(math.Sqrt(d))
}
