package indicators

import "math"

// mean returns mean value for an array of float64 values
func mean(values []float64) float64 {
	var total float64 = 0
	for x := range values {
		total += values[x]
	}
	return total / float64(len(values))
}

func evenSlice(inA, inB []float64) (outA, outB []float64) {
	offsetA := int(math.Max(0, float64(len(inA)-len(inB))))
	return inA[offsetA:], inB[0:]
}

// trueRange returns the true range for high low close
func trueRange(inHigh, inLow, inClose []float64) []float64 {
	outReal := make([]float64, len(inClose))

	startIdx := 1
	outIdx := startIdx
	today := startIdx
	for today < len(inClose) {
		tempLT := inLow[today]
		tempHT := inHigh[today]
		tempCY := inClose[today-1]
		greatest := tempHT - tempLT
		val2 := math.Abs(tempCY - tempHT)
		if val2 > greatest {
			greatest = val2
		}
		val3 := math.Abs(tempCY - tempLT)
		if val3 > greatest {
			greatest = val3
		}
		outReal[outIdx] = greatest
		outIdx++
		today++
	}

	return outReal
}

// variance returns Variance for given time period
func variance(inReal []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal))

	nbInitialElementNeeded := inTimePeriod - 1
	startIdx := nbInitialElementNeeded
	periodTotal1 := 0.0
	periodTotal2 := 0.0
	trailingIdx := startIdx - nbInitialElementNeeded
	i := trailingIdx
	if inTimePeriod > 1 {
		for i < startIdx {
			tempReal := inReal[i]
			periodTotal1 += tempReal
			tempReal *= tempReal
			periodTotal2 += tempReal
			i++
		}
	}
	outIdx := startIdx
	for ok := true; ok; {
		tempReal := inReal[i]
		periodTotal1 += tempReal
		tempReal *= tempReal
		periodTotal2 += tempReal
		meanValue1 := periodTotal1 / float64(inTimePeriod)
		meanValue2 := periodTotal2 / float64(inTimePeriod)
		tempReal = inReal[trailingIdx]
		periodTotal1 -= tempReal
		tempReal *= tempReal
		periodTotal2 -= tempReal
		outReal[outIdx] = meanValue2 - meanValue1*meanValue1
		i++
		trailingIdx++
		outIdx++
		ok = i < len(inReal)
	}
	return outReal
}

// stdDev - Standard Deviation
func stdDev(inReal []float64, inTimePeriod int, inNbDev float64) []float64 {
	outReal := variance(inReal, inTimePeriod)

	if inNbDev != 1.0 {
		for i := 0; i < len(inReal); i++ {
			tempReal := outReal[i]
			if !(tempReal < 0.00000000000001) {
				outReal[i] = math.Sqrt(tempReal) * inNbDev
			} else {
				outReal[i] = 0.0
			}
		}
	} else {
		for i := 0; i < len(inReal); i++ {
			tempReal := outReal[i]
			if !(tempReal < 0.00000000000001) {
				outReal[i] = math.Sqrt(tempReal)
			} else {
				outReal[i] = 0.0
			}
		}
	}
	return outReal
}
