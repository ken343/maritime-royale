package grid

var selection = &Square{
	Coords: Coordinate{
		XPos: -1,
		YPos: -1,
	},
}

func SetSelection(sq *Square) {
	selection = sq
}

func UnsetSelection() {
	selection = &Square{
		Coords: Coordinate{
			XPos: -1,
			YPos: -1,
		},
	}
}

func GetSelection() *Square {
	return selection
}

func ExistsSelection() bool {
	return selection.Coords.XPos >= 0 && selection.Coords.YPos >= 0
}
