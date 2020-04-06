package gameloop

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/elements/objects"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
)

var myScreen *elements.Element
var IsServer = false

func MakeScreen() {
	myScreen = objects.NewScreen(-1280/2, -720/2)
}

//Update is the mainloop designed to be passed into an
//ebiten run function. It is called every tick and thus
//every frame. This is what controls game logic and rendering.
func Update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	world := gamestate.GetEntireWorld()

	myScreen.Update(-myScreen.XPos, -myScreen.YPos)
	myScreen.Draw(screen, 0, 0)

	tileCount := 0
	for _, elem := range world {
		if elem.Active && canView(elem, screen) {

			err := elem.Update(-myScreen.XPos, -myScreen.YPos)
			if err != nil {
				fmt.Println("updating element:", err)
			}

			go elem.UpdateServer()

			err = elem.Draw(screen, -myScreen.XPos, -myScreen.YPos)
			if err != nil {
				fmt.Println("drawing element:", elem)
				return nil
			}

		}
		tileCount++
	}

	msg := fmt.Sprintf(" TPS: %0.2f \n FPS: %0.2f \n Number of Things Drawn: %d", ebiten.CurrentTPS(), ebiten.CurrentFPS(), tileCount)
	ebitenutil.DebugPrint(screen, msg)

	return nil
}

func canView(elem *elements.Element, screen *ebiten.Image) bool {
	w, h := screen.Size()
	buf := 64.0
	if myScreen.XPos <= elem.XPos+buf && elem.XPos <= myScreen.XPos+float64(w)+buf {
		if myScreen.YPos <= elem.YPos+buf && elem.YPos <= myScreen.YPos+float64(h)+buf {
			return true
		}
	}

	return false
}
