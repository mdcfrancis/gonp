package np

import (
	"github.com/pa-m/randomkit"
	"math"
)

func LinSpace(start, end float64, n int) NpArray {
	step := (end - start) / float64(n-1)
	ret := make(NpArray, n)
	for i := 0; i < n; i++ {
		ret[i] = start + float64(i)*step
	}
	return ret
}

func Arrange(start, end, step float64) NpArray {
	n := int((end - start) / step)
	ret := make(NpArray, n)
	for i := 0; i < n; i++ {
		ret[i] = start + float64(i)*step
	}
	return ret
}

func Exp(x NpArray) NpArray {
	ret := make(NpArray, len(x))
	for i := 0; i < len(x); i++ {
		ret[i] = math.Exp(x[i])
	}
	return ret
}

func DivFloat64(num float64, x NpArray) NpArray {
	ret := make(NpArray, len(x))
	for i := 0; i < len(x); i++ {
		ret[i] = num / x[i]
	}
	return ret
}

func Zeros(n int) NpArray {
	ret := make(NpArray, n)
	for i := 0; i < n; i++ {
		ret[i] = 0
	}
	return ret
}
func RandN(rnd *randomkit.RKState, n int) NpArray {
	ret := make([]float64, n)
	for i, _ := range ret {
		ret[i] = rnd.NormFloat64()
	}
	return ret
}

func RandBItMask(n int) int {
	mask := n
	mask |= mask >> 1
	mask |= mask >> 2
	mask |= mask >> 4
	mask |= mask >> 8
	mask |= mask >> 16
	mask |= mask >> 32
	return mask
}

func RandChoice( /*RKState*/ rnd *randomkit.RKState, n int) int {
	mask := RandBItMask(n)
	for {
		candidate := rnd.Uint32() & uint32(mask)
		if candidate < uint32(n) {
			return int(candidate)
		}
	}
}
