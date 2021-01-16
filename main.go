package main

import "app/infrastructure"

func main() {
	db, err := infrastructure.NewDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	r := infrastructure.NewRouter(db)
	r.Run()
}
