package utils

import (
	"reflect"
)

func ContainsAny(slice interface{}, items ...interface{}) bool {
	value := reflect.ValueOf(slice)
	if value.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}

	for _, curItem := range items {
		for i := 0; i < value.Len(); i++ {
			// fmt.Printf("%+v vs %+v\n", value.Index(i).Interface(), curItem)
			if value.Index(i).Interface() == curItem {
				return true
			}
		}
	}
	return false
}

func ContainsAll(slice interface{}, items ...interface{}) bool {
	value := reflect.ValueOf(slice)
	if value.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}
	for _, curItem := range items {
		itemInArray := false
		for i := 0; i < value.Len(); i++ {
			if value.Index(i).Interface() == curItem {
				itemInArray = true
				break
			}
		}
		if !itemInArray {
			return false
		}
	}
	return true
}
