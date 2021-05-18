package envreader

import (
	"strconv"

	"github.com/escaletech/go-escale/messages"
)

func (r *EnvReader) BoolOrDefault(key string, def bool) bool {
	val := r.Env(key)
	if val == "" {
		return def
	}
	
	ret, err := strconv.ParseBool(val)
	if err != nil {
		r.Errs = append(r.Errs, messages.UnableToConvertToBool(key))
	}
	return ret
}