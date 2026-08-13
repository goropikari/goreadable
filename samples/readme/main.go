package readmesample

func BuildReport(a, b, c, d, e, f int) int {
	if a > 0 {
		if b > 0 {
			return a + b + c + d + e + f
		}
	}

	return 0
}
