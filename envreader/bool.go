package envreader

import (
	"fmt"
	"strconv"
)

func (r *EnvReader) BoolOrDefault(key string, def bool) bool {
	var ret bool
	val := r.Env(key)
	fmt.Println(val)
	if val == "" {
		return def
	}
	
	ret, _ = strconv.ParseBool(val)
	return ret
}