package matrix

import (
	"fmt"
	"github.com/adrianh-za/go-max7219"
	"github.com/robertszeker/snake-on-max-7219-go/object"
)

func (m *Matrix) PrintObjects() error {

	buffer := make(map[int]map[int]byte, m.cascaded)
	for cascadedId := 0; cascadedId < m.cascaded; cascadedId++ {
		buffer[cascadedId] = make(map[int]byte, max7219.MAX7219_REG_LASTDIGIT)
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

func (m *Matrix) getCascadedId(point *object.Point) int {
	return (point.X - 1) / 8
}

func (m *Matrix) getPosition(point *object.Point) int {
	return max7219.MAX7219_REG_LASTDIGIT - point.Y
}

func (m *Matrix) getValue(point *object.Point) byte {
	a := max7219.MAX7219_REG_LASTDIGIT - point.X + max7219.MAX7219_REG_LASTDIGIT*m.getCascadedId(point)

	return 1 << a
}
