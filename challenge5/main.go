package main

import (
	"challenge5/acak"
	"challenge5/rapih"
	"fmt"
)

func main(){
	fmt.Println("ini acak \n")
	acak.AcakGoRoutine()
	fmt.Println("ini rapih")
	rapih.RapihGoRoutine()
}
