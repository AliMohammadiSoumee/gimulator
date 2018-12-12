package simulator

import (
	"reflect"
)

type Matcher struct {
	Filter interface{}
}

func (m *Matcher) match(left interface{}) bool {
	if m.Filter == nil {
		return true
	}
	return match(left, m.Filter)
}

func match(left, right interface{}) bool {
	if left == nil || right == nil {
		return left == right
	}
	l, r := reflect.ValueOf(left), reflect.ValueOf(right)

	if l.Type() != r.Type() {
		return false
	}
	return matchValue(l, r)
}

func matchValue(left, right reflect.Value) bool {
	if !left.IsValid() && right.IsValid() {
		return false
	}
	if left.Type() != right.Type() {
		return false
	}

	switch left.Kind() {
	case reflect.Array:
		for i := 0; i < left.Len(); i++ {
			if !matchValue(left.Index(i), right.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Slice:
		if left.IsNil() != right.IsNil() {
			return false
		}
		if left.Len() != right.Len() {
			return false
		}
		if left.Pointer() == right.Pointer() {
			return true
		}
		for i := 0; i < left.Len(); i++ {
			if !matchValue(left.Index(i), right.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Interface:
		if left.IsNil() || right.IsNil() {
			return left.IsNil() == right.IsNil()
		}
		return matchValue(left.Elem(), right.Elem())
	case reflect.Ptr:
		if left.Pointer() == right.Pointer() {
			return true
		}
		return matchValue(left.Elem(), right.Elem())
	case reflect.Map:
		if left.IsNil() != right.IsNil() {
			return false
		}
		if left.Pointer() == right.Pointer() {
			return true
		}
		for _, k := range right.MapKeys() {
			val1 := left.MapIndex(k)
			val2 := right.MapIndex(k)
			if !val1.IsValid() || !val2.IsValid() || !matchValue(val1, val2) {
				return false
			}
		}
		return true
	default:
		return left.Interface() == right.Interface()
	}
}
