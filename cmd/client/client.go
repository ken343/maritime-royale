package main

import (
	_ "image/png"
	"log"
	"time"

	"github.com/ken343/maritime-royale/pkg/elements/firstOrder"
	"github.com/ken343/maritime-royale/pkg/elements/objects"
	"github.com/ken343/maritime-royale/pkg/networking/connection"

	"github.com/hajimehoshi/ebiten"
	"github.com/ken343/maritime-royale/pkg/elements/secondOrder"
	"github.com/ken343/maritime-royale/pkg/gameloop"
	"github.com/ken343/maritime-royale/pkg/gamemap"
	"github.com/ken343/maritime-royale/pkg/gamestate"
	_ "github.com/ken343/maritime-royale/statik"
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
		gamestate.Dial(addr)
	} else {
		gamemap.NewWorld()
		connection.SetID("0")
		newPlayer := objects.NewPlayer(nil)
		newPlayer.ID = "0"
		newPlayer.UniqueName = newPlayer.UniqueName + newPlayer.ID
		gamestate.AddUnitToWorld(newPlayer)
		gamestate.PushChunks()
		gameloop.IsServer = true
	}

	time.Sleep(1 * time.Second)

	gameloop.MakeScreen()

	if err := ebiten.Run(gameloop.Update, screenWidth/screenScale, screenHeight/screenScale, screenScale, "test"); err != nil {
		log.Fatal(err)
	}
}
