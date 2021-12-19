package ta_lib

import "fmt"

type IchimokuCloud struct {
	Price       float64 // Close price of candle
	TenkanSen   float64 // Tenkan-San
	KijunSen    float64 // Kijun-Sen
	SenkouSpanA float64 // Senkou span A
	SenkouSpanB float64 // Senkou span B
	ChikouSpan  float64 // Chico span
}

type Candle struct {
	Timestamp string
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    float64
}

// GetIchimokuCloud - This function will return and IchimokuCloud struct based on provided price data
func GetIchimokuCloud(priceData []Candle) (*IchimokuCloud, error) {

	if len(priceData) < 51 {
		return nil, fmt.Errorf("ichimoku cloud requires at least 52 data points to calculate. you only provided %v", len(priceData))
	}

	ichi := &IchimokuCloud{}
	ichi.Price = priceData[len(priceData)-1].Close
	ichi.CalculateTenkanSen(priceData[len(priceData)-9:])
	ichi.CalculateKijunSen(priceData[len(priceData)-26:])
	ichi.CalculateSenkouSpanA()
	ichi.CalculateSenkouSpanB(priceData[len(priceData)-52:])
	ichi.ChikouSpan = priceData[len(priceData)-26].Close

	return ichi, nil
}

//CalculateTenkanSen Tenkan-sen (red line): Also known as the conversion line.
func (i *IchimokuCloud) CalculateTenkanSen(priceData []Candle) {
	// Calculated by adding the maximum price of the last 9 periods to the minimum price of the last 9 periods,
	// and dividing this value by 2. In terms of moving averages, this can be thought of as a short-term, faster moving line.
	// (Highest High + Lowest Low) / 2
	i.TenkanSen = MinLowPlusMaxHighDiv2(priceData)

}

//CalculateKijunSen Kijun-sen (blue line): Also known as the base line
func (i *IchimokuCloud) CalculateKijunSen(priceData []Candle) {
	// Calculated by adding the maximum price of the last 26 periods to the minimum price of the last 26 periods,
	// and dividing this value by 2. This can be thought of as a slower, long-term version of Tenkan-sen.
	i.KijunSen = MinLowPlusMaxHighDiv2(priceData)

}

//CalculateSenkouSpanA Senkou span A (orange-ish line): Also known as Leading Span A.
func (i *IchimokuCloud) CalculateSenkouSpanA() {
	// Calculated by adding the Tenkan-sen and Kijun sen, and dividing by 2.
	// The line is plotted 26 periods ahead, the opposite direction that the Chiko span is plotted.
	i.SenkouSpanA = (i.TenkanSen + i.KijunSen) / 2
}

//CalculateSenkouSpanB Senkou span B (purple-ish line): Also known as Leading Span B.
func (i *IchimokuCloud) CalculateSenkouSpanB(priceData []Candle) {
	// Calculated by adding the highest high of the last 52 periods and the lowest low of the last 52 periods,
	// and dividing by 2. Like Senkou span A, it is also plotted 26 periods ahead.
	i.SenkouSpanB = MinLowPlusMaxHighDiv2(priceData)
}

//ChikoSpan Chiko span (green line): Also known as the lagging line.
func (i *IchimokuCloud) ChikoSpan() []float64 {
	// This is simply the prices plotted 26 periods behind the actual price.
	// So, on a daily chart, the price for November 26 would be shown on the chart at November 1.
	// This one is slightly more confusing to look at, but it gets easier once you start matching up where the indicator is lagging behind.
	return nil // ToDo: yeah, tbd
}

func (i *IchimokuCloud) AboveKumo() bool {
	// Together, Senkou span A and B make a cloud shaped area that is filled in.
	// This area is known as Kumo. Kumo serves as indicators of support and resistance.
	// If we are above Kumo, it is a bullish indicator.
	return false // Todo
}

func (i *IchimokuCloud) BelowKumo() bool {
	// Together, Senkou span A and B make a cloud shaped area that is filled in.
	// This area is known as Kumo. Kumo serves as indicators of support and resistance.
	// If we are below Kumo, it is a bearish indicator.
	return false // Todo
}

func (i *IchimokuCloud) BuySignal() (buy bool) {

	if i.TenkanSen > i.KijunSen {
		buy = true
	}

	if i.KijunSen < i.TenkanSen {
		buy = false
	}

	return buy
}

func MinLow(input []Candle) float64 {

	var outMin float64

	if len(input) > 0 {
		outMin = input[0].Low
	}
	for i := 1; i < len(input); i++ {
		if input[i].Low < outMin {
			outMin = input[i].Low
		}
	}
	return outMin
}

func MaxHigh(input []Candle) float64 {

	var outMax float64

	if len(input) > 0 {
		outMax = input[0].High
	}
	for i := 1; i < len(input); i++ {
		if input[i].High > outMax {
			outMax = input[i].High
		}
	}
	return outMax
}

func MinLowPlusMaxHighDiv2(input []Candle) float64 {
	return (MinLow(input) + MaxHigh(input)) / 2
}
