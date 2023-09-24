package gameoflife

import "errors"

type MatrixStorage struct {
	rows, cols int
	buffer     []bool
}

func NewOverflowingStorage(rows, cols int) *MatrixStorage {
	return &MatrixStorage{
		rows:   rows,
		cols:   cols,
		buffer: make([]bool, rows*cols),
	}
}

func (s *MatrixStorage) AddPoint(row, col int) {
	s.buffer[row*s.cols+col] = true
}

func (s *MatrixStorage) AddPoints(points []Point) {
	for _, point := range points {
		s.AddPoint(point.Row, point.Col)
	}
}

func (s *MatrixStorage) Next(rules *Rules) {
	//TODO implement me
	panic("implement me")
}

func (s *MatrixStorage) View(originRow, originCol, rows, cols int) ([]bool, error) {
	if originRow < 0 || originCol < 0 || s.rows < originRow+rows || s.cols < originCol+cols {
		return nil, errors.New("view out of matrix")
	}

	offset := (originRow * cols) + originCol

	view := make([]bool, rows*cols)
	for r := 0; r < rows; r++ {
		start := r * cols
		end := (r + 1) * cols
		copy(view[start:end], s.buffer[start+offset:end+offset])
	}

	return view, nil
}
