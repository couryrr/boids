package gameobject

import (
	"iter"

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
			RandomVector2(),
			rl.Vector2Normalize(RandomVector2())))
	}
	return &Flock{
		boids: boids,
	}
}

func (f *Flock) Add(pos rl.Vector2) {
	boid := CreateBoid(len(f.boids), pos, rl.Vector2Zero())
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
