package gameoflife

import "errors"

type Builder struct {
	err     error
	storage Storage
	rules   *Rules
	points  Points
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) WithRules(rules Rules) *Builder {
	b.rules = &rules
	return b
}

func (b *Builder) WithCustomStorage(storage Storage) *Builder {
	if storage == nil {
		b.err = errors.New("storage can not be nil")
	}

	if b.storage != nil {
		b.err = errors.New("can not have multiple storages")
	}

	b.storage = storage
	return b
}

func (b *Builder) WithPoints(points []Point) *Builder {
	if points == nil {
		b.err = errors.New("points can not be nil")
	}
	b.points = points
	return b
}

func (b *Builder) WithOverflowingStorage(rows, cols int) *Builder {
	if b.storage != nil {
		b.err = errors.New("can not have multiple storages")
	}
	b.storage = NewOverflowingStorage(rows, cols)
	return b
}

func (b *Builder) Build() (*GameOfLife, error) {
	if b.storage == nil {
		b.err = errors.New("storage must be defined")
	}

	if b.rules == nil {
		b.err = errors.New("rules must be defined")
	}

	if b.err != nil {
		return nil, b.err
	}

	b.points.prepare()
	b.storage.AddPoints(b.points)

	return &GameOfLife{
		storage: b.storage,
	}, nil
}
