package main

import (
	"time"

	"github.com/jtheiss19/project-undying/pkg/elements/firstOrder"
	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder"
	"github.com/jtheiss19/project-undying/pkg/gamemap"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
	"github.com/jtheiss19/project-undying/pkg/networking/server"
)

const tps = 60

func init() {
	firstOrder.Init()
	secondOrder.Init()
}

func main() {
	go server.Server("8080")

	gamemap.NewWorld()

	var timeSinceLastUpdate int64
	var now int64
	count := 0
	for {

		time.Sleep((1000/tps - time.Duration(timeSinceLastUpdate)) * time.Millisecond)
		if count == 60 {
			//fmt.Println(timeSinceLastUpdate)
			count = 0
		}
		count++
		now = time.Now().UnixNano()

		world := gamestate.GetEntireWorld()

		for _, elem := range world {
			if elem.Active {
				elem.UpdateServer()
				//if err != nil {
				//fmt.Println("updating element:", err)
				//return
				//}
			}
		}

		timeSinceLastUpdate = (time.Now().UnixNano() - now) / 1000000
	}
}
