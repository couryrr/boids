package main

import (
	"github.com/couryrr/boids/internal/game"
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
	game := game.Game{}
	game.Load()
	for !rl.WindowShouldClose() {
		game.Update()
		game.Draw()
	}
}
