package tools

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// IfInt64 转int64
func IfInt64(v interface{}) (int64, error) {
	switch vv := v.(type) {
	case int:
		return int64(vv), nil
	case int8:
		return int64(vv), nil
	case int16:
		return int64(vv), nil
	case int32:
		return int64(vv), nil
	case int64:
		return vv, nil
	case []byte:
		i, err := strconv.ParseInt(string(vv), 10, 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	case json.Number:
		i, err := strconv.ParseInt(string(vv), 10, 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	case string:
		i, err := strconv.ParseInt(vv, 10, 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	}
	return 0, fmt.Errorf("unsupported type: %v", v)
}

// IfFloat 转float
func IfFloat(v interface{}) (float64, error) {
	switch vv := v.(type) {
	case json.Number:
		return vv.Float64()
	case float32:
		return float64(vv), nil
	case float64:
		return vv, nil
	case string:
		if vv == "" {
			return 0, nil
		}
		i, err := strconv.ParseFloat(vv, 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	case []byte:
		if string(vv) == "" {
			return 0, nil
		}
		i, err := strconv.ParseFloat(string(vv), 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	}
	return 0, fmt.Errorf("unsupported type: %v", v)
}
