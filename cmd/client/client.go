package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/JosephZoeller/maritime-royale/pkg/square"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/JosephZoeller/maritime-royale/pkg/mrp"
)

func main() {
	go readMRP("localhost:8080")
	graphics()
}

func readMRP(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err.Error())
		log.Panic()
	}

	var carryOver []byte

	for {
		var message = make([]byte, 0)
		var newMRP mrp.MRP
		var newMRPList []mrp.MRP

		message = carryOver

		for {
			var buf = make([]byte, 1024)
			conn.Read(buf)
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
		if string(mRPItem.Request) == "MAP" {
			var genericSquare = square.SquareGeneric{}
			json.Unmarshal(mRPItem.Body, &genericSquare)
			typeMap := (genericSquare.Terrain.(map[string]interface{}))
			typeStr := fmt.Sprintf("%v", typeMap["Type"])
			switch typeStr {
			case "water":

			case "island":

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

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}
	defer renderer.Destroy()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		//draw to renderer here

		renderer.Present()
	}
}
