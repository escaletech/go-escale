package envreader

import (
	"strconv"
)

func (r *EnvReader) BoolOrDefault(key string, def bool) bool {
	var ret bool
	val := r.Env(key)
	if val == "" {
		return def
	}
	
	ret, _ = strconv.ParseBool(val)
	return ret
}