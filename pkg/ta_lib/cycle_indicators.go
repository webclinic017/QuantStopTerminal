package ta_lib

import "math"

/* Cycle Indicators */

// HtDcPeriod - Hilbert Transform - Dominant Cycle Period (lookback=32)
func HtDcPeriod(inReal []float64) []float64 {

	outReal := make([]float64, len(inReal))

	a := 0.0962
	b := 0.5769
	detrenderOdd := make([]float64, 3)
	detrenderEven := make([]float64, 3)
	q1Odd := make([]float64, 3)
	q1Even := make([]float64, 3)
	jIOdd := make([]float64, 3)
	jIEven := make([]float64, 3)
	jQOdd := make([]float64, 3)
	jQEven := make([]float64, 3)
	rad2Deg := 180.0 / (4.0 * math.Atan(1))
	lookbackTotal := 32
	startIdx := lookbackTotal
	trailingWMAIdx := startIdx - lookbackTotal
	today := trailingWMAIdx
	tempReal := inReal[today]
	today++
	periodWMASub := tempReal
	periodWMASum := tempReal
	tempReal = inReal[today]
	today++
	periodWMASub += tempReal
	periodWMASum += tempReal * 2.0
	tempReal = inReal[today]
	today++
	periodWMASub += tempReal
	periodWMASum += tempReal * 3.0
	trailingWMAValue := 0.0
	i := 9
	smoothedValue := 0.0
	for ok := true; ok; {
		tempReal = inReal[today]
		today++
		periodWMASub += tempReal
		periodWMASub -= trailingWMAValue
		periodWMASum += tempReal * 4.0
		trailingWMAValue = inReal[trailingWMAIdx]
		trailingWMAIdx++
		smoothedValue = periodWMASum * 0.1
		periodWMASum -= periodWMASub
		i--
		ok = i != 0
	}

	hilbertIdx := 0
	detrender := 0.0
	prevDetrenderOdd := 0.0
	prevDetrenderEven := 0.0
	prevDetrenderInputOdd := 0.0
	prevDetrenderInputEven := 0.0
	q1 := 0.0
	prevq1Odd := 0.0
	prevq1Even := 0.0
	prevq1InputOdd := 0.0
	prevq1InputEven := 0.0
	jI := 0.0
	prevJIOdd := 0.0
	prevJIEven := 0.0
	prevJIInputOdd := 0.0
	prevJIInputEven := 0.0
	jQ := 0.0
	prevJQOdd := 0.0
	prevJQEven := 0.0
	prevJQInputOdd := 0.0
	prevJQInputEven := 0.0
	period := 0.0
	outIdx := 32
	previ2 := 0.0
	prevq2 := 0.0
	Re := 0.0
	Im := 0.0
	i2 := 0.0
	q2 := 0.0
	i1ForOddPrev3 := 0.0
	i1ForEvenPrev3 := 0.0
	i1ForOddPrev2 := 0.0
	i1ForEvenPrev2 := 0.0
	smoothPeriod := 0.0
	for today < len(inReal) {
		adjustedPrevPeriod := (0.075 * period) + 0.54
		todayValue := inReal[today]
		periodWMASub += todayValue
		periodWMASub -= trailingWMAValue
		periodWMASum += todayValue * 4.0
		trailingWMAValue = inReal[trailingWMAIdx]
		trailingWMAIdx++
		smoothedValue = periodWMASum * 0.1
		periodWMASum -= periodWMASub
		hilbertTempReal := 0.0
		if (today % 2) == 0 {
			hilbertTempReal = a * smoothedValue
			detrender = -detrenderEven[hilbertIdx]
			detrenderEven[hilbertIdx] = hilbertTempReal
			detrender += hilbertTempReal
			detrender -= prevDetrenderEven
			prevDetrenderEven = b * prevDetrenderInputEven
			detrender += prevDetrenderEven
			prevDetrenderInputEven = smoothedValue
			detrender *= adjustedPrevPeriod
			hilbertTempReal = a * detrender
			q1 = -q1Even[hilbertIdx]
			q1Even[hilbertIdx] = hilbertTempReal
			q1 += hilbertTempReal
			q1 -= prevq1Even
			prevq1Even = b * prevq1InputEven
			q1 += prevq1Even
			prevq1InputEven = detrender
			q1 *= adjustedPrevPeriod
			hilbertTempReal = a * i1ForEvenPrev3
			jI = -jIEven[hilbertIdx]
			jIEven[hilbertIdx] = hilbertTempReal
			jI += hilbertTempReal
			jI -= prevJIEven
			prevJIEven = b * prevJIInputEven
			jI += prevJIEven
			prevJIInputEven = i1ForEvenPrev3
			jI *= adjustedPrevPeriod
			hilbertTempReal = a * q1
			jQ = -jQEven[hilbertIdx]
			jQEven[hilbertIdx] = hilbertTempReal
			jQ += hilbertTempReal
			jQ -= prevJQEven
			prevJQEven = b * prevJQInputEven
			jQ += prevJQEven
			prevJQInputEven = q1
			jQ *= adjustedPrevPeriod
			hilbertIdx++
			if hilbertIdx == 3 {
				hilbertIdx = 0
			}
			q2 = (0.2 * (q1 + jI)) + (0.8 * prevq2)
			i2 = (0.2 * (i1ForEvenPrev3 - jQ)) + (0.8 * previ2)
			i1ForOddPrev3 = i1ForOddPrev2
			i1ForOddPrev2 = detrender
		} else {
			hilbertTempReal = a * smoothedValue
			detrender = -detrenderOdd[hilbertIdx]
			detrenderOdd[hilbertIdx] = hilbertTempReal
			detrender += hilbertTempReal
			detrender -= prevDetrenderOdd
			prevDetrenderOdd = b * prevDetrenderInputOdd
			detrender += prevDetrenderOdd
			prevDetrenderInputOdd = smoothedValue
			detrender *= adjustedPrevPeriod
			hilbertTempReal = a * detrender
			q1 = -q1Odd[hilbertIdx]
			q1Odd[hilbertIdx] = hilbertTempReal
			q1 += hilbertTempReal
			q1 -= prevq1Odd
			prevq1Odd = b * prevq1InputOdd
			q1 += prevq1Odd
			prevq1InputOdd = detrender
			q1 *= adjustedPrevPeriod
			hilbertTempReal = a * i1ForOddPrev3
			jI = -jIOdd[hilbertIdx]
			jIOdd[hilbertIdx] = hilbertTempReal
			jI += hilbertTempReal
			jI -= prevJIOdd
			prevJIOdd = b * prevJIInputOdd
			jI += prevJIOdd
			prevJIInputOdd = i1ForOddPrev3
			jI *= adjustedPrevPeriod
			hilbertTempReal = a * q1
			jQ = -jQOdd[hilbertIdx]
			jQOdd[hilbertIdx] = hilbertTempReal
			jQ += hilbertTempReal
			jQ -= prevJQOdd
			prevJQOdd = b * prevJQInputOdd
			jQ += prevJQOdd
			prevJQInputOdd = q1
			jQ *= adjustedPrevPeriod
			q2 = (0.2 * (q1 + jI)) + (0.8 * prevq2)
			i2 = (0.2 * (i1ForOddPrev3 - jQ)) + (0.8 * previ2)
			i1ForEvenPrev3 = i1ForEvenPrev2
			i1ForEvenPrev2 = detrender
		}
		Re = (0.2 * ((i2 * previ2) + (q2 * prevq2))) + (0.8 * Re)
		Im = (0.2 * ((i2 * prevq2) - (q2 * previ2))) + (0.8 * Im)
		prevq2 = q2
		previ2 = i2
		tempReal = period
		if (Im != 0.0) && (Re != 0.0) {
			period = 360.0 / (math.Atan(Im/Re) * rad2Deg)
		}
		tempReal2 := 1.5 * tempReal
		if period > tempReal2 {
			period = tempReal2
		}
		tempReal2 = 0.67 * tempReal
		if period < tempReal2 {
			period = tempReal2
		}
		if period < 6 {
			period = 6
		} else if period > 50 {
			period = 50
		}
		period = (0.2 * period) + (0.8 * tempReal)
		smoothPeriod = (0.33 * period) + (0.67 * smoothPeriod)
		if today >= startIdx {
			outReal[outIdx] = smoothPeriod
			outIdx++
		}
		today++
	}
	return outReal
}

// HtDcPhase - Hilbert Transform - Dominant Cycle Phase (lookback=63)
func HtDcPhase(inReal []float64) []float64 {

	outReal := make([]float64, len(inReal))
	a := 0.0962
	b := 0.5769
	detrenderOdd := make([]float64, 3)
	detrenderEven := make([]float64, 3)
	q1Odd := make([]float64, 3)
	q1Even := make([]float64, 3)
	jIOdd := make([]float64, 3)
	jIEven := make([]float64, 3)
	jQOdd := make([]float64, 3)
	jQEven := make([]float64, 3)
	smoothPriceIdx := 0
	maxIdxSmoothPrice := (50 - 1)
	smoothPrice := make([]float64, maxIdxSmoothPrice+1)
	tempReal := math.Atan(1)
	rad2Deg := 45.0 / tempReal
	constDeg2RadBy360 := tempReal * 8.0
	lookbackTotal := 63
	startIdx := lookbackTotal
	trailingWMAIdx := startIdx - lookbackTotal
	today := trailingWMAIdx
	tempReal = inReal[today]
	today++
	periodWMASub := tempReal
	periodWMASum := tempReal
	tempReal = inReal[today]
	today++
	periodWMASub += tempReal
	periodWMASum += tempReal * 2.0
	tempReal = inReal[today]
	today++
	periodWMASub += tempReal
	periodWMASum += tempReal * 3.0
	trailingWMAValue := 0.0
	i := 34
	smoothedValue := 0.0
	for ok := true; ok; {
		tempReal = inReal[today]
		today++
		periodWMASub += tempReal
		periodWMASub -= trailingWMAValue
		periodWMASum += tempReal * 4.0
		trailingWMAValue = inReal[trailingWMAIdx]
		trailingWMAIdx++
		smoothedValue = periodWMASum * 0.1
		periodWMASum -= periodWMASub
		i--
		ok = i != 0
	}

	hilbertIdx := 0
	detrender := 0.0
	prevDetrenderOdd := 0.0
	prevDetrenderEven := 0.0
	prevDetrenderInputOdd := 0.0
	prevDetrenderInputEven := 0.0
	q1 := 0.0
	prevq1Odd := 0.0
	prevq1Even := 0.0
	prevq1InputOdd := 0.0
	prevq1InputEven := 0.0
	jI := 0.0
	prevJIOdd := 0.0
	prevJIEven := 0.0
	prevJIInputOdd := 0.0
	prevJIInputEven := 0.0
	jQ := 0.0
	prevJQOdd := 0.0
	prevJQEven := 0.0
	prevJQInputOdd := 0.0
	prevJQInputEven := 0.0
	period := 0.0
	outIdx := 0
	previ2 := 0.0
	prevq2 := 0.0
	Re := 0.0
	Im := 0.0
	i1ForOddPrev3 := 0.0
	i1ForEvenPrev3 := 0.0
	i1ForOddPrev2 := 0.0
	i1ForEvenPrev2 := 0.0
	smoothPeriod := 0.0
	dcPhase := 0.0
	q2 := 0.0
	i2 := 0.0
	for today < len(inReal) {
		adjustedPrevPeriod := (0.075 * period) + 0.54
		todayValue := inReal[today]
		periodWMASub += todayValue
		periodWMASub -= trailingWMAValue
		periodWMASum += todayValue * 4.0
		trailingWMAValue = inReal[trailingWMAIdx]
		trailingWMAIdx++
		smoothedValue = periodWMASum * 0.1
		periodWMASum -= periodWMASub
		hilbertTempReal := 0.0
		smoothPrice[smoothPriceIdx] = smoothedValue
		if (today % 2) == 0 {
			hilbertTempReal = a * smoothedValue
			detrender = -detrenderEven[hilbertIdx]
			detrenderEven[hilbertIdx] = hilbertTempReal
			detrender += hilbertTempReal
			detrender -= prevDetrenderEven
			prevDetrenderEven = b * prevDetrenderInputEven
			detrender += prevDetrenderEven
			prevDetrenderInputEven = smoothedValue
			detrender *= adjustedPrevPeriod
			hilbertTempReal = a * detrender
			q1 = -q1Even[hilbertIdx]
			q1Even[hilbertIdx] = hilbertTempReal
			q1 += hilbertTempReal
			q1 -= prevq1Even
			prevq1Even = b * prevq1InputEven
			q1 += prevq1Even
			prevq1InputEven = detrender
			q1 *= adjustedPrevPeriod
			hilbertTempReal = a * i1ForEvenPrev3
			jI = -jIEven[hilbertIdx]
			jIEven[hilbertIdx] = hilbertTempReal
			jI += hilbertTempReal
			jI -= prevJIEven
			prevJIEven = b * prevJIInputEven
			jI += prevJIEven
			prevJIInputEven = i1ForEvenPrev3
			jI *= adjustedPrevPeriod
			hilbertTempReal = a * q1
			jQ = -jQEven[hilbertIdx]
			jQEven[hilbertIdx] = hilbertTempReal
			jQ += hilbertTempReal
			jQ -= prevJQEven
			prevJQEven = b * prevJQInputEven
			jQ += prevJQEven
			prevJQInputEven = q1
			jQ *= adjustedPrevPeriod
			hilbertIdx++
			if hilbertIdx == 3 {
				hilbertIdx = 0
			}
			q2 = (0.2 * (q1 + jI)) + (0.8 * prevq2)
			i2 = (0.2 * (i1ForEvenPrev3 - jQ)) + (0.8 * previ2)
			i1ForOddPrev3 = i1ForOddPrev2
			i1ForOddPrev2 = detrender
		} else {

			hilbertTempReal = a * smoothedValue
			detrender = -detrenderOdd[hilbertIdx]
			detrenderOdd[hilbertIdx] = hilbertTempReal
			detrender += hilbertTempReal
			detrender -= prevDetrenderOdd
			prevDetrenderOdd = b * prevDetrenderInputOdd
			detrender += prevDetrenderOdd
			prevDetrenderInputOdd = smoothedValue
			detrender *= adjustedPrevPeriod
			hilbertTempReal = a * detrender
			q1 = -q1Odd[hilbertIdx]
			q1Odd[hilbertIdx] = hilbertTempReal
			q1 += hilbertTempReal
			q1 -= prevq1Odd
			prevq1Odd = b * prevq1InputOdd
			q1 += prevq1Odd
			prevq1InputOdd = detrender
			q1 *= adjustedPrevPeriod
			hilbertTempReal = a * i1ForOddPrev3
			jI = -jIOdd[hilbertIdx]
			jIOdd[hilbertIdx] = hilbertTempReal
			jI += hilbertTempReal
			jI -= prevJIOdd
			prevJIOdd = b * prevJIInputOdd
			jI += prevJIOdd
			prevJIInputOdd = i1ForOddPrev3
			jI *= adjustedPrevPeriod
			hilbertTempReal = a * q1
			jQ = -jQOdd[hilbertIdx]
			jQOdd[hilbertIdx] = hilbertTempReal
			jQ += hilbertTempReal
			jQ -= prevJQOdd
			prevJQOdd = b * prevJQInputOdd
			jQ += prevJQOdd
			prevJQInputOdd = q1
			jQ *= adjustedPrevPeriod
			q2 = (0.2 * (q1 + jI)) + (0.8 * prevq2)
			i2 = (0.2 * (i1ForOddPrev3 - jQ)) + (0.8 * previ2)
			i1ForEvenPrev3 = i1ForEvenPrev2
			i1ForEvenPrev2 = detrender
		}
		Re = (0.2 * ((i2 * previ2) + (q2 * prevq2))) + (0.8 * Re)
		Im = (0.2 * ((i2 * prevq2) - (q2 * previ2))) + (0.8 * Im)
		prevq2 = q2
		previ2 = i2
		tempReal = period
		if (Im != 0.0) && (Re != 0.0) {
			period = 360.0 / (math.Atan(Im/Re) * rad2Deg)
		}
		tempReal2 := 1.5 * tempReal
		if period > tempReal2 {
			period = tempReal2
		}
		tempReal2 = 0.67 * tempReal
		if period < tempReal2 {
			period = tempReal2
		}
		if period < 6 {
			period = 6
		} else if period > 50 {
			period = 50
		}
		period = (0.2 * period) + (0.8 * tempReal)
		smoothPeriod = (0.33 * period) + (0.67 * smoothPeriod)
		DCPeriod := smoothPeriod + 0.5
		DCPeriodInt := math.Floor(DCPeriod)
		realPart := 0.0
		imagPart := 0.0
		idx := smoothPriceIdx
		for i := 0; i < int(DCPeriodInt); i++ {
			tempReal = (float64(i) * constDeg2RadBy360) / (DCPeriodInt * 1.0)
			tempReal2 = smoothPrice[idx]
			realPart += math.Sin(tempReal) * tempReal2
			imagPart += math.Cos(tempReal) * tempReal2
			if idx == 0 {
				idx = 50 - 1
			} else {
				idx--
			}
		}
		tempReal = math.Abs(imagPart)
		if tempReal > 0.0 {
			dcPhase = math.Atan(realPart/imagPart) * rad2Deg
		} else if tempReal <= 0.01 {
			if realPart < 0.0 {
				dcPhase -= 90.0
			} else if realPart > 0.0 {
				dcPhase += 90.0
			}
		}
		dcPhase += 90.0
		dcPhase += 360.0 / smoothPeriod
		if imagPart < 0.0 {
			dcPhase += 180.0
		}
		if dcPhase > 315.0 {
			dcPhase -= 360.0
		}
		if today >= startIdx {
			outReal[outIdx] = dcPhase
			outIdx++
		}
		smoothPriceIdx++
		if smoothPriceIdx > maxIdxSmoothPrice {
			smoothPriceIdx = 0
		}

		today++
	}
	return outReal
}

// HtPhasor - Hibert Transform - Phasor Components (lookback=32)
func HtPhasor(inReal []float64) ([]float64, []float64) {

	outInPhase := make([]float64, len(inReal))
	outQuadrature := make([]float64, len(inReal))

	a := 0.0962
	b := 0.5769
	detrenderOdd := make([]float64, 3)
	detrenderEven := make([]float64, 3)
	q1Odd := make([]float64, 3)
	q1Even := make([]float64, 3)
	jIOdd := make([]float64, 3)
	jIEven := make([]float64, 3)
	jQOdd := make([]float64, 3)
	jQEven := make([]float64, 3)
	rad2Deg := 180.0 / (4.0 * math.Atan(1))
	lookbackTotal := 32
	startIdx := lookbackTotal
	trailingWMAIdx := startIdx - lookbackTotal
	today := trailingWMAIdx
	tempReal := inReal[today]
	today++
	periodWMASub := tempReal
	periodWMASum := tempReal
	tempReal = inReal[today]
	today++
	periodWMASub += tempReal
	periodWMASum += tempReal * 2.0
	tempReal = inReal[today]
	today++
	periodWMASub += tempReal
	periodWMASum += tempReal * 3.0
	trailingWMAValue := 0.0
	i := 9
	smoothedValue := 0.0
	for ok := true; ok; {
		tempReal = inReal[today]
		today++
		periodWMASub += tempReal
		periodWMASub -= trailingWMAValue
		periodWMASum += tempReal * 4.0
		trailingWMAValue = inReal[trailingWMAIdx]
		trailingWMAIdx++
		smoothedValue = periodWMASum * 0.1
		periodWMASum -= periodWMASub
		i--
		ok = i != 0
	}
	hilbertIdx := 0
	detrender := 0.0
	prevDetrenderOdd := 0.0
	prevDetrenderEven := 0.0
	prevDetrenderInputOdd := 0.0
	prevDetrenderInputEven := 0.0
	q1 := 0.0
	prevq1Odd := 0.0
	prevq1Even := 0.0
	prevq1InputOdd := 0.0
	prevq1InputEven := 0.0
	jI := 0.0
	prevJIOdd := 0.0
	prevJIEven := 0.0
	prevJIInputOdd := 0.0
	prevJIInputEven := 0.0
	jQ := 0.0
	prevJQOdd := 0.0
	prevJQEven := 0.0
	prevJQInputOdd := 0.0
	prevJQInputEven := 0.0
	period := 0.0
	outIdx := 32
	previ2 := 0.0
	prevq2 := 0.0
	Re := 0.0
	Im := 0.0
	i1ForOddPrev3 := 0.0
	i1ForEvenPrev3 := 0.0
	i1ForOddPrev2 := 0.0
	i1ForEvenPrev2 := 0.0
	i2 := 0.0
	q2 := 0.0
	for today < len(inReal) {
		adjustedPrevPeriod := (0.075 * period) + 0.54
		todayValue := inReal[today]
		periodWMASub += todayValue
		periodWMASub -= trailingWMAValue
		periodWMASum += todayValue * 4.0
		trailingWMAValue = inReal[trailingWMAIdx]
		trailingWMAIdx++
		smoothedValue = periodWMASum * 0.1
		periodWMASum -= periodWMASub
		hilbertTempReal := 0.0
		if (today % 2) == 0 {
			hilbertTempReal = a * smoothedValue
			detrender = -detrenderEven[hilbertIdx]
			detrenderEven[hilbertIdx] = hilbertTempReal
			detrender += hilbertTempReal
			detrender -= prevDetrenderEven
			prevDetrenderEven = b * prevDetrenderInputEven
			detrender += prevDetrenderEven
			prevDetrenderInputEven = smoothedValue
			detrender *= adjustedPrevPeriod
			hilbertTempReal = a * detrender
			q1 = -q1Even[hilbertIdx]
			q1Even[hilbertIdx] = hilbertTempReal
			q1 += hilbertTempReal
			q1 -= prevq1Even
			prevq1Even = b * prevq1InputEven
			q1 += prevq1Even
			prevq1InputEven = detrender
			q1 *= adjustedPrevPeriod

			if today >= startIdx {
				outQuadrature[outIdx] = q1
				outInPhase[outIdx] = i1ForEvenPrev3
				outIdx++
			}
			hilbertTempReal = a * i1ForEvenPrev3
			jI = -jIEven[hilbertIdx]
			jIEven[hilbertIdx] = hilbertTempReal
			jI += hilbertTempReal
			jI -= prevJIEven
			prevJIEven = b * prevJIInputEven
			jI += prevJIEven
			prevJIInputEven = i1ForEvenPrev3
			jI *= adjustedPrevPeriod
			hilbertTempReal = a * q1
			jQ = -jQEven[hilbertIdx]
			jQEven[hilbertIdx] = hilbertTempReal
			jQ += hilbertTempReal
			jQ -= prevJQEven
			prevJQEven = b * prevJQInputEven
			jQ += prevJQEven
			prevJQInputEven = q1
			jQ *= adjustedPrevPeriod
			hilbertIdx++
			if hilbertIdx == 3 {
				hilbertIdx = 0
			}
			q2 = (0.2 * (q1 + jI)) + (0.8 * prevq2)
			i2 = (0.2 * (i1ForEvenPrev3 - jQ)) + (0.8 * previ2)
			i1ForOddPrev3 = i1ForOddPrev2
			i1ForOddPrev2 = detrender
		} else {

			hilbertTempReal = a * smoothedValue
			detrender = -detrenderOdd[hilbertIdx]
			detrenderOdd[hilbertIdx] = hilbertTempReal
			detrender += hilbertTempReal
			detrender -= prevDetrenderOdd
			prevDetrenderOdd = b * prevDetrenderInputOdd
			detrender += prevDetrenderOdd
			prevDetrenderInputOdd = smoothedValue
			detrender *= adjustedPrevPeriod
			hilbertTempReal = a * detrender
			q1 = -q1Odd[hilbertIdx]
			q1Odd[hilbertIdx] = hilbertTempReal
			q1 += hilbertTempReal
			q1 -= prevq1Odd
			prevq1Odd = b * prevq1InputOdd
			q1 += prevq1Odd
			prevq1InputOdd = detrender
			q1 *= adjustedPrevPeriod
			if today >= startIdx {
				outQuadrature[outIdx] = q1
				outInPhase[outIdx] = i1ForOddPrev3
				outIdx++
			}
			hilbertTempReal = a * i1ForOddPrev3
			jI = -jIOdd[hilbertIdx]
			jIOdd[hilbertIdx] = hilbertTempReal
			jI += hilbertTempReal
			jI -= prevJIOdd
			prevJIOdd = b * prevJIInputOdd
			jI += prevJIOdd
			prevJIInputOdd = i1ForOddPrev3
			jI *= adjustedPrevPeriod
			hilbertTempReal = a * q1
			jQ = -jQOdd[hilbertIdx]
			jQOdd[hilbertIdx] = hilbertTempReal
			jQ += hilbertTempReal
			jQ -= prevJQOdd
			prevJQOdd = b * prevJQInputOdd
			jQ += prevJQOdd
			prevJQInputOdd = q1
			jQ *= adjustedPrevPeriod
			q2 = (0.2 * (q1 + jI)) + (0.8 * prevq2)
			i2 = (0.2 * (i1ForOddPrev3 - jQ)) + (0.8 * previ2)
			i1ForEvenPrev3 = i1ForEvenPrev2
			i1ForEvenPrev2 = detrender
		}
		Re = (0.2 * ((i2 * previ2) + (q2 * prevq2))) + (0.8 * Re)
		Im = (0.2 * ((i2 * prevq2) - (q2 * previ2))) + (0.8 * Im)
		prevq2 = q2
		previ2 = i2
		tempReal = period
		if (Im != 0.0) && (Re != 0.0) {
			period = 360.0 / (math.Atan(Im/Re) * rad2Deg)
		}
		tempReal2 := 1.5 * tempReal
		if period > tempReal2 {
			period = tempReal2
		}
		tempReal2 = 0.67 * tempReal
		if period < tempReal2 {
			period = tempReal2
		}
		if period < 6 {
			period = 6
		} else if period > 50 {
			period = 50
		}
		period = (0.2 * period) + (0.8 * tempReal)
		today++
	}
	return outInPhase, outQuadrature
}

// HtSine - Hilbert Transform - SineWave (lookback=63)
func HtSine(inReal []float64) ([]float64, []float64) {

	outSine := make([]float64, len(inReal))
	outLeadSine := make([]float64, len(inReal))

	a := 0.0962
	b := 0.5769
	detrenderOdd := make([]float64, 3)
	detrenderEven := make([]float64, 3)
	q1Odd := make([]float64, 3)
	q1Even := make([]float64, 3)
	jIOdd := make([]float64, 3)
	jIEven := make([]float64, 3)
	jQOdd := make([]float64, 3)
	jQEven := make([]float64, 3)
	smoothPriceIdx := 0
	maxIdxSmoothPrice := (50 - 1)
	smoothPrice := make([]float64, maxIdxSmoothPrice+1)
	tempReal := math.Atan(1)
	rad2Deg := 45.0 / tempReal
	deg2Rad := 1.0 / rad2Deg
	constDeg2RadBy360 := tempReal * 8.0
	lookbackTotal := 63
	startIdx := lookbackTotal
	trailingWMAIdx := startIdx - lookbackTotal
	today := trailingWMAIdx
	tempReal = inReal[today]
	today++
	periodWMASub := tempReal
	periodWMASum := tempReal
	tempReal = inReal[today]
	today++
	periodWMASub += tempReal
	periodWMASum += tempReal * 2.0
	tempReal = inReal[today]
	today++
	periodWMASub += tempReal
	periodWMASum += tempReal * 3.0
	trailingWMAValue := 0.0
	i := 34
	smoothedValue := 0.0
	for ok := true; ok; {
		tempReal = inReal[today]
		today++
		periodWMASub += tempReal
		periodWMASub -= trailingWMAValue
		periodWMASum += tempReal * 4.0
		trailingWMAValue = inReal[trailingWMAIdx]
		trailingWMAIdx++
		smoothedValue = periodWMASum * 0.1
		periodWMASum -= periodWMASub
		i--
		ok = i != 0
	}

	hilbertIdx := 0
	detrender := 0.0
	prevDetrenderOdd := 0.0
	prevDetrenderEven := 0.0
	prevDetrenderInputOdd := 0.0
	prevDetrenderInputEven := 0.0
	q1 := 0.0
	prevq1Odd := 0.0
	prevq1Even := 0.0
	prevq1InputOdd := 0.0
	prevq1InputEven := 0.0
	jI := 0.0
	prevJIOdd := 0.0
	prevJIEven := 0.0
	prevJIInputOdd := 0.0
	prevJIInputEven := 0.0
	jQ := 0.0
	prevJQOdd := 0.0
	prevJQEven := 0.0
	prevJQInputOdd := 0.0
	prevJQInputEven := 0.0
	period := 0.0
	outIdx := 63
	previ2 := 0.0
	prevq2 := 0.0
	Re := 0.0
	Im := 0.0
	i1ForOddPrev3 := 0.0
	i1ForEvenPrev3 := 0.0
	i1ForOddPrev2 := 0.0
	i1ForEvenPrev2 := 0.0
	smoothPeriod := 0.0
	dcPhase := 0.0
	hilbertTempReal := 0.0
	q2 := 0.0
	i2 := 0.0
	for today < len(inReal) {
		adjustedPrevPeriod := (0.075 * period) + 0.54
		todayValue := inReal[today]
		periodWMASub += todayValue
		periodWMASub -= trailingWMAValue
		periodWMASum += todayValue * 4.0
		trailingWMAValue = inReal[trailingWMAIdx]
		trailingWMAIdx++
		smoothedValue = periodWMASum * 0.1
		periodWMASum -= periodWMASub
		smoothPrice[smoothPriceIdx] = smoothedValue
		if (today % 2) == 0 {
			hilbertTempReal = a * smoothedValue
			detrender = -detrenderEven[hilbertIdx]
			detrenderEven[hilbertIdx] = hilbertTempReal
			detrender += hilbertTempReal
			detrender -= prevDetrenderEven
			prevDetrenderEven = b * prevDetrenderInputEven
			detrender += prevDetrenderEven
			prevDetrenderInputEven = smoothedValue
			detrender *= adjustedPrevPeriod
			hilbertTempReal = a * detrender
			q1 = -q1Even[hilbertIdx]
			q1Even[hilbertIdx] = hilbertTempReal
			q1 += hilbertTempReal
			q1 -= prevq1Even
			prevq1Even = b * prevq1InputEven
			q1 += prevq1Even
			prevq1InputEven = detrender
			q1 *= adjustedPrevPeriod
			hilbertTempReal = a * i1ForEvenPrev3
			jI = -jIEven[hilbertIdx]
			jIEven[hilbertIdx] = hilbertTempReal
			jI += hilbertTempReal
			jI -= prevJIEven
			prevJIEven = b * prevJIInputEven
			jI += prevJIEven
			prevJIInputEven = i1ForEvenPrev3
			jI *= adjustedPrevPeriod
			hilbertTempReal = a * q1
			jQ = -jQEven[hilbertIdx]
			jQEven[hilbertIdx] = hilbertTempReal
			jQ += hilbertTempReal
			jQ -= prevJQEven
			prevJQEven = b * prevJQInputEven
			jQ += prevJQEven
			prevJQInputEven = q1
			jQ *= adjustedPrevPeriod
			hilbertIdx++
			if hilbertIdx == 3 {
				hilbertIdx = 0
			}
			q2 = (0.2 * (q1 + jI)) + (0.8 * prevq2)
			i2 = (0.2 * (i1ForEvenPrev3 - jQ)) + (0.8 * previ2)
			i1ForOddPrev3 = i1ForOddPrev2
			i1ForOddPrev2 = detrender
		} else {
			hilbertTempReal = a * smoothedValue
			detrender = -detrenderOdd[hilbertIdx]
			detrenderOdd[hilbertIdx] = hilbertTempReal
			detrender += hilbertTempReal
			detrender -= prevDetrenderOdd
			prevDetrenderOdd = b * prevDetrenderInputOdd
			detrender += prevDetrenderOdd
			prevDetrenderInputOdd = smoothedValue
			detrender *= adjustedPrevPeriod
			hilbertTempReal = a * detrender
			q1 = -q1Odd[hilbertIdx]
			q1Odd[hilbertIdx] = hilbertTempReal
			q1 += hilbertTempReal
			q1 -= prevq1Odd
			prevq1Odd = b * prevq1InputOdd
			q1 += prevq1Odd
			prevq1InputOdd = detrender
			q1 *= adjustedPrevPeriod
			hilbertTempReal = a * i1ForOddPrev3
			jI = -jIOdd[hilbertIdx]
			jIOdd[hilbertIdx] = hilbertTempReal
			jI += hilbertTempReal
			jI -= prevJIOdd
			prevJIOdd = b * prevJIInputOdd
			jI += prevJIOdd
			prevJIInputOdd = i1ForOddPrev3
			jI *= adjustedPrevPeriod
			hilbertTempReal = a * q1
			jQ = -jQOdd[hilbertIdx]
			jQOdd[hilbertIdx] = hilbertTempReal
			jQ += hilbertTempReal
			jQ -= prevJQOdd
			prevJQOdd = b * prevJQInputOdd
			jQ += prevJQOdd
			prevJQInputOdd = q1
			jQ *= adjustedPrevPeriod
			q2 = (0.2 * (q1 + jI)) + (0.8 * prevq2)
			i2 = (0.2 * (i1ForOddPrev3 - jQ)) + (0.8 * previ2)
			i1ForEvenPrev3 = i1ForEvenPrev2
			i1ForEvenPrev2 = detrender
		}
		Re = (0.2 * ((i2 * previ2) + (q2 * prevq2))) + (0.8 * Re)
		Im = (0.2 * ((i2 * prevq2) - (q2 * previ2))) + (0.8 * Im)
		prevq2 = q2
		previ2 = i2
		tempReal = period
		if (Im != 0.0) && (Re != 0.0) {
			period = 360.0 / (math.Atan(Im/Re) * rad2Deg)
		}
		tempReal2 := 1.5 * tempReal
		if period > tempReal2 {
			period = tempReal2
		}
		tempReal2 = 0.67 * tempReal
		if period < tempReal2 {
			period = tempReal2
		}
		if period < 6 {
			period = 6
		} else if period > 50 {
			period = 50
		}
		period = (0.2 * period) + (0.8 * tempReal)
		smoothPeriod = (0.33 * period) + (0.67 * smoothPeriod)
		DCPeriod := smoothPeriod + 0.5
		DCPeriodInt := math.Floor(DCPeriod)
		realPart := 0.0
		imagPart := 0.0
		idx := smoothPriceIdx
		for i := 0; i < int(DCPeriodInt); i++ {
			tempReal = (float64(i) * constDeg2RadBy360) / (DCPeriodInt * 1.0)
			tempReal2 = smoothPrice[idx]
			realPart += math.Sin(tempReal) * tempReal2
			imagPart += math.Cos(tempReal) * tempReal2
			if idx == 0 {
				idx = 50 - 1
			} else {
				idx--
			}
		}
		tempReal = math.Abs(imagPart)
		if tempReal > 0.0 {
			dcPhase = math.Atan(realPart/imagPart) * rad2Deg
		} else if tempReal <= 0.01 {
			if realPart < 0.0 {
				dcPhase -= 90.0
			} else if realPart > 0.0 {
				dcPhase += 90.0
			}
		}
		dcPhase += 90.0
		dcPhase += 360.0 / smoothPeriod
		if imagPart < 0.0 {
			dcPhase += 180.0
		}
		if dcPhase > 315.0 {
			dcPhase -= 360.0
		}
		if today >= startIdx {
			outSine[outIdx] = math.Sin(dcPhase * deg2Rad)
			outLeadSine[outIdx] = math.Sin((dcPhase + 45) * deg2Rad)
			outIdx++
		}
		smoothPriceIdx++
		if smoothPriceIdx > maxIdxSmoothPrice {
			smoothPriceIdx = 0
		}

		today++
	}
	return outSine, outLeadSine
}

// HtTrendMode - Hilbert Transform - Trend vs Cycle Mode (lookback=63)
func HtTrendMode(inReal []float64) []float64 {

	outReal := make([]float64, len(inReal))
	a := 0.0962
	b := 0.5769
	detrenderOdd := make([]float64, 3)
	detrenderEven := make([]float64, 3)
	q1Odd := make([]float64, 3)
	q1Even := make([]float64, 3)
	jIOdd := make([]float64, 3)
	jIEven := make([]float64, 3)
	jQOdd := make([]float64, 3)
	jQEven := make([]float64, 3)
	smoothPriceIdx := 0
	maxIdxSmoothPrice := (50 - 1)
	smoothPrice := make([]float64, maxIdxSmoothPrice+1)
	iTrend1 := 0.0
	iTrend2 := 0.0
	iTrend3 := 0.0
	daysInTrend := 0
	prevdcPhase := 0.0
	dcPhase := 0.0
	prevSine := 0.0
	sine := 0.0
	prevLeadSine := 0.0
	leadSine := 0.0
	tempReal := math.Atan(1)
	rad2Deg := 45.0 / tempReal
	deg2Rad := 1.0 / rad2Deg
	constDeg2RadBy360 := tempReal * 8.0
	lookbackTotal := 63
	startIdx := lookbackTotal
	trailingWMAIdx := startIdx - lookbackTotal
	today := trailingWMAIdx
	tempReal = inReal[today]
	today++
	periodWMASub := tempReal
	periodWMASum := tempReal
	tempReal = inReal[today]
	today++
	periodWMASub += tempReal
	periodWMASum += tempReal * 2.0
	tempReal = inReal[today]
	today++
	periodWMASub += tempReal
	periodWMASum += tempReal * 3.0
	trailingWMAValue := 0.0
	i := 34

	for ok := true; ok; {
		tempReal = inReal[today]
		today++
		periodWMASub += tempReal
		periodWMASub -= trailingWMAValue
		periodWMASum += tempReal * 4.0
		trailingWMAValue = inReal[trailingWMAIdx]
		trailingWMAIdx++
		//smoothedValue := periodWMASum * 0.1
		periodWMASum -= periodWMASub
		i--
		ok = i != 0
	}

	hilbertIdx := 0
	detrender := 0.0
	prevDetrenderOdd := 0.0
	prevDetrenderEven := 0.0
	prevDetrenderInputOdd := 0.0
	prevDetrenderInputEven := 0.0
	q1 := 0.0
	prevq1Odd := 0.0
	prevq1Even := 0.0
	prevq1InputOdd := 0.0
	prevq1InputEven := 0.0
	jI := 0.0
	prevJIOdd := 0.0
	prevJIEven := 0.0
	prevJIInputOdd := 0.0
	prevJIInputEven := 0.0
	jQ := 0.0
	prevJQOdd := 0.0
	prevJQEven := 0.0
	prevJQInputOdd := 0.0
	prevJQInputEven := 0.0
	period := 0.0
	outIdx := 63
	previ2 := 0.0
	prevq2 := 0.0
	Re := 0.0
	Im := 0.0
	i1ForOddPrev3 := 0.0
	i1ForEvenPrev3 := 0.0
	i1ForOddPrev2 := 0.0
	i1ForEvenPrev2 := 0.0
	smoothPeriod := 0.0
	dcPhase = 0.0
	smoothedValue := 0.0
	hilbertTempReal := 0.0
	q2 := 0.0
	i2 := 0.0
	for today < len(inReal) {
		adjustedPrevPeriod := (0.075 * period) + 0.54
		todayValue := inReal[today]
		periodWMASub += todayValue
		periodWMASub -= trailingWMAValue
		periodWMASum += todayValue * 4.0
		trailingWMAValue = inReal[trailingWMAIdx]
		trailingWMAIdx++
		smoothedValue = periodWMASum * 0.1
		periodWMASum -= periodWMASub

		smoothPrice[smoothPriceIdx] = smoothedValue
		if (today % 2) == 0 {
			hilbertTempReal = a * smoothedValue
			detrender = -detrenderEven[hilbertIdx]
			detrenderEven[hilbertIdx] = hilbertTempReal
			detrender += hilbertTempReal
			detrender -= prevDetrenderEven
			prevDetrenderEven = b * prevDetrenderInputEven
			detrender += prevDetrenderEven
			prevDetrenderInputEven = smoothedValue
			detrender *= adjustedPrevPeriod
			hilbertTempReal = a * detrender
			q1 = -q1Even[hilbertIdx]
			q1Even[hilbertIdx] = hilbertTempReal
			q1 += hilbertTempReal
			q1 -= prevq1Even
			prevq1Even = b * prevq1InputEven
			q1 += prevq1Even
			prevq1InputEven = detrender
			q1 *= adjustedPrevPeriod
			hilbertTempReal = a * i1ForEvenPrev3
			jI = -jIEven[hilbertIdx]
			jIEven[hilbertIdx] = hilbertTempReal
			jI += hilbertTempReal
			jI -= prevJIEven
			prevJIEven = b * prevJIInputEven
			jI += prevJIEven
			prevJIInputEven = i1ForEvenPrev3
			jI *= adjustedPrevPeriod
			hilbertTempReal = a * q1
			jQ = -jQEven[hilbertIdx]
			jQEven[hilbertIdx] = hilbertTempReal
			jQ += hilbertTempReal
			jQ -= prevJQEven
			prevJQEven = b * prevJQInputEven
			jQ += prevJQEven
			prevJQInputEven = q1
			jQ *= adjustedPrevPeriod
			hilbertIdx++
			if hilbertIdx == 3 {
				hilbertIdx = 0
			}
			q2 = (0.2 * (q1 + jI)) + (0.8 * prevq2)
			i2 = (0.2 * (i1ForEvenPrev3 - jQ)) + (0.8 * previ2)
			i1ForOddPrev3 = i1ForOddPrev2
			i1ForOddPrev2 = detrender
		} else {
			hilbertTempReal = a * smoothedValue
			detrender = -detrenderOdd[hilbertIdx]
			detrenderOdd[hilbertIdx] = hilbertTempReal
			detrender += hilbertTempReal
			detrender -= prevDetrenderOdd
			prevDetrenderOdd = b * prevDetrenderInputOdd
			detrender += prevDetrenderOdd
			prevDetrenderInputOdd = smoothedValue
			detrender *= adjustedPrevPeriod
			hilbertTempReal = a * detrender
			q1 = -q1Odd[hilbertIdx]
			q1Odd[hilbertIdx] = hilbertTempReal
			q1 += hilbertTempReal
			q1 -= prevq1Odd
			prevq1Odd = b * prevq1InputOdd
			q1 += prevq1Odd
			prevq1InputOdd = detrender
			q1 *= adjustedPrevPeriod
			hilbertTempReal = a * i1ForOddPrev3
			jI = -jIOdd[hilbertIdx]
			jIOdd[hilbertIdx] = hilbertTempReal
			jI += hilbertTempReal
			jI -= prevJIOdd
			prevJIOdd = b * prevJIInputOdd
			jI += prevJIOdd
			prevJIInputOdd = i1ForOddPrev3
			jI *= adjustedPrevPeriod
			hilbertTempReal = a * q1
			jQ = -jQOdd[hilbertIdx]
			jQOdd[hilbertIdx] = hilbertTempReal
			jQ += hilbertTempReal
			jQ -= prevJQOdd
			prevJQOdd = b * prevJQInputOdd
			jQ += prevJQOdd
			prevJQInputOdd = q1
			jQ *= adjustedPrevPeriod
			q2 = (0.2 * (q1 + jI)) + (0.8 * prevq2)
			i2 = (0.2 * (i1ForOddPrev3 - jQ)) + (0.8 * previ2)
			i1ForEvenPrev3 = i1ForEvenPrev2
			i1ForEvenPrev2 = detrender
		}
		Re = (0.2 * ((i2 * previ2) + (q2 * prevq2))) + (0.8 * Re)
		Im = (0.2 * ((i2 * prevq2) - (q2 * previ2))) + (0.8 * Im)
		prevq2 = q2
		previ2 = i2
		tempReal = period
		if (Im != 0.0) && (Re != 0.0) {
			period = 360.0 / (math.Atan(Im/Re) * rad2Deg)
		}
		tempReal2 := 1.5 * tempReal
		if period > tempReal2 {
			period = tempReal2
		}
		tempReal2 = 0.67 * tempReal
		if period < tempReal2 {
			period = tempReal2
		}
		if period < 6 {
			period = 6
		} else if period > 50 {
			period = 50
		}
		period = (0.2 * period) + (0.8 * tempReal)
		smoothPeriod = (0.33 * period) + (0.67 * smoothPeriod)
		prevdcPhase = dcPhase
		DCPeriod := smoothPeriod + 0.5
		DCPeriodInt := math.Floor(DCPeriod)
		realPart := 0.0
		imagPart := 0.0
		idx := smoothPriceIdx
		for i := 0; i < int(DCPeriodInt); i++ {
			tempReal = (float64(i) * constDeg2RadBy360) / (DCPeriodInt * 1.0)
			tempReal2 = smoothPrice[idx]
			realPart += math.Sin(tempReal) * tempReal2
			imagPart += math.Cos(tempReal) * tempReal2
			if idx == 0 {
				idx = 50 - 1
			} else {
				idx--
			}
		}
		tempReal = math.Abs(imagPart)
		if tempReal > 0.0 {
			dcPhase = math.Atan(realPart/imagPart) * rad2Deg
		} else if tempReal <= 0.01 {
			if realPart < 0.0 {
				dcPhase -= 90.0
			} else if realPart > 0.0 {
				dcPhase += 90.0
			}
		}
		dcPhase += 90.0
		dcPhase += 360.0 / smoothPeriod
		if imagPart < 0.0 {
			dcPhase += 180.0
		}
		if dcPhase > 315.0 {
			dcPhase -= 360.0
		}
		prevSine = sine
		prevLeadSine = leadSine
		sine = math.Sin(dcPhase * deg2Rad)
		leadSine = math.Sin((dcPhase + 45) * deg2Rad)
		DCPeriod = smoothPeriod + 0.5
		DCPeriodInt = math.Floor(DCPeriod)
		idx = today
		tempReal = 0.0
		for i := 0; i < int(DCPeriodInt); i++ {
			tempReal += inReal[idx]
			idx--
		}
		if DCPeriodInt > 0 {
			tempReal = tempReal / (DCPeriodInt * 1.0)
		}
		trendline := (4.0*tempReal + 3.0*iTrend1 + 2.0*iTrend2 + iTrend3) / 10.0
		iTrend3 = iTrend2
		iTrend2 = iTrend1
		iTrend1 = tempReal
		trend := 1
		if ((sine > leadSine) && (prevSine <= prevLeadSine)) || ((sine < leadSine) && (prevSine >= prevLeadSine)) {
			daysInTrend = 0
			trend = 0
		}
		daysInTrend++
		if float64(daysInTrend) < (0.5 * smoothPeriod) {
			trend = 0
		}
		tempReal = dcPhase - prevdcPhase
		if (smoothPeriod != 0.0) && ((tempReal > (0.67 * 360.0 / smoothPeriod)) && (tempReal < (1.5 * 360.0 / smoothPeriod))) {
			trend = 0
		}
		tempReal = smoothPrice[smoothPriceIdx]
		if (trendline != 0.0) && (math.Abs((tempReal-trendline)/trendline) >= 0.015) {
			trend = 1
		}
		if today >= startIdx {
			outReal[outIdx] = float64(trend)
			outIdx++
		}
		smoothPriceIdx++
		if smoothPriceIdx > maxIdxSmoothPrice {
			smoothPriceIdx = 0
		}

		today++
	}
	return outReal
}
