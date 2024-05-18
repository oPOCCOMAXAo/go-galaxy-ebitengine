package physics

type Modifier interface {
	Apply([]*Object) []*Object
}
