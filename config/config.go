package config

import (
	"os"
)

var (
	DatabaseURL  string
	JwtSecret    string
	FirebaseJson string
)

func init() {
	DatabaseURL = os.Getenv("DATABASE_URL")
	JwtSecret = os.Getenv("JWT_SECRET")
	FirebaseJson = os.Getenv("FIREBASE_JSON")
}
