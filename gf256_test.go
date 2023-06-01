package shamir

import (
	"testing"
)

// -----------------------------------------------------------------------------

type addMulDivValues struct {
	a byte
	b byte
	r byte
}

type expValues struct {
	x        byte
	exponent int
	r        byte
}

// -----------------------------------------------------------------------------

func TestAddGF256(t *testing.T) {
	values := []addMulDivValues{
		{0x53, 0xCA, 0x99},
		{0xCA, 0x53, 0x99},
	}

	for _, v := range values {
		r := add(v.a, v.b)
		if r != v.r {
			t.Errorf("%v + %v got %v instead of %v", v.a, v.b, r, v.r)
		}
	}
}

func TestMulGF256(t *testing.T) {
	values := []addMulDivValues{
		{0x00, 0x00, 0x00},
		{0x01, 0x5A, 0x5A},
		{0x5A, 0x51, 0xFC},
		{0x03, 0x07, 0x09},
	}

	for _, v := range values {
		r := mul(v.a, v.b)
		if r != v.r {
			t.Errorf("%v * %v got %v instead of %v", v.a, v.b, r, v.r)
		}
	}
}

func TestDivGF256(t *testing.T) {
	values := []addMulDivValues{
		{0x00, 0x07, 0x00},
		{0x03, 0x03, 0x01},
		{0x06, 0x03, 0x02},
	}

	for _, v := range values {
		r := div(v.a, v.b)
		if r != v.r {
			t.Errorf("%v / %v got %v instead of %v", v.a, v.b, r, v.r)
		}
	}
}

func TestPowerGF256(t *testing.T) {
	values := []expValues{
		{0x53, 0x02, 0xB5},
		{0x53, 0x03, 0xC3},
	}

	for _, v := range values {
		r := byte(0x01)
		for i := 0; i < v.exponent; i++ {
			r = mul(r, v.x)
		}
		if r != v.r {
			t.Errorf("%v ^ %v got %v instead of %v", v.x, v.exponent, r, v.r)
		}
	}
}
