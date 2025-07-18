package objects

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Boid struct {
	Id        int
	Radius    float32
	Position  rl.Vector2
	Direction rl.Vector2
	BoundaryV *rl.Vector2
	Speed     float64
}

func CreateBoid(id int, position rl.Vector2, direction rl.Vector2, speed float64) *Boid {
	return &Boid{
		Id:        id,
		Radius:    5,
		Position:  position,
		Direction: direction,
		Speed:     speed,
	}
}

func (b *Boid) GetSteeringForces(factors *Factors, flock *Flock) {
	b.BoundaryV = b.Boundary(factors)
}

func (b *Boid) Boundary(factors *Factors) *rl.Vector2 {
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())

	force := rl.Vector2Zero()

	if b.Position.X <= factors.BoundaryDistance {
		force.X += factors.BoundaryFactor
	}
	if b.Position.X > screenWidth-factors.BoundaryDistance {
		force.X -= factors.BoundaryFactor
	}
	if b.Position.Y <= factors.BoundaryDistance {
		force.Y += factors.BoundaryFactor
	}
	if b.Position.Y > screenHeight-factors.BoundaryDistance {
		force.Y -= factors.BoundaryFactor
	}

	if !rl.Vector2Equals(force, rl.Vector2Zero()) {
		force = rl.Vector2Normalize(force)
	}
	force = rl.Vector2Scale(force, factors.BoundaryScale)

	return &force
}

func (b *Boid) UpdatePosition() {
	res := rl.Vector2Zero()
	res = rl.Vector2Add(res, *b.BoundaryV)
	b.Direction = rl.Vector2Add(b.Direction, res)
	vel := rl.Vector2Scale(rl.Vector2Normalize(b.Direction), float32(b.Speed))
	b.Position = rl.Vector2Add(b.Position, vel)
}

func (b *Boid) Draw() {
	rl.DrawCircleLines(int32(b.Position.X), int32(b.Position.Y), b.Radius, rl.Black)
}

func (b *Boid) DrawDebug(factors *Factors) {
	rl.DrawCircleLines(int32(b.Position.X), int32(b.Position.Y), float32(factors.Separation), rl.Red)
	rl.DrawCircleLines(int32(b.Position.X), int32(b.Position.Y), float32(factors.Fov), rl.Green)
	rl.DrawLineV(b.Position, rl.Vector2Add(b.Position, rl.Vector2Scale(b.Direction, 5)), rl.Purple)
	if b.BoundaryV != nil {
		rl.DrawLineV(b.Position, rl.Vector2Add(b.Position, rl.Vector2Scale(*b.BoundaryV, 5)), rl.Red)
	}
}
