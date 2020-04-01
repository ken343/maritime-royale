package terrain

import (
	"fmt"

	"github.com/ken343/maritime-royale/pkg/screen"
	"github.com/veandco/go-sdl2/sdl"
)

type Terrain interface {
	Draw(renderer *sdl.Renderer, scale int, plrView screen.ViewPort)
}

func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {

	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("loading %v: %v", filename, err))
	}
	defer img.Free()

	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("creating texture from %v: %v", filename, err))
	}

	return tex
}
