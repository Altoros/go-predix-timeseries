package measurement

import (
	"strconv"
)

func FromString(s string) (Measurement, error) {
	var intMeasure int64
	var floatMeasure float64

	intMeasure, err := strconv.ParseInt(s, 10, 32)
	if err == nil {
		return Int(intMeasure), nil
	}
	switch err.(*strconv.NumError).Err {
	case strconv.ErrRange:
		intMeasure, err = strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		return Longint(intMeasure), nil
	case strconv.ErrSyntax:
		if floatMeasure, err = strconv.ParseFloat(s, 32); err == nil {
			return Float(floatMeasure), nil
		}
		switch err.(*strconv.NumError).Err {
		case strconv.ErrRange:
			floatMeasure, err = strconv.ParseFloat(s, 64)
			if err != nil {
				return nil, err
			}
			return Double(floatMeasure), nil
		default:
			return String(s), nil
		}
	default:
		return nil, err
	}
}
