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

		for boid := range flock.All() {
			boid.Steer(flock)
			boid.UpdatePosition()

		}
		for boid :=range flock.All(){
			boid.AddDirection(rl.Vector2Scale(boid.Boundary(),gameobjects.BoundaryScale))
			boid.UpdatePosition()
			boid.Draw()
		}
		rl.EndDrawing()
	}
}
