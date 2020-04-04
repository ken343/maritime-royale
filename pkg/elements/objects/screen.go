package objects

import (
	"github.com/ken343/maritime-royale/pkg/elements"
	"github.com/ken343/maritime-royale/pkg/elements/secondOrder/playerControl"
	"github.com/ken343/maritime-royale/pkg/networking/connection"
)

func NewScreen(xpos float64, ypos float64) *elements.Element {
	screen := &elements.Element{}

	screen.XPos = xpos
	screen.YPos = ypos

	screen.Active = true

	screen.UniqueName = "MySpecialScreen"
	screen.ID = connection.GetID()

	mover := playerControl.NewKeyboardMover(screen, 1)
	screen.AddComponent(mover)

	return screen
}
