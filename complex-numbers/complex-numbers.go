package complexnumbers

import (
	"math"
)

// Number represents a complex number with real and imaginary parts.
type Number struct {
	real float64
	imag float64
}

// Constructor function (optional, but useful)
func New(real, imag float64) Number {
	return Number{real: real, imag: imag}
}

// Real returns the real part of the complex number.
func (n Number) Real() float64 {
	return n.real
}

// Imaginary returns the imaginary part of the complex number.
func (n Number) Imaginary() float64 {
	return n.imag
}

// Add adds two complex numbers.
func (n1 Number) Add(n2 Number) Number {
	return Number{n1.real + n2.real, n1.imag + n2.imag}
}

// Subtract subtracts one complex number from another.
func (n1 Number) Subtract(n2 Number) Number {
	return Number{n1.real - n2.real, n1.imag - n2.imag}
}

// Multiply multiplies two complex numbers.
func (n1 Number) Multiply(n2 Number) Number {
	return Number{
		real: n1.real*n2.real - n1.imag*n2.imag,
		imag: n1.real*n2.imag + n1.imag*n2.real,
	}
}

// Times multiplies a complex number by a real factor.
func (n Number) Times(factor float64) Number {
	return Number{n.real * factor, n.imag * factor}
}

// Divide divides one complex number by another.
func (n1 Number) Divide(n2 Number) Number {
	denom := n2.real*n2.real + n2.imag*n2.imag
	return Number{
		real: (n1.real*n2.real + n1.imag*n2.imag) / denom,
		imag: (n1.imag*n2.real - n1.real*n2.imag) / denom,
	}
}

// Conjugate returns the complex conjugate of the number.
func (n Number) Conjugate() Number {
	return Number{n.real, -n.imag}
}

// Abs returns the magnitude (absolute value) of the complex number.
func (n Number) Abs() float64 {
	return math.Sqrt(n.real*n.real + n.imag*n.imag)
}

// Exp returns the complex exponential of the number: e^(a+bi) = e^a * (cos b + i sin b)
func (n Number) Exp() Number {
	eToA := math.Exp(n.real)
	return Number{
		real: eToA * math.Cos(n.imag),
		imag: eToA * math.Sin(n.imag),
	}
}
