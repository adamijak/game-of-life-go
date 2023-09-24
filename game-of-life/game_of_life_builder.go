package game_of_life

import "errors"

type GameOfLifeBuilder struct {
	err     error
	storage *Storage
}

func NewGameOfLifeBuilder() *GameOfLifeBuilder {
	return &GameOfLifeBuilder{}
}

func (b *GameOfLifeBuilder) WithCustomStorage(storage *Storage) *GameOfLifeBuilder {
	if b.storage != nil {
		b.err = errors.New("Can not ")
	}
	b.storage = storage
	return b
}

func (b *GameOfLifeBuilder) Build() (*GameOfLife, error) {
	return &GameOfLife{
		storage: b.storage,
	}, b.err
}
