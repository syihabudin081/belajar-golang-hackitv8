package main

import (
	"fmt"
	"strings"
)

func main() {

	stringLooping2()
}

// func stringLooping(inputString string) {
// 	words := strings.Split(inputString, " ")
// 	counts := make(map[string]int)

// 	for _, word := range words {
// 		counts[word]++
// 	}

// 	for _, char := range words {
// 		fmt.Printf( char)
// 	}

// 	fmt.Println(counts)

// }

func stringLooping2(){
	text := "selamat malam"
	words := strings.Split(text, " ")
	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}

	for _, char := range text {
		fmt.Printf("%c\n", char)
	}

	fmt.Println(counts)
}
