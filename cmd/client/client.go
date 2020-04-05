package main

import (
	"log"
	"time"

	"github.com/jtheiss19/project-undying/pkg/elements/firstOrder"

	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder"
	"github.com/jtheiss19/project-undying/pkg/gameloop"
	"github.com/jtheiss19/project-undying/pkg/gamemap"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
)

const (
	screenWidth  = 1280
	screenHeight = 720
	screenScale  = 1
	addr         = "localhost:8080" //set to "" to launch game in single player
)

func init() {
	firstOrder.Init()
	secondOrder.Init()
}

func main() {

	if addr != "" {
		gamestate.Dial("localhost:8080")
	} else {
		gamemap.NewWorld()
	}

	time.Sleep(1 * time.Second)

	gameloop.MakeScreen()

	if err := ebiten.Run(gameloop.Update, screenWidth/screenScale, screenHeight/screenScale, screenScale, "test"); err != nil {
		log.Fatal(err)
	}
}
