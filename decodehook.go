package structool

import (
	"errors"
	"reflect"
	"time"
)

var (
	errorInterface = reflect.TypeOf((*error)(nil)).Elem()
)

func DecodeStringToError(next DecodeHookFunc) DecodeHookFunc {
	return func(from, to reflect.Value) (interface{}, error) {
		if from.Kind() != reflect.String {
			return next(from, to)
		}

		value := from.Interface().(string)
		if to.Type().Implements(errorInterface) {
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
