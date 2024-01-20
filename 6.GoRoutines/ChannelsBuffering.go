package main

import "fmt"

func main() {
	messgaes := make(chan string, 2)

	messgaes <- "buffered"
	messgaes <- "channels"
	fmt.Println(<-messgaes)
	fmt.Println(<-messgaes)
}
