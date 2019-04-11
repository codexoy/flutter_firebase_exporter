package csvmaker

import (
	"fmt"
	"reflect"
)

// Slice Slice
func Slice(inter reflect.Value, sliceName string, keys, values *[]string) error {

	for k := 0; k < inter.Len(); k++ {
		switch inter.Index(k).Type().Kind() {
		case reflect.String:
			*keys = append(*keys, fmt.Sprintf("%s[%d]", sliceName, k))
			*values = append(*values, inter.Index(k).Interface().(string))
		case reflect.Bool:

			*keys = append(*keys, fmt.Sprintf("%s[%d]", sliceName, k))
			*values = append(*values, fmt.Sprintf("%t", inter.Index(k).Interface().(bool)))
		case reflect.Int:

			*keys = append(*keys, fmt.Sprintf("%s[%d]", sliceName, k))
			*values = append(*values, fmt.Sprintf("%d", inter.Index(k).Interface().(int)))
		case reflect.Int8:

			*keys = append(*keys, fmt.Sprintf("%s[%d]", sliceName, k))
			*values = append(*values, fmt.Sprintf("%d", inter.Index(k).Interface().(int8)))
		case reflect.Int16:

			*keys = append(*keys, fmt.Sprintf("%s[%d]", sliceName, k))
			*values = append(*values, fmt.Sprintf("%d", inter.Index(k).Interface().(int16)))
		case reflect.Int32:

			*keys = append(*keys, fmt.Sprintf("%s[%d]", sliceName, k))
			*values = append(*values, fmt.Sprintf("%d", inter.Index(k).Interface().(int32)))
		case reflect.Int64:
			*keys = append(*keys, fmt.Sprintf("%s[%d]", sliceName, k))
			*values = append(*values, fmt.Sprintf("%d", inter.Index(k).Interface().(int64)))
		default:
			return fmt.Errorf("case for kind %v is not implemented", inter.Index(k).Type().Kind().String())
		}

	}
	return nil
}
