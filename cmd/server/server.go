package main

import (
	"fmt"
	"time"

	"github.com/jtheiss19/project-undying/pkg/elements/playerControl"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
	"github.com/jtheiss19/project-undying/pkg/server"
)

const tps = 60

func main() {
	go server.Server("8080")

	gamestate.NewWorld()

	var timeSinceLastUpdate int64
	for {

		time.Sleep((1000/tps - time.Duration(timeSinceLastUpdate)) * time.Millisecond)
		now := time.Now().UnixNano()

		for k, elem := range gamestate.GetWorld() {
			if elem.Active {
				err := elem.Update()
				if err != nil {
					fmt.Println("updating element:", err)
					return
				}
			}
			if k == len(gamestate.GetWorld())-1 {
				trc := new(playerControl.Tracker)
				myComp := elem.GetComponent(trc)
				if myComp != nil {
					if myComp.(*playerControl.Tracker).GetContainer() == elem {
						fmt.Println("true")
					} else {
						fmt.Println("false")
					}
				}
				fmt.Println(elem)

			}
		}

		timeSinceLastUpdate = (time.Now().UnixNano() - now) / 1000000
	}
}
