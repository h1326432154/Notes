package leetcode

import (
	"strconv"
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	x := []int{123321, 111, 0}
	for _, v := range x {
		if ok := isPalindrome(v); ok {
			t.Log(v, "是回文")
		} else {
			t.Log(v, "不是回文")
		}
	}
}

// 转成byte逐次比较
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := []byte(strconv.Itoa(x))

	for i := 0; i < len(s)/2-1; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
