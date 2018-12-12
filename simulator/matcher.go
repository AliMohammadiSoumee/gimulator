package simulator

import (
	"reflect"
)

func match(expected, actual interface{}) bool {
	if expected == nil {
		return true
	}

	l, r := reflect.ValueOf(expected), reflect.ValueOf(actual)

	if l.Type() != r.Type() {
		return false
	}
	return matchValue(l, r)
}

func matchValue(expected, actual reflect.Value) bool {
	if expected.Kind() != actual.Kind() {
		return false
	}

	switch expected.Kind() {
	case reflect.Array:
		for i := 0; i < expected.Len(); i++ {
			if !matchValue(expected.Index(i), actual.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Slice:
		if expected.IsNil() != actual.IsNil() {
			return false
		}
		if expected.Len() != actual.Len() {
			return false
		}
		if expected.Pointer() == actual.Pointer() {
			return true
		}
		for i := 0; i < expected.Len(); i++ {
			if !matchValue(expected.Index(i), actual.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Interface:
		if expected.IsNil() || actual.IsNil() {
			return expected.IsNil() == actual.IsNil()
		}
		return matchValue(expected.Elem(), actual.Elem())
	case reflect.Ptr:
		if expected.Pointer() == actual.Pointer() {
			return true
		}
		return matchValue(expected.Elem(), actual.Elem())
	case reflect.Map:
		if expected.IsNil() {
			return true
		}
		if expected.Pointer() == actual.Pointer() {
			return true
		}
		for _, k := range expected.MapKeys() {
			val1 := expected.MapIndex(k)
			val2 := actual.MapIndex(k)
			if !val1.IsValid() || !val2.IsValid() || !matchValue(val1, val2) {
				return false
			}
		}
		return true
	default:
		return expected.Interface() == actual.Interface()
	}
}
