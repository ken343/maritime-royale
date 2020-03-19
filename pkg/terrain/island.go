package terrain

type island struct {
	symbol string
}

func NewIsland() island {
	return island{symbol: "i"}
}

func (s island) Draw() {

}
