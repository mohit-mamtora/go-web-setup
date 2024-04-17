package config

import "os"

var (
	AppKey  string
	AppMode string

	IsDebugMode bool
)

func init() {
	AppKey = os.Getenv("APP_KEY")
	AppMode = os.Getenv("APP_MODE")

	if AppMode == "debug" {
		IsDebugMode = true
	}
}
