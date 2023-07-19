package check

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/mkdr4/validator/internal/consts"
)

func RequiredCompliance(v reflect.Value) bool {
	return !v.IsZero()
}

// work only on Array, Chan, Map, Slice, String
func MinMaxLenCompliance(v reflect.Value, pv string, m consts.Mode) bool {
	vt := v.Kind()
	switch vt {
	case reflect.Array, reflect.Slice, reflect.String, reflect.Chan, reflect.Map:
		if i, err := strconv.Atoi(pv); err == nil {
			if m == consts.MinMode && i <= v.Len() {
				return true
			} else if m == consts.MaxMode && i >= v.Len() {
				return true
			}
		}
	}
	return false
}

func TypeCompliance(v reflect.Value, t string) bool {
	fmt.Println(reflect.ValueOf(v.Elem()).Kind())
	return v.Kind().String() == t
}
