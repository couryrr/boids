package objects

import (
	"github.com/couryrr/boids/internal/simulation/util"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Boid struct {
	Id         int
	Radius     float32
	Position   rl.Vector2
	Direction  rl.Vector2
	BoundaryV  *rl.Vector2
	AvoidanceV *rl.Vector2
	Speed      float64
}

func CreateBoid(id int, position rl.Vector2, direction rl.Vector2, speed float64) *Boid {
	return &Boid{
		Id:         id,
		Radius:     5,
		Position:   position,
		Direction:  direction,
		Speed:      speed,
		BoundaryV:  &rl.Vector2{},
		AvoidanceV: &rl.Vector2{},
	}
}

func (b *Boid) GetSteeringForces(factors *Factors, flock *Flock) {
	*b.BoundaryV = b.Boundary(factors)

	avoidAcc := &util.ForceAccumulator{}
	for neighbor := range flock.All() {
		b.Avoidance(neighbor, factors, avoidAcc)
	}
	value, err := avoidAcc.Value()
	if err == nil {
		*b.AvoidanceV = rl.Vector2Scale(rl.Vector2Normalize(value), factors.AvoidanceScale)
	}
}

func (b *Boid) Boundary(factors *Factors) rl.Vector2 {
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())

	force := rl.Vector2Zero()
	force.X += isColliding(b.Position.X, float32(factors.Fov), factors.BoundaryDistance, factors.BoundaryFactor, screenWidth)
	force.Y += isColliding(b.Position.Y, float32(factors.Fov), factors.BoundaryDistance, factors.BoundaryFactor, screenHeight)

	return rl.Vector2Scale(rl.Vector2Normalize(force), factors.BoundaryScale)
}

func (b *Boid) Avoidance(neighbor *Boid, factors *Factors, accumulator *util.ForceAccumulator) {
	d := rl.Vector2Distance(b.Position, neighbor.Position)
	if d <= float32(factors.Separation) {
		accumulator.Increment(rl.Vector2Subtract(b.Position, neighbor.Position))
	}
}

func (b *Boid) UpdatePosition() {
	if !rl.Vector2Equals(*b.BoundaryV, rl.Vector2Zero()) {
		desiredDir := rl.Vector2Normalize(rl.Vector2Add(b.Direction, *b.BoundaryV))
		b.Direction = rl.Vector2Normalize(rl.Vector2Lerp(b.Direction, desiredDir, 0.1))
	}
	vel := rl.Vector2Scale(b.Direction, float32(b.Speed))
	b.Position = rl.Vector2Add(b.Position, vel)
}

func (b *Boid) Draw() {
	rl.DrawCircleLines(int32(b.Position.X), int32(b.Position.Y), b.Radius, rl.Black)
}

func (b *Boid) DrawDebug(factors *Factors) {
	rl.DrawCircleLines(int32(b.Position.X), int32(b.Position.Y), float32(factors.Separation), rl.Red)
	rl.DrawCircleLines(int32(b.Position.X), int32(b.Position.Y), float32(factors.Fov), rl.Green)
	rl.DrawLineV(b.Position, rl.Vector2Add(b.Position, rl.Vector2Scale(b.Direction, 15)), rl.Purple)
	if b.BoundaryV != nil {
		rl.DrawLineV(b.Position, rl.Vector2Add(b.Position, rl.Vector2Scale(*b.BoundaryV, 15)), rl.Red)
	}
	if b.AvoidanceV != nil {
		rl.DrawLineV(b.Position, rl.Vector2Add(b.Position, rl.Vector2Scale(*b.AvoidanceV, 15)), rl.Blue)
	}
}

func isColliding(part, fov, boundary, factor, screen float32) float32 {
	if part-fov <= boundary {
		d := boundary - (part - fov)
		return factor * d
	}
	if part+fov > screen-boundary {
		d := part + fov - screen - boundary
		return -1 * factor * d
	}
	return 0
}
