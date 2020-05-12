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

// NewRandomPoint generates a random point.
// If a bbox is given the coordinate will be inside this bbox.
func NewRandomPoint(bbox *Bbox) *Point {
	if bbox == nil {
		return NewPoint(
			rand.Float64()*360,
			rand.Float64()*180,
		)

	}

	return NewPoint(
		(rand.Float64()*(bbox.LonMax-bbox.LonMin))+bbox.LonMin,
		(rand.Float64()*(bbox.LatMax-bbox.LatMin))+bbox.LatMin,
	)
}

// NewRandomPoints generates random points as a GeoJSON feature organized in a collection.
// If a bbox is given the points will be generated inside this bbox.
func NewRandomPoints(count uint, bbox *Bbox) *geojson.FeatureCollection {
	container := geojson.NewFeatureCollection()

	for i := uint(0); i < count; i++ {
		container.AddFeature(NewRandomPoint(bbox).AsFeature())
	}

	return container
}
