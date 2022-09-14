package indicators

import (
	"errors"
	"github.com/hannessi/gOanda"
)

type SMMAParams struct {
	// open, close, high, low
	CandlePoint string
	// ask, mid, bid
	RatePoint string
	// Reversed == true if last latest value is not the first index
	Reversed bool
}

func SMMA(period int, candlesticks []gOanda.Candlestick, additionalParams EMAParams) ([]float64, error) {
	return SMMAWithOptions(period, candlesticks, SMMAParams{
		CandlePoint: "close",
		RatePoint:   "mid",
	})
}

func SMMAWithOptions(period int, candlesticks []gOanda.Candlestick, additionalParams SMMAParams) ([]float64, error) {
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

	return smma(period, rates), nil
}

func smma(period int, list []float64) []float64 {
	smmaSlice := make([]float64, 0)

	smmaSlice = append(smmaSlice, list[0])

	for i := 1; i < len(list); i++ {
		partialList := make([]float64,0)
		if i < period {
			partialList = list[:i]
		} else {
			partialList = list[(i-period):i]
		}

		total := sum(partialList)
		smmaSlice = append(smmaSlice, (total-smmaSlice[i-1])/float64(period))
	}

	return smmaSlice
}

func sum(values []float64) float64 {
	s := 0.0
	for _, v := range values {
		s += v
	}
	return s
}