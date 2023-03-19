package acak

import (
	"fmt"
	"time"
)

func printData1(data interface{}) {
	for i := 0; i < 4; i++ {
		fmt.Println(data)
	}
}

func printData2(data interface{}) {
	for i := 0; i < 4; i++ {
		fmt.Println( data)
	}
}

func AcakGoRoutine() {

	data1 := []interface{}{"bisal", "bisa2", "bisa3"}
	data2 := []interface{}{"cobal", "coba2", "coba3"}

	for i := 0; i < 4; i++ {
		go printData1(data1)
		go printData2(data2)
	}

	time.Sleep(1 * time.Second)
}