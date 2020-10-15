package indicators

// Crossover returns true if inA crosses over inB
func Crossover(inA, inB []float64) bool {
	if len(inA) < 3 || len(inB) < 3 {
		return false
	}

	N := len(inA)

	return inA[N-2] <= inB[N-2] && inA[N-1] > inB[N-1]
}

// Crossunder returns true if inA is crossing under inB.
func Crossunder(inA, inB []float64) bool {
	if len(inA) < 3 || len(inB) < 3 {
		return false
	}

	N := len(inA)

	return inA[N-1] <= inB[N-1] && inA[N-2] > inB[N-2]
}
