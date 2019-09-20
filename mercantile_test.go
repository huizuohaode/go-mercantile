package mercantile_test

import (
	"github.com/kasika-technologies/go-mercantile"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBounds(t *testing.T) {
	t.Run("Get the bounding box of a tile", func(t *testing.T) {
		tile := &mercantile.Tile{
			X: 486,
			Y: 332,
			Z: 10,
		}
		lngLatBox := tile.Bounds()

		assert.Equal(t, -9.140625, lngLatBox.West)
		assert.Equal(t, 53.120405283106564, lngLatBox.South)
		assert.Equal(t, -8.7890625, lngLatBox.East)
		assert.Equal(t, 53.33087298301705, lngLatBox.North)
	})
}

func TestTile_XYBounds(t *testing.T) {
	t.Run("Get the web mercator bounding box of a tile", func(t *testing.T) {
		tile := &mercantile.Tile{
			X: 486,
			Y: 332,
			Z: 10,
		}

		bbox := tile.XYBounds()

		assert.Equal(t, -1017529.7205322662, bbox.Left)
		assert.Equal(t, 7005300.768279833, bbox.Bottom)
		assert.Equal(t, -978393.962050256, bbox.Right)
		assert.Equal(t, 7044436.526761842, bbox.Top)
	})
}

func TestUpperLeft(t *testing.T) {
	t.Run("Get the upper left longitude and latitude of a tile", func(t *testing.T) {
		tile := &mercantile.Tile{
			X: 1,
			Y: 1,
			Z: 1,
		}
		lngLat := tile.UpperLeft()

		assert.Equal(t, 0.0, lngLat.Lng)
		assert.Equal(t, 0.0, lngLat.Lat)
	})
}

func TestLngLat_XY(t *testing.T) {
	t.Run("Convert longitude and latitude to web mercator x, y", func(t *testing.T) {
		lngLat := &mercantile.LngLat{
			Lng: -9.140625,
			Lat: 53.33087298301705,
		}

		xy := lngLat.XY(false)

		assert.Equal(t, -1017529.7205322662, xy.X)
		assert.Equal(t, 7044436.526761842, xy.Y)
	})
}

func TestXY_LngLat(t *testing.T) {
	t.Run("Convert web mercator x, y to longitude and latitude", func(t *testing.T) {
		xy := &mercantile.XY{
			X: -1017529.7205322662,
			Y: 7044436.526761842,
		}
		lngLat := xy.LngLat(false)

		assert.Equal(t, -9.140625, lngLat.Lng)
		assert.Equal(t, 53.33087298301704, lngLat.Lat)
	})
}
