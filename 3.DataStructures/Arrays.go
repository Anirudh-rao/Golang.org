package datastructures

import "fmt"

func Arrays() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("Set:", a)
	fmt.Println("Get:", a[4])

	fmt.Println("Length of the Array:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("B Array", b)

	b1 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("New B Array", b1)

	b3 := [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b3)

	var TwoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			TwoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", TwoD)

	twoD := [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)

}
