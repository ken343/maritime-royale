package main

import (
	"fmt"
	_ "image/png"
	"time"

	"github.com/ken343/maritime-royale/pkg/elements/secondOrder"
	"github.com/ken343/maritime-royale/pkg/gamemap"
	"github.com/ken343/maritime-royale/pkg/gamestate"
	"github.com/ken343/maritime-royale/pkg/networking/server"
)

const tps = 60

func init() {
	secondOrder.Init()
}

func main() {
	go server.Server("8080")

	gamemap.NewWorld()

	var timeSinceLastUpdate int64
	for {

		time.Sleep((1000/tps - time.Duration(timeSinceLastUpdate)) * time.Millisecond)
		now := time.Now().UnixNano()

		world := gamestate.GetWorld()

		for _, elem := range world {
			if elem.Active {
				err := elem.UpdateServer(world)
				if err != nil {
					fmt.Println("Error UpdateServer(world):", err)
					return
				}
			}
		}

		timeSinceLastUpdate = (time.Now().UnixNano() - now) / 1000000
	}
}
