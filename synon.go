package synon

import (
	"reflect"
)

func Merge(l, r interface{}) interface{} {
	lvt := reflect.TypeOf(l)
	rvt := reflect.TypeOf(r)

	// Ensure both interfaces are of the same type
	if lvt != rvt {
		return nil
	}
	lv := reflect.ValueOf(l)
	rv := reflect.ValueOf(r)

	switch lv.Kind() {
	case reflect.Map:
		var keys []reflect.Value
		keys = append(keys, lv.MapKeys()...)
		keys = append(keys, rv.MapKeys()...)
		for _, k := range keys {
			if !rv.MapIndex(k).IsValid() {
				continue
			}
			if !lv.MapIndex(k).IsValid() {
				lv.SetMapIndex(k, rv.MapIndex(k))
				continue
			}
			v := Merge(lv.MapIndex(k).Interface(), rv.MapIndex(k).Interface())
			lv.SetMapIndex(k, reflect.ValueOf(v))
		}
		return lv.Interface()
	case reflect.Slice:
		if rv.IsNil() {
			return lv.Interface()
		}
		values := reflect.AppendSlice(lv, rv)
		newValues := reflect.MakeSlice(values.Type(), 0, values.Len())
		var elements = reflect.MakeMap(reflect.TypeOf(map[interface{}]interface{}{}))

		for i := 0; i < values.Len(); i++ {
			key := values.Index(i)
			if elements.MapIndex(key).IsValid() {
				continue
			}
			elements.SetMapIndex(key, values.Index(i))
			newValues = reflect.Append(newValues, values.Index(i))
		}
		return newValues.Interface()
	case reflect.Chan:
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
		return Merge(lv.Interface(), rv.Interface())
	case reflect.Struct:
		newStr := reflect.New(lv.Type()).Elem()
		for i := 0; i < lv.NumField(); i++ {
			vtl := lv.Type()
			if !vtl.Field(i).IsExported() {
				continue // unexported field
			}
			ls := lv.Field(i)
			rs := rv.Field(i)
			v := reflect.ValueOf(Merge(ls.Interface(), rs.Interface()))
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
