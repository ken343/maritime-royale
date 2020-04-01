package graphics

import (
	"github.com/ken343/maritime-royale/pkg/screen"
	"github.com/veandco/go-sdl2/sdl"
)

// DrawSquare does what it says on the tin. Will need to refactor for Ebiten Game Engine
func DrawSquare(renderer *sdl.Renderer, scale int, x int, y int, plrView screen.ViewPort, tex *sdl.Texture) {

	if plrView.Xpos/float64(scale) > float64(x)+1 || float64(x)-1 > (plrView.Width+plrView.Xpos)/float64(scale) {
		return
	}

	if plrView.Ypos/float64(scale) > float64(y)+1 || float64(y)-1 > (plrView.Height+plrView.Ypos)/float64(scale) {
		return
	}

	renderer.Copy(tex,
		&sdl.Rect{X: 0, Y: 0, W: int32(64), H: int32(64)},
		&sdl.Rect{X: int32(x)*int32(scale) - int32(plrView.Xpos), Y: int32(y)*int32(scale) - int32(plrView.Ypos), W: int32(scale), H: int32(scale)})
}
