package indicators

import (
	"errors"
	"github.com/hannessi/gOanda"
)

type EMAParams struct {
	// open, close, high, low
	CandlePoint string
	// ask, mid, bid
	RatePoint string
	// Reversed == true if last latest value id not the first index
	Reversed bool
}

func EMA(period int, candlesticks []gOanda.Candlestick) ([]float64, error) {
	return EMAWithOptions(period, candlesticks, EMAParams{
		CandlePoint: "close",
		RatePoint:   "mid",
	})
}

func EMAWithOptions(period int, candlesticks []gOanda.Candlestick, additionalParams EMAParams) ([]float64, error) {
	rates, err := filterRatesToUse(candlesticks, additionalParams.CandlePoint, additionalParams.RatePoint)
	if err != nil {
		return nil, err
	}

	if len(candlesticks) <= period*3 {
		return nil, errors.New("not enough data points to provide an accurate EMA")
	}

	if !additionalParams.Reversed {
		rates = reverse(rates)
	}

	return ema(period, rates), nil
}

func ema(period int, list []float64) []float64 {
	emaSlice := make([]float64, 0)

	ak := period + 1
	k := float64(2) / float64(ak)

	emaSlice = append(emaSlice, list[0])

	for i := 1; i < len(list); i++ {
		emaSlice = append(emaSlice, (list[i]*k)+(emaSlice[i-1]*(1-k)))
	}

	return emaSlice
}

