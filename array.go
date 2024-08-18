package np

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"math"
)

type NpArray stats.Float64Data // numpy array

func (a NpArray) String() string {
	// format the array as a string
	// for example [5.   3.   3.   3.4  3.8  4.2  4.6  5.   5.4  5.8  5.   5.  ]
	buf := "["
	for _, v := range a {
		buf += fmt.Sprintf("%.8f ", v)
	}
	buf += "]"
	return buf
}

func (a NpArray) Shape() int {
	return len(a)
}

func (a NpArray) Copy() NpArray {
	ret := make(NpArray, len(a))
	copy(ret, a)
	return ret
}

func (a NpArray) AlmostEqual(b NpArray, epsilon float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if math.Abs(a[i]-b[i]) > epsilon {
			return false
		}
	}
	return true
}

func (a NpArray) Shuffle(idx []int) NpArray {
	ret := make(NpArray, len(a))
	for i, v := range idx {
		ret[i] = a[v]
	}
	return ret
}

func (a NpArray) Sum() float64 {
	sum := 0.0
	for _, v := range a {
		sum += v
	}
	return sum
}

func (a NpArray) Mean() float64 {
	m, _ := stats.Mean(stats.Float64Data(a))
	return m
}
func (a NpArray) StandardDeviation() float64 {
	s, _ := stats.StandardDeviation(stats.Float64Data(a))
	return s
}

func (a NpArray) Sub(b NpArray) NpArray {
	if len(a) != len(b) {
		panic(fmt.Errorf("Arrays must be same length %d %d", len(a), len(b)))
	}
	ret := make(NpArray, len(a))
	for i := 0; i < len(a); i++ {
		ret[i] = a[i] - b[i]
	}
	return ret
}

func (a NpArray) Div(b NpArray) NpArray {
	if len(a) != len(b) {
		panic(fmt.Errorf("Arrays must be same length %d %d", len(a), len(b)))
	}
	ret := make(NpArray, len(a))
	for i := 0; i < len(a); i++ {
		ret[i] = a[i] / b[i]
	}
	return ret
}

func (a NpArray) Mul(b NpArray) NpArray {
	if len(a) != len(b) {
		panic(fmt.Errorf("Arrays must be same length %d %d", len(a), len(b)))
	}
	ret := make(NpArray, len(a))
	for i := 0; i < len(a); i++ {
		ret[i] = a[i] * b[i]
	}
	return ret
}

func (a NpArray) PowFloat64(power float64) NpArray {
	ret := make(NpArray, len(a))
	for i := 0; i < len(a); i++ {
		ret[i] = math.Pow(a[i], power)
	}
	return ret
}

func (a NpArray) Add(b NpArray) NpArray {
	if len(a) != len(b) {
		panic(fmt.Errorf("Arrays must be same length %d %d", len(a), len(b)))
	}
	ret := make(NpArray, len(a))
	for i := 0; i < len(a); i++ {
		ret[i] = a[i] + b[i]
	}
	return ret
}

func (a NpArray) Dot(b NpArray) float64 {
	if len(a) != len(b) {
		panic(fmt.Errorf("doProdut : arrays must be same length %d %d", len(a), len(b)))
	}
	ret := 0.0
	for i := 0; i < len(a); i++ {
		ret += a[i] * b[i]
	}
	return ret
}

func (a NpArray) DivFloat64(b float64) NpArray {
	ret := make(NpArray, len(a))
	for i := 0; i < len(a); i++ {
		ret[i] = a[i] / b
	}
	return ret
}

func (a NpArray) MulFloat64(b float64) NpArray {
	ret := make(NpArray, len(a))
	for i := 0; i < len(a); i++ {
		ret[i] = a[i] * b
	}
	return ret
}

func (a NpArray) SubSlice(b NpArray, start, end int) NpArray {
	ret := make(NpArray, end-start)
	for i := start; i < end; i++ {
		ret[i-start] = a[i] - b[i]
	}
	return ret
}

func (a NpArray) DivSlice(b float64, start, end int) NpArray {
	ret := make(NpArray, end-start)
	for i := start; i < end; i++ {
		ret[i-start] = a[i] / b
	}
	return ret
}

func (a NpArray) SubFloat64(b float64) NpArray {
	ret := make(NpArray, len(a))
	for i := 0; i < len(a); i++ {
		ret[i] = a[i] - b
	}
	return ret
}

// interp1D performs a 1D linear interpolation,
// given x and y arrays, and a new x array xi.
func interp1D(x, y, xi NpArray) (NpArray, error) {
	if len(x) != len(y) {
		return nil, fmt.Errorf("x and y must have the same length")
	}

	yi := make(NpArray, len(xi))

	for i, v := range xi {
		if v <= x[0] {
			yi[i] = y[0]
		} else if v >= x[len(x)-1] {
			yi[i] = y[len(y)-1]
		} else {
			for j := 0; j < len(x)-1; j++ {
				if v >= x[j] && v <= x[j+1] {
					yi[i] = y[j] + (y[j+1]-y[j])*(v-x[j])/(x[j+1]-x[j])
					break
				}
			}
		}
	}

	return yi, nil
}

func (a NpArray) LinearInterpolate(scale, new_scale NpArray) NpArray {
	new_a, err := interp1D(scale, a, new_scale)
	if err != nil {
		panic(err)
	}
	return new_a
}

func (a NpArray) Cond(f func(float64) float64) NpArray {
	ret := make(NpArray, len(a))
	for i, v := range a {
		ret[i] = f(v)
	}
	return ret
}

func (a NpArray) Min() float64 {
	min, _ := stats.Min(stats.Float64Data(a))
	return min
}

func (a NpArray) Max() float64 {
	max, _ := stats.Max(stats.Float64Data(a))
	return max
}

func (a NpArray) AddFloat64(b float64) NpArray {
	ret := make(NpArray, len(a))
	for i := 0; i < len(a); i++ {
		ret[i] = a[i] + b
	}
	return ret
}
