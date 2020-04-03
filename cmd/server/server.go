package main

import (
	"fmt"
	"time"

	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder"
	"github.com/jtheiss19/project-undying/pkg/gamemap"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
	"github.com/jtheiss19/project-undying/pkg/networking/server"
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
					fmt.Println("updating element:", err)
					return
				}
			}
		}

		timeSinceLastUpdate = (time.Now().UnixNano() - now) / 1000000
	}
}
