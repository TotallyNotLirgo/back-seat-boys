package context

import (
	"reflect"
)

func (c *Context) Unmarshal(v any) error {
	vPointer, vType, err := c.getReflection(v)
	if err != nil {
		return err
	}
	err = c.GetJSON()
	if err != nil {
		return err
	}
	vNew := reflect.New(vType).Elem()
	vNewType := vNew.Type()
	for i := range vNewType.NumField() {
		err = c.unmarshalKey(vNew.Field(i), vNewType.Field(i))
		if err != nil {
			return err
		}
	}
	vPointer.Elem().Set(vNew)
	return nil
}

func (c *Context) getReflection(v any) (reflect.Value, reflect.Type, error) {
	vPointer := reflect.ValueOf(v)
	zero := reflect.Zero(vPointer.Type())
	if vPointer.Kind() != reflect.Pointer {
		return zero, nil, c.WriteError(500, "value must be a struct pointer")
	}
	vType := vPointer.Type().Elem()
	if vType.Kind() != reflect.Struct {
		return zero, nil, c.WriteError(500, "value must be a struct pointer")
	}
	return vPointer, vType, nil
}

func (c *Context) unmarshalKey(
	field reflect.Value,
	typeField reflect.StructField,
) error {
	key := typeField.Tag.Get("json")
	v, ok := c.Body[key]
	if !ok {
		return c.WriteError(422, "field '%v' missing", key)
	}
	vValue := reflect.ValueOf(v)
	if field.Kind() != vValue.Kind() {
		return c.WriteError(422, "field '%v' not %v", key, field.Kind())
	}
	field.Set(vValue)
	return nil
}
