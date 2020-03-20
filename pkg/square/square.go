package square

import (
	"github.com/JosephZoeller/maritime-royale/pkg/terrain"
	"github.com/JosephZoeller/maritime-royale/pkg/units"
	"github.com/JosephZoeller/maritime-royale/pkg/weather"
)

type Square struct {
	XPos, YPos int
	Terrain    terrain.Terrain
	Unit       units.Unit
	Weather    weather.Weather
}

type SquareGeneric struct {
	XPos, YPos int
	Terrain    interface{}
	Unit       interface{}
	Weather    interface{}
}
