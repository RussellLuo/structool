package structool

import (
	"net"
	"reflect"
	"strconv"
	"time"
)

func EncodeNumberToString(next EncodeHookFunc) EncodeHookFunc {
	return func(value reflect.Value) (interface{}, error) {
		switch v := value.Interface().(type) {
		case int:
			return strconv.Itoa(v), nil
		case int8:
			return strconv.FormatInt(int64(v), 10), nil
		case int16:
			return strconv.FormatInt(int64(v), 10), nil
		case int32:
			return strconv.FormatInt(int64(v), 10), nil
		case int64:
			return strconv.FormatInt(v, 10), nil
		case uint:
			return strconv.FormatUint(uint64(v), 10), nil
		case uint8:
			return strconv.FormatUint(uint64(v), 10), nil
		case uint16:
			return strconv.FormatUint(uint64(v), 10), nil
		case uint32:
			return strconv.FormatUint(uint64(v), 10), nil
		case uint64:
			return strconv.FormatUint(v, 10), nil
		case float32:
			return strconv.FormatFloat(float64(v), 'f', -1, 32), nil
		case float64:
			return strconv.FormatFloat(v, 'f', -1, 64), nil
		}

		return next(value)
	}
}

func EncodeErrorToString(next EncodeHookFunc) EncodeHookFunc {
	return func(value reflect.Value) (interface{}, error) {
		if value.Type().Implements(errorInterface) {
			v := value.Interface()
			if v == nil {
				return "", nil
			}
			return v.(error).Error(), nil
		}

		return next(value)
	}
}

func EncodeTimeToString(layout string) func(EncodeHookFunc) EncodeHookFunc {
	return func(next EncodeHookFunc) EncodeHookFunc {
		return func(value reflect.Value) (interface{}, error) {
			switch v := value.Interface().(type) {
			case time.Time:
				return v.Format(layout), nil
			case *time.Time:
				if v == nil {
					return "", nil
				}
				return v.Format(layout), nil
			}

			return next(value)
		}
	}
}

func EncodeDurationToString(next EncodeHookFunc) EncodeHookFunc {
	return func(value reflect.Value) (interface{}, error) {
		switch v := value.Interface().(type) {
		case time.Duration:
			return v.String(), nil
		case *time.Duration:
			if v == nil {
				return "", nil
			}
			return v.String(), nil
		}

		return next(value)
	}
}

func EncodeIPToString(next EncodeHookFunc) EncodeHookFunc {
	nilToEmpty := func(s string) string {
		if s == "<nil>" {
			return ""
		}
		return s
	}

	return func(value reflect.Value) (interface{}, error) {
		switch v := value.Interface().(type) {
		case net.IP:
			return nilToEmpty(v.String()), nil
		case *net.IP:
			if v == nil {
				return "", nil
			}
			return nilToEmpty(v.String()), nil
		}

		return next(value)
	}
}
