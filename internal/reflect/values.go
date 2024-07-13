package reflect

import (
	"errors"
	"reflect"
)

func (r *Reflect) getStructFieldForValue(s reflect.Value, fieldName string, tag string) (reflect.Value, error) {
	kind := s.Kind()
	if kind == reflect.Ptr || kind == reflect.Interface {
		return r.getStructFieldForValue(s.Elem(), fieldName, tag)
	}

	if kind != reflect.Struct {
		return reflect.Value{}, errors.New("invalid type: " + s.String())
	}

	tp := s.Type()
	for i := 0; i < tp.NumField(); i++ {
		f := tp.Field(i)
		t := f.Tag.Get(tag)
		if t == fieldName {
			return s.Field(i), nil
		}
	}

	return reflect.Value{}, errors.New("field not found")
}
