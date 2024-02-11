package matrix

import (
	"fmt"
	"github.com/robertszeker/snake-on-max-7219-go/object"
)

const LedsInARow = 8

func (m *Matrix) PrintObjects() error {

	buffer := make(map[int]map[int]byte, m.cascaded)
	for cascadedId := 0; cascadedId < m.cascaded; cascadedId++ {
		buffer[cascadedId] = make(map[int]byte, 8)
	}

	for _, obj := range m.objects {
		for _, point := range obj.GetPoints() {
			if existingValue, ok := buffer[m.getCascadedId(point)][m.getPosition(point)]; ok {
				buffer[m.getCascadedId(point)][m.getPosition(point)] = existingValue | m.getValue(point)
			} else {
				buffer[m.getCascadedId(point)][m.getPosition(point)] = m.getValue(point)
			}
		}
	}

	for cascadedId, positionValueMapping := range buffer {
		for position, value := range positionValueMapping {
			err := m.Device.SetBufferLine(
				cascadedId,
				position,
				value,
				true,
			)

			if err != nil {
				return fmt.Errorf("failed to print objects: %w", err)
			}
		}
	}

	return nil
}

func (m *Matrix) printPoint(point *object.Point) error {
	if err := m.Device.SetBufferLine(
		m.getCascadedId(point),
		m.getPosition(point),
		m.getValue(point),
		true,
	); err != nil {
		return fmt.Errorf("failed to set buffer line: %w", err)
	}

	return nil
}

func (m *Matrix) getCascadedId(point *object.Point) int {
	return (point.X - 1) / 8
}

func (m *Matrix) getPosition(point *object.Point) int {
	return LedsInARow - point.Y
}

func (m *Matrix) getValue(point *object.Point) byte {
	a := LedsInARow - point.X + LedsInARow*m.getCascadedId(point)

	return 1 << a
}
