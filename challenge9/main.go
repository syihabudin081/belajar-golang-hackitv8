package main

import (
	"challenge9/database"
	"challenge9/routers"

	_ "github.com/lib/pq"
)

func main() {
	var PORT = ":8080"
	
	database.StartDB() // initialize the database
	
	routers.StartServer().Run(PORT)
}