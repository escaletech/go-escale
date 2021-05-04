package validator

import (
	"os"
)

func (v *Validator) StringRequired(key string) string {
	val := os.Getenv(key)
	if val == "" {
		v.errs = append(v.errs, "missing "+key)
	}

	return val
}
