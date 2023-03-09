package main

import "fmt"

func main() {
	looping()
}

func looping() {
    
    var arrays = [13]int{2,4,6,8,10,12,14,16,18,20} 
	for i := 0; i < len(arrays); i++{

    fmt.Printf("angka ke %d \n",arrays[i])
    
	}
}


