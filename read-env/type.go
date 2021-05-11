package readenv

type EnvReader struct {
	Errs []string
	Env func(key string) string
}