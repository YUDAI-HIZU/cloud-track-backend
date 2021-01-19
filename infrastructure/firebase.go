package infrastructure

import (
	"app/config"
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var app *firebase.App

func init() {
	ctx := context.Background()
	var err error
	app, err = firebase.NewApp(ctx, nil, option.WithCredentialsJSON([]byte(config.FirebaseJson)))
	if err != nil {
		log.Fatalf("failed to new firebase app: %v", err)
	}
}

func NewFirebaseApp() *firebase.App {
	return app
}
