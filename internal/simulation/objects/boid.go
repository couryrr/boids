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
	AlignmentV *rl.Vector2
	CohesionV  *rl.Vector2
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
		AlignmentV: &rl.Vector2{},
		CohesionV:  &rl.Vector2{},
	}
}

func (b *Boid) GetSteeringForces(factors *Factors, flock *Flock) {
	b.AvoidanceV = &rl.Vector2{}
	b.AlignmentV = &rl.Vector2{}
	b.CohesionV = &rl.Vector2{}
	*b.BoundaryV = b.Boundary(factors)

	avoidAcc := &util.ForceAccumulator{}
	aligAcc := &util.ForceAccumulator{}
	cohesionAcc := &util.ForceAccumulator{}
	for neighbor := range flock.All() {
		b.Avoidance(neighbor, factors, avoidAcc)
		b.Alignment(neighbor, factors, aligAcc)
		b.Cohesion(neighbor, factors, cohesionAcc)
	}
	value, err := avoidAcc.Value()
	if err == nil {
		*b.AvoidanceV = rl.Vector2Scale(rl.Vector2Normalize(value), factors.AvoidanceScale)
	}

	value, err = aligAcc.Average()
	if err == nil {
		*b.AlignmentV = rl.Vector2Scale(rl.Vector2Normalize(value), factors.AlignmentScale)
	}

	value, err = cohesionAcc.Average()
	if err == nil {
		*b.CohesionV = rl.Vector2Scale(rl.Vector2Normalize(rl.Vector2Subtract(value, b.Position)), factors.CohesionScale)
	}
}

func (b *Boid) Boundary(factors *Factors) rl.Vector2 {
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())

	force := rl.Vector2Zero()
	if b.Position.X <= factors.BoundaryDistance {
		d := factors.BoundaryDistance - b.Position.X
		force.X += factors.BoundaryFactor * d
	}
	if b.Position.X > screenWidth-factors.BoundaryDistance {
		d := b.Position.X - (screenWidth - factors.BoundaryDistance)
		force.X -= factors.BoundaryFactor * d
	}
	if b.Position.Y <= factors.BoundaryDistance {
		d := factors.BoundaryDistance - b.Position.Y
		force.Y += factors.BoundaryFactor * d
	}
	if b.Position.Y > screenHeight-factors.BoundaryDistance {
		d := b.Position.Y - (screenHeight - factors.BoundaryDistance)
		force.Y -= factors.BoundaryFactor * d
	}

	mag := rl.Vector2Length(force)
	if mag > factors.BoundaryScale {
		force = rl.Vector2Scale(rl.Vector2Normalize(force), factors.BoundaryScale)
	}

	return force
}

func (b *Boid) Avoidance(neighbor *Boid, factors *Factors, accumulator *util.ForceAccumulator) {
	d := rl.Vector2Distance(b.Position, neighbor.Position)
	if b.Id != neighbor.Id && d <= float32(factors.Separation) {
		accumulator.Increment(rl.Vector2Subtract(b.Position, neighbor.Position))
	}
}

func (b *Boid) Alignment(neighbor *Boid, factors *Factors, accumulator *util.ForceAccumulator) {
	d := rl.Vector2Distance(b.Position, neighbor.Position)
	if b.Id != neighbor.Id && d <= float32(factors.Fov) && d > float32(factors.Separation) {
		accumulator.Increment(neighbor.Direction)
	}
}

func (b *Boid) Cohesion(neighbor *Boid, factors *Factors, accumulator *util.ForceAccumulator) {
	d := rl.Vector2Distance(b.Position, neighbor.Position)
	if b.Id != neighbor.Id && d <= float32(factors.Fov) && d > float32(factors.Separation) {
		accumulator.Increment(neighbor.Position)
	}

}
func (b *Boid) UpdatePosition() {
	force := rl.Vector2Add(b.Direction, *b.AvoidanceV)
	force = rl.Vector2Add(force, *b.AlignmentV)
	force = rl.Vector2Add(force, *b.CohesionV)
	force = rl.Vector2Add(force, *b.BoundaryV)
	if !rl.Vector2Equals(force, rl.Vector2Zero()) {
		b.Direction = rl.Vector2Normalize(rl.Vector2Lerp(b.Direction, rl.Vector2Normalize(force), 0.1))
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
	if b.AlignmentV != nil {
		rl.DrawLineV(b.Position, rl.Vector2Add(b.Position, rl.Vector2Scale(*b.AlignmentV, 15)), rl.Green)
	}
	if b.CohesionV != nil {
		rl.DrawLineV(b.Position, rl.Vector2Add(b.Position, rl.Vector2Scale(*b.CohesionV, 15)), rl.Gold)
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
