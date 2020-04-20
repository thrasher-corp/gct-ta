package indicators

// OBV returns the On Balance Volume for the given period
func OBV(closeIn, volIn []float64) []float64 {
	out := make([]float64, len(closeIn))

	out[0] = 0
	for x := 1; x < len(closeIn); x++ {
		// nolint gocritic ifElseChain: switch statement complexity not needed
		if closeIn[x] > closeIn[x-1] {
			out[x] = out[x-1] + volIn[x]
		} else if closeIn[x] < closeIn[x-1] {
			out[x] = out[x-1] - volIn[x]
		} else if closeIn[x] == closeIn[x-1] {
			out[x] = out[x-1]
		}
	}
	return out
}
