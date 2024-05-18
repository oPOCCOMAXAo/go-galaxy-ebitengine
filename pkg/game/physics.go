package game

import (
	"math/rand"

	"github.com/opoccomaxao/go-galaxy-experiment/pkg/emath"
	"github.com/opoccomaxao/go-galaxy-experiment/pkg/physics"
)

func newPhysics(bounds emath.Rect2D) *physics.Engine {
	res := physics.New()

	res.AddModifier(physics.NewGravity(100))
	res.AddModifier(physics.NewBoundRemover(bounds))
	res.AddModifier(physics.NewZeroMassRemover())
	res.AddModifier(physics.NewMotionMod(float64(1) / 60))

	total := 100

	for i := 0; i < total; i++ {
		res.AddObject(randomObject(bounds))
	}

	return res
}

func randomObject(bounds emath.Rect2D) *physics.Object {
	return &physics.Object{
		Position: emath.V2D(
			bounds.Min.X+(bounds.Max.X-bounds.Min.X)*rand.Float64(),
			bounds.Min.Y+(bounds.Max.Y-bounds.Min.Y)*rand.Float64(),
		),
		Velocity: emath.V2D(0, 0),
		Size:     5,
		Mass:     1,
	}
}
