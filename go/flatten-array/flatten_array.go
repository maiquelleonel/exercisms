package flatten

import (
	"reflect"
)

func Flatten(nested interface{}) []interface{} {
	out := []interface{}{}
	rv := reflect.ValueOf(nested)
	if rv.Kind() == reflect.Slice {
		if rv.Len() == 0 {
			return out
		}
		for i := 0; i < rv.Len(); i++ {
			if rv.Index(i).Interface() == nil {
				continue
			}

			if rv.Index(i).Elem().Kind() == reflect.Slice {
				out = append(out, Flatten(rv.Index(i).Interface())...)
			} else {
				out = append(out, rv.Index(i).Interface())
			}

		}

	}

	return out

}
