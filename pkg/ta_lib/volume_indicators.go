package ta_lib

/* Volume Indicators */

// Ad - Chaikin A/D Line
func Ad(inHigh []float64, inLow []float64, inClose []float64, inVolume []float64) []float64 {

	outReal := make([]float64, len(inClose))

	startIdx := 0
	nbBar := len(inClose) - startIdx
	currentBar := startIdx
	outIdx := 0
	ad := 0.0
	for nbBar != 0 {
		high := inHigh[currentBar]
		low := inLow[currentBar]
		tmp := high - low
		close := inClose[currentBar]
		if tmp > 0.0 {
			ad += (((close - low) - (high - close)) / tmp) * (inVolume[currentBar])
		}
		outReal[outIdx] = ad
		outIdx++
		currentBar++
		nbBar--
	}
	return outReal
}

// AdOsc - Chaikin A/D Oscillator
func AdOsc(inHigh []float64, inLow []float64, inClose []float64, inVolume []float64, inFastPeriod int, inSlowPeriod int) []float64 {

	outReal := make([]float64, len(inClose))

	if (inFastPeriod < 2) || (inSlowPeriod < 2) {
		return outReal
	}

	slowestPeriod := 0
	if inFastPeriod < inSlowPeriod {
		slowestPeriod = inSlowPeriod
	} else {
		slowestPeriod = inFastPeriod
	}
	lookbackTotal := slowestPeriod - 1
	startIdx := lookbackTotal
	today := startIdx - lookbackTotal
	ad := 0.0
	fastk := (2.0 / (float64(inFastPeriod) + 1.0))
	oneMinusfastk := 1.0 - fastk
	slowk := (2.0 / (float64(inSlowPeriod) + 1.0))
	oneMinusslowk := 1.0 - slowk
	high := inHigh[today]
	low := inLow[today]
	tmp := high - low
	close := inClose[today]
	if tmp > 0.0 {
		ad += (((close - low) - (high - close)) / tmp) * (inVolume[today])
	}
	today++
	fastEMA := ad
	slowEMA := ad

	for today < startIdx {
		high = inHigh[today]
		low = inLow[today]
		tmp = high - low
		close = inClose[today]
		if tmp > 0.0 {
			ad += (((close - low) - (high - close)) / tmp) * (inVolume[today])
		}
		today++

		fastEMA = (fastk * ad) + (oneMinusfastk * fastEMA)
		slowEMA = (slowk * ad) + (oneMinusslowk * slowEMA)
	}
	outIdx := lookbackTotal
	for today < len(inClose) {
		high = inHigh[today]
		low = inLow[today]
		tmp = high - low
		close = inClose[today]
		if tmp > 0.0 {
			ad += (((close - low) - (high - close)) / tmp) * (inVolume[today])
		}
		today++
		fastEMA = (fastk * ad) + (oneMinusfastk * fastEMA)
		slowEMA = (slowk * ad) + (oneMinusslowk * slowEMA)
		outReal[outIdx] = fastEMA - slowEMA
		outIdx++
	}

	return outReal
}

// Obv - On Balance Volume
func Obv(inReal []float64, inVolume []float64) []float64 {

	outReal := make([]float64, len(inReal))
	startIdx := 0
	prevOBV := inVolume[startIdx]
	prevReal := inReal[startIdx]
	outIdx := 0
	for i := startIdx; i < len(inReal); i++ {
		tempReal := inReal[i]
		if tempReal > prevReal {
			prevOBV += inVolume[i]
		} else if tempReal < prevReal {
			prevOBV -= inVolume[i]
		}
		outReal[outIdx] = prevOBV
		prevReal = tempReal
		outIdx++
	}
	return outReal
}
