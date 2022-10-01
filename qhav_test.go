package qhav

import (
	"math"
	"testing"
)

func approxEqual(want, got, tolerance float64) bool {
	diff := math.Abs(want - got)

	return diff < tolerance
}

func TestHaversine(t *testing.T) {
	test := []struct {
		p1          []float64
		p2          []float64
		unit        string
		expDistance float64
	}{
		{
			p1:          []float64{0.0, 0.0},
			p2:          []float64{0.0, 0.0},
			unit:        "km",
			expDistance: 0,
		},
		{
			p1:          []float64{0.0, 0.0},
			p2:          []float64{1.0, 1.0},
			unit:        "km",
			expDistance: 157.249,
		},
		{
			p1:          []float64{0.0, 0.0, 0.0},
			p2:          []float64{1.0, 1.0, 0.0},
			unit:        "km",
			expDistance: 157.249,
		},
		{
			p1:          []float64{0.0, 0.0},
			p2:          []float64{1.0, 1.0},
			unit:        "meters",
			expDistance: 157249.59,
		},
		{
			p1:          []float64{0.0, 0.0},
			p2:          []float64{1.0, 1.0},
			unit:        "miles",
			expDistance: 97.7106,
		},
	}

	for _, test := range test {
		var u Unit
		switch test.unit {
		case "km":
			u = InKilometers()
		case "meters":
			u = InMeters()
		case "miles":
			u = InMiles()
		}

		got := Haversine(test.p1, test.p2, u)

		if !approxEqual(got, test.expDistance, 0.01) {
			t.Errorf("expected %+v, got: %+v", test.expDistance, got)
		}
	}

}
