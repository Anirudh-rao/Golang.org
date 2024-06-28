package conditonals

import "fmt"

func ForLoop() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1

	}
	for j := 0; j < 3; j++ {
		fmt.Println("The Value of j", j)
	}
	for i := range 3 {
		fmt.Println("The Value of i", i)
	}
	for {
		fmt.Println("Loop without conditons")
		break
	}
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
