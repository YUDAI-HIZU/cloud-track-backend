package main

import "app/infrastructure"

func main() {
	db := infrastructure.NewDatabase()
	defer db.Close()
	app := infrastructure.NewFirebaseApp()
	r := infrastructure.NewRouter(db, app)
	r.Run()
}
