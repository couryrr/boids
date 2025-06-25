package gameobjects

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	BoundaryDistance float32 = 80
	Fov              float64 = 0
	Separation       float64 = 20
	BoundaryFactor   float32 = 0.25
	BoundaryScale    float32 = 1.5
	SeparationScale  float32 = 0.15
	AlignmentScale   float32 = 1
	CohesionScale    float32 = 1
)

type Boid struct {
	Id        int
	Radius    float32
	Position  rl.Vector2 // vector used for the x, y as a point
	Direction rl.Vector2 // normalized vector for the direction
	Speed     float32    // constant motion
}

func RandomVector2() rl.Vector2 {
	return rl.Vector2{
		X: float32(rl.GetRandomValue(0, int32(rl.GetScreenWidth()))),
		Y: float32(rl.GetRandomValue(0, int32(rl.GetScreenHeight()))),
	}
}

func CreateBoid(id int, position rl.Vector2, direction rl.Vector2) *Boid {
	return &Boid{
		Id:        id,
		Radius:    5,
		Position:  position,
		Direction: direction,
		Speed:     2.5,
	}
}

func (b *Boid) AddDirection(dir rl.Vector2) {
	b.Direction = rl.Vector2Normalize(rl.Vector2Add(dir, b.Direction))
}

func (b *Boid) Steer(flock *Flock) {
	culSep := rl.Vector2Zero()
	culDir := rl.Vector2Zero()
	culPos := rl.Vector2Zero()

	for neighbor := range flock.All() {
		// b.Separate(neighbor, &culSep)
		b.Align(neighbor, &culDir)
		b.Cohesion(neighbor, &culPos)
	}

	// TODO: move the scale to inside each method?
	b.AddDirection(rl.Vector2Normalize(rl.Vector2Scale(culSep, SeparationScale)))
	b.AddDirection(rl.Vector2Normalize(rl.Vector2Scale(culDir, AlignmentScale)))
	b.AddDirection(rl.Vector2Normalize(rl.Vector2Scale(culPos, CohesionScale)))
	b.AddDirection(rl.Vector2Scale(b.Boundary(), BoundaryScale))
}

func (b *Boid) Separate(neighbor *Boid, cul *rl.Vector2) {
	distance := rl.Vector2Distance(b.Position, neighbor.Position)
	if b.Id != neighbor.Id && float64(distance) <= Separation {
		*cul = rl.Vector2Subtract(*cul, neighbor.Position)
	}
}

func (b *Boid) Align(neighbor *Boid, cul *rl.Vector2) {
	distance := rl.Vector2Distance(b.Position, neighbor.Position)
	if b.Id != neighbor.Id && distance <= float32(Fov) && float64(distance) > Separation {
		*cul = rl.Vector2Add(*cul, neighbor.Direction)
	}
}

func (b *Boid) Cohesion(neighbor *Boid, cul *rl.Vector2) {
	distance := rl.Vector2Distance(b.Position, neighbor.Position)
	if b.Id != neighbor.Id && distance <= float32(Fov) && float64(distance) > Separation {
		*cul = rl.Vector2Add(*cul, neighbor.Position)
	}
}

func (b Boid) Boundary() rl.Vector2 {
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())

	force := rl.Vector2Zero()
	if b.Position.X <= BoundaryDistance {
		force.X = b.Position.X + BoundaryFactor
	}
	if b.Position.X > screenWidth-BoundaryDistance {
		force.X = BoundaryFactor - b.Position.X
	}
	if b.Position.Y <= BoundaryDistance {
		force.Y = b.Position.Y + BoundaryFactor
	}
	if b.Position.Y > screenHeight-BoundaryDistance {
		force.Y = BoundaryFactor - b.Position.Y
	}

	return rl.Vector2Normalize(force)
}

func (b *Boid) UpdatePosition() {
	b.Position = rl.Vector2Add(b.Position, rl.Vector2Scale(b.Direction, b.Speed))
}

// func (b *Boid) Warp() {
// 	screenHeight := rl.GetScreenHeight()
// 	screenWidth := rl.GetScreenWidth()
// 	if b.Position.Y > float32(screenHeight) {
// 		b.Position.Y = 0
// 	}
// 	if b.Position.Y < 0 {
// 		b.Position.Y = float32(screenHeight)
// 	}
// 	if b.Position.X > float32(screenWidth) {
// 		b.Position.X = 0
// 	}
// 	if b.Position.X < 0 {
// 		b.Position.X = float32(screenWidth)
// 	}
// }

func (b *Boid) Draw() {
	rl.DrawCircleLines(int32(b.Position.X), int32(b.Position.Y), b.Radius, rl.Black)
	// rl.DrawCircleLines(int32(b.Position.X), int32(b.Position.Y), float32(fov), rl.Green)
	// rl.DrawCircleLines(int32(b.Position.X), int32(b.Position.Y), float32(separation), rl.Red)
}
