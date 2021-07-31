package indicators

import "C"
import (
	"errors"
	"github.com/hannessi/gOanda"
	"strings"
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
	rates, err := filterRatesToUse(candlesticks, additionalParams)
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

func filterRatesToUse(candlestick []gOanda.Candlestick, additionalParams EMAParams) ([]float64, error) {
	values := make([]float64, 0)

	t := strings.ToLower(additionalParams.CandlePoint + additionalParams.RatePoint)

	switch t {
	case "openask":
		for _, v := range candlestick {
			values = append(values, v.Ask.O.ToFloat())
		}
	case "openmid":
		for _, v := range candlestick {
			values = append(values, v.Mid.O.ToFloat())
		}
	case "openbid":
		for _, v := range candlestick {
			values = append(values, v.Bid.O.ToFloat())
		}
	case "closeask":
		for _, v := range candlestick {
			values = append(values, v.Ask.C.ToFloat())
		}
	case "closemid":
		for _, v := range candlestick {
			values = append(values, v.Mid.C.ToFloat())
		}
	case "closebid":
		for _, v := range candlestick {
			values = append(values, v.Bid.C.ToFloat())
		}
	case "highask":
		for _, v := range candlestick {
			values = append(values, v.Ask.H.ToFloat())
		}
	case "highmid":
		for _, v := range candlestick {
			values = append(values, v.Mid.H.ToFloat())
		}
	case "highbid":
		for _, v := range candlestick {
			values = append(values, v.Bid.H.ToFloat())
		}
	case "lowask":
		for _, v := range candlestick {
			values = append(values, v.Ask.L.ToFloat())
		}
	case "lowmid":
		for _, v := range candlestick {
			values = append(values, v.Mid.L.ToFloat())
		}
	case "lowbid":
		for _, v := range candlestick {
			values = append(values, v.Bid.L.ToFloat())
		}
	default:
		return nil, errors.New("invalid parameters defined")
	}
	return values, nil
}

func reverse(list []float64) []float64{
	if len(list) == 0 {
		return list
	}
	return append(reverse(list[1:]), list[0])
}
