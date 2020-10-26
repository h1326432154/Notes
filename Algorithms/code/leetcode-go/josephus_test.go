package leetcode

import (
	"fmt"
	"testing"
)

// 约瑟夫环问题
func TestJosephus(t *testing.T) {
	res := lastRemaining(6, 2)
	t.Logf("返回结果为 ： %d", res)
}

func lastRemaining(n int, m int) int {
	fmt.Println("n == ", n, "m == ", m)
	if n == 1 {
		return 0
	}
	return (m + lastRemaining(n-1, m)) % n
}
