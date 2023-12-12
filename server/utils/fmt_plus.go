package utils

import (
	// "fmt"
	"reflect"
	// "strings"
)

//StructToMap
//use Structure convert for map
//obj interface{}
//map[string]interface{}
func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Tag.Get("mapstructure") != "" {
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

//ArrayToString use  number group format for
// func ArrayToString(array []interface{}) string {
// 	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
// }
