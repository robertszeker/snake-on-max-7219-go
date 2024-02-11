package matrix

import (
	"fmt"
	"github.com/adrianh-za/go-max7219"
	"github.com/robertszeker/snake-on-max-7219-go/object"
)

type Matrix struct {
	*max7219.Matrix
	objects  []object.Printable
	cascaded int
}

func NewMatrix(cascaded int) (*Matrix, error) {

	mtx := max7219.NewMatrix(cascaded, max7219.RotateClockwise)
	err := mtx.Open(0, 0, 7)
	if err != nil {
		return nil, fmt.Errorf("failed to open matrix: %w", err)
	}

	return &Matrix{
		Matrix:   mtx,
		cascaded: cascaded,
	}, nil
}

func (m *Matrix) Close() {
	m.Matrix.Close()
}

func (m *Matrix) AddObject(object object.Printable) {
	m.objects = append(m.objects, object)
}
