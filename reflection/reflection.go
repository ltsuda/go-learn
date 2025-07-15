package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	value := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch value.Kind() {
	case reflect.String:
		fn(value.String())
	case reflect.Slice, reflect.Array:
		numberOfValues = value.Len()
		getField = value.Index
	case reflect.Struct:
		numberOfValues = value.NumField()
		getField = value.Field
	case reflect.Map:
		for _, key := range value.MapKeys() {
			walk(value.MapIndex(key).Interface(), fn)
		}
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)

	if value.Kind() == reflect.Pointer {
		value = value.Elem()
	}
	return value
}
