package slicer

import "reflect"

// Verifies if a slice has a determined element of any type (including structs)
// (similar to Javascript's Array.prototype.includes() function)
func Includes(searchedValue, referenceSlice interface{}) bool {
	ref := reflect.ValueOf(referenceSlice)

	for i := 0; i < ref.Len(); i++ {
		refToInterface := ref.Index(i).Interface()

		if reflect.DeepEqual(searchedValue, refToInterface) {
			return true
		}
	}

	return false
}
