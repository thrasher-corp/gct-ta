/*
The MIT License (MIT)

Copyright (c) 2016 Mark Chenoweth
Copyright (c) 2020-2020 Andrew Jackson <andrew.jackson@thrasher.io>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package indicators

// BBANDS returns the  Bollinger Bands for the given period
func BBANDS(inReal []float64, inTimePeriod int, inNbDevUp, inNbDevDn float64, inMAType MaType) (upper, middle, lower []float64) {
	outRealUpperBand := make([]float64, len(inReal))
	outRealMiddleBand := MA(inReal, inTimePeriod, inMAType)
	outRealLowerBand := make([]float64, len(inReal))

	tempBuffer2 := stdDev(inReal, inTimePeriod, 1.0)

	switch inNbDevUp {
	case inNbDevDn:
		if inNbDevUp == 1.0 {
			for i := 0; i < len(inReal); i++ {
				tempReal := tempBuffer2[i]
				tempReal2 := outRealMiddleBand[i]
				outRealUpperBand[i] = tempReal2 + tempReal
				outRealLowerBand[i] = tempReal2 - tempReal
			}
		} else {
			for i := 0; i < len(inReal); i++ {
				tempReal := tempBuffer2[i] * inNbDevUp
				tempReal2 := outRealMiddleBand[i]
				outRealUpperBand[i] = tempReal2 + tempReal
				outRealLowerBand[i] = tempReal2 - tempReal
			}
		}
	case 1.0:
		for i := 0; i < len(inReal); i++ {
			tempReal := tempBuffer2[i]
			tempReal2 := outRealMiddleBand[i]
			outRealUpperBand[i] = tempReal2 + tempReal
			outRealLowerBand[i] = tempReal2 - (tempReal * inNbDevDn)
		}
	default:
		if inNbDevDn == 1.0 {
			for i := 0; i < len(inReal); i++ {
				tempReal := tempBuffer2[i]
				tempReal2 := outRealMiddleBand[i]
				outRealLowerBand[i] = tempReal2 - tempReal
				outRealUpperBand[i] = tempReal2 + (tempReal * inNbDevUp)
			}
		} else {
			for i := 0; i < len(inReal); i++ {
				tempReal := tempBuffer2[i]
				tempReal2 := outRealMiddleBand[i]
				outRealUpperBand[i] = tempReal2 + (tempReal * inNbDevUp)
				outRealLowerBand[i] = tempReal2 - (tempReal * inNbDevDn)
			}
		}
	}

	return outRealUpperBand, outRealMiddleBand, outRealLowerBand
}
