package synon

import (
	"reflect"
)

func Synon(l, r interface{}) interface{} {
	lvt := reflect.TypeOf(l)
	rvt := reflect.TypeOf(r)

	// Ensure both interfaces are of the same type
	if lvt != rvt {
		return nil
	}
	lv := reflect.ValueOf(l)
	rv := reflect.ValueOf(r)

	switch lv.Kind() {
	case reflect.Map, reflect.Array, reflect.Slice, reflect.Chan:
		// TODO: Add key search function
		if rv.Len() > 0 {
			return rv.Interface()
		}
		return lv.Interface()
	case reflect.Ptr:
		if lv.IsNil() {
			return rv.Interface()
		}
		if rv.IsNil() {
			return lv.Interface()
		}
		lv = lv.Elem()
		rv = rv.Elem()
		return Synon(lv.Interface(), rv.Interface())
	case reflect.Struct:
		newStr := reflect.New(lv.Type()).Elem()
		for i := 0; i < lv.NumField(); i++ {
			vtl := lv.Type()
			if !vtl.Field(i).IsExported() {
				continue // unexported field
			}
			ls := lv.Field(i)
			rs := rv.Field(i)
			v := reflect.ValueOf(Synon(ls.Interface(), rs.Interface()))
			if v.IsValid() {
				newStr.Field(i).Set(v)
			}
		}
		return newStr.Addr().Interface()
	default:
		if !rv.IsValid() && !lv.IsValid() {
			return nil
		}
		if rv.IsValid() && rv.IsZero() {
			return lv.Interface()
		}
		return rv.Interface()
	}
}
