package convertor

import "reflect"

type Json map[string]interface{}

func ConvertToTagMap(data interface{}) Json {
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()
	result := make(map[string]interface{}, size)
	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Tag.Get("json")
		value := elem.Field(i).Interface()
		result[field] = value
	}
	return result
}
