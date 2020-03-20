package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/JosephZoeller/maritime-royale/pkg/square"
	"github.com/JosephZoeller/maritime-royale/pkg/terrain"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/JosephZoeller/maritime-royale/pkg/mrp"
)

var renderer *sdl.Renderer
var mapData = map[int]map[int]square.Square{}
var renderCreated = make(chan string)

const maxMapX int = 100
const maxMapY int = 100

func init() {
	for x := 0; x < maxMapX; x++ {
		temp := map[int]square.Square{}
		for y := 0; y < maxMapY; y++ {
			temp[y] = square.Square{
				Terrain: terrain.NewEmpty()}
		}
		mapData[x] = temp
	}
}

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

			var genericSquare = square.SquareGeneric{}

			json.Unmarshal(mRPItem.Body, &genericSquare)

			typeMap := (genericSquare.Terrain.(map[string]interface{}))

			typeStr := fmt.Sprintf("%v", typeMap["Type"])

			switch typeStr {
			case "island":
				mapData[genericSquare.XPos][genericSquare.YPos] = square.Square{
					XPos:    genericSquare.XPos,
					YPos:    genericSquare.YPos,
					Terrain: terrain.NewIsland(renderer, genericSquare.XPos*64, genericSquare.YPos*64)}
			}
		}
	}
}

func graphics() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Maritime Royale",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 800,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}
	defer renderer.Destroy()

	renderCreated <- "Renderer Created Successfully"

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

		for _, line := range mapData {
			for _, squareValue := range line {
				squareValue.Terrain.Draw(renderer)
			}
		}

		renderer.Present()

	}
}
