package np

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNpStack_String(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3},
		NpArray{4, 5, 6},
	}
	expected := "[[1.00000000 2.00000000 3.00000000 ]\n[4.00000000 5.00000000 6.00000000 ]\n]"
	assert.Equal(t, expected, stack.String())
}

func TestNpStack_Shuffle(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3},
		NpArray{4, 5, 6},
	}
	shuffled := stack.Shuffle([]int{1, 0})
	expected := NpStack{
		NpArray{4, 5, 6},
		NpArray{1, 2, 3},
	}
	assert.Equal(t, expected, shuffled)
}

func TestNpStack_Mean(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3},
		NpArray{4, 5, 6},
	}
	expected := NpArray{2, 5}
	assert.Equal(t, expected, stack.Mean())
}

func TestNpStack_StandardDeviation(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3},
		NpArray{4, 5, 6},
	}
	expected := NpArray{0.816496580927726, 0.816496580927726}
	assert.Equal(t, expected, stack.StandardDeviation())
}

func TestNpStack_ScalarStandardDeviation(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3},
		NpArray{4, 5, 6},
	}
	assert.Equal(t, 1.707825127659933, stack.ScalarStandardDeviation())
}

func TestNpStack_ScalarMean(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3},
		NpArray{4, 5, 6},
	}
	assert.Equal(t, 3.5, stack.ScalarMean())
}

func TestNpStack_Sub(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3},
		NpArray{4, 5, 6},
	}
	arr := NpArray{1, 1}
	expected := NpStack{
		NpArray{0, 1, 2},
		NpArray{3, 4, 5},
	}
	assert.Equal(t, expected, stack.Sub(arr))
}

func TestNpStack_SubSlice(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3, 4},
		NpArray{5, 6, 7, 8},
	}
	arr := NpArray{1, 1, 1, 1}
	expected := NpStack{
		NpArray{1, 2},
		NpArray{5, 6},
	}
	assert.Equal(t, expected, stack.SubSlice(arr, 1, 3))
}

func TestNpStack_SubFloat64(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3},
		NpArray{4, 5, 6},
	}
	expected := NpStack{
		NpArray{0, 1, 2},
		NpArray{3, 4, 5},
	}
	assert.Equal(t, expected, stack.SubFloat64(1))
}

func TestNpStack_LinSpace(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3},
		NpArray{4, 5, 6},
	}
	expected := NpStack{
		NpArray{0, 0.3333333333333333, 0.6666666666666666},
		NpArray{0, 0.3333333333333333, 0.6666666666666666},
	}
	assert.Equal(t, expected, stack.LinSpace(0, 1, 3))
}

func TestNpStack_DivFloat64(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3},
		NpArray{4, 5, 6},
	}
	expected := NpStack{
		NpArray{0.5, 1, 1.5},
		NpArray{2, 2.5, 3},
	}
	assert.Equal(t, expected, stack.DivFloat64(2))
}

func TestNpStack_Slice(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3, 4},
		NpArray{5, 6, 7, 8},
	}
	expected := NpStack{
		NpArray{2, 3},
		NpArray{6, 7},
	}
	assert.Equal(t, expected, stack.Slice(1, 3))
}

func TestNpStack_Cond(t *testing.T) {
	stack := NpStack{
		NpArray{1, 0, 3},
		NpArray{4, 5, 6},
	}
	expected := NpStack{
		NpArray{1, 0, 1},
		NpArray{1, 1, 1},
	}
	assert.Equal(t, expected, stack.Cond(func(x float64) float64 {
		if x != 0 {
			return 1
		}
		return 0
	}))
}

func TestNpStack_LinearInterpolate(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3},
		NpArray{4, 5, 6},
	}
	scale := NpArray{0, 0.5, 1}
	newScale := NpArray{0, 0.25, 0.5, 0.75, 1}
	expected := NpStack{
		NpArray{1, 1.5, 2, 2.5, 3},
		NpArray{4, 4.5, 5, 5.5, 6},
	}
	assert.Equal(t, expected, stack.LinearInterpolate(scale, newScale))
}

func TestNpStack_DivSlice(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3},
		NpArray{4, 5, 6},
	}
	expected := NpStack{
		NpArray{0.3333333333333333, 0.6666666666666666},
		NpArray{1.3333333333333333, 1.6666666666666667},
	}
	assert.Equal(t, expected, stack.DivSlice(3, 0, 2))
}

func TestNpStack_Add(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3},
		NpArray{4, 5, 6},
	}
	other := NpStack{
		NpArray{1, 1, 1},
		NpArray{1, 1, 1},
	}
	expected := NpStack{
		NpArray{2, 3, 4},
		NpArray{5, 6, 7},
	}
	assert.Equal(t, expected, stack.Add(other))
}

func TestNpStack_MulFloat64(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3},
		NpArray{4, 5, 6},
	}
	expected := NpStack{
		NpArray{2, 4, 6},
		NpArray{8, 10, 12},
	}
	assert.Equal(t, expected, stack.MulFloat64(2))
}

func TestNpStack_Max(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3},
		NpArray{4, 5, 6},
	}
	expected := NpArray{3, 6}
	assert.Equal(t, expected, stack.Max())
}

func TestNpStack_Min(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3},
		NpArray{4, 5, 6},
	}
	expected := NpArray{1, 4}
	assert.Equal(t, expected, stack.Min())
}

func TestNpStack_Shape(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3},
	}
	shape := stack.Shape()
	assert.Equal(t, 1, shape[0])
	assert.Equal(t, 3, shape[1])
}

func TestNpStack_Column(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3},
	}
	expected := NpArray{1}
	assert.Equal(t, expected, stack.Column(0))
}

func TestNpStack_Div(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3},
	}
	expected := NpStack{
		NpArray{0.5, 1, 1.5},
	}
	assert.Equal(t, expected, stack.Div(NpArray{2}))
}

func TestNpStack_AllMostEqual(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3},
	}
	other := NpStack{
		NpArray{1, 2, 3},
	}
	assert.True(t, stack.AllMostEqual(other, 1e-9))
}

func TestNpStack_AllMostEqualRetunsFalse(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3},
	}
	other := NpStack{
		NpArray{1, 2, 3.1},
	}
	assert.False(t, stack.AllMostEqual(other, 0.01))
}

func TestNpStack_Sum(t *testing.T) {
	stack := NpStack{
		NpArray{1, 2, 3},
	}
	expected := NpArray{6}
	assert.Equal(t, expected, stack.Sum())
}
