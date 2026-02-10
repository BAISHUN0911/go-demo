package common

import (
	"encoding/json"
	"math"
	"strconv"
)

func ToFloat64Round2(v interface{}) (float64, bool) {
	var f float64

	switch val := v.(type) {
	case json.Number:
		n, err := val.Float64()
		if err != nil {
			return 0, false
		}
		f = n
	case float64:
		f = val
	case float32:
		f = float64(val)
	case int:
		f = float64(val)
	case int64:
		f = float64(val)
	case uint:
		f = float64(val)
	case uint64:
		f = float64(val)
	case string:
		n, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return 0, false
		}
		f = n
	default:
		return 0, false
	}

	// 四舍五入保留 2 位小数
	return math.Round(f*100) / 100, true
}
