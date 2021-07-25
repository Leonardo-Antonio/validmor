package validmor

import (
	"errors"
	"reflect"
)

var (
	errRequiredES = errors.New("es un campo requirido y obligatorio, no debe estar vacio")
	errRequiredEN = errors.New("is a required and mandatory field, it must not be empty")
)

type required struct{}

func (r *required) Validate(val interface{}) (bool, error) {

	types := []reflect.Kind{
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.String,
		reflect.Bool,
		reflect.Array,
		reflect.Slice,
		reflect.Float32,
		reflect.Float64,
	}

	for _, t := range types {
		if reflect.ValueOf(val).Kind() == t {
			if reflect.ValueOf(val).IsZero() {
				if currentTypeErrLang == ERR_ES {
					return false, errRequiredES
				}
				return false, errRequiredEN
			}
			return true, nil
		}
	}
	return false, errRequiredEN
}
