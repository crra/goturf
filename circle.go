package goturf

// Ported from: https://github.com/Turfjs/turf/blob/master/packages/turf-circle/index.ts

import (
	"fmt"

	geojson "github.com/paulmach/go.geojson"
)

// NewCircle Takes a point and calculates the circle polygon given a radius in degrees, radians, miles, or kilometers;
func NewCircle(center *Coordinate, radius float64, unit Unit) (*geojson.Geometry, error) {
	steps := 32 // TODO: options object?

	coordinates := []Coordinate{}

	for i := 0; i < steps; i++ {
		bearing := float64(i * -360 / steps)
		next, err := Destination(center, radius, unit, bearing)
		if err != nil {
			return nil, fmt.Errorf("Can't calculate destination, %v", err)
		}
		coordinates = append(coordinates, *next)
	}
	// Close the polygon
	coordinates = append(coordinates, coordinates[0])

	// The golang version of: map(coord.AsPair())
	points := make([][]float64, len(coordinates))
	for i, coord := range coordinates {
		points[i] = coord.AsTuple()
	}

	return geojson.NewPolygonGeometry([][][]float64{points}), nil
}
