package reflect

import (
	"errors"
	"reflect"
)

func (r *Reflect) EnsureIsPointer(v any) error {
	kind := reflect.TypeOf(v).Kind()
	if kind != reflect.Pointer {
		return errors.New("expected pointer")
	}

	return nil
}

func (r *Reflect) GetOriginalType(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return r.GetOriginalType(t.Elem())
	}

	return t
}
