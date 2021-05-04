package validator

import (
	"os"
	"strconv"
)

func (v *Validator) StringRequired(key string) string {
	val := os.Getenv(key)
	if val == "" {
		v.Errs = append(v.Errs, "missing "+key)
	}

	return val
}

func (v *Validator) StringOrDefault(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}

	return val
}

func (v *Validator) BoolOrDefault(key, def string) bool {
	var ret bool
	val := os.Getenv(key)
	if val == "" {
		ret, _ = strconv.ParseBool(def)
		return ret
	}

	ret, _ = strconv.ParseBool(val)
	return ret
}

func (v *Validator) IntOrDefault(key string, def int) int {
	val := os.Getenv(key)
	if val == "" {
		return def
	}

	num, err := strconv.Atoi(val)
	if err != nil {
		v.Errs = append(v.Errs, "unable to convert "+key+" to int")
	}

	return num
}
