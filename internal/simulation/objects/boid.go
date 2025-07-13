package objects

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Boid struct {
	Id            int
	Radius        float32
	Position      rl.Vector2 
	PrevDirection rl.Vector2
	Direction     rl.Vector2 
	SeparateV     rl.Vector2
	AlignV        rl.Vector2
	CohesionV     rl.Vector2
	BoundaryV     rl.Vector2
	Speed         float32 
}

func CreateBoid(id int, position rl.Vector2, direction rl.Vector2, speed float32) *Boid {
	return &Boid{
		Id:        id,
		Radius:    5,
		Position:  position,
		Direction: direction,
		Speed:     .02,
	}
}

func (b *Boid) AddDirection(dir rl.Vector2) {
	b.Direction = rl.Vector2Normalize(rl.Vector2Add(dir, b.Direction))
}

func (b *Boid) GetSteeringForces(factors *Factors, flock *Flock) {
	b.PrevDirection = b.Direction
	// culSep := rl.Vector2Zero()
	//
	// countSep := 0
	// culDir := rl.Vector2Zero()
	// countDir := 0
	// culPos := rl.Vector2Zero()
	// countPos := 0
	//
	// for neighbor := range flock.All() {
	// 	if ShouldSeparate {
	// 		b.Separate(neighbor, &culSep, &countSep)
	// 	}
	//
	// 	if ShouldAlign {
	// 		b.Align(neighbor, &culDir, &countDir)
	// 	}
	//
	// 	if ShouldCohesion {
	// 		b.Cohesion(neighbor, &culPos, &countPos)
	// 	}
	// }
	//
	// if countSep > 0 {
	// 	b.SeparateV = rl.Vector2Scale(culSep, SeparationScale)
	// }
	//
	// if countDir > 0 {
	// 	b.AlignV = rl.Vector2Scale(culDir, AlignmentScale)
	// }
	//
	// if countPos > 0 {
	// 	b.CohesionV = rl.Vector2Scale(culPos, CohesionScale)
	// }
	//
	b.BoundaryV = rl.Vector2Scale(rl.Vector2Normalize(b.Boundary(factors)), factors.BoundaryScale)
}

/*
func (b *Boid) Separate(neighbor *Boid, cul *rl.Vector2, count *int) {
	distance := rl.Vector2Distance(b.Position, neighbor.Position)
	if b.Id != neighbor.Id && float64(distance) <= Separation {
		*count += 1
		*cul = rl.Vector2Subtract(*cul, neighbor.Position)
	}
}

func (b *Boid) Align(neighbor *Boid, cul *rl.Vector2, count *int) {
	distance := rl.Vector2Distance(b.Position, neighbor.Position)
	if b.Id != neighbor.Id && distance <= float32(Fov) && float64(distance) > Separation {
		*count += 1
		*cul = rl.Vector2Add(*cul, neighbor.Direction)
	}
}

func (b *Boid) Cohesion(neighbor *Boid, cul *rl.Vector2, count *int) {
	distance := rl.Vector2Distance(b.Position, neighbor.Position)
	if b.Id != neighbor.Id && distance <= float32(Fov) && float64(distance) > Separation {
		*count += 1
		*cul = rl.Vector2Add(*cul, neighbor.Position)
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
	// rl.TraceLog(rl.LogDebug, "The boid: %v", b)
	// prev := b.Position
	b.Direction = rl.Vector2Add(b.Direction, b.BoundaryV)
	b.Position = rl.Vector2Add(b.Position, rl.Vector2Scale(b.Direction, b.Speed))
	// rl.TraceLog(rl.LogDebug, "The boid at: %v should move to %v	", prev, b.Position)
}

func (b *Boid) Draw() {
	rl.DrawCircleLines(int32(b.Position.X), int32(b.Position.Y), b.Radius, rl.Black)
}

func (b *Boid) DrawDebug(factors *Factors) {
	rl.DrawCircleLines(int32(b.Position.X), int32(b.Position.Y), float32(factors.Separation), rl.Red)
	rl.DrawCircleLines(int32(b.Position.X), int32(b.Position.Y), float32(factors.Fov), rl.Green)
	rl.DrawLineV(b.Position, rl.Vector2Add(b.Position, rl.Vector2Scale(b.Direction, 0.5)), rl.Purple)
	rl.DrawLineV(b.Position, rl.Vector2Add(b.Position, rl.Vector2Scale(b.BoundaryV, 0.5)), rl.Red)
	rl.DrawLineV(b.Position, rl.Vector2Add(b.Position, rl.Vector2Scale(rl.Vector2Add(b.Direction, b.BoundaryV), 0.5)), rl.Gold)
}
