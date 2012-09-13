//Example from http://golang.org/doc/code.html#tmp_13
package newmath

// Approx square root function
func Sqrt(x float64) float64 {
	z := 0.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}
