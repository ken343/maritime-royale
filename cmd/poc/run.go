package main

import (
	"fmt"

	"github.com/JosephZoeller/maritime-royale/pkg/grid"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	waterPOCPath     = "../../assets/sprites/inDev/sprites/waterPOC.bmp"
	destroyerPOCPath = "../../assets/sprites/inDev/sprites/destroyerPOC.bmp"
	waterPNGPath     = "../../assets/sprites/inDev/sprites/water2.png"
	destroyerPNGPath = "../../assets/sprites/inDev/sprites/dest1.png"
	exGridSize       = 10
	exTileSize       = 64
)

var winWidth, winHeight int32 = 640, 640
var winTitle string = "Go-SDL2 Render"


func run() int {
	window, _ := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	defer window.Destroy()

	renderer, _ := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	defer renderer.Destroy()

	waterTileSurface, _ := sdl.LoadBMP(waterPOCPath)
	defer waterTileSurface.Free()

	waterTileTexture, _ := renderer.CreateTextureFromSurface(waterTileSurface)
	defer waterTileTexture.Destroy()

	destroyerTileSurface, _ := sdl.LoadBMP(destroyerPOCPath)
	defer waterTileSurface.Free()

	destroyerTileTexture, _ := renderer.CreateTextureFromSurface(destroyerTileSurface)
	defer waterTileTexture.Destroy()

	grid.FakeInitGridDataDEMO(renderer)

	running := true
	for running {
		// 1. redraw background
		renderer.Clear()
		drawTerrain(renderer, waterTileTexture, winWidth, winHeight, waterTileSurface.W, waterTileSurface.H) 
		drawUnits(renderer, destroyerTileTexture, winWidth, winHeight, destroyerTileSurface.W, destroyerTileSurface.H)
		drawselection(renderer, waterTileSurface.W, waterTileSurface.H)
		drawMoveOptions(renderer, waterTileSurface.W, waterTileSurface.H)
		renderer.Present()

		// 2. poll user input
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			// any one of these ought to trigger an image refresh, no? hmm.
			switch eventType := event.(type) {

			// exit case
			case *sdl.QuitEvent:
				running = false

			// escape key exit case / selecting via keyboard
			case *sdl.KeyboardEvent:
				if eventType.Keysym.Scancode == sdl.SCANCODE_ESCAPE {
					running = false
				} else if eventType.State == sdl.RELEASED {
					keyCode := string(eventType.Keysym.Sym)
					switch keyCode {
					default:
						fmt.Println(keyCode)
					}
				}

			// selecting via mouse
			case *sdl.MouseButtonEvent:
				if eventType.State == sdl.RELEASED {
					fmt.Printf("Mouse Release Event @ {%d, %d}\n", eventType.X, eventType.Y)
					if eventType.Button == sdl.BUTTON_LEFT {
						onMouseLeftRelease(eventType.X, eventType.Y, renderer, waterTileSurface)
					} else if eventType.Button == sdl.BUTTON_RIGHT {
						onMouseRightRelease()
					}
				}

			// zoom in / out
			case *sdl.MouseWheelEvent:
				if eventType.Y > 0 {
					fmt.Println("Scroll Up:", eventType.Y)
				} else {
					fmt.Println("Scroll Down:", eventType.Y)
				}
			}
		}

		sdl.Delay(16)
	}

	return 0
}

func onMouseLeftRelease(mouseX, mouseY int32, renderer *sdl.Renderer, surface *sdl.Surface) {
	if grid.ExistsSelection() {
		srcSq := grid.GetSelection()
		if srcSq.Unit != nil {
			dstSq := grid.GetSquareOf(grid.GetCoordsAt(mouseX, mouseY, surface.W, surface.H))
			if grid.MoveUnit(srcSq, dstSq) == nil {
				grid.SetSelection(dstSq)   // FOR DEMO PURPOSES (allows for immediate selection)
				grid.SetMoveOptions(dstSq) // FOR DEMO PURPOSES (allows for immediate movement)
			}
		}
	} else {
		sq := grid.GetSquareOf(grid.GetCoordsAt(mouseX, mouseY, surface.W, surface.H))
		grid.SetSelection(sq)
		if sq.Unit != nil {
			grid.SetMoveOptions(sq)
		}
	}
}

func onMouseRightRelease() {
	if grid.ExistsSelection() {
		grid.UnsetSelection()
		grid.UnsetMoveOptions()
	}
}

func drawselection(renderer *sdl.Renderer, selWidth, selHeight int32) {
	renderer.SetDrawColor(0x00, 0x00, 0x00, 0x00)

	dst := sdl.Rect{int32(grid.GetSelection().Coords.XPos) * selWidth, int32(grid.GetSelection().Coords.YPos) * selHeight, selWidth, selHeight}
	renderer.DrawRect(&dst)
}

func drawMoveOptions(renderer *sdl.Renderer, sqWidth, sqHeight int32) {
	if grid.CountMoveOptions() > 0 {
		renderer.SetDrawColor(0x00, 0xFF, 0x00, 0xAA)
		for _, squareCoord := range grid.GetMoveOptions() {
			dst := sdl.Rect{int32(squareCoord.XPos) * sqWidth, int32(squareCoord.YPos) * sqHeight, sqWidth, sqHeight}
			renderer.DrawRect(&dst)
		}
	}
}

func drawTerrain(renderer *sdl.Renderer, texture *sdl.Texture, winWidth, winHeight, texWidth, texHeight int32) {
	for x, row := range grid.GetGridData() {
		for y, sq := range row {
			if sq.Terrain != nil {
				dst := sdl.Rect{int32(x) * texWidth, int32(y) * texHeight, texWidth, texHeight}
				renderer.Copy(texture, nil, &dst)
			}
		}
	}

	/* pseudocode for future reference
	for x, row := range mapData {
		for y, sq := range row {
			dst := sdl.Rect{int32(x) * sq.terrainsurface.W, int32(y) * sq.terrainsurface.H,  sq.terrainsurface.W,  sq.terrainsurface.H}
			renderer.Copy(sq.terraintexture, nil, &dst)
		}
	}
	*/
}

func drawUnits(renderer *sdl.Renderer, texture *sdl.Texture, winWidth, winHeight, texWidth, texHeight int32) {
	for x, row := range grid.GetGridData() {
		for y, sq := range row {
			if sq.Unit != nil {
				dst := sdl.Rect{int32(x) * texWidth, int32(y) * texHeight, texWidth, texHeight}
				renderer.Copy(texture, nil, &dst)
			}
		}
	}
}

// END DRAW LOGIC
