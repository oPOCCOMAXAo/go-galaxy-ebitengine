package physics

var _ Modifier = (*Gravity)(nil)

func NewGravity(acceleration float64) *Gravity {
	return &Gravity{Acceleration: acceleration}
}

type Gravity struct {
	Acceleration float64
}

func (g *Gravity) Apply(objects []*Object) []*Object {
	total := len(objects)

	for i := 0; i < total; i++ {
		for j := i + 1; j < total; j++ {
			obj1 := objects[i]
			obj2 := objects[j]

			if obj1.Mass == 0 || obj2.Mass == 0 {
				continue
			}

			dist2 := obj1.Position.Sub(obj2.Position).Len2()
			if dist2 < 1 {
				dist2 = 1
			}

			force := g.Acceleration * obj1.Mass * obj2.Mass / dist2

			direction := obj2.Position.Sub(obj1.Position).Normalize()

			obj1.Acceleration = obj1.Acceleration.Add(direction.Mul(force / obj1.Mass))
			obj2.Acceleration = obj2.Acceleration.Add(direction.Mul(-force / obj2.Mass))
		}
	}

	return objects
}
