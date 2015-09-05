package gola

import (
	"math"
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

func TestVectorAdd(t *testing.T) {
	v1 := New(8.218, -9.341)
	v2 := New(-1.129, 2.111)
	v3 := v1.Add(v2)
	expect := New(
		v1[0]+v2[0],
		v1[1]+v2[1],
	)

	if !v3.Eq(expect) {
		t.Log("Addition failed, expecting %s, got %s", expect, v3)
		t.Fail()
	}
	t.Log(v1, "+", v2, v3)
}

func TestVectorSub(t *testing.T) {
	v1 := New(7.119, 8.215)
	v2 := New(-8.223, 0.878)
	v3 := v1.Sub(v2)
	expect := New(
		v1[0]-v2[0],
		v1[1]-v2[1],
	)
	if !v3.Eq(expect) {
		t.Log("Subtraction failed, expecting %s, got %v", expect, v3)
		t.Fail()
	}
	t.Log(v1, "-", v2, "=", v3)
}

func TestVectorScalarMul(t *testing.T) {
	v := New(1.671, -1.012, -0.318)
	v.ScalarMul(7.41)
	expect := New(
		7.41*1.671,
		7.41*-1.012,
		7.41*-0.318,
	)
	if !v.Eq(expect) {
		t.Logf("Scalar mul failed, expecting %s, got %s", expect, v)
		t.Fail()
	}
	t.Log("1.671,-1.012, -0.318 Scale", 7.41, "=", v)
}

func TestVectorMag(t *testing.T) {
	v := New(-0.221, 7.437)
	expected := math.Sqrt(v[0]*v[0] + v[1]*v[1])
	if v.Mag() != expected {
		t.Logf("Magnitude failed, execpted %d, got %d", expected, v.Mag())
		t.Fail()
	}
	t.Log(v, "Mag() =", v.Mag())
	v = New(8.813, -1.331, -6.247)
	expected = math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
	if v.Mag() != expected {
		t.Logf("Magnitude failed, expected %d, got %d, expected", v.Mag())
		t.Fail()
	}
	t.Log(v, "Mag() = ", v.Mag())
}

func TestVectorUnit(t *testing.T) {
	v := New(5.581, -2.136)
	mag := v.Mag()
	expect := New((1/mag)*v[0], (1/mag)*v[1])
	if !v.Unit().Eq(expect) {
		t.Logf("Vector Unit failed, expecting %s, got %s", expect, v.Unit())
		t.Fail()
	}
	t.Log(v, "Unit() = ", v.Unit())
}
