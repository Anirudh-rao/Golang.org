package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(5)

	const n = 100000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
