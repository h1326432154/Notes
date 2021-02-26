package test

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	// 错误
	// {
	// 	type MyInt int
	// 	var i int = 1
	// 	var j MyInt = i
	// }

	// {
	// 	type MyInt int
	// 	var i int = 1
	// 	var j MyInt = MyInt(i)
	// 	fmt.Print(j)
	// }
	{
		type MyInt int
		var i interface{} = 1
		var j = i.(MyInt)
		fmt.Println(j)
	}
	x := []int{1, 2, 3}
}
