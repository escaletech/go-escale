package validator

type Validator struct {
	Errs []string
	Env  func(key string) string
}
