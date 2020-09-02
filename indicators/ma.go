package indicators

import (
	"log"
	"math"
)

// MaType Moving Average indicator types
type MaType uint

// Moving Average indicator types
const (
	Sma MaType = iota
	Ema
	Dpo
)

func dpo (in []float64, period int) []float64 {
	var out []float64
		out = make([]float64, len(in))
	if len(in) < period {
		return out
	}
	sma := sma(in, period, false)
	for i := range in {
		if i < period {
			continue
		}
		var whatty float64
		whatty = math.Floor(float64(i) - (float64(period)) / 2 + 1)
		log.Printf("i: %v. In: %v. sma: %v inValue: %v, smaValue: %v", i,int64(whatty) ,  i - period, in[int64(whatty)],  sma[i - period])
		price := in[int64(whatty)] - sma[i-period]
		out[i] = price
	}
	return out
}

// DPO returns the Detrended Price Oscillator for a given period
func DPO(in []float64, period int) []float64 {
	return dpo(in, period)
}

// SMA returns the Simple Moving Average for the given period
func sma(in []float64, period int, macd bool) []float64 {
	var out []float64
	if !macd {
		out = make([]float64, len(in))
	}
	if len(in) < period {
		return out
	}
	for i := range in {
		if i+1 >= period {
			avg := mean(in[i+1-period : i+1])
			if macd {
				out = append(out, avg)
				continue
			}
			out[i] = avg
		}
	}
	return out
}

// SMA returns the Simple Moving Average for the given period
func SMA(in []float64, period int) []float64 {
	return sma(in, period, false)
}

// EMA returns the Exponential Moving Average for the given period
func ema(in []float64, period int, macd bool) []float64 {
	var out []float64
	if !macd {
		out = make([]float64, len(in))
	}
	if len(in) < period {
		return out
	}

	smaRet := sma(in, period, macd)
	if macd {
		out = append(out, smaRet[0])
	} else {
		out[period-1] = smaRet[period-1]
	}
	var multiplier = 2.0 / (float64(period) + 1.0)
	for i := period; i < len(in); i++ {
		var lastVal float64
		if macd {
			lastVal = out[len(out)-1]
		} else {
			lastVal = out[i-1]
		}
		ema := (in[i]-lastVal)*multiplier + lastVal
		if macd {
			out = append(out, ema)
			continue
		}
		out[i] = ema
	}
	return out
}

// EMA returns the Exponential Moving Average for the given period
func EMA(in []float64, period int) []float64 {
	return ema(in, period, false)
}

func calcMACD(inA, inB []float64) []float64 {
	inA, inB = evenSlice(inA, inB)
	out := make([]float64, len(inA))
	for i := range inA {
		if inA[i] == 0 || inB[i] == 0 {
			continue
		}
		out[i] = inA[i] - inB[i]
	}
	return out
}

// MACD returns the Moving Average Convergence Divergence indicator
// for the given fastPeriod, slowPeriod and signalPeriod
func MACD(values []float64, fastPeriod, slowPeriod, signalPeriod int) (macdValues, signalPeriodValues, histogramValues []float64) {
	if fastPeriod > len(values) || slowPeriod > len(values) {
		return
	}

	if fastPeriod > slowPeriod {
		return
	}

	if signalPeriod > len(values) {
		return
	}

	fastPeriodValues := ema(values, fastPeriod, true)
	slowPeriodValues := ema(values, slowPeriod, true)
	macdValues = calcMACD(fastPeriodValues, slowPeriodValues)
	signalPeriodValues = ema(macdValues, signalPeriod, true)
	macdValues, signalPeriodValues = evenSlice(macdValues, signalPeriodValues)
	histogramValues = calcMACD(macdValues, signalPeriodValues)

	// find a better solution this is a work around for now to factor in MACD values not matching

	ret := make([]float64, len(values))
	ret2 := make([]float64, len(values))
	ret3 := make([]float64, len(values))
	copy(ret[slowPeriod+(signalPeriod-2):], macdValues)
	copy(ret2[slowPeriod+(signalPeriod-2):], signalPeriodValues)
	copy(ret3[slowPeriod+(signalPeriod-2):], histogramValues)

	return ret, ret2, ret3
}

// MA Moving Average helper
func MA(inReal []float64, inTimePeriod int, inMAType MaType) []float64 {
	if inTimePeriod == 1 {
		return inReal
	}

	outReal := make([]float64, len(inReal))
	switch inMAType {
	case Sma:
		outReal = SMA(inReal, inTimePeriod)
	case Ema:
		outReal = EMA(inReal, inTimePeriod)
	case Dpo:
		outReal = DPO(inReal, inTimePeriod)
	}
	return outReal
}
