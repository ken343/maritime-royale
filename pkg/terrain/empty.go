package terrain

import "github.com/veandco/go-sdl2/sdl"

type empty struct {
}

func NewEmpty() (e empty) {
	return empty{}
}

func (s empty) Draw(renderer *sdl.Renderer) {

}
