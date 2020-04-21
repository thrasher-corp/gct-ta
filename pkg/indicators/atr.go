package indicators

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

// ATR returns the Average True Range for the given period
func ATR(inHigh, inLow, inClose []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inClose))

	if inTimePeriod < 1 || inTimePeriod > len(inHigh) || inTimePeriod > len(inLow) || inTimePeriod > len(inClose) {
		return outReal
	}

	if inTimePeriod == 1 {
		return trueRange(inHigh, inLow, inClose)
	}

	today := inTimePeriod + 1
	tr := trueRange(inHigh, inLow, inClose)
	prevATRTemp := SMA(tr, inTimePeriod)
	prevATR := prevATRTemp[inTimePeriod]
	outReal[inTimePeriod] = prevATR

	for outIdx := inTimePeriod + 1; outIdx < len(inClose); outIdx++ {
		prevATR *= float64(inTimePeriod) - 1.0
		prevATR += tr[today]
		prevATR /= float64(inTimePeriod)
		outReal[outIdx] = prevATR
		today++
	}

	return outReal
}
