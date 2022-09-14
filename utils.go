package indicators

import (
	"errors"
	"github.com/hannessi/gOanda"
	"strings"
)

func filterRatesToUse(candlestick []gOanda.Candlestick, candlePoint, ratePoint string) ([]float64, error) {
	values := make([]float64, 0)

	t := strings.ToLower(candlePoint + ratePoint)

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

func reverse(list []float64) []float64 {
	if len(list) == 0 {
		return list
	}
	return append(reverse(list[1:]), list[0])
}
