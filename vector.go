package gola

import (
	"bytes"
	"strconv"
)

type Vector []float64

func New(elems ...float64) Vector {
	return elems
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
	if len(v) != len(other) {
		return false
	}
	for i, val := range v {
		if val != other[i] {
			return false
		}
	}
	return true
}
