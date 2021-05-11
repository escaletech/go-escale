package readenv

func (r *EnvReader) StringRequired(key string) string {
	val := r.Env(key)
	if val == "" {
		r.Errs = append(r.Errs, "missing "+key)
	}

	return val
}

func (r *EnvReader) StringOrDefault(key, def string) string {
	val := r.Env(key)
	if val == "" {
		return def
	}

	return val
}