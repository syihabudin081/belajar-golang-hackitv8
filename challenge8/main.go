package main

import (
	"challenge8/routers"
	_ "github.com/lib/pq"
)



func main(){

	var PORT = ":8080"
	routers.StartServer().Run(PORT)
}