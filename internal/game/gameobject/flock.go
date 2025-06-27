package gameobject

import (
	"iter"

	"github.com/couryrr/boids/internal/game/util"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Flock struct {
	boids []*Boid
}

func CreateFlock(quantity int) *Flock {
	boids := make([]*Boid, 0, quantity)

	for i := range quantity {
		boids = append(boids, CreateBoid(
			i,
			util.RandomVector2(BoundaryDistance, float32(rl.GetScreenHeight()-int(BoundaryDistance))),
			rl.Vector2Normalize(util.RandomVector2(BoundaryDistance, float32(rl.GetScreenHeight()-int(BoundaryDistance)))), 
			0.2))

	}
	return &Flock{
		boids: boids,
	}
}

func (f *Flock) Add(pos rl.Vector2) {
	boid := CreateBoid(len(f.boids), pos, util.RandomVector2(BoundaryDistance, BoundaryDistance+200), 0.2)
	f.boids = append(f.boids, boid)
}

func (f *Flock) All() iter.Seq[*Boid] {
	return func(yield func(*Boid) bool) {
		for _, v := range f.boids {
			if !yield(v) {
				return
			}
		}
	}
}
