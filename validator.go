package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/mkdr4/validator/internal/check"
	"github.com/mkdr4/validator/internal/consts"
	"github.com/mkdr4/validator/internal/errs"
)

func initTV(data interface{}, validType reflect.Kind) (t reflect.Type, v reflect.Value, err error) {
	t, v = reflect.TypeOf(data), reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != validType {
		return t, v, errors.New("")
	}

	return t, v, nil
}

func Struct(data interface{}) error {
	t, v, err := initTV(data, reflect.Struct)
	if err != nil {
		return fmt.Errorf(errs.VARIABLE_NOT_STRUCT)
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		if value.Kind() == reflect.Struct {
			return Struct(data)
		}

		if validParams := field.Tag.Get("valid"); validParams != "" {
			for _, params := range strings.Split(validParams, ",") {
				ps := strings.Split(params, "=")
				paramsName := ps[0]
				switch paramsName {
				case "required":
					if !check.RequiredCompliance(value) {
						return fmt.Errorf(errs.REQUIRES_VALUE_ABSENT, field.Name)
					}
				case "max":
					if len(ps) > 1 {
						paramsValue := ps[1]
						if !check.MinMaxLenCompliance(value, paramsValue, consts.MaxMode) {
							return fmt.Errorf(errs.LENGTH_VALUE_INCORRECT, field.Name)
						}
					}
				case "min":
					if len(ps) > 1 {
						paramsValue := ps[1]
						if !check.MinMaxLenCompliance(value, paramsValue, consts.MinMode) {
							return fmt.Errorf(errs.LENGTH_VALUE_INCORRECT, field.Name)
						}
					}
				}
			}
		}
	}

	return nil
}
