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

type moneyFlow struct {
	positive float64
	negative float64
}

// MFI returns Money Flow Index for the given period
// nolint gocyclo alternatives are to use reflection this is code increase v performance cost
func MFI(inHigh, inLow, inClose, inVolume []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inClose))
	mflowIdx := 0
	mflow := make([]moneyFlow, inTimePeriod)
	maxIdxMflow := inTimePeriod - 1
	lookbackTotal := inTimePeriod
	startIdx := lookbackTotal
	outIdx := startIdx
	today := startIdx - lookbackTotal
	prevValue := (inHigh[today] + inLow[today] + inClose[today]) / 3.0
	posSumMF := 0.0
	negSumMF := 0.0
	today++
	for i := inTimePeriod; i > 0; i-- {
		tempValue1 := (inHigh[today] + inLow[today] + inClose[today]) / 3.0
		tempValue2 := tempValue1 - prevValue
		prevValue = tempValue1
		tempValue1 *= inVolume[today]
		today++
		//nolint ifElseChain switch case makes execution flow harder to follow with negative numbers
		if tempValue2 < 0 {
			(mflow[mflowIdx]).negative = tempValue1
			negSumMF += tempValue1
			(mflow[mflowIdx]).positive = 0.0
		} else if tempValue2 > 0 {
			(mflow[mflowIdx]).positive = tempValue1
			posSumMF += tempValue1
			(mflow[mflowIdx]).negative = 0.0
		} else {
			(mflow[mflowIdx]).positive = 0.0
			(mflow[mflowIdx]).negative = 0.0
		}
		mflowIdx++
		if mflowIdx > maxIdxMflow {
			mflowIdx = 0
		}
	}
	if today > startIdx {
		tempValue1 := posSumMF + negSumMF
		if tempValue1 < 1.0 {
		} else {
			outReal[outIdx] = 100.0 * (posSumMF / tempValue1)
			outIdx++
		}
	} else {
		for today < startIdx {
			posSumMF -= mflow[mflowIdx].positive
			negSumMF -= mflow[mflowIdx].negative
			tempValue1 := (inHigh[today] + inLow[today] + inClose[today]) / 3.0
			tempValue2 := tempValue1 - prevValue
			prevValue = tempValue1
			tempValue1 *= inVolume[today]
			today++
			//nolint ifElseChain switch case makes execution flow harder to follow with negative numbers
			if tempValue2 < 0 {
				(mflow[mflowIdx]).negative = tempValue1
				negSumMF += tempValue1
				(mflow[mflowIdx]).positive = 0.0
			} else if tempValue2 > 0 {
				(mflow[mflowIdx]).positive = tempValue1
				posSumMF += tempValue1
				(mflow[mflowIdx]).negative = 0.0
			} else {
				(mflow[mflowIdx]).positive = 0.0
				(mflow[mflowIdx]).negative = 0.0
			}
			mflowIdx++
			if mflowIdx > maxIdxMflow {
				mflowIdx = 0
			}
		}
	}
	for today < len(inClose) {
		posSumMF -= (mflow[mflowIdx]).positive
		negSumMF -= (mflow[mflowIdx]).negative
		tempValue1 := (inHigh[today] + inLow[today] + inClose[today]) / 3.0
		tempValue2 := tempValue1 - prevValue
		prevValue = tempValue1
		tempValue1 *= inVolume[today]
		today++
		//nolint ifElseChain switch case makes execution flow harder to follow with negative numbers
		if tempValue2 < 0 {
			(mflow[mflowIdx]).negative = tempValue1
			negSumMF += tempValue1
			(mflow[mflowIdx]).positive = 0.0
		} else if tempValue2 > 0 {
			(mflow[mflowIdx]).positive = tempValue1
			posSumMF += tempValue1
			(mflow[mflowIdx]).negative = 0.0
		} else {
			(mflow[mflowIdx]).positive = 0.0
			(mflow[mflowIdx]).negative = 0.0
		}
		tempValue1 = posSumMF + negSumMF
		if tempValue1 < 1.0 {
			outReal[outIdx] = 0.0
		} else {
			outReal[outIdx] = 100.0 * (posSumMF / tempValue1)
		}
		outIdx++
		mflowIdx++
		if mflowIdx > maxIdxMflow {
			mflowIdx = 0
		}
	}
	return outReal
}
