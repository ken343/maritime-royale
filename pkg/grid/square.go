package grid

import (
	"github.com/JosephZoeller/maritime-royale/pkg/terrain"
	"github.com/JosephZoeller/maritime-royale/pkg/units"
	"github.com/JosephZoeller/maritime-royale/pkg/weather"
)

type Square struct {
	Coords  Coordinate
	Terrain terrain.Terrain
	Unit    units.Unit
	Weather weather.Weather
}

type SquareGeneric struct {
	Coords  Coordinate
	Terrain interface{}
	Unit    interface{}
	Weather interface{}
}
