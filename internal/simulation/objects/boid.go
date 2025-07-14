package objects

import (
	"github.com/couryrr/boids/internal/simulation/util"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Boid struct {
	Id            int
	Radius        float32
	Position      rl.Vector2
	PrevDirection rl.Vector2
	Direction     rl.Vector2
	SeparateV     *rl.Vector2
	AlignV        *rl.Vector2
	CohesionV     *rl.Vector2
	BoundaryV     *rl.Vector2
	Speed         float32
}

func CreateBoid(id int, position rl.Vector2, direction rl.Vector2, speed float32) *Boid {
	return &Boid{
		Id:        id,
		Radius:    5,
		Position:  position,
		Direction: direction,
		Speed:     0.05,
	}
}

func (b *Boid) GetSteeringForces(factors *Factors, flock *Flock) {
	seperationAccumulator := util.ForceAccumulator{}
	// alignAccum := util.ForceAccumulator{}
	// cohesionAccum := util.ForceAccumulator{}

	b.SeparateV = nil
	b.AlignV = nil
	b.CohesionV = nil
	for neighbor := range flock.All() {
		b.Separate(neighbor, factors.Separation, &seperationAccumulator)
		// b.Align(neighbor, factors.Separation, factors.Fov, &alignAccum)
		// b.Cohesion(neighbor, factors.Separation, factors.Fov, &cohesionAccum)
	}

	vec, err := seperationAccumulator.Value()
	if err == nil {
		b.SeparateV = &vec
	}
	/*
		vec, err = alignAccum.Average()
		if err == nil {
			temp := rl.Vector2Subtract(rl.Vector2Normalize(vec), b.Direction)
			b.AlignV = &temp
		}

		vec, err = cohesionAccum.Average()
		if err == nil {
			temp := rl.Vector2Normalize(rl.Vector2Subtract(vec, b.Position))
			b.CohesionV = &temp
		}
	*/
	{
		temp := rl.Vector2Scale(b.Boundary(factors), factors.BoundaryScale)
		b.BoundaryV = &temp
	}
}

func (b *Boid) Separate(neighbor *Boid, Separation float64, fa *util.ForceAccumulator) {
	distance := rl.Vector2Distance(b.Position, neighbor.Position)
	if b.Id != neighbor.Id && float64(distance) <= Separation {
		fa.Increment(rl.Vector2Subtract(b.Position, neighbor.Position))
	}
}

/*
func (b *Boid) Align(neighbor *Boid, separation float64, fov float64, fa *util.ForceAccumulator) {
	distance := rl.Vector2Distance(b.Position, neighbor.Position)
	if b.Id != neighbor.Id && distance <= float32(fov) && float64(distance) > separation {
		fa.Increment(neighbor.Direction)
	}
}

func (b *Boid) Cohesion(neighbor *Boid, separation float64, fov float64, fa *util.ForceAccumulator) {
	distance := rl.Vector2Distance(b.Position, neighbor.Position)
	if b.Id != neighbor.Id && distance <= float32(fov) && float64(distance) > separation {
		fa.Increment(neighbor.Position)
	}
}
*/

func (b *Boid) Boundary(factors *Factors) rl.Vector2 {
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())

	force := rl.Vector2Zero()
	if b.Position.X-float32(factors.Fov) <= factors.BoundaryDistance {
		force.X = b.Position.X + factors.BoundaryFactor
	}
	if b.Position.X+float32(factors.Fov) > screenWidth-factors.BoundaryDistance {
		force.X = factors.BoundaryFactor - b.Position.X
	}
	if b.Position.Y-float32(factors.Fov) <= factors.BoundaryDistance {
		force.Y = b.Position.Y + factors.BoundaryFactor
	}
	if b.Position.Y+float32(factors.Fov) > screenHeight-factors.BoundaryDistance {
		force.Y = factors.BoundaryFactor - b.Position.Y
	}

	return force
}

func (b *Boid) UpdatePosition() {
	res := rl.Vector2Zero()
	if b.SeparateV != nil {
		res = rl.Vector2Add(res, *b.SeparateV)
		rl.TraceLog(rl.LogDebug, "Setting separation: %v", b.SeparateV)
	}
	/*
		if b.AlignV != nil {
			res = rl.Vector2Add(res, *b.AlignV)
		}

		if b.CohesionV != nil {
			res = rl.Vector2Add(res, *b.CohesionV)
		}
	*/
	res = rl.Vector2Add(res, *b.BoundaryV)
	b.Direction = rl.Vector2Add(b.Direction, rl.Vector2Normalize(res))
	b.Position = rl.Vector2Add(b.Position, rl.Vector2Scale(b.Direction, b.Speed))
}

func (b *Boid) Draw() {
	rl.DrawCircleLines(int32(b.Position.X), int32(b.Position.Y), b.Radius, rl.Black)
}

func (b *Boid) DrawDebug(factors *Factors) {
	rl.DrawCircleLines(int32(b.Position.X), int32(b.Position.Y), float32(factors.Separation), rl.Red)
	rl.DrawCircleLines(int32(b.Position.X), int32(b.Position.Y), float32(factors.Fov), rl.Green)
	rl.DrawLineV(b.Position, rl.Vector2Add(b.Position, rl.Vector2Scale(b.Direction, 5)), rl.Purple)
	if b.SeparateV != nil {
		rl.TraceLog(rl.LogDebug, "There is separation: %v", b.SeparateV)
		rl.DrawLineV(b.Position, rl.Vector2Add(b.Position, rl.Vector2Scale(*b.SeparateV, 5)), rl.Blue)
	}
	// rl.DrawLineV(b.Position, rl.Vector2Add(b.Position, rl.Vector2Scale(*b.AlignV, 5)), rl.Orange)
	// rl.DrawLineV(b.Position, rl.Vector2Add(b.Position, rl.Vector2Scale(*b.CohesionV, 5)), rl.Pink)
	if b.BoundaryV != nil {
		rl.DrawLineV(b.Position, rl.Vector2Add(b.Position, rl.Vector2Scale(*b.BoundaryV, 5)), rl.Red)
	}
}
