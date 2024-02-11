package object

type Points []*Point

type Printable interface {
	GetPoints() Points
}

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) *Point {
	return &Point{x, y}
}
