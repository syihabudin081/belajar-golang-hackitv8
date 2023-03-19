// package main

// import "fmt"

// func main() {
// 	looping()
// }

// func looping() {
    
//     var arrays = [13]int{2,4,6,8,10,12,14,16,18,20} 
// 	for i := 0; i < len(arrays); i++{

//     fmt.Printf("angka ke %d \n",arrays[i])
    
// 	}
// }

package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 5, 4, 7, 3, 4, 4, 6, 5, 7, 9, 9}

	jumlah := make(map[int]int)
	for _, value := range arr {
		jumlah[value]++
	}

	for key, value := range jumlah {
		if value > 1 {
			fmt.Print(key, " \n")
		}
	}
	fmt.Println()
}

