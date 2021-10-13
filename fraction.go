package rational

import (
	"fmt"
	"math"
)

type Fraction struct {
	Numerator, Denominator int64
}

// New creates a new *Fraction from a numerator and a denominator
func New(numerator, denominator int64) *Fraction {
	return &Fraction{
		Numerator:   numerator,
		Denominator: denominator,
	}
}

// New creates a new *Fraction from a floating point number.
func NewFromFloat(f float64) *Fraction {
	n := 1
	// math.Trunc returns the integer form of the provided float64
	for f != math.Trunc(f) {
		f *= 10
		n *= 10
	}
	frac := New(int64(f), int64(n))
	frac.Simplify()
	return frac
}

// Fraction returns a floating point version of f
func (f *Fraction) Float() float64 {
	return float64(f.Numerator) / float64(f.Denominator)
}

// String returns a string representation of f
func (f *Fraction) String() string {
	return fmt.Sprintf("%d/%d", f.Numerator, f.Denominator)
}

// Floor returns the equivalent of int64(math.Floor((*Fraction).Float()))
func (f *Fraction) Floor() int64 {
	return int64(f.Float())
}

func gcd(a int64, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Simplify simplifies f
func (f *Fraction) Simplify() {
	greatestCommonDivisor := gcd(f.Numerator, f.Denominator)
	f.Numerator /= greatestCommonDivisor
	f.Denominator /= greatestCommonDivisor
}

// Add adds fb to fa
func (fa *Fraction) Add(fb Fraction) {
	
	// If denominators match, we don't need to do anything special.
	if fa.Denominator != fb.Denominator {
		// If denominators do not match, we need to make ourselves a common denominator.
		faDenom := fa.Denominator
		fbDenom := fb.Denominator

		fa.Numerator *= fbDenom
		fa.Denominator *= fbDenom
		
		fb.Numerator *= faDenom
		fb.Denominator *= faDenom
	}

	// Add numerators

	fa.Numerator += fb.Numerator

	fa.Simplify()
}

// Sub subtracts fb from the fa
func (fa *Fraction) Sub(fb Fraction) {
	nfb := fb
	nfb.Numerator = -fb.Numerator
	fa.Add(nfb)
}

// Mult multiplies fa by fb
func (fa *Fraction) Mult(fb Fraction) {
	fa.Numerator *= fb.Numerator
	fa.Denominator *= fb.Denominator
	fa.Simplify()
}

// Div divides fa by fb
func (fa *Fraction) Div(fb Fraction) {
	nfb := Fraction{
		Numerator:   fb.Denominator,
		Denominator: fb.Numerator,
	}
	fa.Mult(nfb)
}

// LimitDenominator transforms f into a representation that's as close as possible to the original where `(*Fraction).Denominator < max`
func (f *Fraction) LimitDenominator(max int64) {
	if f.Denominator > max {
		x := float64(max) / float64(f.Denominator)
		f.Denominator = int64(float64(f.Denominator) * x)
		f.Numerator = int64(float64(f.Numerator) * x)
	}
}