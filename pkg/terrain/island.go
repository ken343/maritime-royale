package terrain

import "github.com/veandco/go-sdl2/sdl"

type island struct {
	Type string

	tex  *sdl.Texture
	X, Y int
}

var texture *sdl.Texture = nil

func NewIslandServer() (i island) {
	return island{Type: "island"}
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

const playerSize = 64

func (s island) Draw(renderer *sdl.Renderer) {
	// Converting player coordinates to top left of sprite
	x := s.X
	y := s.Y

	renderer.Copy(s.tex,
		&sdl.Rect{X: 0, Y: 0, W: playerSize, H: playerSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: playerSize, H: playerSize})
}
