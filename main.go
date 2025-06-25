package main

import (
	"github.com/couryrr/boids/internal/gameobjects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	width  = 1920
	height = 1080
)

func main() {
	rl.SetTraceLogLevel(rl.LogDebug)
	rl.TraceLog(rl.LogDebug, "Starting game!")
	rl.InitWindow(width, height, "Boids Simulation One Day!")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	flock := gameobjects.CreateFlock(100)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawText("Boid Simulation", 5, 5, 20, rl.LightGray)

		rl.DrawRectangleLines(int32(gameobjects.BoundaryDistance), int32(gameobjects.BoundaryDistance), int32(rl.GetScreenWidth()-int(gameobjects.BoundaryDistance*2)), int32(rl.GetScreenHeight()-int(gameobjects.BoundaryDistance*2)), rl.Red)
		for boid := range flock.All() {
			boid.Steer(flock)
			boid.UpdatePosition()
			boid.Draw()
		}

		rl.EndDrawing()
	}
}
