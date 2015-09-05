package gola

import (
	"bytes"
	"math"
	"strconv"
)

type Vector []float64

func New(elems ...float64) Vector {
	return elems
}

func (v Vector) assertLenMatch(other Vector) {
	if len(v) != len(other) {
		panic("Vector length mismatch")
	}
}

func (v Vector) Copy() (result Vector) {
	result = make([]float64, len(v))
	copy(result, v)
	return
}

func (v Vector) String() string {
	buff := bytes.NewBufferString("[")
	for i, val := range v {
		buff.WriteString(strconv.FormatFloat(val, 'g', -1, 64))
		if i < len(v)-1 {
			buff.WriteRune(',')
		}
	}
	buff.WriteRune(']')
	return buff.String()
}

func (v Vector) Eq(other Vector) bool {
	v.assertLenMatch(other)
	for i, val := range v {
		if val != other[i] {
			return false
		}
	}
	return true
}

func (v Vector) Add(other Vector) (result Vector) {
	v.assertLenMatch(other)
	result = make([]float64, len(v))
	for i, val := range v {
		result[i] = val + other[i]
	}
	return
}

func (v Vector) Sub(other Vector) (result Vector) {
	v.assertLenMatch(other)
	result = make([]float64, len(v))
	for i, val := range v {
		result[i] = val - other[i]
	}
	return
}

func (v Vector) ScalarMul(scale float64) {
	for i := range v {
		v[i] = v[i] * scale
	}
}

func (v Vector) Mag() (result float64) {
	for _, v := range v {
		result += (v * v)
	}
	result = math.Sqrt(result)
	return
}

func (v Vector) Unit() (result Vector) {
	result = v.Copy()
	mag := result.Mag()
	result.ScalarMul(1 / mag)
	return
}

func (v Vector) DotProd(other Vector) (result float64) {
	v.assertLenMatch(other)
	for i, val := range v {
		result += val * other[i]
	}
	return
}

func (v Vector) Angle(other Vector) float64 {
	return math.Acos(v.DotProd(other) / (v.Mag() * other.Mag()))
}
