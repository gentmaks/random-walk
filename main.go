package main

import (
	"fmt"
	"math/rand/v2"
)

type Walker struct {
	x int
	y int
}

func main() {
	grid := make([][]int, 10)
	for i := range grid {
		grid[i] = make([]int, 10)
	}
	walker := Walker{x: 0, y: 0}
	for i := 0; i <= 100; i++ {
		walker.display()
		walker.move()
	}
}

func (w *Walker) display() {
	fmt.Println("The coordinates of the walker are", w.x, w.y)
}

func (w *Walker) move() {
	randomNum := rand.IntN(4)
	switch randomNum {
	case 0:
		w.x++
	case 1:
		w.y++
	case 2:
		w.x--
	case 3:
		w.y--
	}
}
