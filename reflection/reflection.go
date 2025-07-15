package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	value := getValue(x)

	walkValue := func(val reflect.Value) {
		walk(val.Interface(), fn)
	}

	switch value.Kind() {
	case reflect.String:
		fn(value.String())
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			walkValue(value.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			walkValue(value.Field(i))
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			walkValue(value.MapIndex(key))
		}
	case reflect.Chan:
		for {
			if v, ok := value.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valueFromFn := value.Call(nil)
		for _, val := range valueFromFn {
			walkValue(val)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)

	if value.Kind() == reflect.Pointer {
		value = value.Elem()
	}
	return value
}
