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

func TestVectorDotProd(t *testing.T) {
	v1 := New(7.887, 4.138)
	v2 := New(-8.802, 6.776)
	actual := v1.DotProd(v2)
	expect := v1[0]*v2[0] + v1[1]*v2[1]
	if actual != expect {
		t.Logf("DotPoduct failed, expecting %d, got %d", expect, actual)
		t.Fail()
	}
	t.Log(v1, "DotProd", v2, "=", actual)

	v1 = New(-5.955, -4.904, -1.874)
	v2 = New(-4.496, -8.775, 7.103)
	actual = v1.DotProd(v2)
	expect = v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
	if actual != expect {
		t.Logf("DotProdcut failed, expecting %d, got %d", expect, actual)
		t.Fail()
	}

	t.Log(v1, "DotProd", v2, "=", actual)
}

func TestVectorAngle(t *testing.T) {
	v1 := New(3.183, -7.627)
	v2 := New(-2.668, 5.319)
	actual := v1.Angle(v2)
	expect := math.Acos(v1.DotProd(v2) / (v1.Mag() * v2.Mag()))
	if actual != expect {
		t.Logf("Vector angle failed, expecting %d, got %d", expect, actual)
		t.Fail()
	}
	t.Log("Angle between", v1, "and", v2, "=", actual)

	v1 = New(7.35, 0.221, 5.188)
	v2 = New(2.751, 8.259, 3.985)
	actual = v1.Angle(v2)
	expect = math.Acos(v1.DotProd(v2) / (v1.Mag() * v2.Mag()))
	if actual != expect {
		t.Logf("Vector angle failed, exepcting %d, got %d", expect, actual)
		t.Fail()
	}

	t.Log("Angle between", v1, "and", v2, "=", actual)
}

func TestVectorParallel(t *testing.T) {
	v1 := New(-7.579, -7.88)
	v2 := New(22.737, 23.64)
	actual := v1.IsParallel(v2)
	expect := (v1.Angle(v2) == 0 || v1.Angle(v2) == math.Pi)
	if actual != expect {
		t.Logf("Vector parallel failed, expecting %t, got %t", expect, actual)
		t.Fail()
	}
	t.Log(v1, "Parallel to", v2, "is", actual)

	v1 = New(-2.029, 9.97, 4.172)
	v2 = New(-9.231, -6.639, -7.245)
	actual = v1.IsParallel(v2)
	expect = (v1.Angle(v2) == 0 || v1.Angle(v2) == math.Pi)
	if actual != expect {
		t.Logf("Vector parallel failed, expecting %t, got %t", expect, actual)
		t.Fail()
	}
	t.Log(v1, "Parallel to", v2, "is", actual)

	v1 = New(-2.238, -7.284, -1.214)
	v2 = New(-1.821, 1.072, -2.94)
	actual = v1.IsParallel(v2)
	expect = (v1.Angle(v2) == 0 || v1.Angle(v2) == math.Pi)
	if actual != expect {
		t.Logf("Vector parallel failed, expecting %t, got %t", expect, actual)
		t.Fail()
	}
	t.Log(v1, "Parallel to", v2, "is", actual)

	v1 = New(2.118, 4.827)
	v2 = New(0, 0)
	actual = v1.IsParallel(v2)
	expect = (v1.IsZero() || v2.IsZero() || v1.Angle(v2) == 0 || v1.Angle(v2) == math.Pi)
	if actual != expect {
		t.Logf("Vector parallel failed, expecting %t, got %t", expect, actual)
		t.Fail()
	}
	t.Log(v1, "Parallel to", v2, "is", actual)
}

func TestVectorOrthogonality(t *testing.T) {
	v1 := New(-7.579, -7.88)
	v2 := New(22.737, 23.64)
	actual := v1.IsOrthogonal(v2)
	expect := (math.Abs(v1.DotProd(v2)) < zero)
	if actual != expect {
		t.Logf("Vector orthogonal failed, expecting %t, got %t", expect, actual)
		t.Fail()
	}
	t.Log(v1, "Orthogonal to", v2, "is", actual)

	v1 = New(-2.029, 9.97, 4.172)
	v2 = New(-9.231, -6.639, -7.245)
	actual = v1.IsOrthogonal(v2)
	expect = (math.Abs(v1.DotProd(v2)) < zero)
	if actual != expect {
		t.Logf("Vector orthogonal failed, expecting %t, got %t", expect, actual)
		t.Fail()
	}
	t.Log(v1, "Orthogonal to", v2, "is", actual)

	v1 = New(-2.238, -7.284, -1.214)
	v2 = New(-1.821, 1.072, -2.94)
	actual = v1.IsOrthogonal(v2)
	expect = (math.Abs(v1.DotProd(v2)) < zero)
	if actual != expect {
		t.Logf("Vector orthogonal, expecting %t, got %t", expect, actual)
		t.Fail()
	}
	t.Log(v1, "Orthogonal to", v2, "is", actual)

	v1 = New(2.118, 4.827)
	v2 = New(0, 0)
	actual = v1.IsOrthogonal(v2)
	expect = (math.Abs(v1.DotProd(v2)) < zero)
	if actual != expect {
		t.Logf("Vecto orthogonal failed, expecting %t, got %t", expect, actual)
		t.Fail()
	}
	t.Log(v1, "Orthogonal to", v2, "is", actual)
}

func TestVectorProjection(t *testing.T) {
	v := New(3.039, 1.879)
	b := New(0.825, 2.036)
	actual := v.Proj(b)
	expect := b.Unit()
	expect.ScalarMul(v.DotProd(expect))
	if !actual.Eq(expect) {
		t.Logf("Vector projection failed, expeecting %s, got %s", expect, actual)
		t.Fail()
	}
	t.Log(v, "Proj on", b, "=", actual)
}

func TestVectorPerp(t *testing.T) {
	v := New(-9.88, -3.264, -8.159)
	b := New(-2.155, -9.353, -9.473)
	actual := v.Perp(b)
	expect := v.Sub(v.Proj(b))
	if !actual.Eq(expect) {
		t.Logf("Vector perp calculation failed, expecting %s, got %v", expect, actual)
		t.Fail()
	}
	t.Log(v, "Perpendicular to", b, "=", actual)
}

func TestVectorProjectionComponents(t *testing.T) {
	v := New(3.009, -6.172, 3.692, -2.51)
	b := New(6.404, -9.144, 2.759, 8.718)
	projV := v.Proj(b)
	t.Log(v, "Projection on b =", projV)
	perpV := v.Perp(b)
	t.Log(v, "Perpendicular component of b =", perpV)
}

func TestVectorCrossProduct(t *testing.T) {
	v := New(8.462, 7.893, -8.187)
	w := New(6.984, -5.975, 4.778)
	actual := v.CrossProd(w)
	expect := New(
		v[1]*w[2]-w[1]*v[2],
		-1*(v[0]*w[2]-w[0]*v[2]),
		v[0]*w[1]-w[0]*v[1],
	)
	if !actual.Eq(expect) {
		t.Logf("Vector cross product failed, expecting %s, got %s", expect, actual)
		t.Fail()
	}
	t.Log(v, "CrossProd", w, "=", actual)
}

func TestVectorValidateCrossProd(t *testing.T) {
	v := New(-8.987, -9.838, 5.031)
	w := New(-4.268, -1.861, -8.866)
	area := v.CrossProd(w).Mag()
	expect := New(
		v[1]*w[2]-w[1]*v[2],
		-1*(v[0]*w[2]-w[0]*v[2]),
		v[0]*w[1]-w[0]*v[1],
	).Mag()
	if area != expect {
		t.Logf("CrossProd validation failed, expecting %d, got %d", expect, area)
		t.Fail()
	}
	t.Log(v, "magnitude of cross prod", w, "=", area)

	v = New(1.5, 9.547, 3.691)
	w = New(-6.007, 0.124, 5.772)
	triArea := v.CrossProd(w).Mag() * 0.5
	t.Log(v, "X", w, "triangle Area = ", triArea)
}
