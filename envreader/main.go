package envreader

import (
	"os"
	"strings"
)

var Env func(key string) string

func init() {
	Env = os.Getenv
}

// Create a new envReader
func New() *EnvReader {
	return &EnvReader{
		Env: Env,
	}
}

// Check if envReader has errors
func (r *EnvReader) HasErrors() bool {
	return len(r.Errs) > 0
}

// Returns envReader Errs slice items as a string
func (r *EnvReader) GetErrors() string {
	return strings.Join(r.Errs, ", ")
}
