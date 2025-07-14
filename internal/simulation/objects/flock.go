package objects

import (
	"iter"

	"github.com/couryrr/boids/internal/simulation/util"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Flock struct {
	boids []*Boid
}

func CreateFlock(boundry float32, quantity int) *Flock {
	boids := make([]*Boid, 0, quantity)

	for i := range quantity {
		boids = append(boids, CreateBoid(
			i,
			util.RandomVector2(boundry, float32(rl.GetScreenHeight()-int(boundry))),
			rl.Vector2Normalize(util.RandomVector2(boundry, float32(rl.GetScreenHeight()-int(boundry)))),
			0.2))

	}
	return &Flock{
		boids: boids,
	}
}

func (f *Flock) Add(boundry float32, pos rl.Vector2) {
	boid := CreateBoid(len(f.boids), pos, util.RandomVector2(0, 100), 0.2)
	rl.TraceLog(rl.LogDebug, "Made: %v", boid.Direction)
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
