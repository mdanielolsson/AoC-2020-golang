package util

//PowInt returns the base-x exponential of y.
func PowInt(x, y int) int {
	initialX := x
	for i := 2; i <= y; i++ {
		x *= initialX
	}
	return x
}
