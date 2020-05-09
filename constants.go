package goturf

// Ported from: https://github.com/Turfjs/turf/blob/master/packages/turf-helpers/index.ts

import "math"

// Unit represents the available units
type Unit string

// UnitFactors Unit of measurement factors using a spherical (non-ellipsoid) earth radius in radiant
var UnitFactors = make(map[Unit]float64)

// StringUnitFactors ...
var StringUnitFactors = make(map[string]Unit)

const (
	// oneDeg factor to convert Rad/Deg and vice versa
	oneDeg = math.Pi / 180
	// EarthRadiusInMeters of earth in meters
	EarthRadiusInMeters = 6371008.8

	// UnitMeters SI unit meters
	UnitMeters Unit = "meters"
	// UnitKilometers SI unit kilometers
	UnitKilometers Unit = "kilometers"
	// UnitMiles imperial unit miles
	UnitMiles Unit = "miles"
	// UnitCentimeters SI unit centimeters
	UnitCentimeters Unit = "centimeters"
	// UnitInches imperial unit inches
	UnitInches Unit = "inches"
	// UnitFeet imperial unit feet
	UnitFeet Unit = "feet"
	// UnitYards imperial unit yards
	UnitYards Unit = "yards"
	// UnitMillimeters SI unit millimeters
	UnitMillimeters Unit = "millimeters"
	// UnitNauticalMiles nautical miles
	UnitNauticalMiles Unit = "nauticalmiles"
	// UnitDegrees imperial unit degrees
	UnitDegrees Unit = "degrees"
	// UnitRadians unit rad
	UnitRadians Unit = "radians"
)

func init() {
	UnitFactors = map[Unit]float64{
		UnitCentimeters:   EarthRadiusInMeters * 100,
		UnitDegrees:       EarthRadiusInMeters / 111325,
		UnitFeet:          EarthRadiusInMeters * 3.28084,
		UnitInches:        EarthRadiusInMeters * 39.370,
		UnitKilometers:    EarthRadiusInMeters / 1000,
		UnitMeters:        EarthRadiusInMeters,
		UnitMiles:         EarthRadiusInMeters / 1609.344,
		UnitMillimeters:   EarthRadiusInMeters * 1000,
		UnitNauticalMiles: EarthRadiusInMeters / 1852,
		UnitRadians:       1,
		UnitYards:         EarthRadiusInMeters / 1.0936,
	}

	for unit := range UnitFactors {
		StringUnitFactors[string(unit)] = unit
	}
}
