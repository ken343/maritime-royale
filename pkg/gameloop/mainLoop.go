package gameloop

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
)

func Update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	tileCount := 0

	for _, elem := range gamestate.GetWorld() {
		if elem.Active {
			err := elem.Update()
			if err != nil {
				fmt.Println("updating element:", err)
				return nil
			}
			err = elem.Draw(screen)
			if err != nil {
				fmt.Println("drawing element:", elem)
				return nil
			}
			tileCount++
		}
	}

	msg := fmt.Sprintf(" TPS: %0.2f \n FPS: %0.2f \n Number of Things Drawn: %d", ebiten.CurrentTPS(), ebiten.CurrentFPS(), tileCount)
	ebitenutil.DebugPrint(screen, msg)

	return nil
}
