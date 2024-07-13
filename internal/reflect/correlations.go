package reflect

import (
	"errors"
	"reflect"
)

func (r *Reflect) CorrelateRawToValue(raw any, value any, tag string) error {
	return r.correlateRawToValueReflect(reflect.ValueOf(raw), reflect.ValueOf(value), tag)
}

func (r *Reflect) correlateRawToValueReflect(raw reflect.Value, value reflect.Value, tag string) (err error) {
	kind := raw.Kind()
	if kind == reflect.Invalid {
		return nil
	}

	if kind == reflect.Pointer || kind == reflect.Interface {
		return r.correlateRawToValueReflect(raw.Elem(), value, tag)
	}

	switch kind {
	case reflect.Struct:
		value.Set(raw)
	case reflect.Map:
		err = r.correlateMapToValueReflect(raw, value, tag)
	case reflect.Slice, reflect.Array:
		err = r.correlateArrayToValueReflect(raw, value, tag)
	default:
		err = r.correlatePrimitiveToValueReflect(raw, value)
	}

	return
}

func (r *Reflect) CorrelateMapToValue(m map[string]any, v any, tag string) error {
	return r.correlateMapToValueReflect(reflect.ValueOf(m), reflect.ValueOf(v), tag)
}

func (r *Reflect) correlateMapToValueReflect(m reflect.Value, v reflect.Value, tag string) error {
	if m.Kind() != reflect.Map {
		return errors.New("invalid type")
	}

	mapRange := m.MapRange()
	for mapRange.Next() {
		key := mapRange.Key().String()
		field, err := r.getStructFieldForValue(v, key, tag)
		if err != nil {
			return err
		}

		err = r.correlateRawToValueReflect(mapRange.Value(), field, tag)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Reflect) correlateArrayToValueReflect(a reflect.Value, v reflect.Value, tag string) error {
	if a.Kind() != reflect.Slice && a.Kind() != reflect.Array {
		return errors.New("invalid type")
	}

	s := reflect.MakeSlice(r.GetOriginalType(v.Type()), a.Len(), a.Len())
	for i := range a.Len() {
		raw := a.Index(i)
		value := s.Index(i)

		if err := r.correlateRawToValueReflect(raw, value, tag); err != nil {
			return err
		}
	}

	v.Set(s)
	return nil
}

func (r *Reflect) correlatePrimitiveToValueReflect(raw reflect.Value, value reflect.Value) error {
	value.Set(raw.Convert(value.Type()))
	return nil
}
