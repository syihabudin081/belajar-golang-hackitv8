package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i =", i)
	}

	for j := 0; j <= 10; j++ {

		if j < 5 {
			fmt.Println("Nilat j =", j)
		} else if j == 5 {

			s := "CABVBО"
			for k, v := range s {
				fmt.Printf("character %U '%c' starts at byte position %d\n", v, v, k*2)
			}
		} else {

			fmt.Println("Nilai j =", j)
		}
	}
}
