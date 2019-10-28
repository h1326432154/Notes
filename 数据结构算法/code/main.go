package main

import (
	"fmt"
)

func main() {
	i := 0
	arr := [3]int{0}
	for ; i <= 3; i++ {
		arr[i] = 0
		// fmt.Printf("====%d", arr[i])
		fmt.Printf("i = %d | ", i)
		fmt.Printf("arr[i] = %d | ", arr[i])
	}
}
