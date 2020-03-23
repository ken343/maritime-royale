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
	} else {
		i.tex = texture
	}
	i.X = x
	i.Y = y

	i.Type = "island"

	return i
}

func (s island) Draw(renderer *sdl.Renderer, scale int, plrView screen.Screen) {
	// Converting player coordinates to top left of sprite
	x := s.X
	y := s.Y

	if plrView.Xpos > float64(x) || float64(x) > plrView.Width+plrView.Xpos {
		return
	}

	if plrView.Ypos > float64(y) || float64(y) > plrView.Height+plrView.Ypos {
		return
	}

	renderer.Copy(s.tex,
		&sdl.Rect{X: 0, Y: 0, W: int32(scale), H: int32(scale)},
		&sdl.Rect{X: int32(x) * int32(scale), Y: int32(y) * int32(scale), W: int32(scale), H: int32(scale)})
}
