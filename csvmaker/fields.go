package csvmaker

import (
	"fmt"
	"reflect"
)

// GetFields GetFields
func GetFields(i interface{}, keys, values *[]string) error {

	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	for j := 0; j < t.NumField(); j++ {
		switch t.Field(j).Type.Kind() {
		case reflect.Ptr:
			Pointer(v.Field(j), keys, values)
		case reflect.Struct:
			if err := GetFields(v.Field(j), keys, values); err != nil {
				return err
			}
		case reflect.String:
			*keys = append(*keys, t.Field(j).Name)
			*values = append(*values, v.Field(j).Interface().(string))
		case reflect.Bool:
			*keys = append(*keys, t.Field(j).Name)
			*values = append(*values, fmt.Sprintf("%t", v.Field(j).Interface().(bool)))

		default:
			return fmt.Errorf("case for kind %v is not implemented", t.Field(j).Type.Kind().String())
		}
	}
	return nil
}

// Pointer Pointer
func Pointer(inter reflect.Value, keys, values *[]string) error {
	val := inter.Elem()
	if val.CanAddr() {
		if err := GetFields(val.Interface(), keys, values); err != nil {
			return err
		}
	}
	return nil
}
