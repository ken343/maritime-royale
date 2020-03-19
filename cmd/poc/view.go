package main

import (
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	waterPNGPath     = "../../assets/sprites/inDev/sprites/water2.png"
	destroyerPNGPath = "../../assets/sprites/inDev/sprites/dest1.png"
	exGridSize       = 10
	exTileSize       = 64 // could likely gather this information from the file itself without hardcoding
)

func run() error {
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		return err
	}
	defer sdl.Quit()

	// Create a window for us to draw the images on
	window, err := sdl.CreateWindow("Loading test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, exGridSize*exTileSize, exGridSize*exTileSize, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		return err
	}
	// Load a PNG image
	waterImg, err := img.Load(waterPNGPath)
	if err != nil {
		return err
	}
	defer waterImg.Free()

	shipImg, err := img.Load(destroyerPNGPath)
	if err != nil {
		return err
	}
	defer shipImg.Free()

	for i := int32(0); i < exGridSize; i++ {
		for j := int32(0); j < exGridSize; j++ {
			waterImg.BlitScaled(nil, surface, &sdl.Rect{X: i * exTileSize, Y: j * exTileSize, W: exTileSize, H: exTileSize})
		}
	}
	shipImg.BlitScaled(nil, surface, &sdl.Rect{X: 52, Y: 52, W: exTileSize, H: exTileSize}) // cute

	// Update the window surface with what we have drawn
	window.UpdateSurface()

	// Run infinite loop until user closes the window
	running := true
	s := int32(0)
	for running {

		t1 := time.Now().UnixNano()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}

		for i := int32(0); i < exGridSize; i++ {
			for j := int32(0); j < exGridSize; j++ {
				waterImg.BlitScaled(nil, surface, &sdl.Rect{X: i * exTileSize, Y: j * exTileSize, W: exTileSize, H: exTileSize})
			}
		}
		shipImg.BlitScaled(nil, surface, &sdl.Rect{X: s, Y: s, W: exTileSize, H: exTileSize}) // cute

		// Update the window surface with what we have drawn
		window.UpdateSurface()

		speed := 10
		s = s + int32(speed)

		timeFinal := time.Now().UnixNano() - t1
		delay := 16 - (uint32(timeFinal) / 1000000)
		if delay > 16 {
			delay = 0
		}
		sdl.Delay(delay)
	}

	return nil
}
