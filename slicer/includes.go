package slicer

import (
	"errors"
	"reflect"

	"github.com/escaletech/go-escale/messages"
)

// Verifies if a slice has a determined element of any type (including structs)
// (similar to Javascript's Array.prototype.includes() function)
func Includes(searchedValue, referenceSlice interface{}) (*bool, error) {
	if err := validateRefenceSlice(referenceSlice); err != nil {
		return nil, err
	}

	ref := reflect.ValueOf(referenceSlice)

	response := false
	for i := 0; i < ref.Len(); i++ {
		refToInterface := ref.Index(i).Interface()

		if reflect.DeepEqual(searchedValue, refToInterface) {
			response = true
		}
	}

	return &response, nil
}

func validateRefenceSlice(ref interface{}) error {
	r := reflect.TypeOf(ref)
	refType := r.Kind()

	if refType == reflect.Slice || refType == reflect.Array {
		return nil
	}

	return errors.New(messages.RefNotSliceOrArray)
}
