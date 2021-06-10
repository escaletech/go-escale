package httpserver

// Appends "localhost" to address on webserver initialization,
// when app env is "dev"
//
// (prevents macOS annoying Firewall request for authorization
//	everytime the application restarts)
func SetAddr(env, port string) string {
	if env == "dev" {
		return "localhost:" + port
	}

	return ":" + port
}
