package vector

import (
	"github.com/go-gl/mathgl/mgl32"
)

func MA(start *mgl32.Vec3, scale float32, direction *mgl32.Vec3, result *mgl32.Vec3) {
	result[0] = start[0] + scale * direction[0]
	result[1] = start[1] + scale * direction[1]
	result[2] = start[2] + scale * direction[2]
}

func Scale(in *mgl32.Vec3, scale float32, result *mgl32.Vec3) {
	result[0] = in[0] * scale
	result[1] = in[1] * scale
	result[2] = in[2] * scale
}

func Avg(in mgl32.Vec3) float64 {
	return float64((in.X() + in.Y() + in.Z()) / 3)
}