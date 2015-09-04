package gola

import (
	"testing"
)

func TestNewVector(t *testing.T) {
	v := New(1, 2, 4)
	if len(v) != 3 {
		t.Errorf("Expecting vector size %d, but got %d", 3, len(v))
	}
}

func TestVectorString(t *testing.T) {
	v := New(1, 2, 3, 4)
	if v.String() != "[1,2,3,4]" {
		t.Logf("Expecting [1,2,3,4], but got %s", v.String())
		t.Fail()
	}
}

func TestVectorEqual(t *testing.T) {
	v1 := New(45, 44, 90)
	var v2 Vector = []float64{45, 44, 90}
	if !v1.Eq(v2) {
		t.Logf("Vectors are expected to be eqal")
		t.Fail()
	}
}

func TestVectorNotEqua(t *testing.T) {
	v1 := New(12, 56, 7)
	v2 := New(12, 56, 9)
	if v1.Eq(v2) {
		t.Logf("Vectors should not be equal")
		t.Fail()
	}
}
