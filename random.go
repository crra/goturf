package goturf

// Ported from: https://github.com/Turfjs/turf/blob/master/packages/turf-random/index.ts

import (
	"math/rand"
	"time"

	geojson "github.com/paulmach/go.geojson"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// NewRandomCoordinate generates a random coordinate.
// If a bbox is given the coordinate will be inside this bbox.
func NewRandomCoordinate(bbox *Bbox) *Coordinate {
	if bbox == nil {
		return &Coordinate{
			Lon: rand.Float64() * 360,
			Lat: rand.Float64() * 180,
		}
	}

	return &Coordinate{
		Lon: (rand.Float64() * (bbox.LonMax - bbox.LonMin)) + bbox.LonMin,
		Lat: (rand.Float64() * (bbox.LatMax - bbox.LatMin)) + bbox.LatMin,
	}
}

// NewRandomPoints generates random points as a GeoJSON feature oragnized in a collection.
// If a bbox is given the points will be generated inside this bbox.
func NewRandomPoints(count uint, bbox *Bbox) *geojson.FeatureCollection {
	container := geojson.NewFeatureCollection()

	for i := uint(0); i < count; i++ {
		container.AddFeature(NewRandomCoordinate(bbox).AsFeature())
	}

	return container
}
