package goturf

import geojson "github.com/paulmach/go.geojson"

const (
	indexLon = 0
	indexLat = 1
)

// Point is a 2d point
type Point struct {
	lon float64
	lat float64
}

// NewPoint takes lon, lat and wraps it in a Point
func NewPoint(lon, lat float64) *Point {
	return &Point{
		lon: lon,
		lat: lat,
	}
}

// NewPointFromTuple takes a tuple [lon, lat] and wraps it in a Point
func NewPointFromTuple(coords []float64) *Point {
	if len(coords) != 2 {
		return nil
	}

	return NewPoint(coords[indexLon], coords[indexLat])
}

// Lat returns the vertical, latitude coordinate of the point.
func (p *Point) Lat() float64 {
	return p.lat
}

// Lon returns the horizontal, longitude coordinate of the point.
func (p *Point) Lon() float64 {
	return p.lon
}

// AsTuple encodes lon/lat as list in the order 1) lon 2) lat
func (p *Point) AsTuple() []float64 {
	return []float64{p.lon, p.lat}
}

// AsFeature returns the coordinate as geojson feature
func (p *Point) AsFeature() *geojson.Feature {
	return geojson.NewPointFeature(p.AsTuple())
}
