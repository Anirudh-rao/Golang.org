package main

import "fmt"

func zeroval(inval int) {

	inval = 0
}

func zerovalprt(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("intial :", i)
	zeroval(i)
	fmt.Println("zeroval:", i)
	zerovalprt(&i)
	fmt.Println("zerovalprt:", i)
	fmt.Println("Pointer:", &i)
}
