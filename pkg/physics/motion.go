package physics

type MotionMod struct {
	T float64
}

func NewMotionMod(t float64) *MotionMod {
	return &MotionMod{T: t}
}

func (m *MotionMod) Apply(objs []*Object) []*Object {
	for _, obj := range objs {
		obj.Position = obj.Position.Add(obj.Velocity.Mul(m.T))
		obj.Velocity = obj.Velocity.Add(obj.Acceleration.Mul(m.T))
	}

	return objs
}
