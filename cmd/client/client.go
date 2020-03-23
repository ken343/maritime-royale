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
	//dials connection to the gameplay server
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err.Error())
		log.Panic()
	}

	var carryOver []byte

	//waits for a render object to be created and assigned
	fmt.Println(<-renderCreated)
	for {
		var message = make([]byte, 0)
		var newMRP mrp.MRP
		var newMRPList []mrp.MRP

		//if any message was not complete during the last pull
		//from buffer, carryOver stores it. After nil assignment
		//from above, carryOver fills in the beggining couple lines
		message = carryOver

		for {
			var buf = make([]byte, 1024)
			conn.SetReadDeadline(time.Now().Add(10 * time.Second))
			_, err = conn.Read(buf)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			//This is a simple loop used to remove all null chars
			//from the buffer so that they aren't accidentally read
			for k, v := range buf {
				if v == 0 {
					buf = buf[0:k]
					break
				}
			}

			//we append the buffer to a message because it allows for
			//us to pull a larger message if one packet of 1024 bytes
			//was not enough. In that case buffer get overwrittin but
			//message would not change.
			message = append(message, buf...)
			//Here we are checking for multiple MRP's in the buffer,
			//This is useful incase we pull many small MRP's in one,
			//1024 byte buffer.
			messageString := strings.SplitAfter(string(message), "EOF")

			for _, v := range messageString {

				//We just check the segment to see if can be recongnised
				//as a MRP packet. If so we add it to the growing list of
				//MRP's and move on, else we assume the message is incomplete
				//and add it to the carryOver message.
				newMRP, err = mrp.ReadMRP([]byte(v))
				if err == nil {
					newMRPList = append(newMRPList, newMRP)
				} else {
					carryOver = message[len(message)-len(v):]
				}

			}

			//If we have found any MRP's during the above loop
			//we are going to break out and begin processing them
			//before moving back into the list and continue pulling
			if len(newMRPList) != 0 {
				break
			}

		}
		//This handling function is what analyzes the MRP
		//and decides what do.
		handleMRP(newMRPList)
	}
}

func handleMRP(newMRPList []mrp.MRP) {
	//we begin by looping through each MRP
	for _, mRPItem := range newMRPList {

		//We check each kind of request so that we can handle
		//each one uniquely
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
	//Init the enviroment, drivers, graphics.
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

	//Tell the program the renderer is now functional
	renderCreated <- "Renderer Created Successfully"

	plrView := screen.Screen{
		Xpos:   0,
		Ypos:   0,
		Width:  float64(width),
		Height: float64(height),
	}

	sdl.Delay(uint32(2000))

	for {
		//The even poller checks all events, here we are checking for
		//the program to be closed
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		//We set the background color and clear the screen,
		//of all previous graphics. This ensures a clean draw
		//every time
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		for _, terrainSquare := range terrainData {
			terrainSquare.Draw(renderer, 64, plrView)
		}

		//We take the renderer object and present it to the screen.
		renderer.Present()

	}
}
