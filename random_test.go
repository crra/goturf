package goturf_test

import (
	"testing"

	. "github.com/crra/goturf"
	"github.com/mmcloughlin/geohash"
	"github.com/stretchr/testify/assert"
)

func newWorldBbox() *Bbox {
	return &Bbox{
		LonMin: -180,
		LatMin: -90,
		LonMax: 180,
		LatMax: 90,
	}
}

func TestNewRandomPoint(t *testing.T) {
	t.Run("Returns a point with and without bbox", func(t *testing.T) {
		bboxes := []*Bbox{
			newWorldBbox(),
			nil,
		}

		for _, bbox := range bboxes {
			got := NewRandomPoint(bbox)
			assert.NotNil(t, got)
		}
	})

	t.Run("Consecutive points are different", func(t *testing.T) {
		const runs = 5
		// Use a set to collect unique values
		results := map[string]bool{}

		for i := 0; i < runs; i++ {
			point := NewRandomPoint(nil)
			results[geohash.Encode(point.Lat(), point.Lon())] = true
		}

		// Number of elements in the set
		assert.Equal(t, runs, len(results))
	})
}

func TestNewRandomPoints(t *testing.T) {
	t.Run("Right number of points", func(t *testing.T) {
		expectations := []struct {
			number uint
			bbox   *Bbox
		}{
			{0, nil},
			{0, newWorldBbox()},
			{1, nil},
			{1, newWorldBbox()},
		}

		for _, e := range expectations {
			got := NewRandomPoints(e.number, e.bbox)
			assert.NotNil(t, got)

			assert.Equal(t, len(got.Features), int(e.number))

			for _, f := range got.Features {
				assert.NotNil(t, f)
				assert.Equal(t, "Feature", f.Type)

				assert.NotNil(t, f.Geometry)
				assert.Equal(t, "Point", string(f.Geometry.Type))
			}
		}
	})
}
