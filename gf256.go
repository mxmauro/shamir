package shamir

import (
	"crypto/subtle"
)

// -----------------------------------------------------------------------------

func add(a, b byte) byte {
	return a ^ b
}

func mul(x, y byte) byte {
	var z byte

	if y == 0 {
		return 0
	}

	for x > 0 {
		z ^= ^((x & 1) - 1) & y
		x >>= 1
		a := y >> 7
		y <<= 1
		y ^= ^(a - 1) & 0x1B
	}
	return z
}

func inv(x byte) byte {
	if x == 0 {
		return 0
	}
	// Compute x^254
	z := mul(x, x)
	z = mul(z, x)
	z = mul(z, z)
	z = mul(z, x)
	z = mul(z, z)
	z = mul(z, x)
	z = mul(z, z)
	z = mul(z, x)
	z = mul(z, z)
	z = mul(z, x)
	z = mul(z, z)
	z = mul(z, x)
	return mul(z, z)
}

func div(x, y byte) byte {
	if y == 0 {
		panic("division by zero")
	}
	p := mul(x, inv(y))
	return byte(subtle.ConstantTimeSelect(subtle.ConstantTimeByteEq(x, 0), 0, int(p)))
}
