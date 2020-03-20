package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	userwindow "github.com/JosephZoeller/maritime-royale/pkg/userWindow"

	"github.com/JosephZoeller/maritime-royale/pkg/square"
	"github.com/JosephZoeller/maritime-royale/pkg/terrain"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/JosephZoeller/maritime-royale/pkg/mrp"
)

var mapData = map[int]map[int]square.Square{}

var renderer *sdl.Renderer
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
			//Since we now it is map data being sent, we create
			//a generic square such that it can hold the json
			//data.
			var genericSquare = square.SquareGeneric{}

			json.Unmarshal(mRPItem.Body, &genericSquare)

			//We search for the type field in the terrain to know
			//what kind of terrain to create.
			typeMap := (genericSquare.Terrain.(map[string]interface{}))
			typeStr := fmt.Sprintf("%v", typeMap["Type"])

			//We loop through each available kind of terrain and
			//definition here.
			switch typeStr {
			case "island":
				mapData[genericSquare.XPos][genericSquare.YPos] = square.Square{
					XPos:    genericSquare.XPos,
					YPos:    genericSquare.YPos,
					Terrain: terrain.NewIsland(renderer, genericSquare.XPos*64, genericSquare.YPos*64),
				}
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

	//Create a window
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

	//This struct represents what the player can activly
	//see through the window
	plyView := userwindow.Window{
		Xpos:   0,
		Ypos:   0,
		Width:  800,
		Height: 800,
	}

	//Create a renderer object to handle our textures
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}
	defer renderer.Destroy()

	//Tell the program the renderer is now functional
	renderCreated <- "Renderer Created Successfully"

	//Initial env delay to be ran once. Ensures startup
	//on slow machines.
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

		//We go through each tile of the map and draw each one
		//In the future this needs to be edited so that only
		//the elements the user can currently see are drawn.
		for x := int(plyView.Xpos / 64); x < (plyView.Width+int(plyView.Xpos))/64; x++ {
			for y := int(plyView.Ypos / 64); y < (plyView.Height+int(plyView.Ypos))/64; y++ {
				mapData[x][y].Terrain.Draw(renderer)
			}
		}

		//We take the renderer object and present it to the screen.
		renderer.Present()

	}
}
