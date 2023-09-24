package gameoflife

type Storage interface {
	AddPoint(row, col int)
	AddPoints(points []Point)
	Next(rules *Rules)
	View(originRow, originCol, rows, cols int) ([]bool, error)
}
