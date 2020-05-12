package goturf

// Ported from: https://github.com/Turfjs/turf/blob/packages/turf-destination/index.ts

import (
	"fmt"
	"math"
)

// "Packages that fit in tweets don't need to be in the standard library."
// https://twitter.com/rob_pike/status/344225419928694784

// Rad takes degree and returns rad
func Rad(d float64) float64 {
	return d * oneDeg
}

// Deg takes rad and returns degree
func Deg(r float64) float64 {
	return r / oneDeg
}

// LengthToRadians takes a given
func LengthToRadians(distance float64, unit Unit) (float64, error) {
	factor, ok := UnitFactors[unit]
	if !ok {
		return 0, fmt.Errorf("Can't resolve unit, '%s'", unit)
	}

	return distance / factor, nil
}

// Destination  Takes a point and calculates the location of a destination point given a distance in
// degrees, radians, miles, or kilometers; and bearing in degrees.
// This uses the [Haversine formula](http://en.wikipedia.org/wiki/Haversine_formula) to account for global curvature.
func Destination(origin *Point, distance float64, unit Unit, bearingDeg float64) (*Point, error) {
	originLonRad := Rad(origin.Lon())
	originLatRad := Rad(origin.Lat())
	bearingRad := Rad(bearingDeg)
	radians, err := LengthToRadians(distance, unit)
	if err != nil {
		return nil, fmt.Errorf("Can't convert length '%f' with unit '%s', %v", distance, unit, err)
	}

	latitude2 := math.Asin(math.Sin(originLatRad)*math.Cos(radians) + math.Cos(originLatRad)*math.Sin(radians)*math.Cos(bearingRad))
	longitude2 := originLonRad + math.Atan2(math.Sin(bearingRad)*math.Sin(radians)*math.Cos(originLatRad), math.Cos(radians)-math.Sin(originLatRad)*math.Sin(latitude2))

	return NewPoint(Deg(longitude2), Deg(latitude2)), nil
}
