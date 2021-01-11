package main

import (
	"app/database"
	"app/router"
)

func main() {
	database.Init()
	router.Init()
	database.Close()
}
