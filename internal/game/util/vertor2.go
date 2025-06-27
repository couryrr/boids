package util

import rl "github.com/gen2brain/raylib-go/raylib"

func RandomVector2(min, max float32) rl.Vector2 {
	return rl.Vector2{
		X: float32(rl.GetRandomValue(int32(min), int32(max))),
		Y: float32(rl.GetRandomValue(int32(min), int32(max))),
	}
}
