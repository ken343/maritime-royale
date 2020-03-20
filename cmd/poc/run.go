package main

import (
	"errors"
	"fmt"

	"github.com/JosephZoeller/maritime-royale/pkg/terrain"
	"github.com/JosephZoeller/maritime-royale/pkg/units"
	"github.com/JosephZoeller/maritime-royale/pkg/weather"
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

var mapData [10][10]Square

func fakeInitMapData() {
	for x := 0; x < len(mapData); x++ {
		for y := 0; y < len(mapData[x]); y++ {
			p := &mapData[x][y]
			pcoords := &p.coords
			pcoords.x = int32(x)
			pcoords.y = int32(y)
			p.terrain = terrain.NewWater()
		}
	}
	p := &mapData[4][5]
	p.unit = units.NewDestroyer()
}

func updateMap() {
	// JZ: retrieve any changes from server? the whole client/server relationship for this game is still a bit hazy for me tbh

}

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

	fakeInitMapData()

	running := true
	for running {
		// 0. retreive updated map
		updateMap()

		// 1. redraw background
		renderer.Clear()
		drawTerrain(renderer, waterTileTexture, winWidth, winHeight, waterTileSurface.W, waterTileSurface.H) // JZ: TODO a global slice of all clientTerrain types ought to have a surface loaded and a texture created for each type, so this drawing phase can just reference it as its updating the map
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

			// escape key exit case, selecting via keyboard
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
					fmt.Printf("{%d, %d}\n", eventType.X, eventType.Y)
					if eventType.Button == sdl.BUTTON_LEFT {
						mouseOnReleaseLeft(eventType.X, eventType.Y, renderer, waterTileSurface)
					} else {
						mouseOnReleaseRight()
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

func mouseOnReleaseLeft(mouseX, mouseY int32, renderer *sdl.Renderer, surface *sdl.Surface) {
	if mapHasSelection() { // if there is a selection: determine what is selected, then parse when the mouse release ought to do
		srcSq := getSelectionSquare()
		if srcSq.unit != nil {
			dstSq := getSquareOf(getCoordsAt(mouseX, mouseY, surface.W, surface.H))
			err := moveUnit(srcSq, dstSq)
			if err != nil {
				fmt.Println(err)
			} else { // for demonstration purposes. I believe at this point the user needs to commit their move and/or choose a wait/attack/capture action
				selectTileOf(dstSq.coords)
				arbitrateMoveOptions(dstSq)
			}
		}
	} else { // if there is no current selection: select tile of pixel
		selectTileOf(getCoordsAt(mouseX, mouseY, surface.W, surface.H))
		sq := getSelectionSquare()
		if sq.unit != nil {
			arbitrateMoveOptions(sq)
		}
	}
}

func moveUnit(src *Square, dst *Square) error {
	for _, movableCoords := range moveOptions {
		if dst.coords == movableCoords {
			dst.unit = src.unit
			src.unit = nil
			return nil
		}
	}

	return errors.New("Not a valid move.")
}

func mouseOnReleaseRight() {
	if mapHasSelection() { // if there is a selection: determine what is selected, then parse when the mouse release ought to do
		deselect()
		clearMoveOptions()
	}
}

type coordinate struct {
	x, y int32
}

var selection = coordinate{x: -1, y: -1}

func selectTileOf(tileCoords coordinate) {

	p := &selection
	p.x = tileCoords.x
	p.y = tileCoords.y
}

func getCoordsAt(pixelX, pixelY, tileWidth, tileHeight int32) coordinate {
	return coordinate{
		x: pixelX / tileWidth,
		y: pixelY / tileHeight,
	}
}

var moveOptions []coordinate

func arbitrateMoveOptions(sq *Square) {
	clearMoveOptions()
	moveOptions = append(moveOptions, // distance, pathing logic placeholder
		coordinate{x: sq.coords.x + 1, y: sq.coords.y},
		coordinate{x: sq.coords.x - 1, y: sq.coords.y},
		coordinate{x: sq.coords.x, y: sq.coords.y + 1},
		coordinate{x: sq.coords.x, y: sq.coords.y - 1})
}

func clearMoveOptions() {
	moveOptions = make([]coordinate, 0)
}

func deselect() {
	p := &selection
	p.x = -1
	p.y = -1
}

type Square struct {
	coords  coordinate
	terrain terrain.TerrainServer
	unit    units.UnitServer
	weather weather.WeatherServer
}

func getSelectionSquare() *Square {
	return getSquareOf(selection)
}

func getSquareOf(coords coordinate) *Square {
	return &mapData[coords.x][coords.y]
}

func mapHasSelection() bool {
	return selection.x >= 0 && selection.y >= 0
}

func drawselection(renderer *sdl.Renderer, selWidth, selHeight int32) {
	renderer.SetDrawColor(0x00, 0x00, 0x00, 0x00)

	dst := sdl.Rect{selection.x * selWidth, selection.y * selHeight, selWidth, selHeight}
	renderer.DrawRect(&dst)
}

func drawMoveOptions(renderer *sdl.Renderer, sqWidth, sqHeight int32) {
	if len(moveOptions) > 0 {
		renderer.SetDrawColor(0x00, 0xFF, 0x00, 0xAA)
		for _, squareCoord := range moveOptions {
			dst := sdl.Rect{squareCoord.x * sqWidth, squareCoord.y * sqHeight, sqWidth, sqHeight}
			renderer.DrawRect(&dst)
		}
	}
}

func drawTerrain(renderer *sdl.Renderer, texture *sdl.Texture, winWidth, winHeight, texWidth, texHeight int32) {
	for x, row := range mapData {
		for y, sq := range row {
			if sq.terrain != nil {
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
	for x, row := range mapData {
		for y, sq := range row {
			if sq.unit != nil {
				dst := sdl.Rect{int32(x) * texWidth, int32(y) * texHeight, texWidth, texHeight}
				renderer.Copy(texture, nil, &dst)
			}
		}
	}
}
