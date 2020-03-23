package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/JosephZoeller/maritime-royale/pkg/screen"
	"github.com/JosephZoeller/maritime-royale/pkg/terrain"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/JosephZoeller/maritime-royale/pkg/mrp"
)

var renderer *sdl.Renderer
var terrainData = []terrain.Terrain{}
var renderCreated = make(chan string)

const width int = 800
const height int = 800

func main() {
	go graphics()
	readMRP("localhost:8080")

}

func readMRP(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err.Error())
		log.Panic()
	}

	var carryOver []byte
	fmt.Println(<-renderCreated)
	for {
		var message = make([]byte, 0)
		var newMRP mrp.MRP
		var newMRPList []mrp.MRP

		message = carryOver

		for {
			var buf = make([]byte, 1024)
			conn.SetReadDeadline(time.Now().Add(10 * time.Second))
			_, err = conn.Read(buf)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			for k, v := range buf {
				if v == 0 {
					buf = buf[0:k]
					break
				}
			}

			message = append(message, buf...)
			messageString := strings.SplitAfter(string(message), "EOF")

			for _, v := range messageString {

				newMRP, err = mrp.ReadMRP([]byte(v))
				if err == nil {
					newMRPList = append(newMRPList, newMRP)
				} else {
					carryOver = message[len(message)-len(v):]
				}

			}

			if len(newMRPList) != 0 {
				break
			}

		}
		handleMRP(newMRPList)
	}
}

func handleMRP(newMRPList []mrp.MRP) {
	for _, mRPItem := range newMRPList {
		switch string(mRPItem.Request) {
		case "MAP":

			var tempTerrain map[string]interface{}
			json.Unmarshal(mRPItem.Body, &tempTerrain)
			switch tempTerrain["Type"] {
			case "island":
				terrainData =
					append(
						terrainData,
						terrain.NewIsland(renderer, int(tempTerrain["X"].(float64)), int(tempTerrain["Y"].(float64))),
					)
			}
		}
	}
}

func graphics() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}

	window, err :=
		sdl.CreateWindow(
			"Maritime Royale",
			sdl.WINDOWPOS_UNDEFINED,
			sdl.WINDOWPOS_UNDEFINED,
			int32(width),
			int32(height),
			sdl.WINDOW_OPENGL,
		)
	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}
	defer window.Destroy()

	renderer, err =
		sdl.CreateRenderer(
			window,
			-1,
			sdl.RENDERER_ACCELERATED,
		)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}
	defer renderer.Destroy()

	renderCreated <- "Renderer Created Successfully"

	plrView := screen.Screen{
		Xpos:   0,
		Ypos:   0,
		Width:  float64(width),
		Height: float64(height),
	}

	sdl.Delay(uint32(2000))

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		for _, terrainSquare := range terrainData {
			terrainSquare.Draw(renderer, 64, plrView)
		}

		renderer.Present()

	}
}
