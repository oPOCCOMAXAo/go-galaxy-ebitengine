package physics

import "github.com/opoccomaxao/go-galaxy-experiment/pkg/emath"

var _ Modifier = &BoundRemoverMod{}

type BoundRemoverMod struct {
	bounds emath.Rect2D
}

func NewBoundRemover(bounds emath.Rect2D) *BoundRemoverMod {
	return &BoundRemoverMod{
		bounds: bounds,
	}
}

func (b *BoundRemoverMod) Apply(obj []*Object) []*Object {
	res := make([]*Object, 0, len(obj))

	for _, o := range obj {
		if b.bounds.ContainsV(o.Position) {
			res = append(res, o)
		}
	}

	return res
}

var _ Modifier = &ZeroMassRemoverMod{}

type ZeroMassRemoverMod struct{}

func NewZeroMassRemover() *ZeroMassRemoverMod {
	return &ZeroMassRemoverMod{}
}

func (z *ZeroMassRemoverMod) Apply(obj []*Object) []*Object {
	res := make([]*Object, 0, len(obj))

	for _, o := range obj {
		if o.Mass > 0 {
			res = append(res, o)
		}
	}

	return res
}
