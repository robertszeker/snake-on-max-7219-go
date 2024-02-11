package matrix

import (
	"github.com/robertszeker/snake-on-max-7219-go/object"
	"github.com/stretchr/testify/assert"
	"testing"
)

var getCascadedIdTestCases = []struct {
	name               string
	point              *object.Point
	cascaded           int
	expectedCascadedId int
}{
	{
		"hitting first quadrant, lower boundary",
		object.NewPoint(0, 2),
		4,
		0,
	},
	{
		"hitting first quadrant, upper boundary",
		object.NewPoint(7, 2),
		4,
		0,
	},
	{
		"hitting second quadrant, lower boundary",
		object.NewPoint(8, 2),
		4,
		1,
	},
	{
		"hitting second quadrant, upper boundary",
		object.NewPoint(15, 2),
		4,
		1,
	},
	{
		"hitting third quadrant, lower boundary",
		object.NewPoint(16, 2),
		4,
		2,
	},
	{
		"hitting third quadrant, upper boundary",
		object.NewPoint(23, 2),
		4,
		2,
	},
	{
		"hitting fourth quadrant, lower boundary",
		object.NewPoint(24, 2),
		4,
		3,
	},
	{
		"hitting fourth quadrant, upper boundary",
		object.NewPoint(31, 2),
		4,
		3,
	},
}

func TestGetCascaded(t *testing.T) {
	for _, tc := range getCascadedIdTestCases {
		t.Run(tc.name, func(t *testing.T) {
			m := Matrix{cascaded: tc.cascaded}
			actualCascadedId := m.getCascadedId(tc.point)
			assert.Equal(t, tc.expectedCascadedId, actualCascadedId)
		})
	}
}

var getValueTestCases = []struct {
	name         string
	point        *object.Point
	expectedByte byte
}{
	{
		"printing first bit",
		object.NewPoint(0, 2),
		1,
	},
	{
		"printing second bit",
		object.NewPoint(1, 2),
		2,
	},
	{
		"printing third bit",
		object.NewPoint(2, 2),
		4,
	},
}

func TestGetValue(t *testing.T) {
	for _, tc := range getValueTestCases {
		t.Run(tc.name, func(t *testing.T) {
			m := Matrix{cascaded: 4}
			actualByte := m.getValue(tc.point)
			assert.Equal(t, tc.expectedByte, actualByte)
		})
	}
}
