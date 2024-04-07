package flatten

import "reflect"

func Flatten(nested interface{}) []interface{} {
	out := []interface{}{}
	rv := reflect.ValueOf(nested)

	if rv.Len() == 0 {
		return out
	}

	for i := 0; i < rv.Len(); i++ {
		switch rv.Index(i).Elem().Kind() {
		case reflect.Slice:
			out = append(out, Flatten(rv.Index(i).Interface())...)
		case reflect.Int:
			out = append(out, rv.Index(i).Interface())
		}
	}

	return out

}
