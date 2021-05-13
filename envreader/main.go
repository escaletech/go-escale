package envreader

import "strings"

// Create a new envReader
func NewEnvReader(env func(key string) string) *EnvReader {
	return &EnvReader{
		Env: env,
	}
}

// Check if envReader Errs slice is not empty
func (r *EnvReader) HasErrors() bool {
	if len(r.Errs) > 0 {
		return true
	}

	return false
}

// Returns envReader Errs slice items as a string
func (r *EnvReader) GetErrors() string {
	return strings.Join(r.Errs, ", ")
}
