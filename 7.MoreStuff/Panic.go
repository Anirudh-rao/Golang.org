package main

import "os"

func main() {
	panic("a Probem")

	_, err := os.Create("/temp/file")
	if err != nil {
		panic(err)
	}
}
