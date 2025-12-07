package main

import (
	"flag"
	"fmt"
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Walker struct {
	x int
	y int
}

type Point struct {
	x int
	y int
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
	rl.InitWindow(screenWidth, screenHeight, "Random Walk Simulation")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	trails := make([][]Point, numAgents)
	var walkers []Walker
	for i := range numAgents {
		_ = i
		walker := Walker{
			x: screenWidth / cellSize / 2,
			y: screenHeight / cellSize / 2,
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

		// Draw the grid lines
		drawGrid()

		for i := range walkers {
			drawTrail(trails[i])
			drawWalker(walkers[i])
		}

		rl.DrawText(fmt.Sprintf("Walkers: %d", numAgents), 10, 10, 20, rl.Black)
		rl.EndDrawing()
	}
}

// --- Helper Functions for Drawing ---

func drawGrid() {
	// Draw vertical lines
	for i := 0; i < screenWidth/cellSize; i++ {
		x := int32(i * cellSize)
		rl.DrawLineV(rl.NewVector2(float32(x), 0), rl.NewVector2(float32(x), screenHeight), rl.LightGray)
	}
	// Draw horizontal lines
	for j := 0; j < screenHeight/cellSize; j++ {
		y := int32(j * cellSize)
		rl.DrawLineV(rl.NewVector2(0, float32(y)), rl.NewVector2(screenWidth, float32(y)), rl.LightGray)
	}
}

func drawTrail(trail []Point) {
	// Draw a small dot for every past position in the trail
	for _, p := range trail {
		// Convert grid coordinates (p.x, p.y) to screen coordinates
		centerX := int32(p.x*cellSize + cellSize/2)
		centerY := int32(p.y*cellSize + cellSize/2)

		// Draw the point. We use Blue with some transparency (Alpha)
		rl.DrawCircle(centerX, centerY, 2, rl.Fade(rl.Blue, 0.7))
	}
}

func drawWalker(w Walker) {
	// Convert grid coordinates (w.x, w.y) to screen coordinates
	centerX := int32(w.x*cellSize + cellSize/2)
	centerY := int32(w.y*cellSize + cellSize/2)

	// Draw the walker as a larger Red circle
	rl.DrawCircle(centerX, centerY, float32(cellSize/2)-2, rl.Red)
}

// --- Walker Logic ---

func (w *Walker) move() {
	// Move in one of the 8 directions (0-7).
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

	// Keep the walker roughly within the screen boundaries (optional wrapping)
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
