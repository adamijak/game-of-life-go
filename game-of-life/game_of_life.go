package gameoflife

type GameOfLife struct {
	storage Storage
	rules   *Rules
}

func (g *GameOfLife) Next() {
	g.storage.Next(g.rules)
}

func (g *GameOfLife) View(originRow, originCol, rows, cols int) ([]bool, error) {
	return g.storage.View(originRow, originCol, rows, cols)
}
