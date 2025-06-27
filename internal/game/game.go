package game

import (
	"github.com/couryrr/boids/internal/game/gameobject"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	isPlaying bool
	isDebug   bool
	Flock     *gameobject.Flock
}

func (g *Game) Load(flockSize int) {
	g.isPlaying = false
	g.isDebug = true
	g.Flock = gameobject.CreateFlock(flockSize)
}

func (g *Game) Update() {
	if g.isPlaying {
		for boid := range g.Flock.All() {
			boid.GetSteeringForces(g.Flock)
		}
		for boid := range g.Flock.All() {
			boid.UpdatePosition()
		}
	}
	g.Input()
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)
	rl.DrawText("Boid Simulation", 5, 5, 20, rl.LightGray)

	state := "Paused"
	if g.isPlaying {
		state = "Playing"
	}

	rl.DrawText(state, 5, 30, 20, rl.LightGray)

	separateColor, alignColor, cohesionColor := rl.Blue, rl.Blue, rl.Blue
	
	if !gameobject.ShouldSeparate {
		separateColor = rl.Red
	} 	
	
	if !gameobject.ShouldAlign {
		alignColor = rl.Red
	}	

	if !gameobject.ShouldCohesion{
		cohesionColor = rl.Red
	}

	rl.DrawText("Separate", 200, 5, 20, separateColor)
	rl.DrawText("Align", 350, 5, 20, alignColor)
	rl.DrawText("Cohesion", 450, 5, 20, cohesionColor)

	for boid := range g.Flock.All() {
		boid.Draw()
		if g.isDebug {
			boid.DrawDebug()
		}
	}
	if g.isDebug {
		rl.DrawRectangleLines(int32(gameobject.BoundaryDistance), int32(gameobject.BoundaryDistance), int32(rl.GetScreenWidth()-int(gameobject.BoundaryDistance*2)), int32(rl.GetScreenHeight()-int(gameobject.BoundaryDistance*2)), rl.Red)
	}
	rl.EndDrawing()
}
func (g *Game) Click(x, y int32) {
	g.Flock.Add(rl.Vector2{X: float32(x), Y: float32(y)})
}

func (g *Game) Input() {
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		g.Click(rl.GetMouseX(), rl.GetMouseY())
	}
	if rl.IsKeyPressed(rl.KeyF3) {
		g.isDebug = !g.isDebug
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		g.isPlaying = !g.isPlaying
	}
	if rl.IsKeyPressed(rl.KeyA) {
		gameobject.ShouldAlign = !gameobject.ShouldAlign
	}
	if rl.IsKeyPressed(rl.KeyS) {
		gameobject.ShouldSeparate = !gameobject.ShouldSeparate
	}
	if rl.IsKeyPressed(rl.KeyC) {
		gameobject.ShouldCohesion = !gameobject.ShouldCohesion
	}
}
