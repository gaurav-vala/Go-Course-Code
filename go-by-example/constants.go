package main

import "fmt"

const s string = "Constant"

func main() {
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	for j := 0; j <= 5; j++ {
		fmt.Println(j)
	}

	for i := range 5 {
		fmt.Println("Range: ", i)
	}

	for n := range 6 {
		if n%2 != 0 {
			continue
		}
		fmt.Println(n)
	}
}
