package util

import (
	"errors"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ForceAccumulator struct {
	total *rl.Vector2
	count float32
}

func (fa *ForceAccumulator) Average() (rl.Vector2, error) {
	if fa.count == 0 {
		return rl.Vector2Zero(), errors.New("No value")
	}
	return rl.Vector2{
		X: fa.total.X / fa.count,
		Y: fa.total.Y / fa.count,
	}, nil
}

func (fa *ForceAccumulator) Value() (rl.Vector2, error) {
	if fa.count == 0 {
		return rl.Vector2Zero(), errors.New("No value")
	}
	return *fa.total, nil
}

func (fa *ForceAccumulator) Increment(vec rl.Vector2) {
	fa.count += 1
	if fa.total == nil {
		fa.total = &rl.Vector2{
			X: 0,
			Y: 0,
		}
	}
	temp := rl.Vector2Add(*fa.total, vec)
	fa.total = &temp
}

func RandomVector2(min, max float32) rl.Vector2 {
	return rl.Vector2{
		X: float32(rl.GetRandomValue(int32(min), int32(max))),
		Y: float32(rl.GetRandomValue(int32(min), int32(max))),
	}
}
