package object

type Mouse struct {
	point *Point
}

func NewMouse(point *Point) *Mouse {
	return &Mouse{point}
}

func (m *Mouse) GetPoints() Points {
	return []*Point{m.point}
}

var _ Printable = &Mouse{}
