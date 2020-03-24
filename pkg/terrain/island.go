package terrain

import (
	"github.com/JosephZoeller/maritime-royale/pkg/screen"
	"github.com/veandco/go-sdl2/sdl"
)

type island struct {
	Type string

	tex  *sdl.Texture
	X, Y int
}

var texture *sdl.Texture = nil

func NewIslandServer(x int, y int) (i island) {
	return island{Type: "island", X: x, Y: y}
}

func NewIsland(renderer *sdl.Renderer, x int, y int) (i island) {
	if texture == nil {
		i.tex = textureFromBMP(renderer, "../../assets/sprites/inDev/sprites/water2.bmp")
		texture = i.tex
	} else {
		i.tex = texture
	}
	i.X = x
	i.Y = y

	i.Type = "island"

	return i
}

func (s *island) Draw(renderer *sdl.Renderer, scale int, plrView screen.ViewPort) {
	// Converting player coordinates to top left of sprite
	x := s.X
	y := s.Y

	if plrView.Xpos/float64(scale) > float64(x)+1 || float64(x)-1 > (plrView.Width-plrView.Xpos)/float64(scale) {
		return
	}

	if plrView.Ypos/float64(scale) > float64(y)+1 || float64(y)-1 > (plrView.Height-plrView.Ypos)/float64(scale) {
		return
	}

	renderer.Copy(s.tex,
		&sdl.Rect{X: 0, Y: 0, W: int32(scale), H: int32(scale)},
		&sdl.Rect{X: int32(x)*int32(scale) + int32(plrView.Xpos), Y: int32(y)*int32(scale) + int32(plrView.Ypos), W: int32(scale), H: int32(scale)})
}
