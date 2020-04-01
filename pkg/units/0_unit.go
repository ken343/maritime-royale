package units

import (
	"fmt"

	"github.com/ken343/maritime-royale/pkg/mrp"
	"github.com/ken343/maritime-royale/pkg/screen"
	"github.com/veandco/go-sdl2/sdl"
)

// Unit defines the capabilies of the various ingame units.
type Unit interface {
	Draw(renderer *sdl.Renderer, scale int, plrView screen.ViewPort)
	Move(dest string) (mrp.MRP, bool)
	// New(x int, y int, type string) Unit {type will infer facory constructor}
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
