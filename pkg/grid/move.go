package grid

import "errors"

var moveOptions []Coordinate

func SetMoveOptions(sq *Square) {
	UnsetMoveOptions()
	moveOptions = append(moveOptions, // distance, pathing logic and collision placeholder
		Coordinate{XPos: sq.Coords.XPos + 1, YPos: sq.Coords.YPos},
		Coordinate{XPos: sq.Coords.XPos - 1, YPos: sq.Coords.YPos},
		Coordinate{XPos: sq.Coords.XPos, YPos: sq.Coords.YPos + 1},
		Coordinate{XPos: sq.Coords.XPos, YPos: sq.Coords.YPos - 1})
}

func UnsetMoveOptions() {
	moveOptions = make([]Coordinate, 0)
}

func GetMoveOptions() []Coordinate {
	return moveOptions
}

func isValidUnitMove(sq *Square) bool {
	for _, validCoord := range moveOptions {
		if sq.Coords == validCoord {
			return true
		}
	}
	return false
}

func CountMoveOptions() int {
	return len(moveOptions)
}

func MoveUnit(src *Square, dst *Square) error {
	if isValidUnitMove(dst) {
		dst.Unit = src.Unit
		src.Unit = nil
		return nil
	}
	return errors.New("Invalid Move")
}

/* JZ: combines MoveUnit and isValidUnitMove into a single function. either way works but it's cleaner to keep em separate
func MoveUnit(src *Square, dst *Square) error {
	if dst.Unit == nil {
		for _, validCoord := range moveOptions {
			if dst.Coords == validCoord {
				dst.Unit = src.Unit
				src.Unit = nil
				return nil
			}
		}
	}
	return errors.New("Invalid Move")
}
*/

func MoveWeather(src *Square, dst *Square) {
	dst.Weather = src.Weather
	src.Weather = nil
}

func MoveTerrain(src *Square, dst *Square) {
	dst.Terrain = src.Terrain
	src.Terrain = nil
}
