package terrain

import "github.com/veandco/go-sdl2/sdl"

type water struct {
	Type string

	tex  *sdl.Texture
	X, Y int
}

func NewWaterServer() water {
	return water{Type: "water"}
}

func NewWater(renderer *sdl.Renderer, x int, y int) (w water) {
	w.tex = textureFromBMP(renderer, "../../assets/sprites/inDev/sprites/water2.bmp")

	w.X = x
	w.Y = y

	w.Type = "island"

	return w
}
func (s water) Draw(renderer *sdl.Renderer) {

}
