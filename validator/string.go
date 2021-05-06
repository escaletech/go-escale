package validator

import (
	"github.com/escaletech/go-escale/messages"
)

func (v *Validator) StringRequired(key string) string {
	val := v.Env(key)
	if val == "" {
		errMessage := messages.RequiredEnvvar(key)
		v.Errs = append(v.Errs, errMessage)
	}

	return val
}

func (v *Validator) StringOrDefault(key, def string) string {
	val := v.Env(key)
	if val == "" {
		return def
	}

	return val
}
