package main

import "fmt"

func main(){
	challenge1()
}

func challenge1() {
	i := 21
	fmt.Printf("nilai i: %v \n", i)
	fmt.Printf("tipe data dari i: %T \n", i)
	fmt.Printf("tanda persen: %% \n")

	j := true
	fmt.Printf("nilai boolean j: %v \n", j)

	russian := 'Я'
	fmt.Printf("unicode russia: %c \n", russian)

	num := 21
	fmt.Printf("nilai base 10: %d \n", num)
	fmt.Printf("nilai base 8: %o \n", num)
	fmt.Printf("nilai base 16 (kecil): %x \n", num)
	fmt.Printf("nilai base 16 (besar): %X \n", num)

	char := 'Я'
	fmt.Printf("unicode karakter Я: %U \n", char)

	k := 123.456
	fmt.Printf("float: %f \n", k)
	fmt.Printf("float scientific: %e \n", k)
}