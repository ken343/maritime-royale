package screen

import (
	"github.com/veandco/go-sdl2/sdl"
)

type ViewPort struct {
	Xpos   float64
	Ypos   float64
	Width  float64
	Height float64
	Scale  float64
	Speed  float64
}

func NewScreen(Xpos float64, Ypos float64, Width float64, Height float64) (s ViewPort) {
	s.Xpos = Xpos
	s.Ypos = Ypos
	s.Width = Width
	s.Height = Height
	s.Scale = 64
	s.Speed = 0.15
	return s
}

func (s *ViewPort) Update() {
	noramlizedSpeed := s.Scale * s.Speed

	keys := sdl.GetKeyboardState()

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

		switch eventType := event.(type) {

		case *sdl.MouseWheelEvent:
			if eventType.Y > 0 {
				s.Scale += noramlizedSpeed
			} else {
				s.Scale -= noramlizedSpeed
			}

		}
	}

	if keys[sdl.SCANCODE_LEFT] == 1 {

		s.Xpos += noramlizedSpeed

	} else if keys[sdl.SCANCODE_RIGHT] == 1 {

		s.Xpos -= noramlizedSpeed

	}
	if keys[sdl.SCANCODE_UP] == 1 {

		s.Ypos += noramlizedSpeed

	} else if keys[sdl.SCANCODE_DOWN] == 1 {

		s.Ypos -= noramlizedSpeed

	}
}
