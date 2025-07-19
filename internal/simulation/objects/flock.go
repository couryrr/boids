package objects

import (
	"iter"
	"math/rand"

	"github.com/couryrr/boids/internal/simulation/util"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Flock struct {
	boids []*Boid
}

func CreateFlock(boundry float32, quantity int, factors *Factors) *Flock {
	boids := make([]*Boid, 0, quantity)

	for i := range quantity {
		pos := util.RandomVector2(boundry, float32(rl.GetScreenHeight()-int(boundry)))
		direction := rl.Vector2Normalize(util.RandomVector2(boundry, float32(rl.GetScreenHeight()-int(boundry))))
		speed := float64(rand.Int63n(factors.MaxSpeed-factors.MinSpeed) + factors.MaxSpeed)
		boids = append(boids, CreateBoid(
			i,
			pos,
			direction,
			speed))
	}
	return &Flock{
		boids: boids,
	}
}

func (f *Flock) Add(boundry float32, pos rl.Vector2, factors *Factors) {
	direction := rl.Vector2Normalize(util.RandomVector2(0, 100))
	speed := float64(rand.Int63n(factors.MaxSpeed-factors.MinSpeed) + factors.MaxSpeed)
	boid := CreateBoid(len(f.boids), pos, direction, speed)
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
