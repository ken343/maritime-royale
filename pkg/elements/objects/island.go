package objects

import (
	"github.com/ken343/maritime-royale/pkg/elements"
	"github.com/ken343/maritime-royale/pkg/elements/secondOrder/physics"
	"github.com/ken343/maritime-royale/pkg/elements/secondOrder/render"
)

func NewIsland(xpos float64, ypos float64, Name string) *elements.Element {
	Island := &elements.Element{}

	Island.XPos = xpos
	Island.YPos = ypos

	Island.Active = true

	Island.UniqueName = Name

	sr := render.NewSpriteRenderer(Island, "island.png")
	Island.AddComponent(sr)

	coli := physics.NewCollider(Island)
	coli.Radius = 25
	Island.AddComponent(coli)

	return Island
}
