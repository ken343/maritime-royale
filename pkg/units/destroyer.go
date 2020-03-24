package units

import (
	"github.com/JosephZoeller/maritime-royale/pkg/graphics"
	"github.com/JosephZoeller/maritime-royale/pkg/screen"
	"github.com/veandco/go-sdl2/sdl"
)

type destroyer struct {
	Type string

	tex  *sdl.Texture
	X, Y int
}

var destroyerTexture *sdl.Texture = nil

func NewDestroyer(x int, y int) (s destroyer) {
	s.tex = nil

	s.X = x
	s.Y = y

	s.Type = "destroyer"

	return s
}

func (s *destroyer) Draw(renderer *sdl.Renderer, scale int, plrView screen.ViewPort) {
	if renderer == nil {
	} else if destroyerTexture == nil {
		s.tex = textureFromBMP(renderer, "../../assets/sprites/inDev/sprites/destroyerPOC.bmp")
		destroyerTexture = s.tex
	} else {
		s.tex = destroyerTexture
	}
	graphics.DrawSquare(renderer, scale, s.X, s.Y, plrView, destroyerTexture)
}
