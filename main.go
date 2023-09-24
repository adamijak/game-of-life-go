package main

import (
	"fmt"
	gameoflife "game-of-life-go/game-of-life"
	"time"
)

func main() {
	game, _ := gameoflife.
		NewBuilder().
		WithRules(gameoflife.Rules{
			LiveToDie: gameoflife.Neighbours{
				Zero:  true,
				One:   true,
				Four:  true,
				Five:  true,
				Six:   true,
				Seven: true,
				Eight: true,
			},
			DeadToLive: gameoflife.Neighbours{
				Three: true,
			},
		}).
		WithPoints([]gameoflife.Point{
			{Row: 10, Col: 10},
			{Row: 10, Col: 11},
			{Row: 10, Col: 12},
		}).
		WithOverflowingStorage(100, 100).
		Build()

	rows := 50
	cols := 50
	for {
		view, _ := game.View(0, 0, rows, cols)
		printMatrix(view, rows, cols)
		time.Sleep(1000)
		game.Next()
	}
}

func printMatrix(matrix []bool, rows, cols int) {
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			char := "□"
			if matrix[r*cols+c] {
				char = "■"
			}
			fmt.Printf("%s", char)
		}
		fmt.Println()
	}
}
