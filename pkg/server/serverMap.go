package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/ken343/maritime-royale/pkg/mrp"
	"github.com/ken343/maritime-royale/pkg/units"

	"github.com/ken343/maritime-royale/pkg/terrain"
	_ "github.com/ken343/maritime-royale/pkg/tile"
)

//Square contains all the data about a specific square
// These values are now given in game package.

var terrainData = map[string]terrain.Terrain{}
var unitData = map[string]units.Unit{}

const MAPX, MAPY = 50, 50

func init() {
	for x := 0; x < MAPX; x++ {
		for y := 0; y < MAPY; y++ {
			terrainSquare := terrain.NewWater(x, y)
			terrainData[strconv.Itoa(x)+","+strconv.Itoa(y)] = &terrainSquare
		}
	}
	unitSquare := units.NewDestroyer(5, 8)
	unitData["5,8"] = &unitSquare
}

func readMRP(conn net.Conn) {

	var carryOver []byte
	var err error

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
		case "PING":

			myMRP := mrp.NewMRP(
				[]byte("PING"),
				[]byte(strconv.Itoa(int(time.Now().UnixNano()/int64(time.Millisecond)))),
				[]byte("/"),
			)
			conn.Write(mrp.ToByte(myMRP))

		case "MAP":

			sendMap(conn)

		case "UNIT":

			if unitData[string(mRPItem.Body)] != nil {
				_, isPossible := unitData[string(mRPItem.Body)].Move(string(mRPItem.Footers[0]))
				if isPossible {
					unitData[string(mRPItem.Footers[0])] = unitData[string(mRPItem.Body)]
					unitData[string(mRPItem.Body)] = nil
				}
			}

		}
	}
}

func sendMap(conn net.Conn) {

	for _, v := range terrainData {

		body, err := json.Marshal(v)
		if err != nil {
			fmt.Println(err.Error())
			log.Panic()
		}

		var sendingMRP = mrp.NewMRP([]byte("MAP"), body, []byte("/"))

		packet := mrp.ToByte(sendingMRP)

		conn.Write(packet)
	}

	var sendingMRP = mrp.NewMRP([]byte("UNITC"), []byte("clear"), []byte("/"))
	packet := mrp.ToByte(sendingMRP)
	conn.Write(packet)

	for _, v := range unitData {
		body, err := json.Marshal(v)
		if err != nil {
			fmt.Println(err.Error())
			log.Panic()
		}

		var sendingMRP = mrp.NewMRP([]byte("UNIT"), body, []byte("/"))

		packet := mrp.ToByte(sendingMRP)

		conn.Write(packet)
	}
}
