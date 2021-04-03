package config

import "os"

func CORSAllowOrigin() string {
	return os.Getenv("CORS_ALLOW_ORIGIN")
}
