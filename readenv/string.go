package readenv

import "github.com/escaletech/go-escale/messages"

// Validate if env variable exists
func (r *EnvReader) StringRequired(key string) string {
	val := r.Env(key)
	if val == "" {
		r.Errs = append(r.Errs, messages.MissingRequiredEnv(key))
	}

	return val
}

// Retrieve env variable value and if doesn't exists returns a default value.
func (r *EnvReader) StringOrDefault(key, def string) string {
	val := r.Env(key)
	if val == "" {
		return def
	}

	return val
}