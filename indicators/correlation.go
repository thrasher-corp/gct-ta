package indicators

import (
	"math"
)

func sum(input []float64) float64 {
	var sum float64
	for x := range input {
		sum += input[x]
	}
	return sum
}

func correlation(c1, c2 []float64) float64 {
	if len(c1) != len(c2) || len(c1) == 0 || len(c2) == 0 {
		return 0
	}

	sumx, sumy := sum(c1), sum(c2)
	var sumxy, sumpx, sumpy float64
	for i := range c1 {
		sumxy += c1[i] * c2[i]
		sumpx += math.Pow(c1[i], 2)
		sumpy += math.Pow(c2[i], 2)
	}
	l := float64(len(c1))
	return (l*sumxy - (sumx * sumy)) /
		(math.Sqrt((l*sumpx - math.Pow(sumx, 2)) * (l*sumpy - math.Pow(sumy, 2))))
}

func CorrelationCoefficient(c1, c2 []float64, period int) []float64 {
	if len(c1) != len(c2) || len(c1) == 0 || len(c2) == 0 {
		return nil
	}

	if period > len(c1) {
		return nil
	}

	r := make([]float64, len(c1))
	r[period-1] = correlation(c1[:period], c2[:period])
	for x := period; x < len(c1); x++ {
		idx := x - period + 1
		r[x] = correlation(c1[idx:idx+period], c2[idx:idx+period])
	}
	return r
}
