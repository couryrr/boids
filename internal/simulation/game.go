package simulation

import (
	"github.com/couryrr/boids/internal/simulation/objects"
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Simulation struct {
	state     objects.State
	isPlaying bool
	isDebug   bool
	Flock     *objects.Flock
}

var (
	avoid    int32 = 0
	align    int32 = 0
	cohesion int32 = 0
)

func (s *Simulation) Load(flockSize int) {
	s.state = objects.CreateState()
	avoid = int32(s.state.Factors.AvoidanceScale)
	align = int32(s.state.Factors.AlignmentScale)
	cohesion = int32(s.state.Factors.CohesionScale)
	s.isPlaying = false
	s.isDebug = true
	s.Flock = objects.CreateFlock(s.state.Factors.BoundaryDistance, flockSize, &s.state.Factors)
}

func (s *Simulation) Update() {
	s.state.Factors.AvoidanceScale = float32(avoid)
	s.state.Factors.AlignmentScale = float32(align)
	s.state.Factors.CohesionScale = float32(cohesion)

	for boid := range s.Flock.All() {
		boid.GetSteeringForces(&s.state.Factors, s.Flock)
	}
	if s.isPlaying {
		for boid := range s.Flock.All() {
			boid.UpdatePosition()
		}
	}
	s.Input()
}

func (s *Simulation) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)
	rl.DrawText("Boid Simulation", 5, 5, 20, rl.LightGray)

	stateText := "Paused"
	if s.isPlaying {
		stateText = "Playing"
	}

	rl.DrawText(stateText, 5, 30, 20, rl.LightGray)

	gui.Spinner(rl.Rectangle{
		X:      200,
		Y:      5,
		Width:  125,
		Height: 30,
	},
		"", &avoid, 0, 100, true)

	gui.Spinner(rl.Rectangle{
		X:      350,
		Y:      5,
		Width:  125,
		Height: 30,
	},
		"", &align, 0, 100, true)

	gui.Spinner(rl.Rectangle{
		X:      500,
		Y:      5,
		Width:  125,
		Height: 30,
	},
		"", &cohesion, 0, 100, true)

	for boid := range s.Flock.All() {
		boid.Draw()
		if s.isDebug {
			boid.DrawDebug(&s.state.Factors)
		}
	}
	if s.isDebug {
		rl.DrawRectangleLines(int32(s.state.Factors.BoundaryDistance), int32(s.state.Factors.BoundaryDistance), int32(rl.GetScreenWidth()-int(s.state.Factors.BoundaryDistance*2)), int32(rl.GetScreenHeight()-int(s.state.Factors.BoundaryDistance*2)), rl.Red)
	}
	rl.EndDrawing()
}
func (s *Simulation) Click(x, y int32) {
	s.Flock.Add(s.state.Factors.BoundaryDistance, rl.Vector2{X: float32(x), Y: float32(y)}, &s.state.Factors)
}

func (s *Simulation) Input() {
	if rl.IsKeyPressed(rl.KeyF11) {
		display := rl.GetCurrentMonitor()
		if rl.IsWindowFullscreen() {
			rl.SetWindowSize(1920, 1080)
		} else {
			rl.SetWindowSize(rl.GetMonitorWidth(display), rl.GetMonitorHeight(display))
		}
		rl.ToggleFullscreen()
	}
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		s.Click(rl.GetMouseX(), rl.GetMouseY())
	}
	if rl.IsKeyPressed(rl.KeyF3) {
		s.isDebug = !s.isDebug
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		s.isPlaying = !s.isPlaying
	}
	if rl.IsKeyPressed(rl.KeyOne) {
		s.state.ShouldAlign = !s.state.ShouldAlign
	}
	if rl.IsKeyPressed(rl.KeyTwo) {
		s.state.ShouldSeparate = !s.state.ShouldSeparate
	}
	if rl.IsKeyPressed(rl.KeyThree) {
		s.state.ShouldCohesion = !s.state.ShouldCohesion
	}
}
