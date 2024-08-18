package np

import (
	"github.com/montanaflynn/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNpArray_String(t *testing.T) {
	arr := NpArray{5, 3, 3, 3.4, 3.8, 4.2, 4.6, 5, 5.4, 5.8, 5, 5}
	expected := "[5.00000000 3.00000000 3.00000000 3.40000000 3.80000000 4.20000000 4.60000000 5.00000000 5.40000000 5.80000000 5.00000000 5.00000000 ]"
	assert.Equal(t, expected, arr.String())
}

func TestNpArray_Shape(t *testing.T) {
	arr := NpArray{1, 2, 3}
	assert.Equal(t, 3, arr.Shape())
}

func TestNpArray_Copy(t *testing.T) {
	arr := NpArray{1, 2, 3}
	copyArr := arr.Copy()
	assert.Equal(t, arr, copyArr)
	assert.NotSame(t, &arr, &copyArr)
}

func TestNpArray_Shuffle(t *testing.T) {
	arr := NpArray{1, 2, 3}
	shuffled := arr.Shuffle([]int{2, 0, 1})
	expected := NpArray{3, 1, 2}
	assert.Equal(t, expected, shuffled)
}

func TestNpArray_Mean(t *testing.T) {
	arr := NpArray{1, 2, 3}
	expected, _ := stats.Mean(stats.Float64Data(arr))
	assert.Equal(t, expected, arr.Mean())
}

func TestNpArray_StandardDeviation(t *testing.T) {
	arr := NpArray{1, 2, 3}
	expected, _ := stats.StandardDeviation(stats.Float64Data(arr))
	assert.Equal(t, expected, arr.StandardDeviation())
}

func TestNpArray_Sub(t *testing.T) {
	arr1 := NpArray{1, 2, 3}
	arr2 := NpArray{1, 1, 1}
	expected := NpArray{0, 1, 2}
	assert.Equal(t, expected, arr1.Sub(arr2))
}

func TestNpArray_Div(t *testing.T) {
	arr1 := NpArray{1, 2, 3}
	arr2 := NpArray{1, 2, 3}
	expected := NpArray{1, 1, 1}
	assert.Equal(t, expected, arr1.Div(arr2))
}

func TestNpArray_Mul(t *testing.T) {
	arr1 := NpArray{1, 2, 3}
	arr2 := NpArray{1, 2, 3}
	expected := NpArray{1, 4, 9}
	assert.Equal(t, expected, arr1.Mul(arr2))
}

func TestNpArray_Add(t *testing.T) {
	arr1 := NpArray{1, 2, 3}
	arr2 := NpArray{1, 2, 3}
	expected := NpArray{2, 4, 6}
	assert.Equal(t, expected, arr1.Add(arr2))
}

func TestNpArray_DivFloat64(t *testing.T) {
	arr := NpArray{1, 2, 3}
	expected := NpArray{0.5, 1, 1.5}
	assert.Equal(t, expected, arr.DivFloat64(2))
}

func TestNpArray_MulFloat64(t *testing.T) {
	arr := NpArray{1, 2, 3}
	expected := NpArray{2, 4, 6}
	assert.Equal(t, expected, arr.MulFloat64(2))
}

func TestNpArray_SubSlice(t *testing.T) {
	arr1 := NpArray{1, 2, 3, 4}
	arr2 := NpArray{1, 1, 1, 1}
	expected := NpArray{1, 2}
	assert.Equal(t, expected, arr1.SubSlice(arr2, 1, 3))
}

func TestNpArray_SubFloat64(t *testing.T) {
	arr := NpArray{1, 2, 3}
	expected := NpArray{0, 1, 2}
	assert.Equal(t, expected, arr.SubFloat64(1))
}

func TestNpArray_LinearInterpolate(t *testing.T) {
	arr := NpArray{1, 2, 3}
	scale := NpArray{0, 1, 2}
	newScale := NpArray{0.5, 1.5}
	expected := NpArray{1.5, 2.5}
	assert.Equal(t, expected, arr.LinearInterpolate(scale, newScale))
}

func TestNpArray_LinearInterpolatePanics(t *testing.T) {
	arr := NpArray{1, 2, 3}
	scale := NpArray{0, 1}
	newScale := NpArray{0.5, 1.5}
	assert.Panics(t, func() { arr.LinearInterpolate(scale, newScale) })
}

func TestNpArray_Cond(t *testing.T) {
	arr := NpArray{1, 2, 3}
	expected := NpArray{2, 4, 6}
	assert.Equal(t, expected, arr.Cond(func(v float64) float64 { return v * 2 }))
}

func TestNpArray_AddFloat64(t *testing.T) {
	arr := NpArray{1, 2, 3}
	expected := NpArray{2, 3, 4}
	assert.Equal(t, expected, arr.AddFloat64(1))
}

func TestNpArray_DivSlice(t *testing.T) {
	arr := NpArray{1, 2, 3, 4}
	expected := NpArray{1, 1.5}
	assert.Equal(t, expected, arr.DivSlice(2, 1, 3))
}

func TestNpArray_Dot(t *testing.T) {
	arr1 := NpArray{1, 2, 3}
	arr2 := NpArray{1, 2, 3}
	expected := 14.0
	assert.Equal(t, expected, arr1.Dot(arr2))
}

func TestNpArray_Max(t *testing.T) {
	arr := NpArray{1, 2, 3}
	expected := 3.0
	assert.Equal(t, expected, arr.Max())
}

func TestNpArray_PowFloat64(t *testing.T) {
	arr := NpArray{1, 2, 3}
	expected := NpArray{1, 4, 9}
	assert.Equal(t, expected, arr.PowFloat64(2))
}

func TestNpArray_Min(t *testing.T) {
	arr := NpArray{1, 2, 3}
	expected := 1.0
	assert.Equal(t, expected, arr.Min())
}

func TestNpArray_Sum(t *testing.T) {
	arr := NpArray{1, 2, 3}
	expected := 6.0
	assert.Equal(t, expected, arr.Sum())
}

func TestNpArray_AlmostEqualReturnsTrueForEqualArrays(t *testing.T) {
	arr1 := NpArray{1.0, 2.0, 3.0}
	arr2 := NpArray{1.0, 2.0, 3.0}
	assert.True(t, arr1.AlmostEqual(arr2, 0.01))
}

func TestNpArray_AlmostEqualReturnsFalseForDifferentLengthArrays(t *testing.T) {
	arr1 := NpArray{1.0, 2.0, 3.0}
	arr2 := NpArray{1.0, 2.0}
	assert.False(t, arr1.AlmostEqual(arr2, 0.01))
}

func TestNpArray_AlmostEqualReturnsFalseForArraysWithDifferenceGreaterThanEpsilon(t *testing.T) {
	arr1 := NpArray{1.0, 2.0, 3.0}
	arr2 := NpArray{1.0, 2.0, 3.1}
	assert.False(t, arr1.AlmostEqual(arr2, 0.01))
}

func TestNpArray_AlmostEqualReturnsTrueForArraysWithDifferenceLessThanEpsilon(t *testing.T) {
	arr1 := NpArray{1.0, 2.0, 3.0}
	arr2 := NpArray{1.0, 2.0, 3.001}
	assert.True(t, arr1.AlmostEqual(arr2, 0.01))
}

func TestNpArray_DivPanicsForDifferentLengthArrays(t *testing.T) {
	arr1 := NpArray{4.0, 9.0, 16.0}
	arr2 := NpArray{2.0, 3.0}
	assert.Panics(t, func() { arr1.Div(arr2) })
}

func TestNpArray_SubPanicsForDifferentLengthArrays(t *testing.T) {
	arr1 := NpArray{4.0, 9.0, 16.0}
	arr2 := NpArray{2.0, 3.0}
	assert.Panics(t, func() { arr1.Sub(arr2) })
}

func TestNpArray_AddPanicsForDifferentLengthArrays(t *testing.T) {
	arr1 := NpArray{4.0, 9.0, 16.0}
	arr2 := NpArray{2.0, 3.0}
	assert.Panics(t, func() { arr1.Add(arr2) })
}

func TestNpArray_MulPanicsForDifferentLengthArrays(t *testing.T) {
	arr1 := NpArray{4.0, 9.0, 16.0}
	arr2 := NpArray{2.0, 3.0}
	assert.Panics(t, func() { arr1.Mul(arr2) })
}

func TestNpArray_DotPanicsForDifferentLengthArrays(t *testing.T) {
	arr1 := NpArray{4.0, 9.0, 16.0}
	arr2 := NpArray{2.0, 3.0}
	assert.Panics(t, func() { arr1.Dot(arr2) })
}
