package objects

import (
	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/physics"
	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/render"
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
