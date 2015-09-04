package gola

import (
	"bytes"
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

func (v Vector) Scale(scale float64) {
	for i := range v {
		v[i] = v[i] * scale
	}
	return
}
