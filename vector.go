package gola

import (
	"bytes"
	"math"
	"strconv"
)

const (
	zero = 1.0e-10 // zero tolerance
)

// the Vector type is a slice of float64
type Vector []float64

func New(elems ...float64) Vector {
	return elems
}

func (v Vector) assertLenMatch(other Vector) {
	if len(v) != len(other) {
		panic("Vector length mismatch")
	}
}

// Copy returns a new copy of provided vector
func (v Vector) Copy() (result Vector) {
	result = make([]float64, len(v))
	copy(result, v)
	return
}

// String returns a string representation of the Vector
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

// Eq compares vector equality
func (v Vector) Eq(other Vector) bool {
	v.assertLenMatch(other)
	for i, val := range v {
		if val != other[i] {
			return false
		}
	}
	return true
}

// Test for the zero vector
func (v Vector) IsZero() bool {
	return v.Mag() <= zero
}

// Add returns the sum of two vectors
func (v Vector) Add(other Vector) (result Vector) {
	v.assertLenMatch(other)
	result = make([]float64, len(v))
	for i, val := range v {
		result[i] = val + other[i]
	}
	return
}

// Sub returns the subtraction of a vector from another
func (v Vector) Sub(other Vector) (result Vector) {
	v.assertLenMatch(other)
	result = make([]float64, len(v))
	for i, val := range v {
		result[i] = val - other[i]
	}
	return
}

// ScalarMul scales the vector
func (v Vector) ScalarMul(scale float64) {
	for i := range v {
		v[i] = v[i] * scale
	}
}

// Mag computes the magnitude of the vector
func (v Vector) Mag() (result float64) {
	for _, v := range v {
		result += (v * v)
	}
	result = math.Sqrt(result)
	return
}

// Unit returns the normalization of the vector
func (v Vector) Unit() (result Vector) {
	result = v.Copy()
	mag := result.Mag()
	result.ScalarMul(1 / mag)
	return
}

// DotProd calculates the dot product of two vectors
func (v Vector) DotProd(other Vector) (result float64) {
	v.assertLenMatch(other)
	for i, val := range v {
		result += val * other[i]
	}
	return
}

// Angle calculates the angle between two vectors (Rad)
func (v Vector) Angle(other Vector) float64 {
	return math.Acos(v.DotProd(other) / (v.Mag() * other.Mag()))
}

// IsParallel test for parallelism between the vectors
func (v Vector) IsParallel(other Vector) bool {
	if v.IsZero() || other.IsZero() {
		return true
	}
	return v.Angle(other) == 0 || v.Angle(other) == math.Pi
}

// IsOrthogonal tests for orthongonality between the vectors
func (v Vector) IsOrthogonal(other Vector) bool {

	if v.IsZero() || other.IsZero() {
		return true
	}
	return math.Abs(v.DotProd(other)) < zero
}

// Proj returns the projection (or parllel component) v
// unto base vector using: proj(v) = (v * unit(base)) * unit(base)
func (v Vector) Proj(base Vector) (result Vector) {
	baseUnit := base.Unit()
	baseUnit.ScalarMul(v.DotProd(baseUnit))
	return baseUnit
}

// Perp returns the perpendicular component of a given projection
// unto a base vector: perp(v) = v - proj(v)
func (v Vector) Perp(base Vector) (result Vector) {
	return v.Sub(v.Proj(base))
}

// CrossProd calculates the cross product of two vectors
// if A = [a0, a1, a2] and B = [b0, b1, b2], then
// A x B = [
//   A[1]*B[2] - B[1]*A[2],
// -(A[0]*B[2] - B[0]*A[2],
//   A[0]*B[1] - B[0]*A[1]
// ]
func (v Vector) CrossProd(w Vector) (result Vector) {
	v.assertLenMatch(w)
	if len(v) != 3 {
		panic("Vector size must be 3")
	}
	return New(
		v[1]*w[2]-w[1]*v[2],
		-1*(v[0]*w[2]-w[0]*v[2]),
		v[0]*w[1]-w[0]*v[1],
	)
}
