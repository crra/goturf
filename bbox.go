package goturf

// Ported from: https://github.com/Turfjs/turf/blob/master/packages/turf-bbox-polygon/index.ts
// Ported from: https://github.com/Turfjs/turf/blob/master/packages/turf-bbox/index.ts

import (
	"math"

	geojson "github.com/paulmach/go.geojson"
)

var infPos = math.Inf(0)
var infNeg = math.Inf(-1)

// Bbox is a boundingbox
type Bbox struct {
	LonMin float64
	LatMin float64
	LonMax float64
	LatMax float64
}

// AsTuple returns the bbox as (ordered) tuple
func (bbox *Bbox) AsTuple() []float64 {
	return []float64{bbox.LonMin, bbox.LatMin, bbox.LonMax, bbox.LatMax}
}

// AsPolygon expresses the bbox as polygon
func (bbox *Bbox) AsPolygon() (*geojson.Geometry, error) {
	// alias for readability (same as in turfjs)
	west := bbox.LonMin
	south := bbox.LatMin
	east := bbox.LonMax
	north := bbox.LatMax

	lowLeft := (&Coordinate{Lon: west, Lat: south}).AsTuple()
	topLeft := (&Coordinate{Lon: west, Lat: north}).AsTuple()
	topRight := (&Coordinate{Lon: east, Lat: north}).AsTuple()
	lowRight := (&Coordinate{Lon: east, Lat: south}).AsTuple()

	return geojson.NewPolygonGeometry([][][]float64{
		[][]float64{
			lowLeft,
			lowRight,
			topRight,
			topLeft,

			lowLeft, // Closing the polygon
		},
	}), nil
}

// NewBboxFromGeoJSON returns the bbox of all features
func NewBboxFromGeoJSON(thing interface{}) *Bbox {
	result := &Bbox{
		LonMin: infPos,
		LatMin: infPos,
		LonMax: infNeg,
		LatMax: infNeg,
	}

	CoordEach(thing, func(coord *Coordinate, coordIndex, featureIndex, multiFeatureIndex, geometryIndex int) bool {
		result.LonMin = math.Min(result.LonMin, coord.Lon)
		result.LatMin = math.Min(result.LatMin, coord.Lat)

		result.LonMax = math.Max(result.LonMax, coord.Lon)
		result.LatMax = math.Max(result.LatMax, coord.Lat)

		return true
	})

	return result
}
