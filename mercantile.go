package mercantile

import (
	"math"
)

type Bbox struct {
	Left   float64
	Bottom float64
	Right  float64
	Top    float64
}

type LngLatBbox struct {
	West  float64
	South float64
	East  float64
	North float64
}

type Tile struct {
	X int
	Y int
	Z int
}

type LngLat struct {
	Lng float64
	Lat float64
}

type XY struct {
	X float64
	Y float64
}

func (t *Tile) Bounds() *LngLatBbox {
	nt := &Tile{
		X: t.X + 1,
		Y: t.Y + 1,
		Z: t.Z,
	}

	a := t.UpperLeft()
	b := nt.UpperLeft()

	return &LngLatBbox{
		West:  a.Lng,
		South: b.Lat,
		East:  b.Lng,
		North: a.Lat,
	}
}

func (t *Tile) UpperLeft() *LngLat {
	n := math.Pow(2.0, float64(t.Z))
	lonDeg := float64(t.X)/n*360.0 - 180.0
	latRad := math.Atan(math.Sinh(math.Pi * (1 - 2*float64(t.Y)/n)))
	latDeg := latRad * 180.0 / math.Pi
	lngLat := &LngLat{
		Lng: lonDeg,
		Lat: latDeg,
	}
	return lngLat
}

func (lngLat *LngLat) Truncate() {
	if lngLat.Lng > 180.0 {
		lngLat.Lng = 180.0
	} else if lngLat.Lng < -180.0 {
		lngLat.Lng = -180.0
	}

	if lngLat.Lat > 90.0 {
		lngLat.Lng = 90.0
	} else if lngLat.Lat < -90.0 {
		lngLat.Lat = -90.0
	}
}

func (lngLat *LngLat) XY(truncate bool) *XY {
	if truncate {
		lngLat.Truncate()
	}

	x := 6378137.0 * lngLat.Lng * math.Pi / 180.0
	var y float64

	if lngLat.Lat <= -90 {
		y = math.Inf(-1)
	} else if lngLat.Lat >= 90 {
		y = math.Inf(0)
	} else {
		latRad := lngLat.Lat * math.Pi / 180.0
		y = 6378137.0 * math.Log(math.Tan((math.Pi*0.25)+(0.5*latRad)))
	}

	xy := &XY{
		X: x,
		Y: y,
	}
	return xy
}

func (xy *XY) LngLat(truncate bool) *LngLat {
	r2d := 180.0 / math.Pi
	a := 6378137.0

	lngLat := &LngLat{
		Lng: xy.X * r2d / a,
		Lat: ((math.Pi * 0.5) - 2.0*math.Atan(math.Exp(-xy.Y/a))) * r2d,
	}

	if truncate {
		lngLat.Truncate()
	}

	return lngLat
}
