package grid

type Coordinate struct {
	XPos, YPos int
}

func GetCoordsAt(pixelX, pixelY, tileWidth, tileHeight int32) Coordinate {
	x := pixelX / tileWidth
	y := pixelY / tileHeight

	return Coordinate{
		XPos: int(x),
		YPos: int(y),
	}
}