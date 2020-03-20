package terrain

type island struct {
	Type string
}

func NewIsland() island {
	return island{Type: "island"}
}

func (s island) Draw() {

}
