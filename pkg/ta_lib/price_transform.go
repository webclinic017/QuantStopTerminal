package ta_lib

/* Price Transform */

// AvgPrice - Average Price (o+h+l+c)/4
func AvgPrice(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) []float64 {

	outReal := make([]float64, len(inClose))
	outIdx := 0
	startIdx := 0

	for i := startIdx; i < len(inClose); i++ {
		outReal[outIdx] = (inHigh[i] + inLow[i] + inClose[i] + inOpen[i]) / 4
		outIdx++
	}
	return outReal
}

// MedPrice - Median Price (h+l)/2
func MedPrice(inHigh []float64, inLow []float64) []float64 {

	outReal := make([]float64, len(inHigh))
	outIdx := 0
	startIdx := 0

	for i := startIdx; i < len(inHigh); i++ {
		outReal[outIdx] = (inHigh[i] + inLow[i]) / 2.0
		outIdx++
	}
	return outReal
}

// TypPrice - Typical Price (h+l+c)/3
func TypPrice(inHigh []float64, inLow []float64, inClose []float64) []float64 {

	outReal := make([]float64, len(inClose))
	outIdx := 0
	startIdx := 0

	for i := startIdx; i < len(inClose); i++ {
		outReal[outIdx] = (inHigh[i] + inLow[i] + inClose[i]) / 3.0
		outIdx++
	}
	return outReal
}

// WclPrice - Weighted Close Price
func WclPrice(inHigh []float64, inLow []float64, inClose []float64) []float64 {

	outReal := make([]float64, len(inClose))
	outIdx := 0
	startIdx := 0

	for i := startIdx; i < len(inClose); i++ {
		outReal[outIdx] = (inHigh[i] + inLow[i] + (inClose[i] * 2.0)) / 4.0
		outIdx++
	}
	return outReal
}
