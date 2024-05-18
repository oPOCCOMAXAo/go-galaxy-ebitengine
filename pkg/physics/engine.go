package physics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/opoccomaxao/go-galaxy-experiment/pkg/emath"
)

type Engine struct {
	objects []*Object
	mods    []Modifier
}

func New() *Engine {
	return &Engine{}
}

func (e *Engine) AddModifier(mod Modifier) {
	e.mods = append(e.mods, mod)
}

func (e *Engine) AddObject(obj *Object) {
	e.objects = append(e.objects, obj)
}

func (e *Engine) resetAcceleration() {
	for _, obj := range e.objects {
		obj.Acceleration = emath.Vector2D{}
	}
}

func (e *Engine) Update() {
	e.resetAcceleration()

	for _, mod := range e.mods {
		e.objects = mod.Apply(e.objects)
	}
}

func (e *Engine) DrawInto(
	buffer *ebiten.Image,
	viewPort emath.Rect2D,
) {
	buffer.Fill(bgColor)

	for _, object := range e.objects {
		min, max := object.MinMax()

		if !viewPort.ContainsMinMax(max, min) {
			continue
		}

		vector.DrawFilledCircle(
			buffer,
			float32(object.Position.X),
			float32(object.Position.Y),
			float32(object.Size),
			fgColor,
			true,
		)
	}

}
