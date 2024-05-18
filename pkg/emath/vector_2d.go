package emath

import (
	"image"
	"math"
)

type Vector2D struct {
	X, Y float64
}

func V2D(x, y float64) Vector2D {
	return Vector2D{X: x, Y: y}
}

func FromImagePoint(pt image.Point) Vector2D {
	return Vector2D{X: float64(pt.X), Y: float64(pt.Y)}
}

func (v Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{v.X + v2.X, v.Y + v2.Y}
}

func (v Vector2D) Sub(v2 Vector2D) Vector2D {
	return Vector2D{v.X - v2.X, v.Y - v2.Y}
}

func (v Vector2D) Mul(s float64) Vector2D {
	return Vector2D{v.X * s, v.Y * s}
}

func (v Vector2D) Div(s float64) Vector2D {
	return Vector2D{v.X / s, v.Y / s}
}

func (v Vector2D) Len() float64 {
	return math.Hypot(v.X, v.Y)
}

func (v Vector2D) Len2() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vector2D) Normalize() Vector2D {
	l := v.Len()
	if l == 0 {
		return Vector2D{}
	}

	return v.Div(l)
}

func (v Vector2D) Dot(v2 Vector2D) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v Vector2D) Angle(v2 Vector2D) float64 {
	return math.Acos(v.Dot(v2) / (v.Len() * v2.Len()))
}

func (v Vector2D) Rotate(angle float64) Vector2D {
	sin, cos := math.Sincos(angle)
	return Vector2D{
		X: v.X*cos - v.Y*sin,
		Y: v.X*sin + v.Y*cos,
	}
}

func (v Vector2D) AngleTo(v2 Vector2D) float64 {
	return math.Atan2(v2.Y-v.Y, v2.X-v.X)
}

func (v Vector2D) Distance(v2 Vector2D) float64 {
	return v.Sub(v2).Len()
}

func (v Vector2D) Distance2(v2 Vector2D) float64 {
	return v.Sub(v2).Len2()
}
