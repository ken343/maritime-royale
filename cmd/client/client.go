package main

import (
	_ "image/png"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/ken343/maritime-royale/pkg/elements/secondOrder"
	"github.com/ken343/maritime-royale/pkg/gameloop"
	"github.com/ken343/maritime-royale/pkg/gamemap"
	"github.com/ken343/maritime-royale/pkg/gamestate"
)

const (
	screenWidth  = 1280
	screenHeight = 720
	screenScale  = 1
	addr         = "localhost:8080" //set to "" to launch game in single player
)

func init() {
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
