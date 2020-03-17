package server

import (
	"fmt"
	"log"
	"net"

	"github.com/JosephZoeller/maritime-royale/pkg/objects"
)

var mapData = map[int]map[int]objects.Square{}

const MAPX, MAPY = 50, 50

func init() {
	for x := 0; x < MAPX; x++ {
		var temp = map[int]objects.Square{}
		for y := 0; y < MAPY; y++ {
			if ((x*50)+y)%2 == 0 {
				temp[y] =
					objects.Square{
						XPos:    x,
						YPos:    y,
						Content: objects.NewWater()}
			} else {
				temp[y] =
					objects.Square{
						XPos:    x,
						YPos:    y,
						Content: objects.NewIsland()}
			}
		}
		mapData[x] = temp
	}
}

//Server starts a server on the selected port
func Server(port string) {

	server, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err.Error())
		log.Panic()
	}

	newConnSignal := make(chan string)

	for {
		go session(server, newConnSignal)
		fmt.Println(<-newConnSignal)
	}

}

func session(l net.Listener, newConnSignal chan string) {
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err.Error())
		log.Panic()
	}

	newConnSignal <- "New Connection Made"

	closeConnection := make(chan string)

	go sendMap(conn)

	fmt.Println(<-closeConnection)
}

func sendMap(conn net.Conn) {
	var sentMap = ""
	for _, vx := range mapData {
		var line = ""
		for _, vy := range vx {
			line = line + vy.Content.OnDraw()
		}
		sentMap = sentMap + line + "\n"
	}

	conn.Write([]byte(sentMap))
}
