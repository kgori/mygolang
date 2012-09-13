//Basic test suite for newmath package from
//http://golang.org/doc/code.html#tmp_13
package newmath

import "testing"

func TestSqrt4(t *testing.T) {
	const in, out = 4, 2
	if x := Sqrt(in); x != out {
		t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
	}
}

func TestSqrt10000(t *testing.T) {
	const in, out = 10000, 100
	if x := Sqrt(in); x != out {
		t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
	}
}
