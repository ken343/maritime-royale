package terrain

import (
	"github.com/JosephZoeller/maritime-royale/pkg/graphics"
	"github.com/JosephZoeller/maritime-royale/pkg/screen"
	"github.com/veandco/go-sdl2/sdl"
)

type island struct {
	Type string

	tex  *sdl.Texture
	X, Y int
}

var islandTexture *sdl.Texture = nil

func NewIsland(x int, y int) (i island) {
	i.tex = nil

	i.X = x
	i.Y = y

	i.Type = "island"

	return i
}

func (s *island) Draw(renderer *sdl.Renderer, scale int, plrView screen.ViewPort) {
	if renderer == nil {
	} else if islandTexture == nil {
		s.tex = textureFromBMP(renderer, "../../assets/sprites/inDev/sprites/water2.bmp")
		islandTexture = s.tex
	} else {
		s.tex = islandTexture
	}
	graphics.DrawSquare(renderer, scale, s.X, s.Y, plrView, islandTexture)
}
