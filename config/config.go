package config

import "os"

var DatabaseURL string

func init() {
	DatabaseURL = os.Getenv("DATABASE_URL") + "?parseTime=true&loc=Local"
}
