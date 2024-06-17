package config

import "os"

var (
	AppKey  string
	AppMode string

	LogPath string

	IsDebugMode bool
)

func init() {
	AppKey = os.Getenv("APP_KEY")
	AppMode = os.Getenv("APP_MODE")
	LogPath = os.Getenv("LOG_PATH")

	if AppMode == "debug" {
		IsDebugMode = true
	}
}
