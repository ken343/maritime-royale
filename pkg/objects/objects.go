package objects

type content interface {
	OnDraw() string
}

//Square contains all the data about a specific square
type Square struct {
	XPos, YPos int
	Content    content
}
