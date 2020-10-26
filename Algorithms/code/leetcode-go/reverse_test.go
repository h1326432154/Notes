package leetcode

import (
	"strconv"
	"testing"
)

func TestReverse(t *testing.T) {

	in := -321
	res := reverse(in)
	t.Log(res)
}

// reverse 此思路是x不断取余，y加权累加
func reverse(x int) int {
	y := 0
	for x != 0 {
		y = y*10 + x%10
		if !(-(1<<31) <= y && y <= (1<<31)-1) {
			return 0
		}
		x /= 10
	}
	return y
}

// reverse1 转byte交换顺序，再转int
func reverse1(x int) int {
	u := 1
	if x < 0 {
		u = -1
		x = u * x
	}
	var (
		bx = []byte(strconv.Itoa(x))
		l  = len(bx)
	)

	for i := 0; i < l/2; i++ {
		bx[i] = bx[i] ^ bx[l-i-1]
		bx[l-i-1] = bx[i] ^ bx[l-i-1]
		bx[i] = bx[i] ^ bx[l-i-1]
	}
	x, _ = strconv.Atoi(string(bx))
	if !(-(1<<31) <= u*x && u*x <= (1<<31)-1) {
		return 0
	}

	return u * x
}
