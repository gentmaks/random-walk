package main

import (
	"flag"
	"fmt"
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Walker struct {
	x     int
	y     int
	color Color
}

type Point struct {
	x int
	y int
}

type Color struct {
	r uint8
	g uint8
	b uint8
}

const (
	screenWidth  = 800
	screenHeight = 450
	cellSize     = 20
	stepSpeed    = 10
)

func main() {
	var numAgents int
	flag.IntVar(&numAgents, "numAgents", 5, "Enter the number of agents")
	flag.Parse()
	rl.InitWindow(screenWidth, screenHeight, "Random Walk Simulation")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	trails := make([][]Point, numAgents)
	var walkers []Walker
	for i := range numAgents {
		_ = i
		walker := Walker{
			x:     screenWidth / cellSize / 2,
			y:     screenHeight / cellSize / 2,
			color: Color{r: uint8(rand.IntN(256)), g: uint8(rand.IntN(256)), b: uint8(rand.IntN(256))},
		}
		walkers = append(walkers, walker)
		trails[i] = []Point{}
	}

	frameCounter := 0

	for !rl.WindowShouldClose() {
		frameCounter++

		if frameCounter >= stepSpeed {
			for i := range walkers {
				walkers[i].move()
				trails[i] = append(trails[i], Point{x: walkers[i].x, y: walkers[i].y})

			}
			frameCounter = 0
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		drawGrid()

		for i := range walkers {
			drawTrail(trails[i], walkers[i].color)
			drawWalker(walkers[i])
		}

		rl.DrawText(fmt.Sprintf("Walkers: %d", numAgents), 10, 10, 20, rl.Black)
		rl.EndDrawing()
	}
}

func drawGrid() {
	for i := 0; i < screenWidth/cellSize; i++ {
		x := int32(i * cellSize)
		rl.DrawLineV(rl.NewVector2(float32(x), 0), rl.NewVector2(float32(x), screenHeight), rl.LightGray)
	}
	for j := 0; j < screenHeight/cellSize; j++ {
		y := int32(j * cellSize)
		rl.DrawLineV(rl.NewVector2(0, float32(y)), rl.NewVector2(screenWidth, float32(y)), rl.LightGray)
	}
}

func drawTrail(trail []Point, lineColor Color) {
	if len(trail) < 2 {
		return
	}

	color := rl.Color{R: lineColor.r, G: lineColor.g, B: lineColor.b, A: 255}

	for i := 1; i < len(trail); i++ {
		first := trail[i-1]
		second := trail[i]

		x1 := int32(first.x*cellSize + cellSize/2)
		y1 := int32(first.y*cellSize + cellSize/2)
		x2 := int32(second.x*cellSize + cellSize/2)
		y2 := int32(second.y*cellSize + cellSize/2)

		rl.DrawLine(x1, y1, x2, y2, color)

		rl.DrawCircle(x2, y2, 1, color)
	}
}

func drawWalker(w Walker) {
	centerX := int32(w.x*cellSize + cellSize/2)
	centerY := int32(w.y*cellSize + cellSize/2)

	rl.DrawCircle(centerX, centerY, float32(cellSize/2)-2, rl.Red)
}

func (w *Walker) move() {
	randomNum := rand.IntN(8)

	switch randomNum {
	case 0: // Right
		w.x++
	case 1: // Down
		w.y++
	case 2: // Left
		w.x--
	case 3: // Up
		w.y--
	case 4: // Down-Right
		w.x++
		w.y++
	case 5: // Up-Left
		w.x--
		w.y--
	case 6: // Down-Left
		w.x--
		w.y++
	case 7: // Up-Right
		w.x++
		w.y--
	}

	if w.x < 0 {
		w.x = screenWidth/cellSize - 1
	}
	if w.x >= screenWidth/cellSize {
		w.x = 0
	}
	if w.y < 0 {
		w.y = screenHeight/cellSize - 1
	}
	if w.y >= screenHeight/cellSize {
		w.y = 0
	}
}
