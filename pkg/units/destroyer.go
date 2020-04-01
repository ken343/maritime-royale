package units

import (
	"strconv"
	"strings"

	"github.com/ken343/maritime-royale/pkg/graphics"
	"github.com/ken343/maritime-royale/pkg/mrp"
	"github.com/ken343/maritime-royale/pkg/screen"
	"github.com/veandco/go-sdl2/sdl"
)

type destroyer struct {
	Type string

	tex  *sdl.Texture
	X, Y int
}

var destroyerTexture *sdl.Texture = nil

// NewDestroyer creates a new destory at the indicated coordinates.
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

func (s *destroyer) Move(dest string) (mrp.MRP, bool) {

	pos := strings.Split(dest, ",")
	xPos, _ := strconv.Atoi(pos[0])
	yPos, _ := strconv.Atoi(pos[1])

	myMRP := mrp.NewMRP([]byte("UNIT"), []byte(strconv.Itoa(s.X)+","+strconv.Itoa(s.Y)), []byte(dest))

	s.X = xPos
	s.Y = yPos

	return myMRP, true
}
