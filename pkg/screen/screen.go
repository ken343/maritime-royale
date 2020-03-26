package screen

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type ViewPort struct {
	Xpos   float64
	Ypos   float64
	Width  float64
	Height float64
	Scale  float64
	Speed  float64
	Mouse  Mouse
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

// Could theoretically return MRP requests, or data that would initiate the process for a request.
func (s *ViewPort) Update() string {
	noramlizedSpeed := s.Scale * s.Speed

	keys := sdl.GetKeyboardState()

	s.Mouse.Xpos, s.Mouse.Ypos, s.Mouse.State = sdl.GetMouseState()

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

		switch eventType := event.(type) {

		// exit case
		case *sdl.QuitEvent:
			return "EXIT"

		// escape key exit case, selecting via keyboard
		case *sdl.KeyboardEvent:
			if eventType.Keysym.Scancode == sdl.SCANCODE_ESCAPE {
				return "EXIT"
			} else if eventType.State == sdl.RELEASED {
				keyCode := string(eventType.Keysym.Sym)
				switch keyCode {
				default:
					fmt.Println(keyCode)
				}
			}

		// selecting via mouse
		case *sdl.MouseButtonEvent:
			if eventType.State == sdl.RELEASED {
				fmt.Printf("{%d, %d}\n", eventType.X, eventType.Y)
				if eventType.Button == sdl.BUTTON_LEFT {
					//mouseOnReleaseLeft(eventType.X, eventType.Y, renderer, waterTileSurface)
				} else {
					//mouseOnReleaseRight()
				}
			}

		case *sdl.MouseWheelEvent:
			if eventType.Y > 0 {
				s.Scale += noramlizedSpeed
			} else {
				s.Scale -= noramlizedSpeed
			}

		}
	}

	if keys[sdl.SCANCODE_LEFT] == 1 {

		s.Xpos -= noramlizedSpeed

	} else if keys[sdl.SCANCODE_RIGHT] == 1 {

		s.Xpos += noramlizedSpeed

	}
	if keys[sdl.SCANCODE_UP] == 1 {

		s.Ypos -= noramlizedSpeed

	} else if keys[sdl.SCANCODE_DOWN] == 1 {

		s.Ypos += noramlizedSpeed

	}

	if s.Mouse.State == 1 {
		fmt.Println(int((float64(s.Mouse.Xpos)+s.Xpos)/s.Scale), int((float64(s.Mouse.Ypos)+s.Ypos)/s.Scale))
	}

	return ""
}
