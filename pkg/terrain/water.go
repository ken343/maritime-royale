package terrain

import (
	"github.com/JosephZoeller/maritime-royale/pkg/graphics"
	"github.com/JosephZoeller/maritime-royale/pkg/screen"
	"github.com/veandco/go-sdl2/sdl"
)

type water struct {
	Type string

	tex  *sdl.Texture
	X, Y int
}

var waterTexture *sdl.Texture = nil

func NewWater(x int, y int) (w water) {
	w.tex = nil

	w.X = x
	w.Y = y

	w.Type = "water"

	return w
}

func (s *water) Draw(renderer *sdl.Renderer, scale int, plrView screen.ViewPort) {
	if renderer == nil {
	} else if waterTexture == nil {
		s.tex = textureFromBMP(renderer, "../../assets/sprites/inDev/sprites/water2.bmp")
		waterTexture = s.tex
	} else {
		s.tex = waterTexture
	}
	graphics.DrawSquare(renderer, scale, s.X, s.Y, plrView, waterTexture)
}
