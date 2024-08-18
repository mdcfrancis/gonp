package np

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinSpace(t *testing.T) {
	expected := NpArray{0, 0.5, 1}
	assert.Equal(t, expected, LinSpace(0, 1, 3))
}

func TestZeros(t *testing.T) {
	expected := NpArray{0, 0, 0}
	assert.Equal(t, expected, Zeros(3))
}

func TestDivFloat64(t *testing.T) {
	arr := NpArray{1, 2, 3}
	expected := NpArray{1, 0.5, 1.0 / 3.0}
	assert.Equal(t, expected, DivFloat64(1, arr))
}

func TestExp(t *testing.T) {
	arr := NpArray{0, 1, 2}
	expected := NpArray{1, 2.718281828459045, 7.3890560989306495}
	ret := Exp(arr)
	// assert that ret and expected are almost equal
	assert.True(t, ret.AlmostEqual(expected, 0.00001))
}

func TestArrange(t *testing.T) {
	expected := NpArray{0, 0.5}
	assert.Equal(t, expected, Arrange(0, 1, 0.5))
}
