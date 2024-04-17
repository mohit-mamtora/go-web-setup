package config

import "os"

var (
	ServerPort = ":8181"
)

func init() {
	ServerPort = ":" + os.Getenv("APP_PORT")
}
