package grid

/*
import (
	"github.com/JosephZoeller/maritime-royale/pkg/terrain"
	"github.com/JosephZoeller/maritime-royale/pkg/units"
	"github.com/veandco/go-sdl2/sdl"
)
*/
var gridData [][]Square

/*
func FakeInitGridDataDEMO(renderer *sdl.Renderer) {
	gridData = make([][]Square, 10)
	for i := range gridData {
		gridData[i] = make([]Square, 10)
	}

	for x := 0; x < len(gridData); x++ {
		for y := 0; y < len(gridData[x]); y++ {
			p := &gridData[x][y]
			pcoords := &p.Coords
			pcoords.XPos = x
			pcoords.YPos = y
			p.Terrain = terrain.NewWater(renderer, x, y)
		}
	}
	p := &gridData[4][5]
	p.Unit = units.NewDestroyer()
}
*/
func GetGridData() [][]Square {
	return gridData
}

func GetSquareOf(coords Coordinate) *Square {
	return &gridData[coords.XPos][coords.YPos]
}
