package np

import (
	"fmt"
	"math"
)

type NpStack []NpArray

func (m NpStack) String() string {
	ret := "["
	for _, a := range m {
		ret += fmt.Sprintf("%v\n", a)
	}
	ret += "]"
	return ret
}

func (m NpStack) Shape() []int {
	return []int{len(m), len(m[0])}
}

// Shuffle returns a new NpStack with the rows shuffled
func (m NpStack) Shuffle(idx []int) NpStack {
	ret := make(NpStack, len(m))
	for i, a := range idx {
		ret[i] = m[a]
	}
	return ret
}

func (m NpStack) Mean() NpArray {
	ret := make(NpArray, len(m))
	for i, a := range m {
		ret[i] = a.Mean()
	}
	return ret
}

func (m NpStack) ScalarMean() float64 {
	return m.Mean().Mean()
}

func (m NpStack) StandardDeviation() NpArray {
	ret := make(NpArray, len(m))
	for i, a := range m {
		ret[i] = a.StandardDeviation()
	}
	return ret
}

func (m NpStack) ScalarStandardDeviation() float64 {
	std := m.StandardDeviation() // NpArray
	mean := m.Mean()             // NpArray
	scalarMean := m.ScalarMean()
	powStd := std.PowFloat64(2) // NpArray
	meanDelta := mean.SubFloat64(scalarMean)
	powMean := meanDelta.PowFloat64(2)
	combinedVariance := powStd.Add(powMean).DivFloat64(float64(std.Shape()))
	return math.Sqrt(combinedVariance.Sum())
}

func (m NpStack) Sub(b NpArray) NpStack {
	ret := make(NpStack, len(m))
	for i, a := range m {
		ret[i] = a.SubFloat64(b[i])
	}
	return ret
}

func (m NpStack) Div(b NpArray) NpStack {
	ret := make(NpStack, len(m))
	for i, a := range m {
		ret[i] = a.DivFloat64(b[i])
	}
	return ret
}

func (m NpStack) SubSlice(b NpArray, start, end int) NpStack {
	ret := make(NpStack, len(m))
	for i, a := range m {
		ret[i] = a.SubSlice(b, start, end)
	}
	return ret
}

func (m NpStack) DivFloat64(b float64) NpStack {
	ret := make(NpStack, len(m))
	for i, a := range m {
		ret[i] = a.DivFloat64(b)
	}
	return ret
}

func (m NpStack) Slice(start, end int) NpStack {
	ret := make(NpStack, len(m))
	for i, a := range m {
		ret[i] = a[start:end]
	}
	return ret
}

func (m NpStack) Column(i int) NpArray {
	ret := make(NpArray, len(m))
	for j, a := range m {
		ret[j] = a[i]
	}
	return ret
}

func (m NpStack) Cond(f func(float64) float64) NpStack {
	ret := make(NpStack, len(m))
	for i, a := range m {
		ret[i] = a.Cond(f)
	}
	return ret
}

func (m NpStack) LinearInterpolate(scale, new_scale NpArray) NpStack {
	ret := make(NpStack, len(m))
	for i, a := range m {
		ret[i] = a.LinearInterpolate(scale, new_scale)
	}
	return ret
}

func (m NpStack) DivSlice(b float64, start, end int) NpStack {
	ret := make(NpStack, len(m))
	for i, a := range m {
		ret[i] = a.DivSlice(b, start, end)
	}
	return ret
}

func (m NpStack) SubFloat64(b float64) NpStack {
	ret := make(NpStack, len(m))
	for i, a := range m {
		ret[i] = a.SubFloat64(b)
	}
	return ret
}

func (m NpStack) Add(b NpStack) NpStack {
	ret := make(NpStack, len(m))
	for i, a := range m {
		ret[i] = a.Add(b[i])
	}
	return ret
}

func (m NpStack) MulFloat64(b float64) NpStack {
	ret := make(NpStack, len(m))
	for i, a := range m {
		ret[i] = a.MulFloat64(b)
	}
	return ret
}

func (m NpStack) LinSpace(start, end float64, n int) NpStack {
	step := (end - start) / float64(n)
	ret := make(NpStack, len(m))
	for i, _ := range m {
		ret[i] = make(NpArray, n)
		for j := 0; j < n; j++ {
			ret[i][j] = start + float64(j)*step
		}
	}
	return ret
}

func (m NpStack) Max() NpArray {
	ret := make(NpArray, len(m))
	for i, a := range m {
		ret[i] = a.Max()
	}
	return ret
}

func (m NpStack) Min() NpArray {
	ret := make(NpArray, len(m))
	for i, a := range m {
		ret[i] = a.Min()
	}
	return ret
}

func (m NpStack) Sum() NpArray {
	ret := make(NpArray, len(m))
	for i, a := range m {
		ret[i] = a.Sum()
	}
	return ret
}

func (m NpStack) AllMostEqual(b NpStack, epsilon float64) bool {
	for i, a := range m {
		if !a.AlmostEqual(b[i], epsilon) {
			return false
		}
	}
	return true
}
