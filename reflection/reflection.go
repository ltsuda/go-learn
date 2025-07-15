package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	value := getValue(x)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
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
