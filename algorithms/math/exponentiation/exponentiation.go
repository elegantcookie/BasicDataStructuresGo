package exponentiation

// GetPowerOfTwo returns 2^n, where n >= 0
func GetPowerOfTwo(n int) int {
	start := 1 // 2^0 = 1
	for i := 0; i < n; i++ {
		start <<= 1
	}
	return start
}
