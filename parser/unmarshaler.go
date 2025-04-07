package parser

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
)

var (
	ErrStructPointer = errors.Join(
		models.ErrServerError,
		errors.New("value must be a struct pointer"),
	)
)

func (c *Parser) Unmarshal(v any) error {
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

func (c *Parser) getReflection(v any) (reflect.Value, reflect.Type, error) {
	vPointer := reflect.ValueOf(v)
	zero := reflect.Zero(vPointer.Type())
	if vPointer.Kind() != reflect.Pointer {
		c.WriteErrorResponse(ErrStructPointer)
		return zero, nil, ErrStructPointer
	}
	vType := vPointer.Type().Elem()
	if vType.Kind() != reflect.Struct {
		c.WriteErrorResponse(ErrStructPointer)
		return zero, nil, ErrStructPointer
	}
	return vPointer, vType, nil
}

func (c *Parser) unmarshalKey(
	field reflect.Value,
	typeField reflect.StructField,
) error {
	key := typeField.Tag.Get("json")
	v, ok := c.Body[key]
	if !ok {
		err := errors.New(fmt.Sprintf("field '%v' missing", key))
		c.WriteErrorResponse(errors.Join(models.ErrBadRequest, err))
		return err
	}
	vValue := reflect.ValueOf(v)
	if field.Kind() != vValue.Kind() {
		err := errors.New(fmt.Sprintf("field '%v' not %v", key, field.Kind()))
		c.WriteErrorResponse(errors.Join(models.ErrBadRequest, err))
		return err
	}
	field.Set(vValue)
	return nil
}
