package validator

import (
	"os"
	"strconv"
)

func (v *Validator) StringRequired(key string) string {
	val := os.Getenv(key)
	if val == "" {
		v.errs = append(v.errs, "missing "+key)
	}

	return val
}

func (v *Validator) stringOrDefault(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}

	return val
}

func (v *Validator) boolOrDefault(key, def string) bool {
	var ret bool
	val := os.Getenv(key)
	if val == "" {
		ret, _ = strconv.ParseBool(def)
		return ret
	}

	ret, _ = strconv.ParseBool(val)
	return ret
}

func (v *Validator) intOrDefault(key string, def int) int {
	val := os.Getenv(key)
	if val == "" {
		return def
	}

	num, err := strconv.Atoi(val)
	if err != nil {
		v.errs = append(v.errs, "unable to convert "+key+" to int")
	}

	return num
}
