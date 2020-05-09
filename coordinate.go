package goturf

import geojson "github.com/paulmach/go.geojson"

// Coordinate is a 1D point
type Coordinate struct {
	Lon   float64
	Lat   float64
	IsRad bool
}

// Rad Converts a coordinate to rad if encoded as deg
func (c *Coordinate) Rad() *Coordinate {
	if c.IsRad {
		return c
	}

	return &Coordinate{
		Lon:   Rad(c.Lon),
		Lat:   Rad(c.Lat),
		IsRad: true,
	}
}

// NewCoordinateFromTuple takes a tuple [lon, lat] and wraps it in a Coordinate
func NewCoordinateFromTuple(coords []float64) *Coordinate {
	if len(coords) != 2 {
		return nil
	}

	return &Coordinate{
		Lon: coords[0],
		Lat: coords[1],
	}
}

// AsTuple encodes lon/lat as list in the order 1) lon 2) lat
func (c *Coordinate) AsTuple() []float64 {
	return []float64{c.Lon, c.Lat}
}

// AsFeature returns the coordinate as geojson feature
func (c *Coordinate) AsFeature() *geojson.Feature {
	return geojson.NewPointFeature(c.AsTuple())
}
