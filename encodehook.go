package structool

import (
	"reflect"
	"time"
)

func EncodeErrorToString(next EncodeHookFunc) EncodeHookFunc {
	return func(value reflect.Value) (interface{}, error) {
		in := value.Interface()

		if value.Type().Implements(errorInterface) {
			if in == nil {
				return "", nil
			}
			return in.(error).Error(), nil
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
			return v.String(), nil
		}

		return next(value)
	}
}
