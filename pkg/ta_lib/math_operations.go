package ta_lib

import (
	"errors"
	"math"
)

/* Math Operator Functions */

// Add - Vector arithmetic addition
func Add(inReal0 []float64, inReal1 []float64) []float64 {
	outReal := make([]float64, len(inReal0))
	for i := 0; i < len(inReal0); i++ {
		outReal[i] = inReal0[i] + inReal1[i]
	}
	return outReal
}

// Div - Vector arithmetic division
func Div(inReal0 []float64, inReal1 []float64) []float64 {
	outReal := make([]float64, len(inReal0))
	for i := 0; i < len(inReal0); i++ {
		outReal[i] = inReal0[i] / inReal1[i]
	}
	return outReal
}

// DivSlice - Divides vector by a float
func DivSlice(inReal []float64, n float64) []float64 {

	var outReal []float64

	for i := 0; i < len(inReal); i++ {
		outReal = append(outReal, inReal[i]/n)
	}

	return outReal
}

// Max - Highest value over a period
func Max(inReal []float64, inTimePeriod int) []float64 {

	outReal := make([]float64, len(inReal))

	if inTimePeriod < 2 {
		return outReal
	}

	nbInitialElementNeeded := inTimePeriod - 1
	startIdx := nbInitialElementNeeded
	outIdx := startIdx
	today := startIdx
	trailingIdx := startIdx - nbInitialElementNeeded
	highestIdx := -1
	highest := 0.0

	for today < len(outReal) {

		tmp := inReal[today]

		if highestIdx < trailingIdx {
			highestIdx = trailingIdx
			highest = inReal[highestIdx]
			i := highestIdx + 1
			for i <= today {
				tmp = inReal[i]
				if tmp > highest {
					highestIdx = i
					highest = tmp
				}
				i++
			}
		} else if tmp >= highest {
			highestIdx = today
			highest = tmp
		}
		outReal[outIdx] = highest
		outIdx++
		trailingIdx++
		today++
	}

	return outReal
}

func MinFloat64Slice(input []float64) float64 {

	var outMin float64

	if len(input) > 0 {
		outMin = input[0]
	}
	for i := 1; i < len(input); i++ {
		if input[i] < outMin {
			outMin = input[i]
		}
	}
	return outMin
}

func MaxFloat64Slice(input []float64) float64 {

	var outMax float64

	if len(input) > 0 {
		outMax = input[0]
	}
	for i := 1; i < len(input); i++ {
		if input[i] > outMax {
			outMax = input[i]
		}
	}
	return outMax
}

// MaxIndex - Index of highest value over a specified period
func MaxIndex(inReal []float64, inTimePeriod int) []float64 {

	outReal := make([]float64, len(inReal))

	if inTimePeriod < 2 {
		return outReal
	}

	nbInitialElementNeeded := inTimePeriod - 1
	startIdx := nbInitialElementNeeded
	outIdx := startIdx
	today := startIdx
	trailingIdx := startIdx - nbInitialElementNeeded
	highestIdx := -1
	highest := 0.0
	for today < len(inReal) {
		tmp := inReal[today]
		if highestIdx < trailingIdx {
			highestIdx = trailingIdx
			highest = inReal[highestIdx]
			i := highestIdx + 1
			for i <= today {
				tmp := inReal[i]
				if tmp > highest {
					highestIdx = i
					highest = tmp
				}
				i++
			}
		} else if tmp >= highest {
			highestIdx = today
			highest = tmp
		}
		outReal[outIdx] = float64(highestIdx)
		outIdx++
		trailingIdx++
		today++
	}

	return outReal
}

// Min - Lowest value over a period
func Min(inReal []float64, inTimePeriod int) []float64 {

	outReal := make([]float64, len(inReal))

	if inTimePeriod < 2 {
		return outReal
	}

	nbInitialElementNeeded := inTimePeriod - 1
	startIdx := nbInitialElementNeeded
	outIdx := startIdx
	today := startIdx
	trailingIdx := startIdx - nbInitialElementNeeded
	lowestIdx := -1
	lowest := 0.0
	for today < len(outReal) {

		tmp := inReal[today]

		if lowestIdx < trailingIdx {
			lowestIdx = trailingIdx
			lowest = inReal[lowestIdx]
			i := lowestIdx + 1
			for i <= today {
				tmp = inReal[i]
				if tmp < lowest {
					lowestIdx = i
					lowest = tmp
				}
				i++
			}
		} else if tmp <= lowest {
			lowestIdx = today
			lowest = tmp
		}
		outReal[outIdx] = lowest
		outIdx++
		trailingIdx++
		today++
	}

	return outReal
}

// MinIndex - Index of lowest value over a specified period
func MinIndex(inReal []float64, inTimePeriod int) []float64 {

	outReal := make([]float64, len(inReal))

	if inTimePeriod < 2 {
		return outReal
	}

	nbInitialElementNeeded := inTimePeriod - 1
	startIdx := nbInitialElementNeeded
	outIdx := startIdx
	today := startIdx
	trailingIdx := startIdx - nbInitialElementNeeded
	lowestIdx := -1
	lowest := 0.0
	for today < len(inReal) {
		tmp := inReal[today]
		if lowestIdx < trailingIdx {
			lowestIdx = trailingIdx
			lowest = inReal[lowestIdx]
			i := lowestIdx + 1
			for i <= today {
				tmp = inReal[i]
				if tmp < lowest {
					lowestIdx = i
					lowest = tmp
				}
				i++
			}
		} else if tmp <= lowest {
			lowestIdx = today
			lowest = tmp
		}
		outReal[outIdx] = float64(lowestIdx)
		outIdx++
		trailingIdx++
		today++
	}
	return outReal
}

// MinMax - Lowest and highest values over a specified period
func MinMax(inReal []float64, inTimePeriod int) ([]float64, []float64) {

	outMin := make([]float64, len(inReal))
	outMax := make([]float64, len(inReal))

	nbInitialElementNeeded := (inTimePeriod - 1)
	startIdx := nbInitialElementNeeded
	outIdx := startIdx
	today := startIdx
	trailingIdx := startIdx - nbInitialElementNeeded
	highestIdx := -1
	highest := 0.0
	lowestIdx := -1
	lowest := 0.0
	for today < len(inReal) {
		tmpLow, tmpHigh := inReal[today], inReal[today]
		if highestIdx < trailingIdx {
			highestIdx = trailingIdx
			highest = inReal[highestIdx]
			i := highestIdx
			i++
			for i <= today {
				tmpHigh = inReal[i]
				if tmpHigh > highest {
					highestIdx = i
					highest = tmpHigh
				}
				i++
			}
		} else if tmpHigh >= highest {
			highestIdx = today
			highest = tmpHigh
		}
		if lowestIdx < trailingIdx {
			lowestIdx = trailingIdx
			lowest = inReal[lowestIdx]
			i := lowestIdx
			i++
			for i <= today {
				tmpLow = inReal[i]
				if tmpLow < lowest {
					lowestIdx = i
					lowest = tmpLow
				}
				i++
			}
		} else if tmpLow <= lowest {
			lowestIdx = today
			lowest = tmpLow
		}
		outMax[outIdx] = highest
		outMin[outIdx] = lowest
		outIdx++
		trailingIdx++
		today++
	}
	return outMin, outMax
}

// MinMaxIndex - Indexes of lowest and highest values over a specified period
func MinMaxIndex(inReal []float64, inTimePeriod int) ([]float64, []float64) {

	outMinIdx := make([]float64, len(inReal))
	outMaxIdx := make([]float64, len(inReal))

	nbInitialElementNeeded := (inTimePeriod - 1)
	startIdx := nbInitialElementNeeded
	outIdx := startIdx
	today := startIdx
	trailingIdx := startIdx - nbInitialElementNeeded
	highestIdx := -1
	highest := 0.0
	lowestIdx := -1
	lowest := 0.0
	for today < len(inReal) {
		tmpLow, tmpHigh := inReal[today], inReal[today]
		if highestIdx < trailingIdx {
			highestIdx = trailingIdx
			highest = inReal[highestIdx]
			i := highestIdx
			i++
			for i <= today {
				tmpHigh = inReal[i]
				if tmpHigh > highest {
					highestIdx = i
					highest = tmpHigh
				}
				i++
			}
		} else if tmpHigh >= highest {
			highestIdx = today
			highest = tmpHigh
		}
		if lowestIdx < trailingIdx {
			lowestIdx = trailingIdx
			lowest = inReal[lowestIdx]
			i := lowestIdx
			i++
			for i <= today {
				tmpLow = inReal[i]
				if tmpLow < lowest {
					lowestIdx = i
					lowest = tmpLow
				}
				i++
			}
		} else if tmpLow <= lowest {
			lowestIdx = today
			lowest = tmpLow
		}
		outMaxIdx[outIdx] = float64(highestIdx)
		outMinIdx[outIdx] = float64(lowestIdx)
		outIdx++
		trailingIdx++
		today++
	}
	return outMinIdx, outMaxIdx
}

// Mult - Vector arithmetic multiply
func Mult(inReal0 []float64, inReal1 []float64) []float64 {
	outReal := make([]float64, len(inReal0))
	for i := 0; i < len(inReal0); i++ {
		outReal[i] = inReal0[i] * inReal1[i]
	}
	return outReal
}

// Sub - Vector arithmetic subtraction
func Sub(inReal0 []float64, inReal1 []float64) []float64 {
	outReal := make([]float64, len(inReal0))
	for i := 0; i < len(inReal0); i++ {
		outReal[i] = inReal0[i] - inReal1[i]
	}
	return outReal
}

// Sum - Vector summation
func Sum(inReal []float64, inTimePeriod int) []float64 {

	outReal := make([]float64, len(inReal))

	lookbackTotal := inTimePeriod - 1
	startIdx := lookbackTotal
	periodTotal := 0.0
	trailingIdx := startIdx - lookbackTotal
	i := trailingIdx
	if inTimePeriod > 1 {
		for i < startIdx {
			periodTotal += inReal[i]
			i++
		}
	}
	outIdx := startIdx
	for i < len(inReal) {
		periodTotal += inReal[i]
		tempReal := periodTotal
		periodTotal -= inReal[trailingIdx]
		outReal[outIdx] = tempReal
		i++
		trailingIdx++
		outIdx++
	}

	return outReal
}

// HeikinashiCandles - from candle values extracts heikinashi candle values.
// Returns highs, opens, closes and lows of the heikinashi candles (in this order).
//
//    NOTE: The number of Heikin-Ashi candles will always be one less than the number of provided candles, due to the fact
//          that a previous candle is necessary to calculate the Heikin-Ashi candle, therefore the first provided candle is not considered
//          as "current candle" in the algorithm, but only as "previous candle".
func HeikinashiCandles(highs []float64, opens []float64, closes []float64, lows []float64) ([]float64, []float64, []float64, []float64) {
	N := len(highs)

	heikinHighs := make([]float64, N)
	heikinOpens := make([]float64, N)
	heikinCloses := make([]float64, N)
	heikinLows := make([]float64, N)

	for currentCandle := 1; currentCandle < N; currentCandle++ {
		previousCandle := currentCandle - 1

		heikinHighs[currentCandle] = math.Max(highs[currentCandle], math.Max(opens[currentCandle], closes[currentCandle]))
		heikinOpens[currentCandle] = (opens[previousCandle] + closes[previousCandle]) / 2
		heikinCloses[currentCandle] = (highs[currentCandle] + opens[currentCandle] + closes[currentCandle] + lows[currentCandle]) / 4
		heikinLows[currentCandle] = math.Min(highs[currentCandle], math.Min(opens[currentCandle], closes[currentCandle]))
	}

	return heikinHighs, heikinOpens, heikinCloses, heikinLows
}

// Hlc3 returns the Hlc3 values
//
//     NOTE: Every Hlc item is defined as follows : (high + low + close) / 3
//           It is used as AvgPrice candle.
func Hlc3(highs []float64, lows []float64, closes []float64) []float64 {
	N := len(highs)
	result := make([]float64, N)
	for i := range highs {
		result[i] = (highs[i] + lows[i] + closes[i]) / 3
	}

	return result
}

// Crossover returns true if series1 is crossing over series2.
//
//    NOTE: Usually this is used with Media Average Series to check if it crosses for buy signals.
//          It assumes first values are the most recent.
//          The crossover function does not use most recent value, since usually it's not a complete candle.
//          The second recent values and the previous are used, instead.
func Crossover(series1 []float64, series2 []float64) bool {
	if len(series1) < 3 || len(series2) < 3 {
		return false
	}

	N := len(series1)

	return series1[N-2] <= series2[N-2] && series1[N-1] > series2[N-1]
}

// Crossunder returns true if series1 is crossing under series2.
//
//    NOTE: Usually this is used with Media Average Series to check if it crosses for sell signals.
func Crossunder(series1 []float64, series2 []float64) bool {
	if len(series1) < 3 || len(series2) < 3 {
		return false
	}

	N := len(series1)

	return series1[N-1] <= series2[N-1] && series1[N-2] > series2[N-2]
}

// GroupCandles groups a set of candles in another set of candles, based on a grouping factor.
// This is pretty useful if you want to transform, for example, 15min candles into 1h candles using the same data.
// This avoids calling the exchange multiple times for multiple contexts.
//
// Example:
//     To transform 15 minute candles in 30 minutes candles you have a grouping factor = 2
//     To transform 15 minute candles in 1 hour candles you have a grouping factor = 4
//     To transform 30 minute candles in 1 hour candles you have a grouping factor = 2
func GroupCandles(highs []float64, opens []float64, closes []float64, lows []float64, groupingFactor int) ([]float64, []float64, []float64, []float64, error) {
	N := len(highs)
	if groupingFactor == 0 {
		return nil, nil, nil, nil, errors.New("grouping factor must be > 0")
	} else if groupingFactor == 1 {
		return highs, opens, closes, lows, nil // no need to group in this case, return the parameters.
	}
	if N%groupingFactor > 0 {
		return nil, nil, nil, nil, errors.New("cannot group properly, need a groupingFactor which is a factor of the number of candles")
	}

	groupedN := N / groupingFactor

	groupedHighs := make([]float64, groupedN)
	groupedOpens := make([]float64, groupedN)
	groupedCloses := make([]float64, groupedN)
	groupedLows := make([]float64, groupedN)

	lastOfCurrentGroup := groupingFactor - 1

	k := 0
	for i := 0; i < N; i += groupingFactor { // scan all param candles
		groupedOpens[k] = opens[i]
		groupedCloses[k] = closes[i+lastOfCurrentGroup]

		groupedHighs[k] = highs[i]
		groupedLows[k] = lows[i]

		endOfCurrentGroup := i + lastOfCurrentGroup
		for j := i + 1; j <= endOfCurrentGroup; j++ { // group high lows candles here
			if lows[j] < groupedLows[k] {
				groupedLows[k] = lows[j]
			}
			if highs[j] > groupedHighs[k] {
				groupedHighs[k] = highs[j]
			}
		}
		k++
	}
	return groupedHighs, groupedOpens, groupedCloses, groupedLows, nil
}
