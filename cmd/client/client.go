package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/JosephZoeller/maritime-royale/pkg/screen"
	"github.com/JosephZoeller/maritime-royale/pkg/terrain"
	"github.com/JosephZoeller/maritime-royale/pkg/units"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/JosephZoeller/maritime-royale/pkg/mrp"
)

var terrainData = []terrain.Terrain{}
var unitData = map[string]units.Unit{}

var renderer *sdl.Renderer
var renderCreated = make(chan string)

var plrView = screen.ViewPort{}

const width int = 800
const height int = 800

func main() {
	//dials connection to the gameplay server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err.Error())
		log.Panic()
	}

	go readMRP(conn)
	graphics(conn)

}

func readMRP(conn net.Conn) {
	var err error
	go returnPing(conn)

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
			//conn.SetReadDeadline(time.Now().Add(10 * time.Second))
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
		handleMRP(newMRPList, conn)
	}
}

func handleMRP(newMRPList []mrp.MRP, conn net.Conn) {
	//we begin by looping through each MRP
	for _, mRPItem := range newMRPList {

		//We check each kind of request so that we can handle
		//each one uniquely
		switch string(mRPItem.Request) {
		case "MAP":

			var tempTerrain map[string]interface{}
			json.Unmarshal(mRPItem.Body, &tempTerrain)
			switch tempTerrain["Type"] {
			case "water":
				water := terrain.NewWater(int(tempTerrain["X"].(float64)), int(tempTerrain["Y"].(float64)))
				terrainData =
					append(
						terrainData,
						&water,
					)
			}

		case "UNIT":

			var tempUnit map[string]interface{}
			json.Unmarshal(mRPItem.Body, &tempUnit)
			switch tempUnit["Type"] {
			case "destroyer":
				destroyer := units.NewDestroyer(int(tempUnit["X"].(float64)), int(tempUnit["Y"].(float64)))
				unitData[strconv.Itoa(int(tempUnit["X"].(float64)))+","+strconv.Itoa(int(tempUnit["Y"].(float64)))] = &destroyer
			}

		case "UNITC":
			unitData = map[string]units.Unit{}
			fmt.Println("clear")

		case "PING":
			delay, _ := strconv.Atoi(string(mRPItem.Body))
			delay = int(time.Now().UnixNano()/int64(time.Millisecond)) - delay
			fmt.Println("Ping:", delay, "ms")
			go returnPing(conn)
		}

	}
}

func returnPing(conn net.Conn) {
	time.Sleep(5 * time.Second)
	myMRP := mrp.NewMRP([]byte("PING"), []byte("this is a ping"), []byte("/"))
	conn.Write(mrp.MRPToByte(myMRP))

	//This needs to be in its own call at some point
	myMRP = mrp.NewMRP([]byte("MAP"), []byte("Gimme dat map"), []byte("/"))
	conn.Write(mrp.MRPToByte(myMRP))
}

func graphics(conn net.Conn) {
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

	plrView = screen.NewScreen(
		0,
		0,
		float64(width),
		float64(height),
	)

	sdl.Delay(uint32(2000))

	var isSelected string = ""

	for {

		//We set the background color and clear the screen,
		//of all previous graphics. This ensures a clean draw
		//every time
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		unitCommands := plrView.Update()

		for _, v := range unitCommands {
			if plrView.Mouse.State == 1 {
				if isSelected == "" && unitData[v] != nil {
					isSelected = v
				} else if isSelected != "" {
					myMRP, isPossible := unitData[isSelected].Move(v)
					if isPossible {
						conn.Write(mrp.MRPToByte(myMRP))
					}
					isSelected = ""
				}
			}
		}

		for _, terrainSquare := range terrainData {
			terrainSquare.Draw(renderer, int(plrView.Scale), plrView)
		}

		for _, unitSquare := range unitData {

			unitSquare.Draw(renderer, int(plrView.Scale), plrView)

		}

		//We take the renderer object and present it to the screen.
		renderer.Present()

		sdl.Delay(2)

	}
}
