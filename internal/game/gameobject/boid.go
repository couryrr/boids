package gameobject

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// FIXME: just doing thing badly for the moment...
var (
	BoundaryDistance float32 = 80
	Fov              float64 = 125
	Separation       float64 = 50
	BoundaryFactor   float32 = 1
	BoundaryScale    float32 = 50
	SeparationScale  float32 = 0.15
	AlignmentScale   float32 = 0.25
	CohesionScale    float32 = 0.25
	ShouldSeparate   bool    = true
	ShouldAlign      bool    = true
	ShouldCohesion   bool    = true
)

type Boid struct {
	Id            int
	Radius        float32
	Position      rl.Vector2 // vector used for the x, y as a point
	PrevDirection rl.Vector2
	Direction     rl.Vector2 // normalized vector for the direction
	SeparateV     rl.Vector2
	AlignV        rl.Vector2
	CohesionV     rl.Vector2
	BoundaryV     rl.Vector2
	Speed         float32 // constant motion
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

func (b *Boid) GetSteeringForces(flock *Flock) {
	// b.PrevDirection = b.Direction
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
	b.BoundaryV = rl.Vector2Scale(rl.Vector2Normalize(b.Boundary()), BoundaryScale)
}

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

func (b Boid) Boundary() rl.Vector2 {
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())

	force := rl.Vector2Zero()
	if b.Position.X-float32(Fov) <= BoundaryDistance {
		force.X = b.Position.X + BoundaryFactor
	}
	if b.Position.X+float32(Fov) > screenWidth-BoundaryDistance {
		force.X = BoundaryFactor - b.Position.X
	}
	if b.Position.Y-float32(Fov) <= BoundaryDistance {
		force.Y = b.Position.Y + BoundaryFactor
	}
	if b.Position.Y+float32(Fov) > screenHeight-BoundaryDistance {
		force.Y = BoundaryFactor - b.Position.Y
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

func (b *Boid) DrawDebug() {
	rl.DrawCircleLines(int32(b.Position.X), int32(b.Position.Y), float32(Separation), rl.Red)
	rl.DrawCircleLines(int32(b.Position.X), int32(b.Position.Y), float32(Fov), rl.Green)
	rl.DrawLineV(b.Position, rl.Vector2Add(b.Position, rl.Vector2Scale(b.Direction, 0.5)), rl.Purple)
	rl.DrawLineV(b.Position, rl.Vector2Add(b.Position, rl.Vector2Scale(b.BoundaryV, 0.5)), rl.Red)
	rl.DrawLineV(b.Position, rl.Vector2Add(b.Position, rl.Vector2Scale(rl.Vector2Add(b.Direction, b.BoundaryV), 0.5)), rl.Gold)
}
