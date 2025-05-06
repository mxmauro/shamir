// See the LICENSE file for license details.

package shamir

import (
	"crypto/rand"
)

// -----------------------------------------------------------------------------

type polynomial []byte

type point struct {
	x, y byte
}

// -----------------------------------------------------------------------------

func newRandomPolynomial(secret byte, degree byte) polynomial {
	p := make(polynomial, degree+1)
	p[0] = secret
	if _, err := rand.Read(p[1:]); err != nil {
		return nil
	}
	return p
}

func evaluatePolynomial(p polynomial, x byte) byte {
	if x == 0 {
		// Return the origin.
		return p[0]
	}

	// Horner's method.
	degree := len(p) - 1
	out := p[degree]
	for i := degree - 1; i >= 0; i-- {
		out = add(mul(out, x), p[i])
	}

	return out
}

func interpolateAt0(points []point) byte {
	var result byte

	for i, ptI := range points {
		v := byte(1)
		for j, ptJ := range points {
			if i != j {
				num := ptJ.x
				denom := add(ptI.x, ptJ.x)
				frac := div(num, denom)
				v = mul(v, frac)
			}
		}
		result = add(result, mul(v, ptI.y))
	}

	return result
}
