package envreader

import (
	"strconv"

	"github.com/escaletech/go-escale/messages"
)

func (r *EnvReader) IntOrDefault(key string, def int) int {
	val := r.Env(key)
	if val == "" {
		return def
	}

	num, err := strconv.Atoi(val)
	if err != nil {
		r.Errs = append(r.Errs, messages.UnableToConvertToInt(key))
	}

	return num
}