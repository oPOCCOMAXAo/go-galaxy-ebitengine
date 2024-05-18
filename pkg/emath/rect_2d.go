package emath

type Rect2D struct {
	Min Vector2D
	Max Vector2D
}

func (r Rect2D) ContainsV(p Vector2D) bool {
	return p.X >= r.Min.X && p.X <= r.Max.X && p.Y >= r.Min.Y && p.Y <= r.Max.Y
}

func (r Rect2D) ContainsMinMax(min, max Vector2D) bool {
	return min.X >= r.Min.X && min.Y >= r.Min.Y && max.X <= r.Max.X && max.Y <= r.Max.Y
}
