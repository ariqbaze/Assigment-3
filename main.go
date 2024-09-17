package main

import (
	"Assigment-3/database"
	"Assigment-3/router"
)

func main() {
	database.StartDB()

	r := router.StartApp()

	r.Run(":8080")
}
