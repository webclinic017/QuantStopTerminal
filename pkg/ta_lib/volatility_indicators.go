package ta_lib

import "math"

/* Volatility Indicators */

// Atr - Average True Range
func Atr(inHigh []float64, inLow []float64, inClose []float64, inTimePeriod int) []float64 {

	outReal := make([]float64, len(inClose))

	inTimePeriodF := float64(inTimePeriod)

	if inTimePeriod < 1 {
		return outReal
	}

	if inTimePeriod <= 1 {
		return TRange(inHigh, inLow, inClose)
	}

	outIdx := inTimePeriod
	today := inTimePeriod + 1

	tr := TRange(inHigh, inLow, inClose)
	prevATRTemp := Sma(tr, inTimePeriod)
	prevATR := prevATRTemp[inTimePeriod]
	outReal[inTimePeriod] = prevATR

	for outIdx = inTimePeriod + 1; outIdx < len(inClose); outIdx++ {
		prevATR *= inTimePeriodF - 1.0
		prevATR += tr[today]
		prevATR /= inTimePeriodF
		outReal[outIdx] = prevATR
		today++
	}

	return outReal
}

// Natr - Normalized Average True Range
func Natr(inHigh []float64, inLow []float64, inClose []float64, inTimePeriod int) []float64 {

	outReal := make([]float64, len(inClose))

	if inTimePeriod < 1 {
		return outReal
	}

	if inTimePeriod <= 1 {
		return TRange(inHigh, inLow, inClose)
	}

	inTimePeriodF := float64(inTimePeriod)
	outIdx := inTimePeriod
	today := inTimePeriod

	tr := TRange(inHigh, inLow, inClose)
	prevATRTemp := Sma(tr, inTimePeriod)
	prevATR := prevATRTemp[inTimePeriod]

	tempValue := inClose[today]
	if tempValue != 0.0 {
		outReal[outIdx] = (prevATR / tempValue) * 100.0
	} else {
		outReal[outIdx] = 0.0
	}

	for outIdx = inTimePeriod + 1; outIdx < len(inClose); outIdx++ {
		today++
		prevATR *= inTimePeriodF - 1.0
		prevATR += tr[today]
		prevATR /= inTimePeriodF
		tempValue = inClose[today]
		if tempValue != 0.0 {
			outReal[outIdx] = (prevATR / tempValue) * 100.0
		} else {
			outReal[0] = 0.0
		}
	}

	return outReal
}

// TRange - True Range
func TRange(inHigh []float64, inLow []float64, inClose []float64) []float64 {

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
