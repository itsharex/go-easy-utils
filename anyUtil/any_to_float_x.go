package anyUtil

import (
	"github.com/jefferyjob/go-easy-utils"
	"math"
	"reflect"
	"strconv"
)

// AnyToFloat32 将给定的值转换为float32
func AnyToFloat32(value interface{}) (float32, error) {
	f64, err := AnyToFloat64(value)
	if err != nil {
		return 0, err
	}
	if f64 < -math.MaxFloat32 || f64 > math.MaxFloat32 {
		return 0, go_easy_utils.ErrValOut
	}
	return float32(f64), nil
}

// AnyToFloat64 将给定的值转换为float64
func AnyToFloat64(i interface{}) (float64, error) {
	if i == nil {
		return 0, nil
	}

	// 检查解引用后的值是否为 nil
	if reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil() {
		return 0, nil
	}

	v := reflect.ValueOf(i)
	// 处理指针类型
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		if reflect.ValueOf(i).IsNil() {
			return 0, nil
		}
		v = reflect.ValueOf(i).Elem()
	}

	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return v.Float(), nil
	case reflect.String:
		floatValue, err := strconv.ParseFloat(v.String(), 64)
		if err != nil {
			return 0, go_easy_utils.ErrSyntax
		}
		return floatValue, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(v.Uint()), nil
	case reflect.Complex64, reflect.Complex128:
		return real(v.Complex()), nil
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, go_easy_utils.ErrType
	}
}
