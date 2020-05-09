package goturf

// Ported from: https://github.com/Turfjs/turf/blob/master/packages/turf-meta/index.js

import (
	"log"

	geojson "github.com/paulmach/go.geojson"
)

// CoordinateCallback is used for walking all coordinates
type CoordinateCallback func(coordinate *Coordinate, coordIndex, featureIndex, multiFeatureIndex, geometryIndex int) bool

// CoordEach visits all Coordinates in a geojson strtuct
func CoordEach(thing interface{}, callback CoordinateCallback) bool {
	// A empty thing is fine
	if thing == nil {
		return true
	}

	// This logic may look a little weird. The reason why it is that way
	// is because it's trying to be fast. GeoJSON supports multiple kinds
	// of objects at its root: FeatureCollection, Features, Geometries.
	// This function has the responsibility of handling all of them, and that
	// means that some of the `for` loops you see below actually just don't apply
	// to certain inputs. For instance, if you give this just a
	// Point geometry, then both loops are short-circuited and all we do
	// is gradually rename the input until it's called 'geometry'.
	//
	// This also aims to allocate as few resources as possible: just a
	// few numbers and booleans, rather than any temporary arrays as would
	// be required with the normalization approach.

	// GeoJSON root nodes
	featureCollection, isFeatureCollection := thing.(*geojson.FeatureCollection)
	feature, isFeature := thing.(*geojson.Feature)
	geometry, _ := thing.(*geojson.Geometry)

	coordIndex := 0

	// At least once, more if the 'thing' is a container ('FeatureCollection')
	stop := 1
	if isFeatureCollection {
		stop = len(featureCollection.Features)
	}
	for featureIndex := 0; featureIndex < stop; featureIndex++ {
		// geometry is already provided as root
		if isFeatureCollection {
			geometry = featureCollection.Features[featureIndex].Geometry
		} else if isFeature {
			geometry = feature.Geometry
		}

		if geometry == nil {
			log.Printf("Unknown case for: %v\n", thing)
			break
		}

		isGeometryCollection := geometry.Type == geojson.GeometryCollection

		// At least once, more for simple collections
		stopG := 1
		if isGeometryCollection {
			stopG = len(geometry.Geometries)
		}
		for geomIndex := 0; geomIndex < stopG; geomIndex++ {
			multiFeatureIndex := 0
			geometryIndex := 0

			currentGeometry := geometry
			if isGeometryCollection {
				currentGeometry = geometry.Geometries[geomIndex]
			}

			if currentGeometry == nil {
				continue
			}

			switch currentGeometry.Type {
			case geojson.GeometryPoint:
				if callback(NewCoordinateFromTuple(currentGeometry.Point), coordIndex, featureIndex, multiFeatureIndex, geometryIndex) == false {
					return false
				}
				coordIndex++
				multiFeatureIndex++

			case geojson.GeometryLineString,
				geojson.GeometryMultiPoint:
				var points [][]float64
				if currentGeometry.Type == geojson.GeometryLineString {
					points = currentGeometry.LineString
				} else if currentGeometry.Type == geojson.GeometryMultiPoint {
					points = currentGeometry.MultiPoint
				}

				for _, point := range points {
					if callback(NewCoordinateFromTuple(point), coordIndex, featureIndex, multiFeatureIndex, geometryIndex) == false {
						return false
					}
					coordIndex++
					if currentGeometry.Type == geojson.GeometryMultiPoint {
						multiFeatureIndex++
					}
				}
				if currentGeometry.Type == geojson.GeometryLineString {
					multiFeatureIndex++
				}

			case geojson.GeometryPolygon,
				geojson.GeometryMultiLineString:

				var lines [][][]float64
				if currentGeometry.Type == geojson.GeometryPolygon {
					lines = currentGeometry.Polygon
				} else if currentGeometry.Type == geojson.GeometryMultiLineString {
					lines = currentGeometry.MultiLineString
				}

				for _, line := range lines {
					for _, point := range line {
						if callback(NewCoordinateFromTuple(point), coordIndex, featureIndex, multiFeatureIndex, geometryIndex) == false {
							return false
						}
						coordIndex++
						if currentGeometry.Type == geojson.GeometryMultiLineString {
							multiFeatureIndex++
						}
						if currentGeometry.Type == geojson.GeometryPolygon {
							geometryIndex++
						}
					}
				}

				if currentGeometry.Type == geojson.GeometryPolygon {
					multiFeatureIndex++
				}

			case geojson.GeometryMultiPolygon:
				for _, poly := range currentGeometry.MultiPolygon {
					for _, line := range poly {
						for _, point := range line {
							if callback(NewCoordinateFromTuple(point), coordIndex, featureIndex, multiFeatureIndex, geometryIndex) == false {
								return false
							}
							coordIndex++
						}
						geometryIndex++
					}
					multiFeatureIndex++
				}

			case geojson.GeometryCollection:
				for _, geo := range currentGeometry.Geometries {
					if CoordEach(geo, callback) == false {
						return false
					}
				}
			}
		}
	}

	return false
}
