package structool

import (
	"errors"
	"net"
	"reflect"
	"strconv"
	"time"
)

var (
	errorInterface = reflect.TypeOf((*error)(nil)).Elem()
)

func DecodeStringToNumber(next DecodeHookFunc) DecodeHookFunc {
	return func(from, to reflect.Value) (interface{}, error) {
		if from.Kind() != reflect.String {
			return next(from, to)
		}

		value := from.Interface().(string)

		switch to.Interface().(type) {
		case int:
			v, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			return v, nil
		case int8:
			v, err := strconv.ParseInt(value, 10, 8)
			if err != nil {
				return nil, err
			}
			return int8(v), nil
		case int16:
			v, err := strconv.ParseInt(value, 10, 16)
			if err != nil {
				return nil, err
			}
			return int16(v), nil
		case int32:
			v, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				return nil, err
			}
			return int32(v), nil
		case int64:
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return nil, err
			}
			return v, nil
		case uint:
			v, err := strconv.ParseUint(value, 10, 0)
			if err != nil {
				return nil, err
			}
			return v, nil
		case uint8:
			v, err := strconv.ParseUint(value, 10, 8)
			if err != nil {
				return nil, err
			}
			return uint8(v), nil
		case uint16:
			v, err := strconv.ParseUint(value, 10, 16)
			if err != nil {
				return nil, err
			}
			return uint16(v), nil
		case uint32:
			v, err := strconv.ParseUint(value, 10, 32)
			if err != nil {
				return nil, err
			}
			return uint32(v), nil
		case uint64:
			v, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return nil, err
			}
			return v, nil
		case float32:
			v, err := strconv.ParseFloat(value, 32)
			if err != nil {
				return nil, err
			}
			return float32(v), nil
		case float64:
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, err
			}
			return v, nil
		}

		return next(from, to)
	}
}

func DecodeStringToError(next DecodeHookFunc) DecodeHookFunc {
	return func(from, to reflect.Value) (interface{}, error) {
		if from.Kind() != reflect.String {
			return next(from, to)
		}

		value := from.Interface().(string)

		if to.Type().Implements(errorInterface) {
			if value == "" {
				return nil, nil
			}
			return errors.New(value), nil
		}

		return next(from, to)
	}
}

func DecodeStringToTime(layout string) func(DecodeHookFunc) DecodeHookFunc {
	return func(next DecodeHookFunc) DecodeHookFunc {
		return func(from, to reflect.Value) (interface{}, error) {
			if from.Kind() != reflect.String {
				return next(from, to)
			}

			value := from.Interface().(string)

			switch to.Interface().(type) {
			case time.Time:
				return time.Parse(layout, value)
			case *time.Time:
				t, err := time.Parse(layout, value)
				if err != nil {
					return nil, err
				}
				return &t, nil
			}

			return next(from, to)
		}
	}
}

func DecodeStringToDuration(next DecodeHookFunc) DecodeHookFunc {
	return func(from, to reflect.Value) (interface{}, error) {
		if from.Kind() != reflect.String {
			return next(from, to)
		}

		value := from.Interface().(string)

		switch to.Interface().(type) {
		case time.Duration:
			return time.ParseDuration(value)
		case *time.Duration:
			d, err := time.ParseDuration(value)
			if err != nil {
				return nil, err
			}
			return &d, nil
		}

		return next(from, to)
	}
}

func DecodeStringToIP(next DecodeHookFunc) DecodeHookFunc {
	return func(from, to reflect.Value) (interface{}, error) {
		if from.Kind() != reflect.String {
			return next(from, to)
		}

		value := from.Interface().(string)

		switch to.Interface().(type) {
		case net.IP:
			return net.ParseIP(value), nil
		case *net.IP:
			ip := net.ParseIP(value)
			return &ip, nil
		}

		return next(from, to)
	}
}
