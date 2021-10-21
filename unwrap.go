package goutils

import "reflect"

func UnwrapType(raw reflect.Type) reflect.Type {
	if raw == nil {
		return raw
	} else if raw.Kind() == reflect.Ptr {
		return UnwrapType(raw.Elem())
	} else {
		return raw
	}
}

func UnwrapValue(raw reflect.Value) reflect.Value {
	if !raw.IsValid() {
		return raw
	} else if raw.Kind() == reflect.Ptr {
		return UnwrapValue(raw.Elem())
	} else {
		return raw
	}
}
