package physics

import "github.com/opoccomaxao/go-galaxy-experiment/pkg/emath"

type Object struct {
	Position     emath.Vector2D
	Velocity     emath.Vector2D
	Acceleration emath.Vector2D
	Size         float64
	Mass         float64
}

func (o *Object) MinMax() (emath.Vector2D, emath.Vector2D) {
	sizeV := emath.Vector2D{X: o.Size, Y: o.Size}

	return o.Position.Sub(sizeV), o.Position.Add(sizeV)
}
